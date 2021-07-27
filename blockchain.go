package main

type Blockchain struct {
    blocks  []*Block
}

func NewBlockchain() *Blockchain {
    return &Blockchain{blocks:[]*Block{NewGenesisBlock()}}
}

func (bc *Blockchain) Addblock(data string) {
	PrevBlock := bc.blocks[len(bc.blocks) -1]
	newBlcok := NewBlock(data, PrevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *block {
	return NewBlock(data:"Genesis Block", []byte{})
}