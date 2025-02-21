package poller

import (
	"testing"

	types "github.com/babylonlabs-io/babylon/x/checkpointing/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
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

func Fuzz_Nosy_MockBabylonQueryClient_RawCheckpointList__(f *testing.F) {
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
		var status types.CheckpointStatus
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

		m := NewMockBabylonQueryClient(ctrl)
		m.RawCheckpointList(status, pagination)
	})
}

// skipping Fuzz_Nosy_MockBabylonQueryClientMockRecorder_RawCheckpointList__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_Poller_GetSealedCheckpointChan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pl *Poller
		fill_err = tp.Fill(&pl)
		if fill_err != nil {
			return
		}
		if pl == nil {
			return
		}

		pl.GetSealedCheckpointChan()
	})
}

func Fuzz_Nosy_Poller_PollSealedCheckpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var pl *Poller
		fill_err = tp.Fill(&pl)
		if fill_err != nil {
			return
		}
		if pl == nil {
			return
		}

		pl.PollSealedCheckpoints()
	})
}

// skipping Fuzz_Nosy_BabylonQueryClient_RawCheckpointList__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/submitter/poller.BabylonQueryClient
