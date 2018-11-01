package core

import (
	"fmt"
	"log"
)

type BlockChain struct {
	Blocks []*Block
}

//创建区块链
func GenBlockChain(genesisData string) *BlockChain {
	genesisBlock := GenGenesisBlock(genesisData)
	blockchain := BlockChain{}
	blockchain.AddBlock(&genesisBlock)
	return &blockchain
}

//输入data，创建一个新区块并添加到区块链中
func (bc *BlockChain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenBlock(*preBlock, data)
	bc.AddBlock(&newBlock)
}

//添加区块
func (bc *BlockChain) AddBlock(b *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, b)
		return
	}
	if isValid(*b, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, b)
	} else {
		log.Fatal("invalid Block")
	}
}

//打印区块链
func (bc *BlockChain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %d\n", block.Timestamp)
		fmt.Printf("PrevBlockHash: %s\n", block.PrevBlockHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Println()
	}
}

//添加新区块之前校验是否合法
func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index != oldBlock.Index+1 {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	if CalHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
