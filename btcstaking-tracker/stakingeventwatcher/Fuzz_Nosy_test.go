package stakingeventwatcher

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/vigilante/config"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	gomock "github.com/golang/mock/gomock"
	"github.com/lightningnetwork/lnd/chainntnfs"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"go.etcd.io/etcd/client/v2"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_BabylonClientAdapter_ActivateDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var babylonClient *client.Client
		fill_err = tp.Fill(&babylonClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BTCStakingTrackerConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var proof *types.BTCSpvProof
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if babylonClient == nil || cfg == nil || proof == nil {
			return
		}

		bca := NewBabylonClientAdapter(babylonClient, cfg)
		bca.ActivateDelegation(ctx, stakingTxHash, proof)
	})
}

func Fuzz_Nosy_BabylonClientAdapter_BtcClientTipHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var babylonClient *client.Client
		fill_err = tp.Fill(&babylonClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BTCStakingTrackerConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if babylonClient == nil || cfg == nil {
			return
		}

		bca := NewBabylonClientAdapter(babylonClient, cfg)
		bca.BtcClientTipHeight()
	})
}

func Fuzz_Nosy_BabylonClientAdapter_DelegationsByStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var babylonClient *client.Client
		fill_err = tp.Fill(&babylonClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BTCStakingTrackerConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var status types.BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		var limit uint64
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if babylonClient == nil || cfg == nil {
			return
		}

		bca := NewBabylonClientAdapter(babylonClient, cfg)
		bca.DelegationsByStatus(status, offset, limit)
	})
}

func Fuzz_Nosy_BabylonClientAdapter_IsDelegationActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var babylonClient *client.Client
		fill_err = tp.Fill(&babylonClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BTCStakingTrackerConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		if babylonClient == nil || cfg == nil {
			return
		}

		bca := NewBabylonClientAdapter(babylonClient, cfg)
		bca.IsDelegationActive(stakingTxHash)
	})
}

func Fuzz_Nosy_BabylonClientAdapter_IsDelegationVerified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var babylonClient *client.Client
		fill_err = tp.Fill(&babylonClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BTCStakingTrackerConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		if babylonClient == nil || cfg == nil {
			return
		}

		bca := NewBabylonClientAdapter(babylonClient, cfg)
		bca.IsDelegationVerified(stakingTxHash)
	})
}

func Fuzz_Nosy_BabylonClientAdapter_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var babylonClient *client.Client
		fill_err = tp.Fill(&babylonClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BTCStakingTrackerConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		if babylonClient == nil || cfg == nil {
			return
		}

		bca := NewBabylonClientAdapter(babylonClient, cfg)
		bca.Params()
	})
}

func Fuzz_Nosy_BabylonClientAdapter_QueryHeaderDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var babylonClient *client.Client
		fill_err = tp.Fill(&babylonClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BTCStakingTrackerConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var headerHash *chainhash.Hash
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		if babylonClient == nil || cfg == nil || headerHash == nil {
			return
		}

		bca := NewBabylonClientAdapter(babylonClient, cfg)
		bca.QueryHeaderDepth(headerHash)
	})
}

func Fuzz_Nosy_BabylonClientAdapter_ReportUnbonding__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var babylonClient *client.Client
		fill_err = tp.Fill(&babylonClient)
		if fill_err != nil {
			return
		}
		var cfg *config.BTCStakingTrackerConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var stakeSpendingTx *wire.MsgTx
		fill_err = tp.Fill(&stakeSpendingTx)
		if fill_err != nil {
			return
		}
		var inclusionProof *types.InclusionProof
		fill_err = tp.Fill(&inclusionProof)
		if fill_err != nil {
			return
		}
		if babylonClient == nil || cfg == nil || stakeSpendingTx == nil || inclusionProof == nil {
			return
		}

		bca := NewBabylonClientAdapter(babylonClient, cfg)
		bca.ReportUnbonding(ctx, stakingTxHash, stakeSpendingTx, inclusionProof)
	})
}

func Fuzz_Nosy_MockBabylonNodeAdapter_ActivateDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var proof *types.BTCSpvProof
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if ctrl == nil || proof == nil {
			return
		}

		m := NewMockBabylonNodeAdapter(ctrl)
		m.ActivateDelegation(ctx, stakingTxHash, proof)
	})
}

func Fuzz_Nosy_MockBabylonNodeAdapter_BtcClientTipHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonNodeAdapter(ctrl)
		m.BtcClientTipHeight()
	})
}

func Fuzz_Nosy_MockBabylonNodeAdapter_DelegationsByStatus__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var status types.BTCDelegationStatus
		fill_err = tp.Fill(&status)
		if fill_err != nil {
			return
		}
		var offset uint64
		fill_err = tp.Fill(&offset)
		if fill_err != nil {
			return
		}
		var limit uint64
		fill_err = tp.Fill(&limit)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonNodeAdapter(ctrl)
		m.DelegationsByStatus(status, offset, limit)
	})
}

func Fuzz_Nosy_MockBabylonNodeAdapter_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonNodeAdapter(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockBabylonNodeAdapter_IsDelegationActive__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonNodeAdapter(ctrl)
		m.IsDelegationActive(stakingTxHash)
	})
}

func Fuzz_Nosy_MockBabylonNodeAdapter_IsDelegationVerified__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonNodeAdapter(ctrl)
		m.IsDelegationVerified(stakingTxHash)
	})
}

func Fuzz_Nosy_MockBabylonNodeAdapter_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonNodeAdapter(ctrl)
		m.Params()
	})
}

func Fuzz_Nosy_MockBabylonNodeAdapter_QueryHeaderDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var headerHash *chainhash.Hash
		fill_err = tp.Fill(&headerHash)
		if fill_err != nil {
			return
		}
		if ctrl == nil || headerHash == nil {
			return
		}

		m := NewMockBabylonNodeAdapter(ctrl)
		m.QueryHeaderDepth(headerHash)
	})
}

func Fuzz_Nosy_MockBabylonNodeAdapter_ReportUnbonding__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var stakeSpendingTx *wire.MsgTx
		fill_err = tp.Fill(&stakeSpendingTx)
		if fill_err != nil {
			return
		}
		var inclusionProof *types.InclusionProof
		fill_err = tp.Fill(&inclusionProof)
		if fill_err != nil {
			return
		}
		if ctrl == nil || stakeSpendingTx == nil || inclusionProof == nil {
			return
		}

		m := NewMockBabylonNodeAdapter(ctrl)
		m.ReportUnbonding(ctx, stakingTxHash, stakeSpendingTx, inclusionProof)
	})
}

// skipping Fuzz_Nosy_MockBabylonNodeAdapterMockRecorder_ActivateDelegation__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBabylonNodeAdapterMockRecorder_BtcClientTipHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBabylonNodeAdapterMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.BtcClientTipHeight()
	})
}

// skipping Fuzz_Nosy_MockBabylonNodeAdapterMockRecorder_DelegationsByStatus__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBabylonNodeAdapterMockRecorder_IsDelegationActive__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBabylonNodeAdapterMockRecorder_IsDelegationVerified__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBabylonNodeAdapterMockRecorder_Params__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBabylonNodeAdapterMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.Params()
	})
}

// skipping Fuzz_Nosy_MockBabylonNodeAdapterMockRecorder_QueryHeaderDepth__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBabylonNodeAdapterMockRecorder_ReportUnbonding__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_StakingEventWatcher_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		if sew == nil {
			return
		}

		sew.Start()
	})
}

func Fuzz_Nosy_StakingEventWatcher_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		if sew == nil {
			return
		}

		sew.Stop()
	})
}

func Fuzz_Nosy_StakingEventWatcher_activateBtcDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var proof *types.BTCSpvProof
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		var inclusionBlockHash chainhash.Hash
		fill_err = tp.Fill(&inclusionBlockHash)
		if fill_err != nil {
			return
		}
		var requiredDepth uint32
		fill_err = tp.Fill(&requiredDepth)
		if fill_err != nil {
			return
		}
		if sew == nil || proof == nil {
			return
		}

		sew.activateBtcDelegation(stakingTxHash, proof, inclusionBlockHash, requiredDepth)
	})
}

func Fuzz_Nosy_StakingEventWatcher_buildSpendingTxProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		var spendingTx *wire.MsgTx
		fill_err = tp.Fill(&spendingTx)
		if fill_err != nil {
			return
		}
		if sew == nil || spendingTx == nil {
			return
		}

		sew.buildSpendingTxProof(spendingTx)
	})
}

// skipping Fuzz_Nosy_StakingEventWatcher_checkBabylonDelegations__ because parameters include func, chan, or unsupported interface: func(del github.com/babylonlabs-io/vigilante/btcstaking-tracker/stakingeventwatcher.Delegation)

func Fuzz_Nosy_StakingEventWatcher_checkBtcForStakingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		if sew == nil {
			return
		}

		sew.checkBtcForStakingTx()
	})
}

func Fuzz_Nosy_StakingEventWatcher_fetchDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		if sew == nil {
			return
		}

		sew.fetchDelegations()
	})
}

func Fuzz_Nosy_StakingEventWatcher_handleNewBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		var blockNotifier *chainntnfs.BlockEpochEvent
		fill_err = tp.Fill(&blockNotifier)
		if fill_err != nil {
			return
		}
		if sew == nil || blockNotifier == nil {
			return
		}

		sew.handleNewBlocks(blockNotifier)
	})
}

func Fuzz_Nosy_StakingEventWatcher_handleUnbondedDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		if sew == nil {
			return
		}

		sew.handleUnbondedDelegations()
	})
}

func Fuzz_Nosy_StakingEventWatcher_handlerVerifiedDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		if sew == nil {
			return
		}

		sew.handlerVerifiedDelegations()
	})
}

func Fuzz_Nosy_StakingEventWatcher_latency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		var method string
		fill_err = tp.Fill(&method)
		if fill_err != nil {
			return
		}
		if sew == nil {
			return
		}

		sew.latency(method)
	})
}

func Fuzz_Nosy_StakingEventWatcher_quitContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		if sew == nil {
			return
		}

		sew.quitContext()
	})
}

func Fuzz_Nosy_StakingEventWatcher_reportUnbondingToBabylon__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var stakeSpendingTx *wire.MsgTx
		fill_err = tp.Fill(&stakeSpendingTx)
		if fill_err != nil {
			return
		}
		var proof *types.InclusionProof
		fill_err = tp.Fill(&proof)
		if fill_err != nil {
			return
		}
		if sew == nil || stakeSpendingTx == nil || proof == nil {
			return
		}

		sew.reportUnbondingToBabylon(ctx, stakingTxHash, stakeSpendingTx, proof)
	})
}

func Fuzz_Nosy_StakingEventWatcher_syncedWithBabylon__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		if sew == nil {
			return
		}

		sew.syncedWithBabylon()
	})
}

func Fuzz_Nosy_StakingEventWatcher_waitForRequiredDepth__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var inclusionBlockHash *chainhash.Hash
		fill_err = tp.Fill(&inclusionBlockHash)
		if fill_err != nil {
			return
		}
		var requiredDepth uint32
		fill_err = tp.Fill(&requiredDepth)
		if fill_err != nil {
			return
		}
		if sew == nil || inclusionBlockHash == nil {
			return
		}

		sew.waitForRequiredDepth(ctx, stakingTxHash, inclusionBlockHash, requiredDepth)
	})
}

func Fuzz_Nosy_StakingEventWatcher_waitForStakeSpendInclusionProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var spendingTx *wire.MsgTx
		fill_err = tp.Fill(&spendingTx)
		if fill_err != nil {
			return
		}
		if sew == nil || spendingTx == nil {
			return
		}

		sew.waitForStakeSpendInclusionProof(ctx, spendingTx)
	})
}

func Fuzz_Nosy_StakingEventWatcher_watchForSpend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var sew *StakingEventWatcher
		fill_err = tp.Fill(&sew)
		if fill_err != nil {
			return
		}
		var spendEvent *chainntnfs.SpendEvent
		fill_err = tp.Fill(&spendEvent)
		if fill_err != nil {
			return
		}
		var td *TrackedDelegation
		fill_err = tp.Fill(&td)
		if fill_err != nil {
			return
		}
		if sew == nil || spendEvent == nil || td == nil {
			return
		}

		sew.watchForSpend(spendEvent, td)
	})
}

func Fuzz_Nosy_TrackedDelegation_Clone__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var td *TrackedDelegation
		fill_err = tp.Fill(&td)
		if fill_err != nil {
			return
		}
		if td == nil {
			return
		}

		td.Clone()
	})
}

func Fuzz_Nosy_TrackedDelegations_AddDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTx *wire.MsgTx
		fill_err = tp.Fill(&stakingTx)
		if fill_err != nil {
			return
		}
		var stakingOutputIdx uint32
		fill_err = tp.Fill(&stakingOutputIdx)
		if fill_err != nil {
			return
		}
		var unbondingOutput *wire.TxOut
		fill_err = tp.Fill(&unbondingOutput)
		if fill_err != nil {
			return
		}
		var delegationStartHeight uint32
		fill_err = tp.Fill(&delegationStartHeight)
		if fill_err != nil {
			return
		}
		var shouldUpdate bool
		fill_err = tp.Fill(&shouldUpdate)
		if fill_err != nil {
			return
		}
		if stakingTx == nil || unbondingOutput == nil {
			return
		}

		td := NewTrackedDelegations()
		td.AddDelegation(stakingTx, stakingOutputIdx, unbondingOutput, delegationStartHeight, shouldUpdate)
	})
}

func Fuzz_Nosy_TrackedDelegations_AddEmptyDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var txHash chainhash.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}

		td := NewTrackedDelegations()
		td.AddEmptyDelegation(txHash)
	})
}

func Fuzz_Nosy_TrackedDelegations_DelegationsIter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, chunkSize int) {
		td := NewTrackedDelegations()
		td.DelegationsIter(chunkSize)
	})
}

func Fuzz_Nosy_TrackedDelegations_GetDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}

		td := NewTrackedDelegations()
		td.GetDelegation(stakingTxHash)
	})
}

func Fuzz_Nosy_TrackedDelegations_HasDelegationChanged__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var newDelegation *newDelegation
		fill_err = tp.Fill(&newDelegation)
		if fill_err != nil {
			return
		}
		if newDelegation == nil {
			return
		}

		td := NewTrackedDelegations()
		td.HasDelegationChanged(stakingTxHash, newDelegation)
	})
}

func Fuzz_Nosy_TrackedDelegations_RemoveDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}

		td := NewTrackedDelegations()
		td.RemoveDelegation(stakingTxHash)
	})
}

func Fuzz_Nosy_TrackedDelegations_UpdateActivation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx chainhash.Hash
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var inProgress bool
		fill_err = tp.Fill(&inProgress)
		if fill_err != nil {
			return
		}

		td := NewTrackedDelegations()
		td.UpdateActivation(tx, inProgress)
	})
}

// skipping Fuzz_Nosy_BabylonNodeAdapter_ActivateDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/stakingeventwatcher.BabylonNodeAdapter

// skipping Fuzz_Nosy_BabylonNodeAdapter_BtcClientTipHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/stakingeventwatcher.BabylonNodeAdapter

// skipping Fuzz_Nosy_BabylonNodeAdapter_DelegationsByStatus__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/stakingeventwatcher.BabylonNodeAdapter

// skipping Fuzz_Nosy_BabylonNodeAdapter_IsDelegationActive__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/stakingeventwatcher.BabylonNodeAdapter

// skipping Fuzz_Nosy_BabylonNodeAdapter_IsDelegationVerified__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/stakingeventwatcher.BabylonNodeAdapter

// skipping Fuzz_Nosy_BabylonNodeAdapter_Params__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/stakingeventwatcher.BabylonNodeAdapter

// skipping Fuzz_Nosy_BabylonNodeAdapter_QueryHeaderDepth__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/stakingeventwatcher.BabylonNodeAdapter

// skipping Fuzz_Nosy_BabylonNodeAdapter_ReportUnbonding__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/stakingeventwatcher.BabylonNodeAdapter

func Fuzz_Nosy_getStakingTxInputIdx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var td *TrackedDelegation
		fill_err = tp.Fill(&td)
		if fill_err != nil {
			return
		}
		if tx == nil || td == nil {
			return
		}

		getStakingTxInputIdx(tx, td)
	})
}
