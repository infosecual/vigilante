package relayer

import (
	"testing"

	checkpointingtypes "github.com/babylonlabs-io/babylon/x/checkpointing/types"
	"github.com/babylonlabs-io/vigilante/types"
	"github.com/btcsuite/btcd/btcutil"
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

func Fuzz_Nosy_Relayer_ChainTwoTxAndSend__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var data1 []byte
		fill_err = tp.Fill(&data1)
		if fill_err != nil {
			return
		}
		var data2 []byte
		fill_err = tp.Fill(&data2)
		if fill_err != nil {
			return
		}
		if rl == nil {
			return
		}

		rl.ChainTwoTxAndSend(data1, data2)
	})
}

func Fuzz_Nosy_Relayer_GetChangeAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		if rl == nil {
			return
		}

		rl.GetChangeAddress()
	})
}

func Fuzz_Nosy_Relayer_MaybeResubmitSecondCheckpointTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var ckpt *checkpointingtypes.RawCheckpointWithMetaResponse
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if rl == nil || ckpt == nil {
			return
		}

		rl.MaybeResubmitSecondCheckpointTx(ckpt)
	})
}

func Fuzz_Nosy_Relayer_SendCheckpointToBTC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var ckpt *checkpointingtypes.RawCheckpointWithMetaResponse
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if rl == nil || ckpt == nil {
			return
		}

		rl.SendCheckpointToBTC(ckpt)
	})
}

// skipping Fuzz_Nosy_Relayer_buildAndSendTx__ because parameters include func, chan, or unsupported interface: func() (*github.com/babylonlabs-io/vigilante/types.BtcTxInfo, error)

func Fuzz_Nosy_Relayer_buildChainedDataTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		var prevTx *wire.MsgTx
		fill_err = tp.Fill(&prevTx)
		if fill_err != nil {
			return
		}
		if rl == nil || prevTx == nil {
			return
		}

		rl.buildChainedDataTx(d2, prevTx)
	})
}

func Fuzz_Nosy_Relayer_buildDataTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var d2 []byte
		fill_err = tp.Fill(&d2)
		if fill_err != nil {
			return
		}
		if rl == nil {
			return
		}

		rl.buildDataTx(d2)
	})
}

func Fuzz_Nosy_Relayer_calcMinRelayFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var txVirtualSize int64
		fill_err = tp.Fill(&txVirtualSize)
		if fill_err != nil {
			return
		}
		if rl == nil {
			return
		}

		rl.calcMinRelayFee(txVirtualSize)
	})
}

func Fuzz_Nosy_Relayer_calculateBumpedFee__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var ckptInfo *types.CheckpointInfo
		fill_err = tp.Fill(&ckptInfo)
		if fill_err != nil {
			return
		}
		if rl == nil || ckptInfo == nil {
			return
		}

		rl.calculateBumpedFee(ckptInfo)
	})
}

func Fuzz_Nosy_Relayer_convertCkptToTwoTxAndSubmit__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var ckpt *checkpointingtypes.RawCheckpointResponse
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if rl == nil || ckpt == nil {
			return
		}

		rl.convertCkptToTwoTxAndSubmit(ckpt)
	})
}

func Fuzz_Nosy_Relayer_encodeCheckpointData__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var ckpt *checkpointingtypes.RawCheckpointResponse
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if rl == nil || ckpt == nil {
			return
		}

		rl.encodeCheckpointData(ckpt)
	})
}

func Fuzz_Nosy_Relayer_finalizeTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if rl == nil || tx == nil {
			return
		}

		rl.finalizeTransaction(tx)
	})
}

func Fuzz_Nosy_Relayer_getFeeRate__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		if rl == nil {
			return
		}

		rl.getFeeRate()
	})
}

func Fuzz_Nosy_Relayer_logAndRecordCheckpointMetrics__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var tx1 *types.BtcTxInfo
		fill_err = tp.Fill(&tx1)
		if fill_err != nil {
			return
		}
		var tx2 *types.BtcTxInfo
		fill_err = tp.Fill(&tx2)
		if fill_err != nil {
			return
		}
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}
		if rl == nil || tx1 == nil || tx2 == nil {
			return
		}

		rl.logAndRecordCheckpointMetrics(tx1, tx2, epochNum)
	})
}

func Fuzz_Nosy_Relayer_maybeResendSecondTxOfCheckpointToBTC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var tx2 *types.BtcTxInfo
		fill_err = tp.Fill(&tx2)
		if fill_err != nil {
			return
		}
		var bumpedFee btcutil.Amount
		fill_err = tp.Fill(&bumpedFee)
		if fill_err != nil {
			return
		}
		if rl == nil || tx2 == nil {
			return
		}

		rl.maybeResendSecondTxOfCheckpointToBTC(tx2, bumpedFee)
	})
}

func Fuzz_Nosy_Relayer_retrySendTx2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var ckpt *checkpointingtypes.RawCheckpointResponse
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if rl == nil || ckpt == nil {
			return
		}

		rl.retrySendTx2(ckpt)
	})
}

func Fuzz_Nosy_Relayer_sendTxToBTC__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if rl == nil || tx == nil {
			return
		}

		rl.sendTxToBTC(tx)
	})
}

func Fuzz_Nosy_Relayer_shouldResendCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var ckptInfo *types.CheckpointInfo
		fill_err = tp.Fill(&ckptInfo)
		if fill_err != nil {
			return
		}
		var bumpedFee btcutil.Amount
		fill_err = tp.Fill(&bumpedFee)
		if fill_err != nil {
			return
		}
		if rl == nil || ckptInfo == nil {
			return
		}

		rl.shouldResendCheckpoint(ckptInfo, bumpedFee)
	})
}

func Fuzz_Nosy_Relayer_shouldSendCompleteCkpt__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var ckptEpoch uint64
		fill_err = tp.Fill(&ckptEpoch)
		if fill_err != nil {
			return
		}
		if rl == nil {
			return
		}

		rl.shouldSendCompleteCkpt(ckptEpoch)
	})
}

func Fuzz_Nosy_Relayer_shouldSendTx2__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var ckptEpoch uint64
		fill_err = tp.Fill(&ckptEpoch)
		if fill_err != nil {
			return
		}
		if rl == nil {
			return
		}

		rl.shouldSendTx2(ckptEpoch)
	})
}

func Fuzz_Nosy_Relayer_signTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var rl *Relayer
		fill_err = tp.Fill(&rl)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if rl == nil || tx == nil {
			return
		}

		rl.signTx(tx)
	})
}

func Fuzz_Nosy_calculateTxVirtualSize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		calculateTxVirtualSize(tx)
	})
}

// skipping Fuzz_Nosy_maybeResendFromStore__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/submitter/relayer.GetLatestCheckpointFunc

func Fuzz_Nosy_rpcHostURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, host string, walletName string) {
		rpcHostURL(host, walletName)
	})
}
