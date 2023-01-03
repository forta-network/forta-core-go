package registry

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
)

func equalsAddress(address common.Address, addr string) bool {
	return strings.EqualFold(address.Hex(), addr)
}

func getTopic(log types.Log) string {
	if len(log.Topics) == 0 {
		return ""
	}
	return strings.ToLower(log.Topics[0].Hex())
}

func getLoggerForLog(le types.Log) *log.Entry {
	return log.WithFields(log.Fields{
		"address":     strings.ToLower(le.Address.Hex()),
		"txHash":      le.TxHash.Hex(),
		"blockNumber": le.BlockNumber,
		"blockHash":   le.BlockHash.Hex(),
	})
}
