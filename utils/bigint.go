package utils

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func Hex(i *big.Int) string {
	return common.BytesToHash(i.Bytes()).Hex()
}

func HexAddr(i *big.Int) string {
	return common.BytesToAddress(i.Bytes()).Hex()
}

func IntArray(i []*big.Int) []int64 {
	var res []int64
	for _, val := range i {
		res = append(res, val.Int64())
	}
	return res
}
