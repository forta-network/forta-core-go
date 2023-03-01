package ens

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"github.com/wealdtech/go-ens/v3"
)

func TestENS(t *testing.T) {
	r := require.New(t)

	testName := "test.name"
	knownAddr := common.HexToAddress("01")
	unknownAddr := ens.UnknownAddress

	ensStore := NewENStoreWithResolver(ResolverFunc(func(input string) (common.Address, error) {
		if input == testName {
			return knownAddr, nil
		}
		return unknownAddr, errors.New("unknown address")
	}))

	addr, err := ensStore.Resolve(testName)
	r.NoError(err)
	r.Equal(knownAddr, addr)

	addr, err = ensStore.Resolve("bad name")
	r.Error(err)
	r.Equal(unknownAddr, addr)
}

func TestThing(t *testing.T) {
	e, err := DialENSStoreAt("https://polygon-rpc.com", "0x08f42fcc52a9C2F391bF507C4E8688D0b53e1bd7")

	assert.NoError(t, err)
	c, err := e.ResolveRegistryContracts()
	assert.NoError(t, err)

	t.Log(c.RewardsDistributor.Hex())
}
