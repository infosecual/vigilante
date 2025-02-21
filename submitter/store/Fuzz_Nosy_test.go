package store

import (
	"testing"

	"github.com/babylonlabs-io/vigilante/proto"
	"github.com/btcsuite/btcd/wire"
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

func Fuzz_Nosy_StoredCheckpoint_FromProto__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx1 *wire.MsgTx
		fill_err = tp.Fill(&tx1)
		if fill_err != nil {
			return
		}
		var tx2 *wire.MsgTx
		fill_err = tp.Fill(&tx2)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		var protoTx *proto.StoredCheckpoint
		fill_err = tp.Fill(&protoTx)
		if fill_err != nil {
			return
		}
		if tx1 == nil || tx2 == nil || protoTx == nil {
			return
		}

		s := NewStoredCheckpoint(tx1, tx2, epoch)
		s.FromProto(protoTx)
	})
}

func Fuzz_Nosy_StoredCheckpoint_ToProto__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx1 *wire.MsgTx
		fill_err = tp.Fill(&tx1)
		if fill_err != nil {
			return
		}
		var tx2 *wire.MsgTx
		fill_err = tp.Fill(&tx2)
		if fill_err != nil {
			return
		}
		var epoch uint64
		fill_err = tp.Fill(&epoch)
		if fill_err != nil {
			return
		}
		if tx1 == nil || tx2 == nil {
			return
		}

		s := NewStoredCheckpoint(tx1, tx2, epoch)
		s.ToProto()
	})
}

func Fuzz_Nosy_SubmitterStore_LatestCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SubmitterStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.LatestCheckpoint()
	})
}

func Fuzz_Nosy_SubmitterStore_PutCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SubmitterStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ckpt *StoredCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if s == nil || ckpt == nil {
			return
		}

		s.PutCheckpoint(ckpt)
	})
}

func Fuzz_Nosy_SubmitterStore_createBuckets__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SubmitterStore
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

func Fuzz_Nosy_SubmitterStore_get__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SubmitterStore
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

func Fuzz_Nosy_SubmitterStore_put__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *SubmitterStore
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var key []byte
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		var val []byte
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
