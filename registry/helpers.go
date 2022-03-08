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

func isEvent(le types.Log, topic string) bool {
	if len(le.Topics) == 0 {
		return false
	}
	return strings.EqualFold(le.Topics[0].Hex(), topic)
}

func getLoggerForLog(le types.Log) *log.Entry {
	return log.WithFields(log.Fields{
		"address":     strings.ToLower(le.Address.Hex()),
		"txHash":      le.TxHash.Hex(),
		"blockNumber": le.BlockNumber,
		"blockHash":   le.BlockHash.Hex(),
	})
}
