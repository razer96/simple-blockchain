package blockchain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	a := assert.New(t)

	testData := "test data"
	bc := NewBlockchain()

	bc.AddBlock(testData)

	a.Equal("Genesis block", string(bc.blocks[0].Data))
	a.Equal(bc.blocks[0].Hash, bc.blocks[1].PrevBlockHash)
}

func TestAddBlock(t *testing.T) {
	a := assert.New(t)

	testData := "test data"
	bc := NewBlockchain()

	bc.AddBlock(testData)

	a.Equal(2, len(bc.blocks))
}