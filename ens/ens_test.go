package ens

import (
	"errors"
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
