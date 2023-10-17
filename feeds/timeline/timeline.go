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
	maxMinutes   int
	blockMinutes []*Minute // when the block was produced
	localMinutes []*Minute // when we handled the block
	delay        *time.Duration
	mu           sync.RWMutex
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

	currSize := len(bt.blockMinutes)
	if currSize > bt.maxMinutes {
		extra := currSize - bt.maxMinutes
		bt.blockMinutes = bt.blockMinutes[extra:] // release oldest
	}

	currSize = len(bt.localMinutes)
	if currSize > bt.maxMinutes {
		extra := currSize - bt.maxMinutes
		bt.localMinutes = bt.localMinutes[extra:] // release oldest
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
	delay := time.Since(*blockTs)
	bt.delay = &delay

	localMinuteTs := time.Now().Truncate(time.Minute)

	blockMinuteTs := blockTs.Truncate(time.Minute)
	blockNum, err := hexutil.DecodeUint64(evt.Block.Number)
	if err != nil {
		log.WithError(err).Error("failed to decode block number")
	}

	var foundBlockMinute bool
	for _, minute := range bt.blockMinutes {
		if minute.Timestamp.Equal(blockMinuteTs) {
			if blockNum > minute.HighestBlockNumber {
				minute.HighestBlockNumber = blockNum
			}
			foundBlockMinute = true
			break
		}
	}
	if !foundBlockMinute {
		bt.blockMinutes = append(bt.blockMinutes, &Minute{
			HighestBlockNumber: blockNum,
			Timestamp:          blockMinuteTs,
		})
	}

	var foundLocalMinute bool
	for _, minute := range bt.localMinutes {
		if minute.Timestamp.Equal(localMinuteTs) {
			if blockNum > minute.HighestBlockNumber {
				minute.HighestBlockNumber = blockNum
			}
			foundLocalMinute = true
			break
		}
	}
	if !foundLocalMinute {
		bt.localMinutes = append(bt.localMinutes, &Minute{
			HighestBlockNumber: blockNum,
			Timestamp:          localMinuteTs,
		})
	}

	return nil
}

func (bt *BlockTimeline) GetDelay() (time.Duration, bool) {
	bt.mu.RLock()
	defer bt.mu.RUnlock()

	if bt.delay == nil {
		return 0, false
	}
	return *bt.delay, true
}

// GetGlobalHighest returns the global highest block number within the same minute of the given timestamp.
func (bt *BlockTimeline) GetGlobalHighest(ts time.Time) (uint64, bool) {
	return bt.getHighest(bt.blockMinutes, ts)
}

// GetLocalHighest returns the local highest block number within the same minute of the given timestamp.
func (bt *BlockTimeline) GetLocalHighest(ts time.Time) (uint64, bool) {
	return bt.getHighest(bt.localMinutes, ts)
}

func (bt *BlockTimeline) getHighest(minutes []*Minute, ts time.Time) (uint64, bool) {
	bt.mu.RLock()
	defer bt.mu.RUnlock()

	ts = ts.Truncate(time.Minute)
	for _, minute := range minutes {
		if minute.Timestamp.Equal(ts) {
			return minute.HighestBlockNumber, true
		}
	}
	return 0, false
}

// CalculateLag calculates the block number lag in a given minute.
func (bt *BlockTimeline) CalculateLag(ts time.Time) (int64, bool) {
	highestGlobal, ok := bt.GetGlobalHighest(ts)
	if !ok {
		return 0, false
	}
	highestLocal, ok := bt.GetLocalHighest(ts)
	if !ok {
		return 0, false
	}
	return int64(highestGlobal) - int64(highestLocal), true
}

// Size returns the minute count.
func (bt *BlockTimeline) Size() int {
	bt.mu.RLock()
	defer bt.mu.RUnlock()

	return len(bt.blockMinutes)
}
