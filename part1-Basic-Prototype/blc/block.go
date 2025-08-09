package blc

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	//1.区块高度
	//Height int64
	//2.上一个区块的HASH
	PrevHash []byte
	//3.交易数据
	Data []byte
	//4.时间戳
	Timestamp int64
	//5.Hash
	Hash []byte
}

func (block Block) setHash() {
	timeStr := strconv.FormatInt(block.Timestamp, 2)
	timeByte := []byte(timeStr)
	headers := bytes.Join([][]byte{block.PrevHash, timeByte, block.Data, block.Hash}, []byte{})
	hash := sha256.Sum256(headers)
	block.Hash = hash[:]
}

func newBlock(data string, prevBlockHash []byte) *Block {
	block := Block{prevBlockHash, []byte(data), time.Now().Unix(), []byte{}}
	block.setHash()
	return &block
}

func newGenesisBlock() *Block {
	return newBlock("Genesis Block", []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
}
