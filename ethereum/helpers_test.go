package ethereum

import (
	"testing"

	"github.com/ethereum/go-ethereum"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/stretchr/testify/assert"
)

func TestToFilterArg(t *testing.T) {
	fromBlock, _ := utils.HexToBigInt("0x25beafb")
	toBlock := fromBlock

	q := ethereum.FilterQuery{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Addresses: nil,
		Topics:    nil,
	}

	result, err := toFilterArg(q)
	assert.NoError(t, err)

	m, ok := result.(map[string]interface{})
	assert.True(t, ok)

	_, addrExists := m["addresses"]
	assert.False(t, addrExists)

	_, topicsExists := m["topics"]
	assert.False(t, topicsExists)
}
