package registry

import (
	"time"

	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/utils"
	log "github.com/sirupsen/logrus"
)

type Source struct {
	TxHash             string    `json:"txHash"`
	BlockNumberDecimal int64     `json:"blockNumberDecimal"`
	BlockNumber        string    `json:"blockNumber"`
	BlockHash          string    `json:"blockHash"`
	Timestamp          time.Time `json:"timestamp"`
}

func SourceFromBlock(txHash string, blk *domain.Block) Source {
	if blk == nil {
		return Source{
			TxHash: txHash,
		}
	}

	ts := utils.HexToInt64(blk.Timestamp)
	return Source{
		BlockNumberDecimal: utils.HexToInt64(blk.Number),
		BlockNumber:        blk.Number,
		BlockHash:          blk.Hash,
		Timestamp:          time.Unix(ts, 0).UTC(),
		TxHash:             txHash,
	}
}

func MessageFrom(txHash string, blk *domain.Block, action string) Message {
	return Message{
		Timestamp: time.Now().UTC(),
		Action:    action,
		Source:    SourceFromBlock(txHash, blk),
	}
}

type Message struct {
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
	Source    Source    `json:"source"`
}

// ActionName implements the message interface.
func (m Message) ActionName() string {
	return m.Action
}

// Info implements the message interface.
func (m Message) Info() Message {
	return m
}

// MessageInterface implements a message interface.
type MessageInterface interface {
	Info() Message
	ActionName() string
	LogFields() log.Fields
}
