package core

import (
	"time"
)

type Block struct {
	Timestamp int64 //区块创建的时间戳
	Data []byte //区块包含的数据
	PrevBlockHash []byte //前一个区块的哈希值
	Hash []byte //区块自身的哈希值，用于校验区块数据有效
	Nonce int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(),
		[]byte(data),
		prevBlockHash,
		[]byte{},
		0,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}


func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
