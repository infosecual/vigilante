package config

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

func Fuzz_Nosy_DBConfig_GetDBBackend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, homePath string) {
		cfg := DefaultDBConfigWithHomePath(homePath)
		cfg.GetDBBackend()
	})
}

func Fuzz_Nosy_DBConfig_ToBoltBackendConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, homePath string) {
		cfg := DefaultDBConfigWithHomePath(homePath)
		cfg.ToBoltBackendConfig()
	})
}

func Fuzz_Nosy_DBConfig_Validate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, homePath string) {
		cfg := DefaultDBConfigWithHomePath(homePath)
		cfg.Validate()
	})
}

func Fuzz_Nosy_DataDir__(f *testing.F) {
	f.Fuzz(func(t *testing.T, homePath string) {
		DataDir(homePath)
	})
}

func Fuzz_Nosy_isOneOf__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var v string
		fill_err = tp.Fill(&v)
		if fill_err != nil {
			return
		}
		var list []string
		fill_err = tp.Fill(&list)
		if fill_err != nil {
			return
		}

		isOneOf(v, list)
	})
}
