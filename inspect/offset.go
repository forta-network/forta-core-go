package inspect

import (
	"context"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/montanaflynn/stats"
	"golang.org/x/sync/errgroup"
)

type offsetStats struct {
	Mean   float64
	Median float64
	Max    float64
}

func calculateOffsetStats(
	ctx context.Context, primaryClient *ethclient.Client,
	secondaryClient *ethclient.Client,
) (offsetStats, error) {
	ds, err := collectOffsetData(ctx, primaryClient, secondaryClient)
	if err != nil {
		return offsetStats{}, err
	}
	return extractStats(ds), nil
}

func collectOffsetData(ctx context.Context, primaryClient *ethclient.Client, secondaryClient *ethclient.Client) (
	[]float64, error,
) {
	var dataPoints []float64
	t := time.NewTicker(time.Millisecond * 300)
	for {
		select {
		case <-ctx.Done():
			return dataPoints, nil
		case <-t.C:
			g, ctx := errgroup.WithContext(ctx)

			var (
				primaryAPIBlock, secondaryAPIBlock uint64
			)

			// fetch latest block for primary api
			g.Go(
				func() (err error) {
					primaryAPIBlock, err = primaryClient.BlockNumber(ctx)
					if errors.Is(err, context.DeadlineExceeded) {
						return nil
					}
					return
				},
			)

			// fetch latest block for secondary api
			g.Go(
				func() (err error) {
					secondaryAPIBlock, err = secondaryClient.BlockNumber(ctx)
					if errors.Is(err, context.DeadlineExceeded) {
						return nil
					}
					return
				},
			)

			err := g.Wait()
			if err != nil && err != context.DeadlineExceeded {
				return dataPoints, err
			}

			var offset = float64(primaryAPIBlock) - float64(secondaryAPIBlock)
			dataPoints = append(dataPoints, offset)
		}
	}
}

func extractStats(ds []float64) offsetStats {
	os := offsetStats{}
	os.Mean, _ = stats.Mean(ds)
	os.Median, _ = stats.Median(ds)
	os.Max, _ = stats.Max(ds)

	return os
}