package store

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

func Fuzz_Nosy_MonitorStore_LatestEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MonitorStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.LatestEpoch()
	})
}

func Fuzz_Nosy_MonitorStore_LatestHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MonitorStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.LatestHeight()
	})
}

func Fuzz_Nosy_MonitorStore_PutLatestEpoch__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MonitorStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.PutLatestEpoch(epoch)
	})
}

func Fuzz_Nosy_MonitorStore_PutLatestHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MonitorStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.PutLatestHeight(height)
	})
}

func Fuzz_Nosy_MonitorStore_createBuckets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MonitorStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.createBuckets()
	})
}

func Fuzz_Nosy_MonitorStore_get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MonitorStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var bucketName []byte
		fill_err = tp.Fill(&bucketName)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.get(key, bucketName)
	})
}

func Fuzz_Nosy_MonitorStore_put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *MonitorStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var val uint64
		fill_err = tp.Fill(&val)
		if fill_err != nil {
			return
		}
		var bucketName []byte
		fill_err = tp.Fill(&bucketName)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.put(key, val, bucketName)
	})
}

func Fuzz_Nosy_uint64FromBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		uint64FromBytes(b)
	})
}

func Fuzz_Nosy_uint64ToBytes__(f *testing.F) {
	f.Fuzz(func(t *testing.T, v uint64) {
		uint64ToBytes(v)
	})
}
