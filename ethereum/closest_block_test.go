package ethereum

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/forta-network/forta-core-go/utils"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
	"time"
)

func TestRandomBlocks(t *testing.T) {

	ctx := context.Background()
	rawurl := "https://rpc.ankr.com/eth"
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		t.Error(fmt.Errorf("failed to dial: %v", err))
		return
	}

	for _, i := range []int64{16390016, 16390015, 8145008, 4072504, 2036252, 2036262} {

		b, err := client.BlockByNumber(ctx, big.NewInt(i))
		if err != nil {
			t.Error(fmt.Errorf("failed to get latest block number: %v", err))
			return
		}

		log.Infof("Got block : %d  -  %d  %s", b.Number(), b.Time(), time.Unix(int64(b.Time()), 0).String())
	}

}

func TestGetClosestBlock(t *testing.T) {

	ankrUrl := "https://rpc.ankr.com/eth"
	ctx := context.Background()
	eth, err := NewStreamEthClient(ctx, "ankr", ankrUrl)
	if err != nil {
		t.Error(err)
		return
	}

	tests := []struct {
		name       string
		activeTime time.Time
		want       int64
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name:       "Most recent",
			activeTime: time.Unix(1673517970, 0),
			want:       16390015,
			wantErr:    assert.NoError,
		},
		{
			name:       "Further away",
			activeTime: time.Unix(1563052013, 0),
			want:       8145008,
			wantErr:    assert.NoError,
		},
		{
			name:       "Exact time",
			activeTime: time.Unix(1501002503, 0),
			want:       4072504,
			wantErr:    assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := eth.GetClosestBlock(ctx, tt.activeTime)
			if !tt.wantErr(t, err, fmt.Sprintf("GetClosestBlock( @ %v)", tt.activeTime)) {
				return
			}
			result, err := utils.HexToBigInt(got.Number)
			if !tt.wantErr(t, err, fmt.Sprintf("HexToBigInt( %v)", got.Number)) {
				return
			}
			assert.Equalf(t, big.NewInt(tt.want), result, "GetClosestBlock( @ %v)", tt.activeTime)
		})
	}
}

func TestGetClosestBlockAfter(t *testing.T) {

	ankrUrl := "https://rpc.ankr.com/eth"
	ctx := context.Background()
	eth, err := NewStreamEthClient(ctx, "ankr", ankrUrl)
	if err != nil {
		t.Error(err)
		return
	}

	tests := []struct {
		name       string
		activeTime time.Time
		want       int64
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name:       "Exact current",
			activeTime: time.Unix(1470690792, 0),
			want:       2036252,
			wantErr:    assert.NoError,
		},

		{
			name:       "After current",
			activeTime: time.Unix(1470690790, 0),
			want:       2036252,
			wantErr:    assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := eth.GetClosestBlockAfter(ctx, tt.activeTime)
			if !tt.wantErr(t, err, fmt.Sprintf("GetClosestBlockAfter( @ %v)", tt.activeTime)) {
				return
			}
			result, err := utils.HexToBigInt(got.Number)
			if !tt.wantErr(t, err, fmt.Sprintf("HexToBigInt( %v)", got.Number)) {
				return
			}
			assert.Equalf(t, big.NewInt(tt.want), result, "GetClosestBlockAfter(@  %v)", tt.activeTime)
		})
	}
}

func TestGetClosestBlockBefore(t *testing.T) {

	ankrUrl := "https://rpc.ankr.com/eth"
	ctx := context.Background()
	eth, err := NewStreamEthClient(ctx, "ankr", ankrUrl)
	if err != nil {
		t.Error(err)
		return
	}

	tests := []struct {
		name       string
		activeTime time.Time
		want       int64
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name:       "Before Current",
			activeTime: time.Unix(1470691013, 0),
			want:       2036262,
			wantErr:    assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := eth.GetClosestBlockBefore(ctx, tt.activeTime)
			if !tt.wantErr(t, err, fmt.Sprintf("GetClosestBlockBefore( @ %v)", tt.activeTime)) {
				return
			}

			result, err := utils.HexToBigInt(got.Number)
			if !tt.wantErr(t, err, fmt.Sprintf("HexToBigInt( %v)", got.Number)) {
				return
			}

			assert.Equalf(t, big.NewInt(tt.want), result, "GetClosestBlockBefore( @ %v)", tt.activeTime)
		})
	}
}
