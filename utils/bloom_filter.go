package utils

import (
	"bytes"
	"encoding/base64"

	"github.com/bits-and-blooms/bloom"
	"github.com/forta-network/forta-core-go/protocol"
)

const (
	NumMaxAddressesPerAlert = 50
)

func RecreateBloomFilter(bf *protocol.BloomFilter) (*bloom.BloomFilter, error) {
	b, err := base64.StdEncoding.DecodeString(bf.Bitset)
	if err != nil {
		return nil, err
	}

	m, err := HexToBigInt(bf.M)
	if err != nil {
		return nil, err
	}

	k, err := HexToBigInt(bf.K)
	if err != nil {
		return nil, err
	}

	bFilter := bloom.New(uint(m.Uint64()), uint(k.Uint64()))
	if _, err := bFilter.ReadFrom(bytes.NewBuffer(b)); err != nil {
		return nil, err
	}

	return bFilter, nil
}

