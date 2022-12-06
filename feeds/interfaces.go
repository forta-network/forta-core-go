package feeds

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/forta-network/forta-core-go/clients/health"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/protocol"
)

// BlockFeed is a subscribable feed of blocks.
type BlockFeed interface {
	Start()
	StartRange(start int64, end int64, rate int64)
	IsStarted() bool
	Subscribe(handler func(evt *domain.BlockEvent) error) <-chan error
	health.Reporter
}

// TransactionFeed is a subscribable feed of transactions.
type TransactionFeed interface {
	ForEachTransaction(blockHandler func(evt *domain.BlockEvent) error, txHandler func(evt *domain.TransactionEvent) error) error
}

// AlertFeed is a subscribable feed of alerts.
type AlertFeed interface {
	Start()
	StartRange(start uint64, end uint64, rate int64)
	AddSubscription(subscription *protocol.CombinerBotSubscription)
	RemoveSubscription(subscription *protocol.CombinerBotSubscription)
	Subscriptions() []*protocol.CombinerBotSubscription
	ForEachAlert(alertHandler func(evt *domain.AlertEvent) error) error
	RegisterHandler(alertHandler func(evt *domain.AlertEvent) error) <-chan error
	health.Reporter
}

// LogFeed is a feed of logs
type LogFeed interface {
	ForEachLog(handler func(blk *domain.Block, logEntry types.Log) error, finishBlockHandler func(blk *domain.Block) error) error
	GetLogsForLastBlocks(blocksAgo int64) ([]types.Log, error)
	GetLogsForRange(blockStart *big.Int, blockEnd *big.Int) ([]types.Log, error)
}
