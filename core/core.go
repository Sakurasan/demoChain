package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        int64  //区块编号
	TimeStrap    int64  //时间戳
	PreBlockHash string //上一个区块哈希
	Hash         string //当前哈希

	Data string //区块数据

}

func CalculateHash(b *Block) string {
	blockdata := string(b.Index) + string(b.TimeStrap) + b.PreBlockHash + b.Data
	hash := sha256.Sum256([]byte(blockdata))
	return hex.EncodeToString(hash[:])
}

func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.TimeStrap = time.Now().Unix()
	newBlock.PreBlockHash = preBlock.Hash

	newBlock.Data = data
	newBlock.Hash = CalculateHash(&newBlock)
	return newBlock
}

func GenerateBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.TimeStrap = time.Now().Unix()
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "创世块")
}
