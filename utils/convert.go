package utils

import (
	"encoding/hex"
	"math/big"
	"strconv"
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

func PoolIDToBigInt(s string) *big.Int {
	i, _ := strconv.ParseUint(s, 10, 64)
	return big.NewInt(0).SetUint64(i)
}

func PoolIDToString(i *big.Int) string {
	return i.String()
}

func PoolIDHexToBigInt(hex string) *big.Int {
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

// IsValidBotID tells if given bot ID is valid.
func IsValidBotID(botID string) bool {
	if len(botID) < 2 || botID[:2] != "0x" {
		return false
	}

	b, err := hex.DecodeString(botID[2:])
	return len(b) == 32 && err == nil
}
