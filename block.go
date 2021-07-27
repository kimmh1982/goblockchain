package main

import "time" 
import "bytes" 
import "cryto/sha256"
import "strconv"

type Block struct {
	Timestamp	int64
	Data	[]byte
	PrevBlockHash	[]byte
	Hash	[]byte
}

func NewBlock(data string, PrevBlockHash []byte) *Block {

	block := &Block{Timestamp: time.Now().Unix(), Data:[]byte(data), PrevBlockHash: PrevBlockHash, Hash:[]byte{}}
	block.SetHash()

	return block

}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, base:10))
	headers := bytes.Join([][]byte{b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}


