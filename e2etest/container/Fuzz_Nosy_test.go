package container

import (
	"testing"

	"github.com/ory/dockertest/v3/docker"
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

func Fuzz_Nosy_Manager_ClearResources__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		m, err := NewManager(t1)
		if err != nil {
			return
		}
		m.ClearResources()
	})
}

func Fuzz_Nosy_Manager_ExecBitcoindCliCmd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var command []string
		fill_err = tp.Fill(&command)
		if fill_err != nil {
			return
		}
		if t1 == nil || t2 == nil {
			return
		}

		m, err := NewManager(t1)
		if err != nil {
			return
		}
		m.ExecBitcoindCliCmd(t2, command)
	})
}

func Fuzz_Nosy_Manager_ExecCmd__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var containerName string
		fill_err = tp.Fill(&containerName)
		if fill_err != nil {
			return
		}
		var command []string
		fill_err = tp.Fill(&command)
		if fill_err != nil {
			return
		}
		if t1 == nil || t2 == nil {
			return
		}

		m, err := NewManager(t1)
		if err != nil {
			return
		}
		m.ExecCmd(t2, containerName, command)
	})
}

func Fuzz_Nosy_Manager_RunBabylondResource__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var mounthPath string
		fill_err = tp.Fill(&mounthPath)
		if fill_err != nil {
			return
		}
		var baseHeaderHex string
		fill_err = tp.Fill(&baseHeaderHex)
		if fill_err != nil {
			return
		}
		var slashingPkScript string
		fill_err = tp.Fill(&slashingPkScript)
		if fill_err != nil {
			return
		}
		var epochInterval uint
		fill_err = tp.Fill(&epochInterval)
		if fill_err != nil {
			return
		}
		if t1 == nil || t2 == nil {
			return
		}

		m, err := NewManager(t1)
		if err != nil {
			return
		}
		m.RunBabylondResource(t2, mounthPath, baseHeaderHex, slashingPkScript, epochInterval)
	})
}

func Fuzz_Nosy_Manager_RunBitcoindResource__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *testing.T
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var t2 *testing.T
		fill_err = tp.Fill(&t2)
		if fill_err != nil {
			return
		}
		var bitcoindCfgPath string
		fill_err = tp.Fill(&bitcoindCfgPath)
		if fill_err != nil {
			return
		}
		if t1 == nil || t2 == nil {
			return
		}

		m, err := NewManager(t1)
		if err != nil {
			return
		}
		m.RunBitcoindResource(t2, bitcoindCfgPath)
	})
}

func Fuzz_Nosy_noRestart__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var config *docker.HostConfig
		fill_err = tp.Fill(&config)
		if fill_err != nil {
			return
		}
		if config == nil {
			return
		}

		noRestart(config)
	})
}
