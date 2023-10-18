package timeline

import (
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/protocol/settings"
	log "github.com/sirupsen/logrus"
)

// BlockTimeline implements a block feed subscriber and keeps track of the
// latest block in every minute.
type BlockTimeline struct {
	threshold    int
	maxMinutes   int
	chainMinutes []*Minute // when the block was produced
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
func NewBlockTimeline(chainID, maxMinutes int) *BlockTimeline {
	bt := &BlockTimeline{
		threshold:  settings.GetChainSettings(chainID).BlockThreshold,
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

	currSize := len(bt.chainMinutes)
	if currSize > bt.maxMinutes {
		extra := currSize - bt.maxMinutes
		bt.chainMinutes = bt.chainMinutes[extra:] // release oldest
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
	for _, minute := range bt.chainMinutes {
		if minute.Timestamp.Equal(blockMinuteTs) {
			if blockNum > minute.HighestBlockNumber {
				minute.HighestBlockNumber = blockNum
			}
			foundBlockMinute = true
			break
		}
	}
	if !foundBlockMinute {
		bt.chainMinutes = append(bt.chainMinutes, &Minute{
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

	sort.Slice(bt.chainMinutes, func(i, j int) bool {
		return bt.chainMinutes[i].Timestamp.Before(bt.chainMinutes[j].Timestamp)
	})
	sort.Slice(bt.localMinutes, func(i, j int) bool {
		return bt.localMinutes[i].Timestamp.Before(bt.localMinutes[j].Timestamp)
	})

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

func (bt *BlockTimeline) getHighest(minutes []*Minute, ts time.Time) (uint64, bool) {
	ts = ts.Truncate(time.Minute)
	for _, minute := range minutes {
		if minute.Timestamp.Equal(ts) {
			return minute.HighestBlockNumber, true
		}
	}
	return 0, false
}

func (bt *BlockTimeline) getLatestUpTo(minutes []*Minute, ts time.Time) (uint64, bool) {
	ts = ts.Truncate(time.Minute)
	var foundMinute *Minute
	for _, minute := range minutes {
		if minute.Timestamp.After(ts) {
			break
		}
		foundMinute = minute
	}
	if foundMinute != nil {
		return foundMinute.HighestBlockNumber, true
	}
	return 0, false
}

// CalculateLag calculates the block number lag by taking the average of each minute.
func (bt *BlockTimeline) CalculateLag() (float64, bool) {
	bt.mu.RLock()
	defer bt.mu.RUnlock()

	var (
		total float64
		count float64
	)
	for i, chainMinute := range bt.chainMinutes {
		// exclude the last minute
		if i == len(bt.chainMinutes)-1 {
			break
		}
		// avoid calculation if we can't find a highest
		localMinuteHighest, ok := bt.getLatestUpTo(bt.localMinutes, chainMinute.Timestamp)
		if !ok {
			continue
		}
		total += float64(chainMinute.HighestBlockNumber - localMinuteHighest)
		count++
	}
	if count == 0 {
		return 0, false
	}
	return total / count, true
}

// EstimateBlockScore estimates the block score based on the lag and the block threshold.
func (bt *BlockTimeline) EstimateBlockScore() (float64, bool) {
	lag, ok := bt.CalculateLag()
	if !ok {
		return 0, false
	}
	estimate := (float64(bt.threshold) - float64(lag)) / float64(bt.threshold)
	if estimate < 0 {
		estimate = 0
	}
	if estimate > 1 {
		estimate = 1
	}
	return estimate, true
}
