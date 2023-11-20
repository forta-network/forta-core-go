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
	r.NoError(blockTimeline.HandleBlock(blockForTimestamp(min1Ts, "0x1")))
	// replace the first minute number
	r.NoError(blockTimeline.HandleBlock(blockForTimestamp(min1Ts, "0x2")))
	// add new one for the next minute
	r.NoError(blockTimeline.HandleBlock(blockForTimestamp(min2Ts, "0x3")))
	// replace the second minute number
	r.NoError(blockTimeline.HandleBlock(blockForTimestamp(min2Ts, "0x4")))
	// replace the first minute number
	r.NoError(blockTimeline.HandleBlock(blockForTimestamp(min1Ts, "0x5")))
	// replace the second minute number
	r.NoError(blockTimeline.HandleBlock(blockForTimestamp(min2Ts, "0x6")))
	// add a third minute
	r.NoError(blockTimeline.HandleBlock(blockForTimestamp(min3Ts, "0x7")))

	// verify state
	r.EqualValues(5, blockTimeline.chainMinutes[0].HighestBlockNumber)
	r.EqualValues(6, blockTimeline.chainMinutes[1].HighestBlockNumber)
	r.EqualValues(7, blockTimeline.chainMinutes[2].HighestBlockNumber)

	// cleanup should remove the first minute because of the max minutes num
	blockTimeline.doCleanup()

	// verify state
	r.EqualValues(6, blockTimeline.chainMinutes[0].HighestBlockNumber)
	r.EqualValues(7, blockTimeline.chainMinutes[1].HighestBlockNumber)

	highestGlobal, ok := blockTimeline.getLatestUpTo(blockTimeline.chainMinutes, min2)
	r.True(ok)
	r.Equal(uint64(6), highestGlobal)
}

func TestTimeline_CalculateLag(t *testing.T) {
	r := require.New(t)

	blockTimeline := NewBlockTimeline(1, 1000000)
	blockTimeline.threshold = 10

	start := time.Now().UTC().Truncate(time.Minute)

	// minute: 1
	// local:  1
	// chain:  2
	// lag:    1
	min1 := start
	blockTimeline.localMinutes = append(blockTimeline.localMinutes, &Minute{
		Timestamp:          min1,
		HighestBlockNumber: 1,
	})
	blockTimeline.chainMinutes = append(blockTimeline.chainMinutes, &Minute{
		Timestamp:          min1,
		HighestBlockNumber: 2,
	})
	_, ok := blockTimeline.CalculateLag()
	r.False(ok)

	// minute: 2
	// local:  3
	// chain:  8
	// lag:    5
	min2 := min1.Add(time.Minute)
	blockTimeline.localMinutes = append(blockTimeline.localMinutes, &Minute{
		Timestamp:          min2,
		HighestBlockNumber: 3,
	})
	blockTimeline.chainMinutes = append(blockTimeline.chainMinutes, &Minute{
		Timestamp:          min2,
		HighestBlockNumber: 8,
	})
	lag, ok := blockTimeline.CalculateLag()
	r.True(ok)
	r.Equal(float64(1), lag) // because excludes the last minute: (2-1)/1

	// lags for a while: these minutes have block minutes but no local processing minutes
	// minute: 3
	// chain:  12
	// lag:    9 <-- using previous local (12 - 3)
	blockTimeline.chainMinutes = append(blockTimeline.chainMinutes, &Minute{
		Timestamp:          min1.Add(time.Minute * 2),
		HighestBlockNumber: 12,
	})
	// minute: 3
	// chain:  16
	// lag:    13 <-- using previous local (16 - 3)
	blockTimeline.chainMinutes = append(blockTimeline.chainMinutes, &Minute{
		Timestamp:          min1.Add(time.Minute * 3),
		HighestBlockNumber: 16,
	})
	// minute: 4
	// chain:  18
	// lag:    15 <-- using previous local (18 - 3)
	blockTimeline.chainMinutes = append(blockTimeline.chainMinutes, &Minute{
		Timestamp:          min1.Add(time.Minute * 4),
		HighestBlockNumber: 18,
	})

	// catches up in this minute - processes up to 20
	// minute: 6
	// local:  20
	// chain:  22
	// lag:    2
	min6 := min1.Add(time.Minute * 5)
	blockTimeline.localMinutes = append(blockTimeline.localMinutes, &Minute{
		Timestamp:          min6,
		HighestBlockNumber: 20,
	})
	blockTimeline.chainMinutes = append(blockTimeline.chainMinutes, &Minute{
		Timestamp:          min6,
		HighestBlockNumber: 22,
	})

	// this last minute doesn't matter as last minutes are excluded from calculation
	// minute: 7
	// local:  22
	// chain:  26
	// lag:    4 <-- shouldn't matter
	min7 := min1.Add(time.Minute * 6)
	blockTimeline.localMinutes = append(blockTimeline.localMinutes, &Minute{
		Timestamp:          min7,
		HighestBlockNumber: 22,
	})
	blockTimeline.chainMinutes = append(blockTimeline.chainMinutes, &Minute{
		Timestamp:          min7,
		HighestBlockNumber: 26,
	})

	// we are iterating by the block minutes during the calculation so
	// this local minute doesn't matter
	// minute: 8
	// local:  24
	// lag:    2 <-- shouldn't matter, using previous local (24-22)
	blockTimeline.localMinutes = append(blockTimeline.localMinutes, &Minute{
		Timestamp:          min1.Add(time.Minute * 7),
		HighestBlockNumber: 24,
	})

	// the final calculation
	lag, ok = blockTimeline.CalculateLag()
	r.True(ok)
	r.Equal(float64(1+5+9+13+15+2)/float64(6), lag)
	estimate, ok := blockTimeline.EstimateBlockScore()
	r.True(ok)
	r.Equal(0.25, estimate)

	testDelay := time.Second
	blockTimeline.delay = &testDelay
	delay, ok := blockTimeline.GetDelay()
	r.True(ok)
	r.Equal(testDelay, delay)
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
		t.Skip("skipping timeline experiment")
	}

	//logrus.SetLevel(logrus.ErrorLevel)

	r := require.New(t)
	ctx, cancel := context.WithCancel(context.Background())
	ethClient, err := ethereum.NewStreamEthClient(ctx, "", os.Getenv("JSON_RPC_API"))
	ethClient.SetRetryInterval(time.Second * 2)
	r.NoError(err)
	blockFeed, err := feeds.NewBlockFeed(ctx, ethClient, ethClient, feeds.BlockFeedConfig{
		ChainID: big.NewInt(137),
		//DisableLogs: true,
	})
	r.NoError(err)

	blockTimeline := &BlockTimeline{}
	errCh := blockFeed.Subscribe(func(evt *domain.BlockEvent) error {
		// blockTs, _ := evt.Block.GetTimestamp()
		// delay := time.Since(*blockTs)
		// fmt.Println("delay:", delay)
		return blockTimeline.HandleBlock(evt)
	})

	go blockFeed.Start()

	go func() {
		longTicker := time.After(time.Minute * 10)
		shortTicker := time.NewTicker(time.Minute).C
		for {
			select {
			case <-longTicker:
				cancel()
				return
			case <-shortTicker:
				lag, ok := blockTimeline.CalculateLag()
				if !ok {
					continue
				}
				fmt.Println("lag at", time.Now().Truncate(time.Minute).Format(time.RFC3339), ":", lag, "blocks")
			}
		}
	}()

	<-errCh
}
