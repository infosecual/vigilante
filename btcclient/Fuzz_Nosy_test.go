package btcclient

import (
	"testing"

	"github.com/babylonlabs-io/vigilante/config"
	"github.com/babylonlabs-io/vigilante/types"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/lightningnetwork/lnd/chainntnfs"
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

func Fuzz_Nosy_Client_FindTailBlocksByHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		var baseHeight uint32
		fill_err = tp.Fill(&baseHeight)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.FindTailBlocksByHeight(baseHeight)
	})
}

func Fuzz_Nosy_Client_FundRawTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
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
		if cfg == nil || parentLogger == nil || tx == nil || isWitness == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.FundRawTransaction(tx, opts, isWitness)
	})
}

func Fuzz_Nosy_Client_GetBTCConfig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.GetBTCConfig()
	})
}

func Fuzz_Nosy_Client_GetBestBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.GetBestBlock()
	})
}

func Fuzz_Nosy_Client_GetBlockByHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		var blockHash *chainhash.Hash
		fill_err = tp.Fill(&blockHash)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil || blockHash == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.GetBlockByHash(blockHash)
	})
}

func Fuzz_Nosy_Client_GetBlockByHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.GetBlockByHeight(height)
	})
}

func Fuzz_Nosy_Client_GetHighUTXOAndSum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.GetHighUTXOAndSum()
	})
}

func Fuzz_Nosy_Client_GetNetParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.GetNetParams()
	})
}

func Fuzz_Nosy_Client_GetNewAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		var account string
		fill_err = tp.Fill(&account)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.GetNewAddress(account)
	})
}

func Fuzz_Nosy_Client_GetRawTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		var txHash *chainhash.Hash
		fill_err = tp.Fill(&txHash)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil || txHash == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.GetRawTransaction(txHash)
	})
}

func Fuzz_Nosy_Client_GetTipBlockVerbose__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.GetTipBlockVerbose()
	})
}

func Fuzz_Nosy_Client_GetWalletLockTime__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.GetWalletLockTime()
	})
}

func Fuzz_Nosy_Client_GetWalletPass__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.GetWalletPass()
	})
}

func Fuzz_Nosy_Client_ListReceivedByAddress__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.ListReceivedByAddress()
	})
}

func Fuzz_Nosy_Client_ListUnspent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.ListUnspent()
	})
}

func Fuzz_Nosy_Client_SendRawTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
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
		if cfg == nil || parentLogger == nil || tx == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.SendRawTransaction(tx, allowHighFees)
	})
}

func Fuzz_Nosy_Client_SignRawTransactionWithWallet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil || tx == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.SignRawTransactionWithWallet(tx)
	})
}

func Fuzz_Nosy_Client_Stop__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.Stop()
	})
}

func Fuzz_Nosy_Client_TxDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
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
		if cfg == nil || parentLogger == nil || txHash == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.TxDetails(txHash, pkScript)
	})
}

func Fuzz_Nosy_Client_WalletPassphrase__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
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
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.WalletPassphrase(passphrase, timeoutSecs)
	})
}

func Fuzz_Nosy_Client_getBestBlockHashWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.getBestBlockHashWithRetry()
	})
}

func Fuzz_Nosy_Client_getBestIndexedBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.getBestIndexedBlock()
	})
}

func Fuzz_Nosy_Client_getBlockCountWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.getBlockCountWithRetry()
	})
}

func Fuzz_Nosy_Client_getBlockHashWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.getBlockHashWithRetry(height)
	})
}

func Fuzz_Nosy_Client_getBlockVerboseWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		var hash *chainhash.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil || hash == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.getBlockVerboseWithRetry(hash)
	})
}

func Fuzz_Nosy_Client_getBlockWithRetry__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		var hash *chainhash.Hash
		fill_err = tp.Fill(&hash)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil || hash == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.getBlockWithRetry(hash)
	})
}

func Fuzz_Nosy_Client_getChainBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		var baseHeight uint32
		fill_err = tp.Fill(&baseHeight)
		if fill_err != nil {
			return
		}
		var tipBlock *types.IndexedBlock
		fill_err = tp.Fill(&tipBlock)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil || tipBlock == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.getChainBlocks(baseHeight, tipBlock)
	})
}

func Fuzz_Nosy_Client_getTxDetails__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cfg *config.Config
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var parentLogger *zap.Logger
		fill_err = tp.Fill(&parentLogger)
		if fill_err != nil {
			return
		}
		var req chainntnfs.ConfRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		var msg string
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if cfg == nil || parentLogger == nil {
			return
		}

		c, err := NewWallet(cfg, parentLogger)
		if err != nil {
			return
		}
		c.getTxDetails(req, msg)
	})
}

func Fuzz_Nosy_EmptyHintCache_CommitConfirmHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EmptyHintCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var _x2 uint32
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 []chainntnfs.ConfRequest
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.CommitConfirmHint(_x2, _x3...)
	})
}

func Fuzz_Nosy_EmptyHintCache_CommitSpendHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EmptyHintCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var _x2 uint32
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 []chainntnfs.SpendRequest
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.CommitSpendHint(_x2, _x3...)
	})
}

func Fuzz_Nosy_EmptyHintCache_PurgeConfirmHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EmptyHintCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var _x2 []chainntnfs.ConfRequest
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.PurgeConfirmHint(_x2...)
	})
}

func Fuzz_Nosy_EmptyHintCache_PurgeSpendHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EmptyHintCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var _x2 []chainntnfs.SpendRequest
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.PurgeSpendHint(_x2...)
	})
}

func Fuzz_Nosy_EmptyHintCache_QueryConfirmHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EmptyHintCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var _x2 chainntnfs.ConfRequest
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.QueryConfirmHint(_x2)
	})
}

func Fuzz_Nosy_EmptyHintCache_QuerySpendHint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c *EmptyHintCache
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}
		var _x2 chainntnfs.SpendRequest
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		if c == nil {
			return
		}

		c.QuerySpendHint(_x2)
	})
}

// skipping Fuzz_Nosy_BTCClient_FindTailBlocksByHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCClient

// skipping Fuzz_Nosy_BTCClient_GetBestBlock__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCClient

// skipping Fuzz_Nosy_BTCClient_GetBlockByHash__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCClient

// skipping Fuzz_Nosy_BTCClient_GetBlockByHeight__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCClient

// skipping Fuzz_Nosy_BTCClient_GetRawTransaction__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCClient

// skipping Fuzz_Nosy_BTCClient_GetTransaction__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCClient

// skipping Fuzz_Nosy_BTCClient_GetTxOut__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCClient

// skipping Fuzz_Nosy_BTCClient_SendRawTransaction__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCClient

// skipping Fuzz_Nosy_BTCClient_Stop__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCClient

// skipping Fuzz_Nosy_BTCClient_TxDetails__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCClient

// skipping Fuzz_Nosy_BTCClient_WaitForShutdown__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCClient

// skipping Fuzz_Nosy_BTCWallet_FundRawTransaction__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_GetBTCConfig__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_GetHighUTXOAndSum__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_GetNetParams__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_GetNewAddress__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_GetRawTransaction__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_GetWalletLockTime__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_GetWalletPass__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_ListReceivedByAddress__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_ListUnspent__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_SendRawTransaction__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_SignRawTransactionWithWallet__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_Stop__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_TxDetails__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

// skipping Fuzz_Nosy_BTCWallet_WalletPassphrase__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/vigilante/btcclient.BTCWallet

func Fuzz_Nosy_BuildDialer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, rpcHost string) {
		BuildDialer(rpcHost)
	})
}

func Fuzz_Nosy_rpcHostURL__(f *testing.F) {
	f.Fuzz(func(t *testing.T, host string, walletName string) {
		rpcHostURL(host, walletName)
	})
}
