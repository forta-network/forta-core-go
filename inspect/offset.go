package inspect

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/montanaflynn/stats"
	"golang.org/x/sync/errgroup"
)

type offsetStats struct {
	Mean        float64
	Median      float64
	Max         float64
	SampleCount float64
}

func calculateOffsetStats(
	ctx context.Context, primaryClient *ethclient.Client,
	secondaryClient *ethclient.Client,
) (offsetStats, error) {
	ds, err := collectOffsetData(ctx, primaryClient, secondaryClient)
	if err != nil {
		return offsetStats{}, err
	}

	return extractStats(ds)
}

// collectOffsetData measures how long does it take to receive a recently created block and compares given eth clients.
// The idea is to mimic the behavior of Scanner feed and Bot proxy query.
func collectOffsetData(ctx context.Context, primaryClient *ethclient.Client, secondaryClient *ethclient.Client) (
	[]float64, error,
) {
	maxDuration := time.Second * 20

	var dataPoints []float64
	t := time.NewTicker(time.Millisecond * 300)

	// initialize data

	// get the next block in line, so that we can observe exactly how much delay both APIs have after
	// the block creation
	blockToQuery, err := primaryClient.BlockNumber(ctx)
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		return nil, err
	}

	blockToQuery++

	for {
		select {
		case <-ctx.Done():
			return dataPoints, nil
		case <-t.C:
			g, ctx := errgroup.WithContext(ctx)

			var (
				primaryAPIDelay, secondaryAPIDelay int64
			)

			// fetch latest block for primary api
			g.Go(
				func() (err error) {
					c, cancel := context.WithTimeout(ctx, maxDuration)
					defer cancel()
					primaryAPIDelay, err = measureBlockDelay(c, primaryClient, blockToQuery)
					return
				},
			)

			// fetch latest block for secondary api
			g.Go(
				func() (err error) {
					c, cancel := context.WithTimeout(ctx, maxDuration)
					defer cancel()

					secondaryAPIDelay, err = measureBlockDelay(c, secondaryClient, blockToQuery)
					return
				},
			)

			err = g.Wait()
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					return dataPoints, nil
				}
				return dataPoints, err
			}

			var offset = float64(secondaryAPIDelay) - float64(primaryAPIDelay)

			dataPoints = append(dataPoints, offset)
			blockToQuery++
		}
	}
}
func measureBlockDelay(ctx context.Context, client *ethclient.Client, blockNum uint64) (int64, error) {
	t := time.Millisecond * 200

	start := time.Now()

	for {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			_, err := client.BlockByNumber(ctx, big.NewInt(int64(blockNum)))
			if err == nil {
				goto exit
			}
			time.Sleep(t)
		}
	}

exit:
	return time.Since(start).Milliseconds(), nil
}
func extractStats(ds []float64) (offsetStats, error) {
	os := offsetStats{}
	var err error

	os.Mean, err = stats.Mean(ds)
	if err != nil {
		return os, err
	}

	os.Median, err = stats.Median(ds)
	if err != nil {
		return os, err
	}

	os.Max, err = stats.Max(ds)
	if err != nil {
		return os, err
	}

	os.SampleCount = float64(len(ds))

	return os, err
}
