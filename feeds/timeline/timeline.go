package timeline

import (
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/forta-network/forta-core-go/domain"
	log "github.com/sirupsen/logrus"
)

// BlockTimeline implements a block feed subscriber and keeps track of the
// latest block in every minute.
type BlockTimeline struct {
	maxMinutes int
	minutes    []*Minute
	mu         sync.RWMutex
}

// Minute represents a minute in a chain.
type Minute struct {
	HighestBlockNumber uint64
	Timestamp          time.Time
}

// NewBlockTimeline creates a new block timeline.
func NewBlockTimeline(maxMinutes int) *BlockTimeline {
	bt := &BlockTimeline{
		maxMinutes: maxMinutes,
	}

	go bt.cleanup()

	return bt
}

func (bt *BlockTimeline) cleanup() {
	ticker := time.NewTicker(time.Minute)
	for {
		<-ticker.C
		bt.doCleanup()
	}
}

func (bt *BlockTimeline) doCleanup() {
	bt.mu.Lock()
	defer bt.mu.Unlock()

	currSize := len(bt.minutes)
	if currSize > bt.maxMinutes {
		extra := currSize - bt.maxMinutes
		bt.minutes = bt.minutes[extra:] // release oldest
	}
}

// HandleBlock handles a block incoming from block feed.
func (bt *BlockTimeline) HandleBlock(evt *domain.BlockEvent) error {
	bt.mu.Lock()
	defer bt.mu.Unlock()

	blockTs, err := evt.Block.GetTimestamp()
	if err != nil {
		log.WithError(err).Error("failed to get block timestamp")
		return nil
	}
	blockMinuteTs := blockTs.Truncate(time.Minute)
	blockNum, err := hexutil.DecodeUint64(evt.Block.Number)
	if err != nil {
		log.WithError(err).Error("failed to decode block number")
	}
	for _, minute := range bt.minutes {
		if minute.Timestamp.Equal(blockMinuteTs) {
			if blockNum > minute.HighestBlockNumber {
				minute.HighestBlockNumber = blockNum
			}
			return nil // we found the minute
		}
	}
	// could not find the minute - append it
	bt.minutes = append(bt.minutes, &Minute{
		HighestBlockNumber: blockNum,
		Timestamp:          blockMinuteTs,
	})
	return nil
}

// GetHighestBlockNumber returns the highest block number within the same minute of the given timestamp.
func (bt *BlockTimeline) GetHighestBlockNumber(ts time.Time) (uint64, bool) {
	bt.mu.RLock()
	defer bt.mu.RUnlock()

	ts = ts.Truncate(time.Minute)
	for _, minute := range bt.minutes {
		if minute.Timestamp.Equal(ts) {
			return minute.HighestBlockNumber, true
		}
	}
	return 0, false
}
