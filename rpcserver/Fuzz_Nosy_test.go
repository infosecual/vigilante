package rpcserver

import (
	"testing"

	btcstakingtracker "github.com/babylonlabs-io/vigilante/btcstaking-tracker"
	"github.com/babylonlabs-io/vigilante/config"
	"github.com/babylonlabs-io/vigilante/monitor"
	"github.com/babylonlabs-io/vigilante/reporter"
	"github.com/babylonlabs-io/vigilante/submitter"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

func Fuzz_Nosy_Server_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.GRPCConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		var submitter *submitter.Submitter
		fill_err = tp.Fill(&submitter)
		if fill_err != nil {
			return
		}
		var reporter *reporter.Reporter
		fill_err = tp.Fill(&reporter)
		if fill_err != nil {
			return
		}
		var monitor *monitor.Monitor
		fill_err = tp.Fill(&monitor)
		if fill_err != nil {
			return
		}
		var bstracker *btcstakingtracker.BTCStakingTracker
		fill_err = tp.Fill(&bstracker)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil || submitter == nil || reporter == nil || monitor == nil || bstracker == nil {
			return
		}

		s, err := New(cfg, parentLogger, submitter, reporter, monitor, bstracker)
		if err != nil {
			return
		}
		s.Start()
	})
}

// skipping Fuzz_Nosy_service_Version__ because parameters include func, chan, or unsupported interface: golang.org/x/net/context.Context

func Fuzz_Nosy_StartVigilanteService__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var gs *grpc.Server
		fill_err = tp.Fill(&gs)
		if fill_err != nil {
			return
		}
		if gs == nil {
			return
		}

		StartVigilanteService(gs)
	})
}
