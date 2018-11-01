package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index         int64  //区块编号
	Timestamp     int64  //区块时间戳
	PrevBlockHash string //上个区块的hash
	Hash          string //当前区块的hash
	Data          string //区块数据
}

//计算当前区块的hash
func CalHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashEncodeBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashEncodeBytes[:])
}

//生成一个新区快
func GenBlock(prevBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.PrevBlockHash = prevBlock.Hash
	newBlock.Data = data
	newBlock.Hash = CalHash(newBlock)
	return newBlock
}

//生成一个创世区块
func GenGenesisBlock(genesisData string) Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenBlock(preBlock, genesisData)
}
