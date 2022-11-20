package timeline

import (
	"testing"

	"github.com/forta-network/forta-core-go/domain"
	"github.com/stretchr/testify/require"
)

func TestTimeline(t *testing.T) {
	r := require.New(t)

	blockTimeline := &BlockTimeline{
		maxMinutes: 2,
	}

	// add first minute block number
	blockTimeline.HandleBlock(blockForTimestamp("0x1000000000", "0x1"))
	// replace the first minute number
	blockTimeline.HandleBlock(blockForTimestamp("0x1000000000", "0x2"))
	// add new one for the next minute
	blockTimeline.HandleBlock(blockForTimestamp("0x2000000000", "0x3"))
	// replace the second minute number
	blockTimeline.HandleBlock(blockForTimestamp("0x2000000000", "0x4"))
	// replace the first minute number
	blockTimeline.HandleBlock(blockForTimestamp("0x1000000000", "0x5"))
	// replace the second minute number
	blockTimeline.HandleBlock(blockForTimestamp("0x2000000000", "0x6"))
	// add a third minute
	blockTimeline.HandleBlock(blockForTimestamp("0x3000000000", "0x7"))

	// verify state
	r.EqualValues(5, blockTimeline.minutes[0].HighestBlockNumber)
	r.EqualValues(6, blockTimeline.minutes[1].HighestBlockNumber)
	r.EqualValues(7, blockTimeline.minutes[2].HighestBlockNumber)

	// cleanup should remove the first minute because of the max minutes num
	blockTimeline.doCleanup()

	// verify state
	r.EqualValues(6, blockTimeline.minutes[0].HighestBlockNumber)
	r.EqualValues(7, blockTimeline.minutes[1].HighestBlockNumber)
}

func blockForTimestamp(ts, blockNumber string) *domain.BlockEvent {
	return &domain.BlockEvent{
		Block: &domain.Block{
			Timestamp: ts,
			Number:    blockNumber,
		},
	}
}
