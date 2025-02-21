package btcslasher

import (
	"testing"

	bbn "github.com/babylonlabs-io/babylon/types"
	bstypes "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/mock/gomock"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
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

func Fuzz_Nosy_BTCSlasher_Bootstrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Bootstrap(startHeight)
	})
}

func Fuzz_Nosy_BTCSlasher_LoadParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.LoadParams()
	})
}

func Fuzz_Nosy_BTCSlasher_SlashFinalityProvider__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var extractedFpBTCSK *btcec.PrivateKey
		fill_err = tp.Fill(&extractedFpBTCSK)
		if fill_err != nil {
			return
		}
		if bs == nil || extractedFpBTCSK == nil {
			return
		}

		bs.SlashFinalityProvider(extractedFpBTCSK)
	})
}

func Fuzz_Nosy_BTCSlasher_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Start()
	})
}

func Fuzz_Nosy_BTCSlasher_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Stop()
	})
}

func Fuzz_Nosy_BTCSlasher_WaitForShutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.WaitForShutdown()
	})
}

func Fuzz_Nosy_BTCSlasher_equivocationTracker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.equivocationTracker()
	})
}

func Fuzz_Nosy_BTCSlasher_getAllActiveAndUnbondedBTCDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if bs == nil || fpBTCPK == nil {
			return
		}

		bs.getAllActiveAndUnbondedBTCDelegations(fpBTCPK)
	})
}

// skipping Fuzz_Nosy_BTCSlasher_handleAllEvidences__ because parameters include func, chan, or unsupported interface: func(evidences []*github.com/babylonlabs-io/babylon/x/finality/types.EvidenceResponse) error

func Fuzz_Nosy_BTCSlasher_handleEvidence__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var evt *coretypes.ResultEvent
		fill_err = tp.Fill(&evt)
		if fill_err != nil {
			return
		}
		var isConsumer bool
		fill_err = tp.Fill(&isConsumer)
		if fill_err != nil {
			return
		}
		if bs == nil || evt == nil {
			return
		}

		bs.handleEvidence(evt, isConsumer)
	})
}

func Fuzz_Nosy_BTCSlasher_isTaprootOutputSpendable__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var txBytes []byte
		fill_err = tp.Fill(&txBytes)
		if fill_err != nil {
			return
		}
		var outIdx uint32
		fill_err = tp.Fill(&outIdx)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.isTaprootOutputSpendable(txBytes, outIdx)
	})
}

func Fuzz_Nosy_BTCSlasher_isTxSubmittedToBitcoin__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var txHash *chainhash.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if bs == nil || txHash == nil {
			return
		}

		bs.isTxSubmittedToBitcoin(txHash)
	})
}

func Fuzz_Nosy_BTCSlasher_quitContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.quitContext()
	})
}

func Fuzz_Nosy_BTCSlasher_sendSlashingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var extractedfpBTCSK *btcec.PrivateKey
		fill_err = tp.Fill(&extractedfpBTCSK)
		if fill_err != nil {
			return
		}
		var del *bstypes.BTCDelegationResponse
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}
		var isUnbondingSlashingTx bool
		fill_err = tp.Fill(&isUnbondingSlashingTx)
		if fill_err != nil {
			return
		}
		if bs == nil || fpBTCPK == nil || extractedfpBTCSK == nil || del == nil {
			return
		}

		bs.sendSlashingTx(fpBTCPK, extractedfpBTCSK, del, isUnbondingSlashingTx)
	})
}

func Fuzz_Nosy_BTCSlasher_slashBTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var extractedfpBTCSK *btcec.PrivateKey
		fill_err = tp.Fill(&extractedfpBTCSK)
		if fill_err != nil {
			return
		}
		var del *bstypes.BTCDelegationResponse
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}
		if bs == nil || fpBTCPK == nil || extractedfpBTCSK == nil || del == nil {
			return
		}

		bs.slashBTCDelegation(fpBTCPK, extractedfpBTCSK, del)
	})
}

func Fuzz_Nosy_BTCSlasher_slashingEnforcer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BTCSlasher
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.slashingEnforcer()
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_BTCCheckpointParams__(f *testing.F) {
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

		m := NewMockBabylonQueryClient(ctrl)
		m.BTCCheckpointParams()
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_BTCStakingParamsByVersion__(f *testing.F) {
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
		var version uint32
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonQueryClient(ctrl)
		m.BTCStakingParamsByVersion(version)
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_EXPECT__(f *testing.F) {
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

		m := NewMockBabylonQueryClient(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_FinalityProviderDelegations__(f *testing.F) {
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
		var fpBTCPKHex string
		fill_err = tp.Fill(&fpBTCPKHex)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if ctrl == nil || pagination == nil {
			return
		}

		m := NewMockBabylonQueryClient(ctrl)
		m.FinalityProviderDelegations(fpBTCPKHex, pagination)
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_IsRunning__(f *testing.F) {
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

		m := NewMockBabylonQueryClient(ctrl)
		m.IsRunning()
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_ListEvidences__(f *testing.F) {
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
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if ctrl == nil || pagination == nil {
			return
		}

		m := NewMockBabylonQueryClient(ctrl)
		m.ListEvidences(startHeight, pagination)
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_Subscribe__(f *testing.F) {
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
		var subscriber string
		fill_err = tp.Fill(&subscriber)
		if fill_err != nil {
			return
		}
		var query string
		fill_err = tp.Fill(&query)
		if fill_err != nil {
			return
		}
		var outCapacity []int
		fill_err = tp.Fill(&outCapacity)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonQueryClient(ctrl)
		m.Subscribe(subscriber, query, outCapacity...)
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_UnsubscribeAll__(f *testing.F) {
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
		var subscriber string
		fill_err = tp.Fill(&subscriber)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonQueryClient(ctrl)
		m.UnsubscribeAll(subscriber)
	})
}

func Fuzz_Nosy_MockBabylonQueryClientMockRecorder_BTCCheckpointParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBabylonQueryClientMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.BTCCheckpointParams()
	})
}

// skipping Fuzz_Nosy_MockBabylonQueryClientMockRecorder_BTCStakingParamsByVersion__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBabylonQueryClientMockRecorder_FinalityProviderDelegations__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBabylonQueryClientMockRecorder_IsRunning__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBabylonQueryClientMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.IsRunning()
	})
}

// skipping Fuzz_Nosy_MockBabylonQueryClientMockRecorder_ListEvidences__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBabylonQueryClientMockRecorder_Subscribe__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBabylonQueryClientMockRecorder_UnsubscribeAll__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_BabylonQueryClient_BTCCheckpointParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/btcslasher.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_BTCStakingParamsByVersion__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/btcslasher.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_FinalityProviderDelegations__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/btcslasher.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_IsRunning__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/btcslasher.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_ListEvidences__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/btcslasher.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_Subscribe__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/btcslasher.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_UnsubscribeAll__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/btcslasher.BabylonQueryClient

func Fuzz_Nosy_findFPIdxInWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var fpBTCPK *bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		var fpBtcPkList []bbn.BIP340PubKey
		fill_err = tp.Fill(&fpBtcPkList)
		if fill_err != nil {
			return
		}
		if fpBTCPK == nil {
			return
		}

		findFPIdxInWitness(fpBTCPK, fpBtcPkList)
	})
}
