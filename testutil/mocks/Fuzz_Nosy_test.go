package mocks

import (
	"testing"

	btcjson "github.com/btcsuite/btcd/btcjson"
	chainhash "github.com/btcsuite/btcd/chaincfg/chainhash"
	wire "github.com/btcsuite/btcd/wire"
	gomock "github.com/golang/mock/gomock"
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

func Fuzz_Nosy_MockBTCClient_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCClient(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockBTCClient_FindTailBlocksByHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCClient(ctrl)
		m.FindTailBlocksByHeight(height)
	})
}

func Fuzz_Nosy_MockBTCClient_GetBestBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCClient(ctrl)
		m.GetBestBlock()
	})
}

func Fuzz_Nosy_MockBTCClient_GetBlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var blockHash *chainhash.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if ctrl == nil || blockHash == nil {
			return
		}

		m := NewMockBTCClient(ctrl)
		m.GetBlockByHash(blockHash)
	})
}

func Fuzz_Nosy_MockBTCClient_GetBlockByHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCClient(ctrl)
		m.GetBlockByHeight(height)
	})
}

func Fuzz_Nosy_MockBTCClient_GetRawTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var txHash *chainhash.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if ctrl == nil || txHash == nil {
			return
		}

		m := NewMockBTCClient(ctrl)
		m.GetRawTransaction(txHash)
	})
}

func Fuzz_Nosy_MockBTCClient_GetTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var txHash *chainhash.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if ctrl == nil || txHash == nil {
			return
		}

		m := NewMockBTCClient(ctrl)
		m.GetTransaction(txHash)
	})
}

func Fuzz_Nosy_MockBTCClient_GetTxOut__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var txHash *chainhash.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		var index uint32
		fill_err = tp.Fill(&index)
		if fill_err != nil {
			return
		}
		var mempool bool
		fill_err = tp.Fill(&mempool)
		if fill_err != nil {
			return
		}
		if ctrl == nil || txHash == nil {
			return
		}

		m := NewMockBTCClient(ctrl)
		m.GetTxOut(txHash, index, mempool)
	})
}

func Fuzz_Nosy_MockBTCClient_SendRawTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var allowHighFees bool
		fill_err = tp.Fill(&allowHighFees)
		if fill_err != nil {
			return
		}
		if ctrl == nil || tx == nil {
			return
		}

		m := NewMockBTCClient(ctrl)
		m.SendRawTransaction(tx, allowHighFees)
	})
}

func Fuzz_Nosy_MockBTCClient_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCClient(ctrl)
		m.Stop()
	})
}

func Fuzz_Nosy_MockBTCClient_TxDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var txHash *chainhash.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		var pkScript []byte
		fill_err = tp.Fill(&pkScript)
		if fill_err != nil {
			return
		}
		if ctrl == nil || txHash == nil {
			return
		}

		m := NewMockBTCClient(ctrl)
		m.TxDetails(txHash, pkScript)
	})
}

func Fuzz_Nosy_MockBTCClient_WaitForShutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCClient(ctrl)
		m.WaitForShutdown()
	})
}

// skipping Fuzz_Nosy_MockBTCClientMockRecorder_FindTailBlocksByHeight__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBTCClientMockRecorder_GetBestBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBTCClientMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.GetBestBlock()
	})
}

// skipping Fuzz_Nosy_MockBTCClientMockRecorder_GetBlockByHash__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBTCClientMockRecorder_GetBlockByHeight__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBTCClientMockRecorder_GetRawTransaction__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBTCClientMockRecorder_GetTransaction__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBTCClientMockRecorder_GetTxOut__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBTCClientMockRecorder_SendRawTransaction__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBTCClientMockRecorder_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBTCClientMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.Stop()
	})
}

// skipping Fuzz_Nosy_MockBTCClientMockRecorder_TxDetails__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBTCClientMockRecorder_WaitForShutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBTCClientMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.WaitForShutdown()
	})
}

func Fuzz_Nosy_MockBTCWallet_EXPECT__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.EXPECT()
	})
}

func Fuzz_Nosy_MockBTCWallet_FundRawTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var opts btcjson.FundRawTransactionOpts
		fill_err = tp.Fill(&opts)
		if fill_err != nil {
			return
		}
		var isWitness *bool
		fill_err = tp.Fill(&isWitness)
		if fill_err != nil {
			return
		}
		if ctrl == nil || tx == nil || isWitness == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.FundRawTransaction(tx, opts, isWitness)
	})
}

func Fuzz_Nosy_MockBTCWallet_GetBTCConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.GetBTCConfig()
	})
}

func Fuzz_Nosy_MockBTCWallet_GetHighUTXOAndSum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.GetHighUTXOAndSum()
	})
}

func Fuzz_Nosy_MockBTCWallet_GetNetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.GetNetParams()
	})
}

func Fuzz_Nosy_MockBTCWallet_GetNewAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var account string
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.GetNewAddress(account)
	})
}

func Fuzz_Nosy_MockBTCWallet_GetRawTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var txHash *chainhash.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if ctrl == nil || txHash == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.GetRawTransaction(txHash)
	})
}

func Fuzz_Nosy_MockBTCWallet_GetWalletLockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.GetWalletLockTime()
	})
}

func Fuzz_Nosy_MockBTCWallet_GetWalletPass__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.GetWalletPass()
	})
}

func Fuzz_Nosy_MockBTCWallet_ListReceivedByAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.ListReceivedByAddress()
	})
}

func Fuzz_Nosy_MockBTCWallet_ListUnspent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.ListUnspent()
	})
}

func Fuzz_Nosy_MockBTCWallet_SendRawTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var allowHighFees bool
		fill_err = tp.Fill(&allowHighFees)
		if fill_err != nil {
			return
		}
		if ctrl == nil || tx == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.SendRawTransaction(tx, allowHighFees)
	})
}

func Fuzz_Nosy_MockBTCWallet_SignRawTransactionWithWallet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if ctrl == nil || tx == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.SignRawTransactionWithWallet(tx)
	})
}

func Fuzz_Nosy_MockBTCWallet_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.Stop()
	})
}

func Fuzz_Nosy_MockBTCWallet_TxDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var txHash *chainhash.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		var pkScript []byte
		fill_err = tp.Fill(&pkScript)
		if fill_err != nil {
			return
		}
		if ctrl == nil || txHash == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.TxDetails(txHash, pkScript)
	})
}

func Fuzz_Nosy_MockBTCWallet_WalletPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ctrl *gomock.Controller
		fill_err = tp.Fill(&ctrl)
		if fill_err != nil {
			return
		}
		var passphrase string
		fill_err = tp.Fill(&passphrase)
		if fill_err != nil {
			return
		}
		var timeoutSecs int64
		fill_err = tp.Fill(&timeoutSecs)
		if fill_err != nil {
			return
		}
		if ctrl == nil {
			return
		}

		m := NewMockBTCWallet(ctrl)
		m.WalletPassphrase(passphrase, timeoutSecs)
	})
}

// skipping Fuzz_Nosy_MockBTCWalletMockRecorder_FundRawTransaction__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBTCWalletMockRecorder_GetBTCConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBTCWalletMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.GetBTCConfig()
	})
}

func Fuzz_Nosy_MockBTCWalletMockRecorder_GetHighUTXOAndSum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBTCWalletMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.GetHighUTXOAndSum()
	})
}

func Fuzz_Nosy_MockBTCWalletMockRecorder_GetNetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBTCWalletMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.GetNetParams()
	})
}

// skipping Fuzz_Nosy_MockBTCWalletMockRecorder_GetNewAddress__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBTCWalletMockRecorder_GetRawTransaction__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBTCWalletMockRecorder_GetWalletLockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBTCWalletMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.GetWalletLockTime()
	})
}

func Fuzz_Nosy_MockBTCWalletMockRecorder_GetWalletPass__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBTCWalletMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.GetWalletPass()
	})
}

func Fuzz_Nosy_MockBTCWalletMockRecorder_ListReceivedByAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBTCWalletMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.ListReceivedByAddress()
	})
}

func Fuzz_Nosy_MockBTCWalletMockRecorder_ListUnspent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBTCWalletMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.ListUnspent()
	})
}

// skipping Fuzz_Nosy_MockBTCWalletMockRecorder_SendRawTransaction__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBTCWalletMockRecorder_SignRawTransactionWithWallet__ because parameters include func, chan, or unsupported interface: interface{}

func Fuzz_Nosy_MockBTCWalletMockRecorder_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var mr *MockBTCWalletMockRecorder
		fill_err = tp.Fill(&mr)
		if fill_err != nil {
			return
		}
		if mr == nil {
			return
		}

		mr.Stop()
	})
}

// skipping Fuzz_Nosy_MockBTCWalletMockRecorder_TxDetails__ because parameters include func, chan, or unsupported interface: interface{}

// skipping Fuzz_Nosy_MockBTCWalletMockRecorder_WalletPassphrase__ because parameters include func, chan, or unsupported interface: interface{}
