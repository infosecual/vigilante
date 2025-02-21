package btcscanner

import (
	"testing"

	"github.com/babylonlabs-io/vigilante/types"
	"github.com/btcsuite/btcd/wire"
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

func Fuzz_Nosy_BtcScanner_Bootstrap__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Bootstrap()
	})
}

func Fuzz_Nosy_BtcScanner_GetBaseHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.GetBaseHeight()
	})
}

func Fuzz_Nosy_BtcScanner_GetBtcClient__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.GetBtcClient()
	})
}

func Fuzz_Nosy_BtcScanner_GetCheckpointsChan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.GetCheckpointsChan()
	})
}

func Fuzz_Nosy_BtcScanner_GetConfirmedBlocksChan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.GetConfirmedBlocksChan()
	})
}

func Fuzz_Nosy_BtcScanner_GetK__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.GetK()
	})
}

func Fuzz_Nosy_BtcScanner_GetUnconfirmedBlockCache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.GetUnconfirmedBlockCache()
	})
}

func Fuzz_Nosy_BtcScanner_SetBaseHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var h uint32
		fill_err = tp.Fill(&h)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.SetBaseHeight(h)
	})
}

// skipping Fuzz_Nosy_BtcScanner_SetBtcClient__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCClient

// skipping Fuzz_Nosy_BtcScanner_SetConfirmedBlocksChan__ because parameters include func, chan, or unsupported interface: chan *github.com/babylonlabs-io/vigilante/types.IndexedBlock

func Fuzz_Nosy_BtcScanner_SetK__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var k uint32
		fill_err = tp.Fill(&k)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.SetK(k)
	})
}

func Fuzz_Nosy_BtcScanner_SetLogger__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var logger *zap.SugaredLogger
		fill_err = tp.Fill(&logger)
		if fill_err != nil {
			return
		}
		if bs == nil || logger == nil {
			return
		}

		bs.SetLogger(logger)
	})
}

func Fuzz_Nosy_BtcScanner_SetUnconfirmedBlockCache__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var c *types.BTCCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		if bs == nil || c == nil {
			return
		}

		bs.SetUnconfirmedBlockCache(c)
	})
}

func Fuzz_Nosy_BtcScanner_Start__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var startHeight uint32
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Start(startHeight)
	})
}

func Fuzz_Nosy_BtcScanner_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.Stop()
	})
}

func Fuzz_Nosy_BtcScanner_bootstrapAndBlockEventHandler__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.bootstrapAndBlockEventHandler()
	})
}

func Fuzz_Nosy_BtcScanner_handleNewBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		var header *wire.BlockHeader
		fill_err = tp.Fill(&header)
		if fill_err != nil {
			return
		}
		if bs == nil || header == nil {
			return
		}

		bs.handleNewBlock(height, header)
	})
}

func Fuzz_Nosy_BtcScanner_matchAndPop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.matchAndPop()
	})
}

func Fuzz_Nosy_BtcScanner_sendConfirmedBlocksToChan__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var blocks []*types.IndexedBlock
		fill_err = tp.Fill(&blocks)
		if fill_err != nil {
			return
		}
		if bs == nil {
			return
		}

		bs.sendConfirmedBlocksToChan(blocks)
	})
}

func Fuzz_Nosy_BtcScanner_tryToExtractCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var block *types.IndexedBlock
		fill_err = tp.Fill(&block)
		if fill_err != nil {
			return
		}
		if bs == nil || block == nil {
			return
		}

		bs.tryToExtractCheckpoint(block)
	})
}

func Fuzz_Nosy_BtcScanner_tryToExtractCkptSegment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var bs *BtcScanner
		fill_err = tp.Fill(&bs)
		if fill_err != nil {
			return
		}
		var b *types.IndexedBlock
		fill_err = tp.Fill(&b)
		if fill_err != nil {
			return
		}
		if bs == nil || b == nil {
			return
		}

		bs.tryToExtractCkptSegment(b)
	})
}

// skipping Fuzz_Nosy_Scanner_GetBaseHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor/btcscanner.Scanner

// skipping Fuzz_Nosy_Scanner_GetCheckpointsChan__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor/btcscanner.Scanner

// skipping Fuzz_Nosy_Scanner_GetConfirmedBlocksChan__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor/btcscanner.Scanner

// skipping Fuzz_Nosy_Scanner_Start__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor/btcscanner.Scanner

// skipping Fuzz_Nosy_Scanner_Stop__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/monitor/btcscanner.Scanner
