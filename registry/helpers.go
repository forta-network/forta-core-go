package registry

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"strings"
)

func equalsAddress(address common.Address, addr string) bool {
	return strings.ToLower(address.Hex()) == strings.ToLower(addr)
}

func isEvent(le types.Log, topic string) bool {
	if len(le.Topics) == 0 {
		return false
	}
	return strings.ToLower(le.Topics[0].Hex()) == strings.ToLower(topic)
}

func getLoggerForLog(le types.Log) *log.Entry {
	return log.WithFields(log.Fields{
		"address":     strings.ToLower(le.Address.Hex()),
		"txHash":      le.TxHash.Hex(),
		"blockNumber": le.BlockNumber,
		"blockHash":   le.BlockHash.Hex(),
	})
}
