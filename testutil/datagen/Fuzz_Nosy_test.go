package datagen

import (
	"math/rand"
	"testing"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
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

func Fuzz_Nosy_GenRandomBabylonTxPair__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		GenRandomBabylonTxPair(r)
	})
}

func Fuzz_Nosy_GenRandomBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var numBabylonTxs int
		fill_err = tp.Fill(&numBabylonTxs)
		if fill_err != nil {
			return
		}
		var prevHash *chainhash.Hash
		fill_err = tp.Fill(&prevHash)
		if fill_err != nil {
			return
		}
		if r == nil || prevHash == nil {
			return
		}

		GenRandomBlock(r, numBabylonTxs, prevHash)
	})
}

func Fuzz_Nosy_GenRandomBlockchainWithBabylonTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var n uint64
		fill_err = tp.Fill(&n)
		if fill_err != nil {
			return
		}
		var partialPercentage float32
		fill_err = tp.Fill(&partialPercentage)
		if fill_err != nil {
			return
		}
		var fullPercentage float32
		fill_err = tp.Fill(&fullPercentage)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		GenRandomBlockchainWithBabylonTx(r, n, partialPercentage, fullPercentage)
	})
}

func Fuzz_Nosy_GetRandomIndexedBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var numBlocks uint64
		fill_err = tp.Fill(&numBlocks)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		GetRandomIndexedBlocks(r, numBlocks)
	})
}

func Fuzz_Nosy_GetRandomIndexedBlocksFromHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rand.Rand
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var numBlocks uint64
		fill_err = tp.Fill(&numBlocks)
		if fill_err != nil {
			return
		}
		var rootHeight int32
		fill_err = tp.Fill(&rootHeight)
		if fill_err != nil {
			return
		}
		var rootHash chainhash.Hash
		fill_err = tp.Fill(&rootHash)
		if fill_err != nil {
			return
		}
		if r == nil {
			return
		}

		GetRandomIndexedBlocksFromHeight(r, numBlocks, rootHeight, rootHash)
	})
}
