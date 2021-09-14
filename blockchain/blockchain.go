package blockchain

import (
	"github.com/razer96/simple-blockchain/block"
)

type Blockchain struct {
	blocks []*block.Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := block.NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func (bc *Blockchain) Blocks() []*block.Block {
	return bc.blocks
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		blocks: []*block.Block{block.NewGenesisBlock()},
	}
}
