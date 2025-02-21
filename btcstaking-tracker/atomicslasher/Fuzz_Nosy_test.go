package atomicslasher

import (
	"context"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/cosmos/cosmos-sdk/types/query"
	gomock "github.com/golang/mock/gomock"
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

func Fuzz_Nosy_AtomicSlasher_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var as *AtomicSlasher
		fill_err = tp.Fill(&as)
		if fill_err != nil {
			return
		}
		if as == nil {
			return
		}

		as.Start()
	})
}

func Fuzz_Nosy_AtomicSlasher_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var as *AtomicSlasher
		fill_err = tp.Fill(&as)
		if fill_err != nil {
			return
		}
		if as == nil {
			return
		}

		as.Stop()
	})
}

func Fuzz_Nosy_AtomicSlasher_btcDelegationTracker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var as *AtomicSlasher
		fill_err = tp.Fill(&as)
		if fill_err != nil {
			return
		}
		if as == nil {
			return
		}

		as.btcDelegationTracker()
	})
}

func Fuzz_Nosy_AtomicSlasher_quitContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var as *AtomicSlasher
		fill_err = tp.Fill(&as)
		if fill_err != nil {
			return
		}
		if as == nil {
			return
		}

		as.quitContext()
	})
}

func Fuzz_Nosy_AtomicSlasher_selectiveSlashingReporter__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var as *AtomicSlasher
		fill_err = tp.Fill(&as)
		if fill_err != nil {
			return
		}
		if as == nil {
			return
		}

		as.selectiveSlashingReporter()
	})
}

func Fuzz_Nosy_AtomicSlasher_slashingTxTracker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var as *AtomicSlasher
		fill_err = tp.Fill(&as)
		if fill_err != nil {
			return
		}
		if as == nil {
			return
		}

		as.slashingTxTracker()
	})
}

func Fuzz_Nosy_BTCDelegationIndex_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var trackedDel *TrackedDelegation
		fill_err = tp.Fill(&trackedDel)
		if fill_err != nil {
			return
		}
		if trackedDel == nil {
			return
		}

		bdi := NewBTCDelegationIndex()
		bdi.Add(trackedDel)
	})
}

func Fuzz_Nosy_BTCDelegationIndex_FindSlashedBTCDelegation__(f *testing.F) {
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

		bdi := NewBTCDelegationIndex()
		bdi.FindSlashedBTCDelegation(txHash)
	})
}

func Fuzz_Nosy_BTCDelegationIndex_Get__(f *testing.F) {
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

		bdi := NewBTCDelegationIndex()
		bdi.Get(stakingTxHash)
	})
}

func Fuzz_Nosy_BTCDelegationIndex_Remove__(f *testing.F) {
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

		bdi := NewBTCDelegationIndex()
		bdi.Remove(stakingTxHash)
	})
}

func Fuzz_Nosy_BabylonAdapter_BTCDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ba *BabylonAdapter
		fill_err = tp.Fill(&ba)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if ba == nil {
			return
		}

		ba.BTCDelegation(ctx, stakingTxHashHex)
	})
}

func Fuzz_Nosy_BabylonAdapter_BTCStakingParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ba *BabylonAdapter
		fill_err = tp.Fill(&ba)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var version uint32
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		if ba == nil {
			return
		}

		ba.BTCStakingParams(ctx, version)
	})
}

// skipping Fuzz_Nosy_BabylonAdapter_HandleAllBTCDelegations__ because parameters include func, chan, or unsupported interface: func(btcDel *github.com/babylonlabs-io/babylon/x/btcstaking/types.BTCDelegationResponse) error

func Fuzz_Nosy_BabylonAdapter_IsFPSlashed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ba *BabylonAdapter
		fill_err = tp.Fill(&ba)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var fpBTCPK *types.BIP340PubKey
		fill_err = tp.Fill(&fpBTCPK)
		if fill_err != nil {
			return
		}
		if ba == nil || fpBTCPK == nil {
			return
		}

		ba.IsFPSlashed(_x2, fpBTCPK)
	})
}

func Fuzz_Nosy_BabylonAdapter_ReportSelectiveSlashing__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ba *BabylonAdapter
		fill_err = tp.Fill(&ba)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var fpBTCSK *btcec.PrivateKey
		fill_err = tp.Fill(&fpBTCSK)
		if fill_err != nil {
			return
		}
		if ba == nil || fpBTCSK == nil {
			return
		}

		ba.ReportSelectiveSlashing(ctx, stakingTxHash, fpBTCSK)
	})
}

func Fuzz_Nosy_MockBabylonClient_BTCDelegation__(f *testing.F) {
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
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonClient(ctrl)
		m.BTCDelegation(stakingTxHashHex)
	})
}

func Fuzz_Nosy_MockBabylonClient_BTCDelegations__(f *testing.F) {
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
		var pagination *query.PageRequest
		fill_err = tp.Fill(&pagination)
		if fill_err != nil {
			return
		}
		if ctrl == nil || pagination == nil {
			return
		}

		m := NewMockBabylonClient(ctrl)
		m.BTCDelegations(status, pagination)
	})
}

func Fuzz_Nosy_MockBabylonClient_BTCStakingParamsByVersion__(f *testing.F) {
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

		m := NewMockBabylonClient(ctrl)
		m.BTCStakingParamsByVersion(version)
	})
}

func Fuzz_Nosy_MockBabylonClient_EXPECT__(f *testing.F) {
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

		m := NewMockBabylonClient(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockBabylonClient_FinalityProvider__(f *testing.F) {
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
		var fpBtcPkHex string
		fill_err = tp.Fill(&fpBtcPkHex)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonClient(ctrl)
		m.FinalityProvider(fpBtcPkHex)
	})
}

func Fuzz_Nosy_MockBabylonClient_MustGetAddr__(f *testing.F) {
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

		m := NewMockBabylonClient(ctrl)
		m.MustGetAddr()
	})
}

// skipping Fuzz_Nosy_MockBabylonClient_ReliablySendMsg__ because parameters include func, chan, or unsupported interface: github.com/cosmos/cosmos-sdk/types.Msg

// skipping Fuzz_Nosy_MockBabylonClientMockRecorder_BTCDelegation__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBabylonClientMockRecorder_BTCDelegations__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBabylonClientMockRecorder_BTCStakingParamsByVersion__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBabylonClientMockRecorder_FinalityProvider__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBabylonClientMockRecorder_MustGetAddr__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBabylonClientMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.MustGetAddr()
	})
}

// skipping Fuzz_Nosy_MockBabylonClientMockRecorder_ReliablySendMsg__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_SlashingTxInfo_IsSlashStakingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var path SlashingPath
		fill_err = tp.Fill(&path)
		if fill_err != nil {
			return
		}
		var stakingTxHash chainhash.Hash
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		var slashingMsgTx *wire.MsgTx
		fill_err = tp.Fill(&slashingMsgTx)
		if fill_err != nil {
			return
		}
		if slashingMsgTx == nil {
			return
		}

		s := NewSlashingTxInfo(path, stakingTxHash, slashingMsgTx)
		s.IsSlashStakingTx()
	})
}

// skipping Fuzz_Nosy_BabylonClient_BTCDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/atomicslasher.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_BTCDelegations__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/atomicslasher.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_BTCStakingParamsByVersion__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/atomicslasher.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_FinalityProvider__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/atomicslasher.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_MustGetAddr__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/atomicslasher.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_ReliablySendMsg__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker/atomicslasher.BabylonClient

func Fuzz_Nosy_parseSlashingTxWitness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var witnessStack wire.TxWitness
		fill_err = tp.Fill(&witnessStack)
		if fill_err != nil {
			return
		}
		var covPKs []types.BIP340PubKey
		fill_err = tp.Fill(&covPKs)
		if fill_err != nil {
			return
		}
		var fpPKs []types.BIP340PubKey
		fill_err = tp.Fill(&fpPKs)
		if fill_err != nil {
			return
		}

		parseSlashingTxWitness(witnessStack, covPKs, fpPKs)
	})
}

func Fuzz_Nosy_tryExtractFPSK__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var covSigMap map[string]*types.BIP340Signature
		fill_err = tp.Fill(&covSigMap)
		if fill_err != nil {
			return
		}
		var fpIdx int
		fill_err = tp.Fill(&fpIdx)
		if fill_err != nil {
			return
		}
		var fpPK *types.BIP340PubKey
		fill_err = tp.Fill(&fpPK)
		if fill_err != nil {
			return
		}
		var covASigLists []*types.CovenantAdaptorSignatures
		fill_err = tp.Fill(&covASigLists)
		if fill_err != nil {
			return
		}
		if fpPK == nil {
			return
		}

		tryExtractFPSK(covSigMap, fpIdx, fpPK, covASigLists)
	})
}
