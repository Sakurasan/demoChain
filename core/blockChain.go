package core

import (
	"encoding/json"
	"log"
)

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	generateBlock := GenerateBlock()
	blockChin := new(BlockChain)
	blockChin.AppendBlock(&generateBlock)
	return blockChin
}

func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlock(&newBlock)

}

func (bc *BlockChain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(newBlock, bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Println("非法的区块")
	}

}

func isValid(newBlock *Block, oldBlock *Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PreBlockHash != oldBlock.Hash {
		return false
	}
	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

func (bc *BlockChain) Print() {
	// for _,b :=range bc{
	// 	log.Println("Index:",)
	// }
	bj, _ := json.Marshal(bc)
	log.Println(string(bj))
}
