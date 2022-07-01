package utils

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func BigIntFromIntString(str string) *big.Int {
	res := big.NewInt(0)
	res.SetString(str, 10)
	return res
}

func BigIntToHex(i *big.Int) string {
	return hexutil.EncodeBig(i)
}

func ScannerIDBigIntToHex(i *big.Int) string {
	return strings.ToLower(common.BigToAddress(i).Hex())
}

func ScannerIDHexToBigInt(scannerID string) *big.Int {
	return common.HexToHash(scannerID).Big()
}

func AgentBigIntToHex(i *big.Int) string {
	return BytesToHex(i.Bytes())
}

func AgentHexToBigInt(hex string) *big.Int {
	return common.HexToHash(hex).Big()
}

func HexToBigInt(hex string) (*big.Int, error) {
	return hexutil.DecodeBig(hex)
}

func Bytes32ToHex(b [32]byte) string {
	return common.BytesToHash(b[:]).Hex()
}

func BytesToHex(b []byte) string {
	return common.BytesToHash(b).Hex()
}

func HexToInt64(hex string) int64 {
	bigInt, err := HexToBigInt(hex)
	if err != nil {
		log.WithField("hex", hex).Error("could not convert hex to number")
		return 0
	}
	return bigInt.Int64()
}
