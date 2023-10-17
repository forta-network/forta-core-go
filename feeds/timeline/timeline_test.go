package timeline

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/forta-network/forta-core-go/feeds"
	"github.com/stretchr/testify/require"
)

func TestTimeline_GlobalHighest(t *testing.T) {
	r := require.New(t)

	blockTimeline := &BlockTimeline{
		maxMinutes: 2,
	}

	min1 := time.Now().UTC().Truncate(time.Minute)
	min2 := min1.Add(time.Minute)
	min1Ts := hexutil.EncodeUint64(uint64(min1.Unix()))
	min2Ts := hexutil.EncodeUint64(uint64(min2.Unix()))
	min3Ts := hexutil.EncodeUint64(uint64(min1.Add(time.Minute * 2).Unix()))

	// add first minute block number
	blockTimeline.HandleBlock(blockForTimestamp(min1Ts, "0x1"))
	// replace the first minute number
	blockTimeline.HandleBlock(blockForTimestamp(min1Ts, "0x2"))
	// add new one for the next minute
	blockTimeline.HandleBlock(blockForTimestamp(min2Ts, "0x3"))
	// replace the second minute number
	blockTimeline.HandleBlock(blockForTimestamp(min2Ts, "0x4"))
	// replace the first minute number
	blockTimeline.HandleBlock(blockForTimestamp(min1Ts, "0x5"))
	// replace the second minute number
	blockTimeline.HandleBlock(blockForTimestamp(min2Ts, "0x6"))
	// add a third minute
	blockTimeline.HandleBlock(blockForTimestamp(min3Ts, "0x7"))

	// verify state
	r.EqualValues(5, blockTimeline.blockMinutes[0].HighestBlockNumber)
	r.EqualValues(6, blockTimeline.blockMinutes[1].HighestBlockNumber)
	r.EqualValues(7, blockTimeline.blockMinutes[2].HighestBlockNumber)

	// cleanup should remove the first minute because of the max minutes num
	blockTimeline.doCleanup()

	// verify state
	r.EqualValues(6, blockTimeline.blockMinutes[0].HighestBlockNumber)
	r.EqualValues(7, blockTimeline.blockMinutes[1].HighestBlockNumber)

	highestGlobal, ok := blockTimeline.GetGlobalHighest(min2)
	r.True(ok)
	r.Equal(uint64(6), highestGlobal)
}

func TestTimeline_CalculateLag(t *testing.T) {
	r := require.New(t)

	blockTimeline := &BlockTimeline{}

	currMin := time.Now().UTC().Truncate(time.Minute)
	blockTimeline.localMinutes = append(blockTimeline.localMinutes, &Minute{
		Timestamp:          currMin,
		HighestBlockNumber: 1,
	})
	blockTimeline.blockMinutes = append(blockTimeline.blockMinutes, &Minute{
		Timestamp:          currMin,
		HighestBlockNumber: 2,
	})
	lag, ok := blockTimeline.CalculateLag(currMin)
	r.True(ok)
	r.Equal(int64(1), lag)
}

func blockForTimestamp(ts, blockNumber string) *domain.BlockEvent {
	return &domain.BlockEvent{
		Block: &domain.Block{
			Timestamp: ts,
			Number:    blockNumber,
		},
	}
}

func TestRealTimeLag(t *testing.T) {
	if os.Getenv("TIMELINE_EXPERIMENT") != "1" {
		return
	}
	r := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	ethClient, err := ethereum.NewStreamEthClient(ctx, "", os.Getenv("JSON_RPC_API"))
	ethClient.SetRetryInterval(time.Second * 2)
	r.NoError(err)
	blockFeed, err := feeds.NewBlockFeed(ctx, ethClient, ethClient, feeds.BlockFeedConfig{
		ChainID:     big.NewInt(137),
		DisableLogs: true,
	})
	r.NoError(err)

	blockTimeline := &BlockTimeline{}
	errCh := blockFeed.Subscribe(func(evt *domain.BlockEvent) error {
		blockTs, _ := evt.Block.GetTimestamp()
		delay := time.Since(*blockTs)
		fmt.Println("delay:", delay)
		return blockTimeline.HandleBlock(evt)
	})

	go blockFeed.Start()

	go func() {
		<-time.After(time.Minute * 4)
		cancel()
	}()

	<-errCh

	for _, minute := range blockTimeline.blockMinutes {
		lag, ok := blockTimeline.CalculateLag(minute.Timestamp)
		if !ok {
			continue
		}
		fmt.Println(minute.Timestamp.Format(time.RFC3339), ":", lag)
	}
}
