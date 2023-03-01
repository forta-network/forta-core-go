package utils

import (
	"bytes"
	"encoding/base64"
	"math/big"

	"github.com/bits-and-blooms/bloom"
	"github.com/forta-network/forta-core-go/protocol"
)

const (
	NumMaxAddressesPerAlert  = 50
	AddressBloomFilterFPRate = 1e-3
)

// CreateBloomFilter creates a new bloom filter using given items and false positive rate.
func CreateBloomFilter(items []string, fpRate float64) (*protocol.BloomFilter, error) {
	// create bloom filter from all addresses
	bf := bloom.NewWithEstimates(uint(len(items)), fpRate)
	for _, address := range items {
		bf.Add([]byte(address))
	}

	// extract bitset from bloom filter
	var b bytes.Buffer

	_, err := bf.WriteTo(&b)
	if err != nil {
		return nil, err
	}

	// create bloom filter
	bitset := base64.StdEncoding.EncodeToString(b.Bytes())

	kBigInt := new(big.Int).SetUint64(uint64(bf.K()))
	mBigInt := new(big.Int).SetUint64(uint64(bf.Cap()))

	kHexStr := BigIntToHex(kBigInt)
	mHexStr := BigIntToHex(mBigInt)

	return &protocol.BloomFilter{
		K:         kHexStr,
		M:         mHexStr,
		Bitset:    bitset,
		ItemCount: uint32(len(items)),
	}, nil
}

func CreateBloomFilterFromProto(bf *protocol.BloomFilter) (*bloom.BloomFilter, error) {
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

