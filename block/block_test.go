package block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBlock(t *testing.T) {
	a := assert.New(t)
	testData := "my transaction"
	prevBlockHash :=  []byte("previous transaction")
	block := NewBlock(testData, prevBlockHash)

	a.NotNil(block)
	a.Equal(testData, string(block.Data))
	a.Equal(prevBlockHash, block.PrevBlockHash)
}