//go:build ignore
// +build ignore

// TODO: Re-enable these tests by removing above build tag.
// They are disabled for now due to performance issues and flakiness.

package inspect

import (
	"context"
	"math"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
)

func TestRunScanNodeInspection(t *testing.T) {
	type args struct {
		ctx     context.Context
		nodeURL string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]float64
		wantErr bool
	}{
		{
			name: "test a valid node",
			args: args{
				nodeURL: "https://cloudflare-eth.com",
				ctx:     context.Background(),
			},
			want: map[string]float64{
				MetricContainerScanAPIAccessible:     StateSuccess,
				MetricContainerScanAPIModuleWeb3:     StateSuccess,
				MetricContainerScanAPIModuleEth:      StateSuccess,
				MetricContainerScanAPIModuleNet:      StateSuccess,
				MetricContainerScanAPIChainID:        1.0,
				MetricContainerScanAPIHistorySupport: 0.0,
			},
		},
		{
			name: "detects on bad node access",
			args: args{
				nodeURL: "https://polygon-bdbdfgbd.co",
				ctx:     context.Background(),
			},
			want: map[string]float64{
				MetricContainerScanAPIAccessible:     StateError,
				MetricContainerScanAPIChainID:        -1.0,
				MetricContainerScanAPIModuleWeb3:     StateError,
				MetricContainerScanAPIModuleEth:      StateError,
				MetricContainerScanAPIModuleNet:      StateError,
				MetricContainerScanAPIHistorySupport: StateError,
			},
			wantErr: true,
		},
		{
			name: "test polygon",
			args: args{
				nodeURL: "https://polygon-rpc.com",
				ctx:     context.Background(),
			},
			want: map[string]float64{
				MetricContainerScanAPIAccessible:     StateSuccess,
				MetricContainerScanAPIChainID:        137.0,
				MetricContainerScanAPIModuleWeb3:     StateSuccess,
				MetricContainerScanAPIModuleEth:      StateSuccess,
				MetricContainerScanAPIModuleNet:      StateSuccess,
				MetricContainerScanAPIHistorySupport: 0.0,
			},
			wantErr: false,
		},
		{
			name: "bad url",
			args: args{
				nodeURL: "badurl",
				ctx:     context.Background(),
			},
			want: map[string]float64{
				MetricContainerScanAPIAccessible:     StateError,
				MetricContainerScanAPIChainID:        StateError,
				MetricContainerScanAPIModuleWeb3:     StateError,
				MetricContainerScanAPIModuleEth:      StateError,
				MetricContainerScanAPIModuleNet:      StateError,
				MetricContainerScanAPIHistorySupport: StateError,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := RunScanNodeInspection(tt.args.ctx, tt.args.nodeURL)
				if (err != nil) != tt.wantErr {
					t.Errorf("RunScanNodeInspection() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("RunScanNodeInspection() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_findOldestSupportedBlock(t *testing.T) {
	// TODO: add a non-genesis case with mock client
	cli, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	lat, err := cli.BlockNumber(ctx)
	if err != nil {
		t.Fatal(err)
	}

	type args struct {
		ctx    context.Context
		client *ethclient.Client
		high   uint64
		low    uint64
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			name: "api supports genesis block",
			args: args{
				ctx:    ctx,
				client: cli,
				low:    0,
				high:   lat,
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				m := make(map[uint64]bool)
				got := findOldestSupportedBlock(
					tt.args.ctx, m,
					tt.args.client, tt.args.low, tt.args.high,
				)
				t.Log(float64(math.MaxUint64))
				if (err != nil) != tt.wantErr {
					t.Errorf("findOldestSupportedBlock() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("findOldestSupportedBlock() got = %v, want %v", got, tt.want)
				}

			},
		)
	}
}
