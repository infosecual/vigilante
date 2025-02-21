package reporter

import (
	"context"
	"testing"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	gomock "github.com/golang/mock/gomock"
	"github.com/lightningnetwork/lnd/chainntnfs"
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

func Fuzz_Nosy_MockBabylonClient_BTCBaseHeader__(f *testing.F) {
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
		m.BTCBaseHeader()
	})
}

func Fuzz_Nosy_MockBabylonClient_BTCCheckpointParams__(f *testing.F) {
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
		m.BTCCheckpointParams()
	})
}

func Fuzz_Nosy_MockBabylonClient_BTCHeaderChainTip__(f *testing.F) {
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
		m.BTCHeaderChainTip()
	})
}

func Fuzz_Nosy_MockBabylonClient_ContainsBTCBlock__(f *testing.F) {
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
		var blockHash *chainhash.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if ctrl == nil || blockHash == nil {
			return
		}

		m := NewMockBabylonClient(ctrl)
		m.ContainsBTCBlock(blockHash)
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

func Fuzz_Nosy_MockBabylonClient_GetConfig__(f *testing.F) {
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
		m.GetConfig()
	})
}

func Fuzz_Nosy_MockBabylonClient_InsertBTCSpvProof__(f *testing.F) {
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
		var msg *types.MsgInsertBTCSpvProof
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if ctrl == nil || msg == nil {
			return
		}

		m := NewMockBabylonClient(ctrl)
		m.InsertBTCSpvProof(ctx, msg)
	})
}

func Fuzz_Nosy_MockBabylonClient_InsertHeaders__(f *testing.F) {
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
		var msgs *types.MsgInsertHeaders
		fill_err = tp.Fill(&msgs)
		if fill_err != nil {
			return
		}
		if ctrl == nil || msgs == nil {
			return
		}

		m := NewMockBabylonClient(ctrl)
		m.InsertHeaders(ctx, msgs)
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

func Fuzz_Nosy_MockBabylonClient_Stop__(f *testing.F) {
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
		m.Stop()
	})
}

func Fuzz_Nosy_MockBabylonClientMockRecorder_BTCBaseHeader__(f *testing.F) {
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

		mr.BTCBaseHeader()
	})
}

func Fuzz_Nosy_MockBabylonClientMockRecorder_BTCCheckpointParams__(f *testing.F) {
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

		mr.BTCCheckpointParams()
	})
}

func Fuzz_Nosy_MockBabylonClientMockRecorder_BTCHeaderChainTip__(f *testing.F) {
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

		mr.BTCHeaderChainTip()
	})
}

// skipping Fuzz_Nosy_MockBabylonClientMockRecorder_ContainsBTCBlock__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBabylonClientMockRecorder_GetConfig__(f *testing.F) {
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

		mr.GetConfig()
	})
}

// skipping Fuzz_Nosy_MockBabylonClientMockRecorder_InsertBTCSpvProof__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBabylonClientMockRecorder_InsertHeaders__ because parameters include func, chan, or unsupported interface: interface{}

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

func Fuzz_Nosy_MockBabylonClientMockRecorder_Stop__(f *testing.F) {
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

		mr.Stop()
	})
}

func Fuzz_Nosy_Reporter_ProcessCheckpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var signer string
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var ibs []*types.IndexedBlock
		fill_err = tp.Fill(&ibs)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.ProcessCheckpoints(signer, ibs)
	})
}

func Fuzz_Nosy_Reporter_ProcessHeaders__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var signer string
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var ibs []*types.IndexedBlock
		fill_err = tp.Fill(&ibs)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.ProcessHeaders(signer, ibs)
	})
}

func Fuzz_Nosy_Reporter_ShuttingDown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.ShuttingDown()
	})
}

func Fuzz_Nosy_Reporter_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Start()
	})
}

func Fuzz_Nosy_Reporter_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.Stop()
	})
}

func Fuzz_Nosy_Reporter_WaitForShutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.WaitForShutdown()
	})
}

func Fuzz_Nosy_Reporter_blockEventHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var blockNotifier *chainntnfs.BlockEpochEvent
		fill_err = tp.Fill(&blockNotifier)
		if fill_err != nil {
			return
		}
		if r == nil || blockNotifier == nil {
			return
		}

		r.blockEventHandler(blockNotifier)
	})
}

func Fuzz_Nosy_Reporter_bootstrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.bootstrap()
	})
}

func Fuzz_Nosy_Reporter_bootstrapWithRetries__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.bootstrapWithRetries()
	})
}

func Fuzz_Nosy_Reporter_checkConsistency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.checkConsistency()
	})
}

func Fuzz_Nosy_Reporter_checkHeaderConsistency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var consistencyCheckHeight uint32
		fill_err = tp.Fill(&consistencyCheckHeight)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.checkHeaderConsistency(consistencyCheckHeight)
	})
}

func Fuzz_Nosy_Reporter_extractCheckpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ib *types.IndexedBlock
		fill_err = tp.Fill(&ib)
		if fill_err != nil {
			return
		}
		if r == nil || ib == nil {
			return
		}

		r.extractCheckpoints(ib)
	})
}

func Fuzz_Nosy_Reporter_getHeaderMsgsToSubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var signer string
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		var ibs []*types.IndexedBlock
		fill_err = tp.Fill(&ibs)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.getHeaderMsgsToSubmit(signer, ibs)
	})
}

func Fuzz_Nosy_Reporter_handleNewBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var header *wire.BlockHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if r == nil || header == nil {
			return
		}

		r.handleNewBlock(height, header)
	})
}

func Fuzz_Nosy_Reporter_initBTCCache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.initBTCCache()
	})
}

func Fuzz_Nosy_Reporter_matchAndSubmitCheckpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var signer string
		fill_err = tp.Fill(&signer)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.matchAndSubmitCheckpoints(signer)
	})
}

func Fuzz_Nosy_Reporter_processNewBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var ib *types.IndexedBlock
		fill_err = tp.Fill(&ib)
		if fill_err != nil {
			return
		}
		if r == nil || ib == nil {
			return
		}

		r.processNewBlock(ib)
	})
}

func Fuzz_Nosy_Reporter_quitChan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.quitChan()
	})
}

func Fuzz_Nosy_Reporter_reporterQuitCtx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.reporterQuitCtx()
	})
}

func Fuzz_Nosy_Reporter_submitHeaderMsgs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var msg *types.MsgInsertHeaders
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if r == nil || msg == nil {
			return
		}

		r.submitHeaderMsgs(msg)
	})
}

func Fuzz_Nosy_Reporter_waitUntilBTCSync__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *Reporter
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		r.waitUntilBTCSync()
	})
}

// skipping Fuzz_Nosy_BabylonClient_BTCBaseHeader__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/reporter.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_BTCCheckpointParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/reporter.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_BTCHeaderChainTip__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/reporter.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_ContainsBTCBlock__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/reporter.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_GetConfig__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/reporter.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_InsertBTCSpvProof__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/reporter.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_InsertHeaders__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/reporter.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_MustGetAddr__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/reporter.BabylonClient

// skipping Fuzz_Nosy_BabylonClient_Stop__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/reporter.BabylonClient

// skipping Fuzz_Nosy_chunkBy__ because parameters include func, chan, or unsupported interface: []T
