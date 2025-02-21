package metrics

import (
	"testing"

	"github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/prometheus/client_golang/prometheus"
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

func Fuzz_Nosy_SlasherMetrics_RecordSlashedDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var registry *prometheus.Registry
		fill_err = tp.Fill(&registry)
		if fill_err != nil {
			return
		}
		var del *types.BTCDelegationResponse
		fill_err = tp.Fill(&del)
		if fill_err != nil {
			return
		}
		if registry == nil || del == nil {
			return
		}

		sm := newSlasherMetrics(registry)
		sm.RecordSlashedDelegation(del)
	})
}

func Fuzz_Nosy_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var reg *prometheus.Registry
		fill_err = tp.Fill(&reg)
		if fill_err != nil {
			return
		}
		if reg == nil {
			return
		}

		Start(addr, reg)
	})
}

func Fuzz_Nosy_start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var addr string
		fill_err = tp.Fill(&addr)
		if fill_err != nil {
			return
		}
		var reg *prometheus.Registry
		fill_err = tp.Fill(&reg)
		if fill_err != nil {
			return
		}
		if reg == nil {
			return
		}

		start(addr, reg)
	})
}
