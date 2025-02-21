package btcstakingtracker

import (
	"testing"

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

func Fuzz_Nosy_BTCStakingTracker_Bootstrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tracker *BTCStakingTracker
		fill_err = tp.Fill(&tracker)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		if tracker == nil {
			return
		}

		tracker.Bootstrap(startHeight)
	})
}

func Fuzz_Nosy_BTCStakingTracker_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tracker *BTCStakingTracker
		fill_err = tp.Fill(&tracker)
		if fill_err != nil {
			return
		}
		if tracker == nil {
			return
		}

		tracker.Start()
	})
}

func Fuzz_Nosy_BTCStakingTracker_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tracker *BTCStakingTracker
		fill_err = tp.Fill(&tracker)
		if fill_err != nil {
			return
		}
		if tracker == nil {
			return
		}

		tracker.Stop()
	})
}

// skipping Fuzz_Nosy_IAtomicSlasher_Start__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker.IAtomicSlasher

// skipping Fuzz_Nosy_IAtomicSlasher_Stop__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker.IAtomicSlasher

// skipping Fuzz_Nosy_IBTCSlasher_Bootstrap__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker.IBTCSlasher

// skipping Fuzz_Nosy_IBTCSlasher_Start__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker.IBTCSlasher

// skipping Fuzz_Nosy_IBTCSlasher_Stop__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcstaking-tracker.IBTCSlasher
