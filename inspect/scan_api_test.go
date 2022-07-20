//go:build ignore
// +build ignore

// TODO: Re-enable these tests by removing above build tag.
// They are disabled for now due to performance issues and flakiness.

package inspect

import (
	"context"
	"math"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

type Env struct {
}

var (
	APINoTrace = "https://cloudflare-eth.com"
)

func init() {

}

//
func TestRunAllInspections(t *testing.T) {
	type args struct {
		ctx     context.Context
		nodeURL string
	}
	tests := []struct {
		name      string
		args      args
		expected  map[string]float64
		expectErr bool
	}{
		{
			name: "test node w/o trace",
			args: args{
				nodeURL: "https://cloudflare-eth.com",
				ctx:     context.Background(),
			},
			expected: map[string]float64{
				MetricContainerScanAPIAccessible:     ResultSuccess,
				MetricContainerScanAPIChainID:        1,
				MetricContainerScanAPIModuleWeb3:     ResultSuccess,
				MetricContainerScanAPIModuleEth:      ResultSuccess,
				MetricContainerScanAPIModuleNet:      ResultSuccess,
				MetricContainerScanAPIHistorySupport: ResultUnknown,
			},
			expectErr: true,
		},
		{
			name: "detects on bad node access",
			args: args{
				nodeURL: "https://polygon-bdbdfgbd.co",
				ctx:     context.Background(),
			},
			expected: map[string]float64{
				MetricContainerScanAPIAccessible:     ResultFailure,
				MetricContainerScanAPIChainID:        ResultFailure,
				MetricContainerScanAPIModuleWeb3:     ResultFailure,
				MetricContainerScanAPIModuleEth:      ResultFailure,
				MetricContainerScanAPIModuleNet:      ResultFailure,
				MetricContainerScanAPIHistorySupport: ResultFailure,
			},
			expectErr: true,
		},
		{
			name: "test polygon",
			args: args{
				nodeURL: "https://polygon-rpc.com",
				ctx:     context.Background(),
			},
			expected: map[string]float64{
				MetricContainerScanAPIAccessible:     ResultSuccess,
				MetricContainerScanAPIChainID:        137,
				MetricContainerScanAPIModuleWeb3:     ResultSuccess,
				MetricContainerScanAPIModuleEth:      ResultSuccess,
				MetricContainerScanAPIModuleNet:      ResultSuccess,
				MetricContainerScanAPIHistorySupport: ResultUnknown,
			},
			expectErr: false,
		},
		{
			name: "bad url",
			args: args{
				nodeURL: "badurl",
				ctx:     context.Background(),
			},
			expected: map[string]float64{
				MetricContainerScanAPIAccessible:     ResultFailure,
				MetricContainerScanAPIChainID:        ResultFailure,
				MetricContainerScanAPIModuleWeb3:     ResultFailure,
				MetricContainerScanAPIModuleEth:      ResultFailure,
				MetricContainerScanAPIModuleNet:      ResultFailure,
				MetricContainerScanAPIHistorySupport: ResultFailure,
			},
			expectErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				r := require.New(t)

				results, err := RunAllInspections(tt.args.ctx, tt.args.nodeURL)
				if tt.expectErr {
					r.Error(err)
				} else {
					r.NoError(err)
				}

				r.Equal(tt.expected, results)
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
