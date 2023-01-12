package utils

import (
	"context"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/ethereum"
	"math/big"
	"time"
)

// GetClosestBlock Obtains the block closest to the time given
func GetClosestBlock(ctx context.Context, eth ethereum.Client, activeTime time.Time) (*domain.Block, error) {

	minBlockNumber := big.NewInt(0)
	maxBlockNumber, err := eth.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	closestBlockNumber := new(big.Int).Add(maxBlockNumber, minBlockNumber)
	closestBlockNumber = closestBlockNumber.Div(closestBlockNumber, big.NewInt(2))
	closestBlock, err := eth.BlockByNumber(ctx, closestBlockNumber)
	if err != nil {
		return nil, err
	}

	for minBlockNumber.Cmp(maxBlockNumber) <= 0 {
		closestBlockTime, err := closestBlock.GetTimestamp()
		if err != nil {
			return nil, err
		}

		if closestBlockTime.Equal(activeTime) {
			break
		} else if closestBlockTime.After(activeTime) {
			maxBlockNumber = maxBlockNumber.Sub(closestBlockNumber, big.NewInt(1))
		} else {
			minBlockNumber = minBlockNumber.Add(closestBlockNumber, big.NewInt(1))
		}

		closestBlockNumber = new(big.Int).Add(maxBlockNumber, minBlockNumber)
		closestBlockNumber = closestBlockNumber.Div(closestBlockNumber, big.NewInt(2))
		closestBlock, err = eth.BlockByNumber(ctx, closestBlockNumber)
		if err != nil {
			return nil, err
		}
	}

	return closestBlock, nil
}

// GetClosestBlockBefore Obtains the closest block only before the current active time
func GetClosestBlockBefore(ctx context.Context, eth ethereum.Client, activeTime time.Time) (*domain.Block, error) {
	closestBlock, err := GetClosestBlock(ctx, eth, activeTime)
	if err != nil {
		return nil, err
	}

	closestBlockTime, err := closestBlock.GetTimestamp()
	if err != nil {
		return nil, err
	}

	if closestBlockTime.Equal(activeTime) || closestBlockTime.Before(activeTime) {
		return closestBlock, nil
	}
	closestBlockNumber, err := HexToBigInt(closestBlock.Number)
	if err != nil {
		return nil, err
	}
	closestBlockNumberBefore := big.NewInt(0).Sub(closestBlockNumber, big.NewInt(1))
	return eth.BlockByNumber(ctx, closestBlockNumberBefore)
}

// GetClosestBlockAfter Obtains the closest block only after the current active time
func GetClosestBlockAfter(ctx context.Context, eth ethereum.Client, activeTime time.Time) (*domain.Block, error) {
	closestBlock, err := GetClosestBlock(ctx, eth, activeTime)
	if err != nil {
		return nil, err
	}

	closestBlockTime, err := closestBlock.GetTimestamp()
	if err != nil {
		return nil, err
	}

	if closestBlockTime.Equal(activeTime) || closestBlockTime.After(activeTime) {
		return closestBlock, nil
	}
	closestBlockNumber, err := HexToBigInt(closestBlock.Number)
	if err != nil {
		return nil, err
	}
	closestBlockNumberBefore := big.NewInt(0).Add(closestBlockNumber, big.NewInt(1))
	return eth.BlockByNumber(ctx, closestBlockNumberBefore)
}
