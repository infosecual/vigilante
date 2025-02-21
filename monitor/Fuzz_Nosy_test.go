package monitor

import (
	"testing"

	"github.com/babylonlabs-io/babylon/x/checkpointing/types"
	checkpointingtypes "github.com/babylonlabs-io/babylon/x/checkpointing/types"
	vigilantetypes "github.com/babylonlabs-io/vigilante/types"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	query "github.com/cosmos/cosmos-sdk/types/query"
	"github.com/golang/mock/gomock"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"go.uber.org/zap"
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

func Fuzz_Nosy_MockBabylonQueryClient_BTCHeaderChainTip__(f *testing.F) {
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
		m.BTCHeaderChainTip()
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_BlsPublicKeyList__(f *testing.F) {
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
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
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
		m.BlsPublicKeyList(epochNumber, pagination)
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_ContainsBTCBlock__(f *testing.F) {
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

		m := NewMockBabylonQueryClient(ctrl)
		m.ContainsBTCBlock(blockHash)
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_CurrentEpoch__(f *testing.F) {
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
		m.CurrentEpoch()
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

func Fuzz_Nosy_MockBabylonQueryClient_EndedEpochBTCHeight__(f *testing.F) {
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
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonQueryClient(ctrl)
		m.EndedEpochBTCHeight(epochNum)
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

func Fuzz_Nosy_MockBabylonQueryClient_RawCheckpoint__(f *testing.F) {
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
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonQueryClient(ctrl)
		m.RawCheckpoint(epochNumber)
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_ReportedCheckpointBTCHeight__(f *testing.F) {
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
		var hashStr string
		fill_err = tp.Fill(&hashStr)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBabylonQueryClient(ctrl)
		m.ReportedCheckpointBTCHeight(hashStr)
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_Start__(f *testing.F) {
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
		m.Start()
	})
}

func Fuzz_Nosy_MockBabylonQueryClient_Stop__(f *testing.F) {
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
		m.Stop()
	})
}

func Fuzz_Nosy_MockBabylonQueryClientMockRecorder_BTCHeaderChainTip__(f *testing.F) {
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

		mr.BTCHeaderChainTip()
	})
}

// skipping Fuzz_Nosy_MockBabylonQueryClientMockRecorder_BlsPublicKeyList__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBabylonQueryClientMockRecorder_ContainsBTCBlock__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBabylonQueryClientMockRecorder_CurrentEpoch__(f *testing.F) {
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

		mr.CurrentEpoch()
	})
}

// skipping Fuzz_Nosy_MockBabylonQueryClientMockRecorder_EndedEpochBTCHeight__ because parameters include func, chan, or unsupported interface: interface{}

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

// skipping Fuzz_Nosy_MockBabylonQueryClientMockRecorder_RawCheckpoint__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBabylonQueryClientMockRecorder_ReportedCheckpointBTCHeight__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBabylonQueryClientMockRecorder_Start__(f *testing.F) {
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

		mr.Start()
	})
}

func Fuzz_Nosy_MockBabylonQueryClientMockRecorder_Stop__(f *testing.F) {
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

		mr.Stop()
	})
}

func Fuzz_Nosy_Monitor_CheckLiveness__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var cr *vigilantetypes.CheckpointRecord
		fill_err = tp.Fill(&cr)
		if fill_err != nil {
			return
		}
		if m == nil || cr == nil {
			return
		}

		m.CheckLiveness(cr)
	})
}

func Fuzz_Nosy_Monitor_FindTipConfirmedEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.FindTipConfirmedEpoch()
	})
}

func Fuzz_Nosy_Monitor_GetCurrentEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.GetCurrentEpoch()
	})
}

func Fuzz_Nosy_Monitor_Metrics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Metrics()
	})
}

func Fuzz_Nosy_Monitor_QueryInfoForNextEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.QueryInfoForNextEpoch(epoch)
	})
}

func Fuzz_Nosy_Monitor_SetLogger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var logger *zap.SugaredLogger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if m == nil || logger == nil {
			return
		}

		m.SetLogger(logger)
	})
}

func Fuzz_Nosy_Monitor_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var baseHeight uint32
		fill_err = tp.Fill(&baseHeight)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Start(baseHeight)
	})
}

func Fuzz_Nosy_Monitor_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.Stop()
	})
}

func Fuzz_Nosy_Monitor_UpdateEpochInfo__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.UpdateEpochInfo(epoch)
	})
}

func Fuzz_Nosy_Monitor_VerifyCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var btcCkpt *checkpointingtypes.RawCheckpoint
		fill_err = tp.Fill(&btcCkpt)
		if fill_err != nil {
			return
		}
		if m == nil || btcCkpt == nil {
			return
		}

		m.VerifyCheckpoint(btcCkpt)
	})
}

func Fuzz_Nosy_Monitor_addCheckpointToCheckList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ckpt *vigilantetypes.CheckpointRecord
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if m == nil || ckpt == nil {
			return
		}

		m.addCheckpointToCheckList(ckpt)
	})
}

func Fuzz_Nosy_Monitor_checkHeaderConsistency__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var header *wire.BlockHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if m == nil || header == nil {
			return
		}

		m.checkHeaderConsistency(header)
	})
}

func Fuzz_Nosy_Monitor_handleNewConfirmedCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var ckpt *vigilantetypes.CheckpointRecord
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if m == nil || ckpt == nil {
			return
		}

		m.handleNewConfirmedCheckpoint(ckpt)
	})
}

func Fuzz_Nosy_Monitor_handleNewConfirmedHeader__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var block *vigilantetypes.IndexedBlock
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if m == nil || block == nil {
			return
		}

		m.handleNewConfirmedHeader(block)
	})
}

func Fuzz_Nosy_Monitor_queryBTCHeaderChainTipWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.queryBTCHeaderChainTipWithRetry()
	})
}

func Fuzz_Nosy_Monitor_queryBlsPublicKeyListWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.queryBlsPublicKeyListWithRetry(epoch)
	})
}

func Fuzz_Nosy_Monitor_queryContainsBTCBlockWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var blockHash *chainhash.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if m == nil || blockHash == nil {
			return
		}

		m.queryContainsBTCBlockWithRetry(blockHash)
	})
}

func Fuzz_Nosy_Monitor_queryCurrentEpochWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.queryCurrentEpochWithRetry()
	})
}

func Fuzz_Nosy_Monitor_queryEndedEpochBTCHeightWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.queryEndedEpochBTCHeightWithRetry(epoch)
	})
}

func Fuzz_Nosy_Monitor_queryRawCheckpointWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.queryRawCheckpointWithRetry(epoch)
	})
}

func Fuzz_Nosy_Monitor_queryReportedCheckpointBTCHeightWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var hashStr string
		fill_err = tp.Fill(&hashStr)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.queryReportedCheckpointBTCHeightWithRetry(hashStr)
	})
}

func Fuzz_Nosy_Monitor_runBTCScanner__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.runBTCScanner(startHeight)
	})
}

func Fuzz_Nosy_Monitor_runLivenessChecker__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var m *Monitor
		fill_err = tp.Fill(&m)
		if fill_err != nil {
			return
		}
		if m == nil {
			return
		}

		m.runLivenessChecker()
	})
}

// skipping Fuzz_Nosy_BabylonQueryClient_BTCHeaderChainTip__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_BlsPublicKeyList__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_ContainsBTCBlock__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_CurrentEpoch__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_EndedEpochBTCHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_IsRunning__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_RawCheckpoint__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_ReportedCheckpointBTCHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_Start__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor.BabylonQueryClient

// skipping Fuzz_Nosy_BabylonQueryClient_Stop__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor.BabylonQueryClient

func Fuzz_Nosy_convertFromBlsPublicKeyListResponse__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var valBLSKeys []*types.BlsPublicKeyListResponse
		fill_err = tp.Fill(&valBLSKeys)
		if fill_err != nil {
			return
		}

		convertFromBlsPublicKeyListResponse(valBLSKeys)
	})
}

func Fuzz_Nosy_minBTCHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, h1 uint32, h2 uint32) {
		minBTCHeight(h1, h2)
	})
}
