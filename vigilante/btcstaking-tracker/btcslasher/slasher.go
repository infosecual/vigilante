package btcslasher

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/babylonlabs-io/vigilante/types"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"golang.org/x/sync/semaphore"

	bbn "github.com/babylonlabs-io/babylon/types"
	bstypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/babylonlabs-io/vigilante/btcclient"
	"github.com/babylonlabs-io/vigilante/metrics"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"go.uber.org/zap"
)

const (
	txSubscriberName          = "tx-subscriber"
	messageActionName         = "/babylon.finality.v1.MsgAddFinalitySig"
	consumerMessageActionName = "/babylon.finality.v1.MsgEquivocationEvidence"
	evidenceEventName         = "babylon.finality.v1.EventSlashedFinalityProvider.evidence"
)

type BTCSlasher struct {
	logger *zap.SugaredLogger

	// connect to BTC node
	BTCClient btcclient.BTCClient
	// BBNQuerier queries epoch info from Babylon
	BBNQuerier BabylonQueryClient

	// parameters
	netParams              *chaincfg.Params
	btcFinalizationTimeout uint32
	retrySleepTime         time.Duration
	maxRetrySleepTime      time.Duration
	maxRetryTimes          uint
	// channel for finality signature messages, which might include
	// equivocation evidences
	finalitySigChan <-chan coretypes.ResultEvent
	// channel for consumer fp equivocation evidences
	equivocationEvidenceChan <-chan coretypes.ResultEvent
	// channel for SKs of slashed finality providers
	slashedFPSKChan chan *btcec.PrivateKey
	// channel for receiving the slash result of each BTC delegation
	slashResultChan chan *SlashResult

	maxSlashingConcurrency int64

	metrics *metrics.SlasherMetrics

	startOnce sync.Once
	stopOnce  sync.Once
	wg        sync.WaitGroup
	quit      chan struct{}
}

func New(
	parentLogger *zap.Logger,
	btcClient btcclient.BTCClient,
	bbnQuerier BabylonQueryClient,
	netParams *chaincfg.Params,
	retrySleepTime time.Duration,
	maxRetrySleepTime time.Duration,
	maxRetryTimes uint,
	maxSlashingConcurrency uint8,
	slashedFPSKChan chan *btcec.PrivateKey,
	metrics *metrics.SlasherMetrics,
) (*BTCSlasher, error) {
	logger := parentLogger.With(zap.String("module", "slasher")).Sugar()

	return &BTCSlasher{
		logger:                 logger,
		BTCClient:              btcClient,
		BBNQuerier:             bbnQuerier,
		netParams:              netParams,
		retrySleepTime:         retrySleepTime,
		maxRetrySleepTime:      maxRetrySleepTime,
		maxRetryTimes:          maxRetryTimes,
		maxSlashingConcurrency: int64(maxSlashingConcurrency),
		slashedFPSKChan:        slashedFPSKChan,
		slashResultChan:        make(chan *SlashResult, 1000),
		quit:                   make(chan struct{}),
		metrics:                metrics,
	}, nil
}

func (bs *BTCSlasher) quitContext() (context.Context, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	bs.wg.Add(1)
	go func() {
		defer cancel()
		defer bs.wg.Done()

		select {
		case <-bs.quit:

		case <-ctx.Done():
		}
	}()

	return ctx, cancel
}

func (bs *BTCSlasher) LoadParams() error {
	if bs.btcFinalizationTimeout != 0 {
		// already loaded, skip
		return nil
	}

	btccParamsResp, err := bs.BBNQuerier.BTCCheckpointParams()
	if err != nil {
		return err
	}
	bs.btcFinalizationTimeout = btccParamsResp.Params.CheckpointFinalizationTimeout

	return nil
}

func (bs *BTCSlasher) Start() error {
	var startErr error
	bs.startOnce.Do(func() {
		// load module parameters
		if err := bs.LoadParams(); err != nil {
			startErr = err

			return
		}

		// start the subscriber to slashing events
		// NOTE: at this point monitor has already started the Babylon querier routine
		queryName := fmt.Sprintf("tm.event = 'Tx' AND message.action='%s'", messageActionName)
		// subscribe to babylon fp slashing events
		bs.finalitySigChan, startErr = bs.BBNQuerier.Subscribe(txSubscriberName, queryName)
		if startErr != nil {
			return
		}
		// subscribe to consumer fp slashing events
		queryName = fmt.Sprintf("tm.event = 'Tx' AND message.action='%s'", consumerMessageActionName)
		bs.equivocationEvidenceChan, startErr = bs.BBNQuerier.Subscribe(txSubscriberName, queryName)
		if startErr != nil {
			return
		}

		// BTC slasher has started
		bs.logger.Debugf("slasher routine has started subscribing %s", queryName)

		// start slasher
		bs.wg.Add(2)
		go bs.equivocationTracker()
		go bs.slashingEnforcer()

		bs.logger.Info("the BTC slasher has started")
	})

	return startErr
}

// slashingEnforcer is a routine that keeps receiving finality providers
// to be slashed and slashes their BTC delegations on Bitcoin
func (bs *BTCSlasher) slashingEnforcer() {
	defer bs.wg.Done()

	bs.logger.Info("slashing enforcer has started")

	// start handling incoming slashing events
	for {
		select {
		case <-bs.quit:
			bs.logger.Debug("handle delegations loop quit")

			return
		case fpBTCSK, ok := <-bs.slashedFPSKChan:
			if !ok {
				// slasher receives the channel from outside, so its lifecycle
				// is out of slasher's control. So we need to ensure the channel
				// is not closed yet
				bs.logger.Debug("slashedFKSK channel is already closed, terminating the slashing enforcer")

				return
			}
			// slash all the BTC delegations of this finality provider
			fpBTCPKHex := bbn.NewBIP340PubKeyFromBTCPK(fpBTCSK.PubKey()).MarshalHex()
			bs.logger.Infof("slashing finality provider %s", fpBTCPKHex)

			if err := bs.SlashFinalityProvider(fpBTCSK); err != nil {
				bs.logger.Errorf("failed to slash finality provider %s: %v", fpBTCPKHex, err)
			}
		case slashRes := <-bs.slashResultChan:
			if slashRes.Err != nil {
				bs.logger.Errorf(
					"failed to slash BTC delegation with staking tx hash %s under finality provider %s: %v",
					slashRes.Del.StakingTxHex,
					slashRes.Del.FpBtcPkList[0].MarshalHex(), // TODO: work with restaking
					slashRes.Err,
				)
			} else {
				bs.logger.Infof(
					"successfully slash BTC delegation with staking tx hash %s under finality provider %s",
					slashRes.Del.StakingTxHex,
					slashRes.Del.FpBtcPkList[0].MarshalHex(), // TODO: work with restaking
				)

				// record the metrics of the slashed delegation
				bs.metrics.RecordSlashedDelegation(slashRes.Del)
			}
		}
	}
}

func (bs *BTCSlasher) handleEvidence(evt *coretypes.ResultEvent, isConsumer bool) {
	evidence := filterEvidence(evt)

	if evidence == nil {
		return
	}

	fpBTCPKHex := evidence.FpBtcPk.MarshalHex()
	fpType := "babylon"
	if isConsumer {
		fpType = "consumer"
	}
	bs.logger.Infof("new equivocating %s finality provider %s to be slashed", fpType, fpBTCPKHex)
	bs.logger.Debugf("found equivocation evidence of %s finality provider %s: %v", fpType, fpBTCPKHex, evidence)

	// extract the SK of the slashed finality provider
	fpBTCSK, err := evidence.ExtractBTCSK()
	if err != nil {
		bs.logger.Errorf("failed to extract BTC SK of the slashed %s finality provider %s: %v", fpType, fpBTCPKHex, err)

		return
	}

	bs.slashedFPSKChan <- fpBTCSK
}

// equivocationTracker is a routine to track the equivocation events on Babylon,
// extract equivocating finality providers' SKs, and sen to slashing enforcer
// routine
func (bs *BTCSlasher) equivocationTracker() {
	defer bs.wg.Done()

	bs.logger.Info("equivocation tracker has started")

	// start handling incoming slashing events
	for {
		select {
		case <-bs.quit:
			bs.logger.Debug("handle delegations loop quit")

			return
		case resultEvent := <-bs.finalitySigChan:
			bs.handleEvidence(&resultEvent, false)
		case resultEvent := <-bs.equivocationEvidenceChan:
			bs.handleEvidence(&resultEvent, true)
		}
	}
}

// SlashFinalityProvider slashes all BTC delegations under a given finality provider
// the checkBTC option indicates whether to check the slashing tx's input is still spendable
// on Bitcoin (including mempool txs).
func (bs *BTCSlasher) SlashFinalityProvider(extractedFpBTCSK *btcec.PrivateKey) error {
	fpBTCPK := bbn.NewBIP340PubKeyFromBTCPK(extractedFpBTCSK.PubKey())
	bs.logger.Infof("start slashing finality provider %s", fpBTCPK.MarshalHex())

	// get all active and unbonded BTC delegations at the current BTC height
	// Some BTC delegations could be expired in Babylon's view but not expired in
	// Bitcoin's view. We will not slash such BTC delegations since they don't have
	// voting power (thus don't affect consensus) in Babylon
	activeBTCDels, unbondedBTCDels, err := bs.getAllActiveAndUnbondedBTCDelegations(fpBTCPK)
	if err != nil {
		return fmt.Errorf("failed to get BTC delegations under finality provider %s: %w", fpBTCPK.MarshalHex(), err)
	}

	// Initialize a mutex protected *btcec.PrivateKey
	safeExtractedFpBTCSK := types.NewPrivateKeyWithMutex(extractedFpBTCSK)

	// Initialize a semaphore to control the number of concurrent operations
	sem := semaphore.NewWeighted(bs.maxSlashingConcurrency)
	activeBTCDels = append(activeBTCDels, unbondedBTCDels...)
	delegations := activeBTCDels

	// try to slash both staking and unbonding txs for each BTC delegation
	// sign and submit slashing tx for each active and unbonded delegation
	for _, del := range delegations {
		bs.wg.Add(1)
		go func(d *bstypes.BTCDelegationResponse) {
			defer bs.wg.Done()
			ctx, cancel := bs.quitContext()
			defer cancel()

			// Acquire the semaphore before interacting with the BTC node
			if err := sem.Acquire(ctx, 1); err != nil {
				bs.logger.Errorf("failed to acquire semaphore: %v", err)

				return
			}
			defer sem.Release(1)

			safeExtractedFpBTCSK.UseKey(func(key *secp256k1.PrivateKey) {
				bs.slashBTCDelegation(fpBTCPK, key, d)
			})
		}(del)
	}

	bs.metrics.SlashedFinalityProvidersCounter.Inc()

	return nil
}

func (bs *BTCSlasher) WaitForShutdown() {
	bs.wg.Wait()
}

func (bs *BTCSlasher) Stop() error {
	var stopErr error
	bs.stopOnce.Do(func() {
		bs.logger.Info("stopping slasher")

		// close subscriber
		if err := bs.BBNQuerier.UnsubscribeAll(txSubscriberName); err != nil {
			bs.logger.Errorf("failed to unsubscribe from %s: %v", txSubscriberName, err)
		}

		// notify all subroutines
		close(bs.quit)
		bs.wg.Wait()

		bs.logger.Info("stopped slasher")
	})

	return stopErr
}
