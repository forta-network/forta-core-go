package registry

import (
	"errors"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/goccy/go-json"
	log "github.com/sirupsen/logrus"
	"time"
)

type Source struct {
	BlockNumber string    `json:"blockNumber"`
	BlockHash   string    `json:"blockHash"`
	Timestamp   time.Time `json:"timestamp"`
}

func SourceFromBlock(blk *domain.Block) Source {
	if blk == nil {
		return Source{}
	}
	ts := utils.HexToInt64(blk.Timestamp)
	return Source{
		BlockNumber: blk.Number,
		BlockHash:   blk.Hash,
		Timestamp:   time.Unix(ts, 0),
	}
}

type Message struct {
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
	Source    Source    `json:"source"`
}

func ParseMessage(msg string) (*Message, error) {
	var r Message
	err := json.Unmarshal([]byte(msg), &r)
	if err != nil {
		return nil, err
	}
	if r.Action == "" {
		log.WithFields(log.Fields{
			"body": msg,
		}).Error("action is not populated")
		return nil, errors.New("action is not populated")
	}
	return &r, nil
}
