package types

import (
	"testing"

	"github.com/babylonlabs-io/babylon/btctxformatter"
	"github.com/babylonlabs-io/babylon/x/checkpointing/types"
	checkpointingtypes "github.com/babylonlabs-io/babylon/x/checkpointing/types"
	"github.com/boljen/go-bitmap"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
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

func Fuzz_Nosy_BTCCache_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxEntries uint32
		fill_err = tp.Fill(&maxEntries)
		if fill_err != nil {
			return
		}
		var ib *IndexedBlock
		fill_err = tp.Fill(&ib)
		if fill_err != nil {
			return
		}
		if ib == nil {
			return
		}

		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.Add(ib)
	})
}

func Fuzz_Nosy_BTCCache_FindBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxEntries uint32, blockHeight uint32) {
		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.FindBlock(blockHeight)
	})
}

func Fuzz_Nosy_BTCCache_First__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxEntries uint32) {
		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.First()
	})
}

func Fuzz_Nosy_BTCCache_GetAllBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxEntries uint32) {
		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.GetAllBlocks()
	})
}

func Fuzz_Nosy_BTCCache_GetLastBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxEntries uint32, stopHeight uint32) {
		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.GetLastBlocks(stopHeight)
	})
}

func Fuzz_Nosy_BTCCache_Init__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxEntries uint32
		fill_err = tp.Fill(&maxEntries)
		if fill_err != nil {
			return
		}
		var ibs []*IndexedBlock
		fill_err = tp.Fill(&ibs)
		if fill_err != nil {
			return
		}

		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.Init(ibs)
	})
}

func Fuzz_Nosy_BTCCache_RemoveAll__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxEntries uint32) {
		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.RemoveAll()
	})
}

func Fuzz_Nosy_BTCCache_RemoveLast__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxEntries uint32) {
		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.RemoveLast()
	})
}

func Fuzz_Nosy_BTCCache_Resize__(f *testing.F) {
	f.Fuzz(func(t *testing.T, m1 uint32, m2 uint32) {
		b, err := NewBTCCache(m1)
		if err != nil {
			return
		}
		b.Resize(m2)
	})
}

func Fuzz_Nosy_BTCCache_Size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxEntries uint32) {
		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.Size()
	})
}

func Fuzz_Nosy_BTCCache_Tip__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxEntries uint32) {
		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.Tip()
	})
}

func Fuzz_Nosy_BTCCache_Trim__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxEntries uint32) {
		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.Trim()
	})
}

func Fuzz_Nosy_BTCCache_TrimConfirmedBlocks__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxEntries uint32, k int) {
		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.TrimConfirmedBlocks(k)
	})
}

func Fuzz_Nosy_BTCCache_add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var maxEntries uint32
		fill_err = tp.Fill(&maxEntries)
		if fill_err != nil {
			return
		}
		var ib *IndexedBlock
		fill_err = tp.Fill(&ib)
		if fill_err != nil {
			return
		}
		if ib == nil {
			return
		}

		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.add(ib)
	})
}

func Fuzz_Nosy_BTCCache_size__(f *testing.F) {
	f.Fuzz(func(t *testing.T, maxEntries uint32) {
		b, err := NewBTCCache(maxEntries)
		if err != nil {
			return
		}
		b.size()
	})
}

func Fuzz_Nosy_CheckpointCache_AddCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag btctxformatter.BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version btctxformatter.FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var ckpt *Ckpt
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		c := NewCheckpointCache(tag, version)
		c.AddCheckpoint(ckpt)
	})
}

func Fuzz_Nosy_CheckpointCache_AddSegment__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag btctxformatter.BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version btctxformatter.FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}
		var ckptSeg *CkptSegment
		fill_err = tp.Fill(&ckptSeg)
		if fill_err != nil {
			return
		}
		if ckptSeg == nil {
			return
		}

		c := NewCheckpointCache(tag, version)
		c.AddSegment(ckptSeg)
	})
}

func Fuzz_Nosy_CheckpointCache_HasCheckpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag btctxformatter.BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version btctxformatter.FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}

		c := NewCheckpointCache(tag, version)
		c.HasCheckpoints()
	})
}

func Fuzz_Nosy_CheckpointCache_Match__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag btctxformatter.BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version btctxformatter.FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}

		c := NewCheckpointCache(tag, version)
		c.Match()
	})
}

func Fuzz_Nosy_CheckpointCache_NumCheckpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag btctxformatter.BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version btctxformatter.FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}

		c := NewCheckpointCache(tag, version)
		c.NumCheckpoints()
	})
}

func Fuzz_Nosy_CheckpointCache_NumSegments__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag btctxformatter.BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version btctxformatter.FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}

		c := NewCheckpointCache(tag, version)
		c.NumSegments()
	})
}

func Fuzz_Nosy_CheckpointCache_PopEarliestCheckpoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag btctxformatter.BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version btctxformatter.FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}

		c := NewCheckpointCache(tag, version)
		c.PopEarliestCheckpoint()
	})
}

func Fuzz_Nosy_CheckpointCache_sortCheckpoints__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tag btctxformatter.BabylonTag
		fill_err = tp.Fill(&tag)
		if fill_err != nil {
			return
		}
		var version btctxformatter.FormatVersion
		fill_err = tp.Fill(&version)
		if fill_err != nil {
			return
		}

		c := NewCheckpointCache(tag, version)
		c.sortCheckpoints()
	})
}

func Fuzz_Nosy_CheckpointRecord_EpochNum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *checkpointingtypes.RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		cr := NewCheckpointRecord(ckpt, height)
		cr.EpochNum()
	})
}

func Fuzz_Nosy_CheckpointRecord_ID__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckpt *checkpointingtypes.RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		var height uint32
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		cr := NewCheckpointRecord(ckpt, height)
		cr.ID()
	})
}

func Fuzz_Nosy_CheckpointsBookkeeper_Add__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var cr *CheckpointRecord
		fill_err = tp.Fill(&cr)
		if fill_err != nil {
			return
		}
		if cr == nil {
			return
		}

		cb := NewCheckpointsBookkeeper()
		cb.Add(cr)
	})
}

func Fuzz_Nosy_CheckpointsBookkeeper_Remove__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id string) {
		cb := NewCheckpointsBookkeeper()
		cb.Remove(id)
	})
}

func Fuzz_Nosy_CheckpointsBookkeeper_has__(f *testing.F) {
	f.Fuzz(func(t *testing.T, id string) {
		cb := NewCheckpointsBookkeeper()
		cb.has(id)
	})
}

func Fuzz_Nosy_Ckpt_MustGenSPVProofs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var ckptSeg1 *CkptSegment
		fill_err = tp.Fill(&ckptSeg1)
		if fill_err != nil {
			return
		}
		var ckptSeg2 *CkptSegment
		fill_err = tp.Fill(&ckptSeg2)
		if fill_err != nil {
			return
		}
		var epochNumber uint64
		fill_err = tp.Fill(&epochNumber)
		if fill_err != nil {
			return
		}
		if ckptSeg1 == nil || ckptSeg2 == nil {
			return
		}

		ckpt := NewCkpt(ckptSeg1, ckptSeg2, epochNumber)
		ckpt.MustGenSPVProofs()
	})
}

func Fuzz_Nosy_EpochInfo_Equal__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}
		var valSet types.ValidatorWithBlsKeySet
		fill_err = tp.Fill(&valSet)
		if fill_err != nil {
			return
		}
		var epochInfo *EpochInfo
		fill_err = tp.Fill(&epochInfo)
		if fill_err != nil {
			return
		}
		if epochInfo == nil {
			return
		}

		ei := NewEpochInfo(epochNum, valSet)
		ei.Equal(epochInfo)
	})
}

func Fuzz_Nosy_EpochInfo_GetEpochNumber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}
		var valSet types.ValidatorWithBlsKeySet
		fill_err = tp.Fill(&valSet)
		if fill_err != nil {
			return
		}

		ei := NewEpochInfo(epochNum, valSet)
		ei.GetEpochNumber()
	})
}

func Fuzz_Nosy_EpochInfo_GetSignersKeySetWithPowerSum__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}
		var valSet types.ValidatorWithBlsKeySet
		fill_err = tp.Fill(&valSet)
		if fill_err != nil {
			return
		}
		var bm bitmap.Bitmap
		fill_err = tp.Fill(&bm)
		if fill_err != nil {
			return
		}

		ei := NewEpochInfo(epochNum, valSet)
		ei.GetSignersKeySetWithPowerSum(bm)
	})
}

func Fuzz_Nosy_EpochInfo_GetTotalPower__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}
		var valSet types.ValidatorWithBlsKeySet
		fill_err = tp.Fill(&valSet)
		if fill_err != nil {
			return
		}

		ei := NewEpochInfo(epochNum, valSet)
		ei.GetTotalPower()
	})
}

func Fuzz_Nosy_EpochInfo_VerifyMultiSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var epochNum uint64
		fill_err = tp.Fill(&epochNum)
		if fill_err != nil {
			return
		}
		var valSet types.ValidatorWithBlsKeySet
		fill_err = tp.Fill(&valSet)
		if fill_err != nil {
			return
		}
		var ckpt *checkpointingtypes.RawCheckpoint
		fill_err = tp.Fill(&ckpt)
		if fill_err != nil {
			return
		}
		if ckpt == nil {
			return
		}

		ei := NewEpochInfo(epochNum, valSet)
		ei.VerifyMultiSig(ckpt)
	})
}

func Fuzz_Nosy_GenesisInfo_GetBLSKeySet__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		gi, err := GetGenesisInfoFromFile(filePath)
		if err != nil {
			return
		}
		gi.GetBLSKeySet()
	})
}

func Fuzz_Nosy_GenesisInfo_GetBaseBTCHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		gi, err := GetGenesisInfoFromFile(filePath)
		if err != nil {
			return
		}
		gi.GetBaseBTCHeight()
	})
}

func Fuzz_Nosy_GenesisInfo_GetCheckpointTag__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		gi, err := GetGenesisInfoFromFile(filePath)
		if err != nil {
			return
		}
		gi.GetCheckpointTag()
	})
}

func Fuzz_Nosy_GenesisInfo_GetEpochInterval__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string) {
		gi, err := GetGenesisInfoFromFile(filePath)
		if err != nil {
			return
		}
		gi.GetEpochInterval()
	})
}

func Fuzz_Nosy_GenesisInfo_SetBaseBTCHeight__(f *testing.F) {
	f.Fuzz(func(t *testing.T, filePath string, height uint32) {
		gi, err := GetGenesisInfoFromFile(filePath)
		if err != nil {
			return
		}
		gi.SetBaseBTCHeight(height)
	})
}

func Fuzz_Nosy_IndexedBlock_BlockHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
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
		var txs []*btcutil.Tx
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		ib := NewIndexedBlock(height, header, txs)
		ib.BlockHash()
	})
}

func Fuzz_Nosy_IndexedBlock_GenSPVProof__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
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
		var txs []*btcutil.Tx
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		var txIdx int
		fill_err = tp.Fill(&txIdx)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		ib := NewIndexedBlock(height, header, txs)
		ib.GenSPVProof(txIdx)
	})
}

func Fuzz_Nosy_IndexedBlock_MsgBlock__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
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
		var txs []*btcutil.Tx
		fill_err = tp.Fill(&txs)
		if fill_err != nil {
			return
		}
		if header == nil {
			return
		}

		ib := NewIndexedBlock(height, header, txs)
		ib.MsgBlock()
	})
}

func Fuzz_Nosy_PrivateKeyWithMutex_GetKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var key *secp256k1.PrivateKey
		fill_err = tp.Fill(&key)
		if fill_err != nil {
			return
		}
		if key == nil {
			return
		}

		p := NewPrivateKeyWithMutex(key)
		p.GetKey()
	})
}

// skipping Fuzz_Nosy_PrivateKeyWithMutex_UseKey__ because parameters include func, chan, or unsupported interface: func(key *github.com/decred/dcrd/dcrec/secp256k1/v4.PrivateKey)

func Fuzz_Nosy_UTXO_GetOutPoint__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *btcjson.ListUnspentResult
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var net *chaincfg.Params
		fill_err = tp.Fill(&net)
		if fill_err != nil {
			return
		}
		if r == nil || net == nil {
			return
		}

		u, err := NewUTXO(r, net)
		if err != nil {
			return
		}
		u.GetOutPoint()
	})
}

func Fuzz_Nosy_SupportedBtcNetwork_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c SupportedBtcNetwork
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.String()
	})
}

func Fuzz_Nosy_GetWrappedTxs__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var msg *wire.MsgBlock
		fill_err = tp.Fill(&msg)
		if fill_err != nil {
			return
		}
		if msg == nil {
			return
		}

		GetWrappedTxs(msg)
	})
}
