package regmsg

import (
	"time"

	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/utils"
	log "github.com/sirupsen/logrus"
)

// Source contains message source info.
type Source struct {
	TxHash             string    `json:"txHash"`
	BlockNumberDecimal int64     `json:"blockNumberDecimal"`
	BlockNumber        string    `json:"blockNumber"`
	BlockHash          string    `json:"blockHash"`
	Timestamp          time.Time `json:"timestamp"`
}

// SourceFromBlock creates message source from given block.
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

// From creates a message from given inputs.
func From(txHash string, blk *domain.Block, action string) Message {
	return Message{
		Timestamp: time.Now().UTC(),
		Action:    action,
		Source:    SourceFromBlock(txHash, blk),
	}
}

// Message contains base message info. In practice, the content works as a common message
// header for all message types.
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

// Interface implements a message interface. One needs to wrap the Message type
// and implement extra method(s) to satisfy this interface.
type Interface interface {
	Info() Message
	ActionName() string
	LogFields() log.Fields
}

// HandlerFunc is a handler func type.
type HandlerFunc[I Interface] func(logger *log.Entry, msg I) error

// HandlerFuncs is an alias for the array of handler funcs.
type HandlerFuncs[I Interface] []HandlerFunc[I]

// Handlers conveniently stacks up given handler funcs into an array.
func Handlers[I Interface](args ...HandlerFunc[I]) HandlerFuncs[I] {
	return args
}

// HandlerInterface is an interface definition alternative to the handler func.
type HandlerInterface[I Interface] interface {
	Handle(logger *log.Entry, msg I) error
}
