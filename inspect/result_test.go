package inspect

import (
	"context"
	"testing"
)

//
func TestRunAllInspections(t *testing.T) {
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
			name: "test node w/o trace",
			args: args{
				nodeURL: "https://cloudflare-eth.com",
				ctx:     context.Background(),
			},
			wantErr: true,
			want: map[string]float64{
				MetricContainerScanAPIAccessible:     StateSuccess,
				MetricContainerScanAPIChainID:        1,
				MetricContainerScanAPIModuleWeb3:     StateSuccess,
				MetricContainerScanAPIModuleEth:      StateSuccess,
				MetricContainerScanAPIModuleNet:      StateSuccess,
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
				MetricContainerScanAPIChainID:        StateError,
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
				MetricContainerScanAPIChainID:        137,
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
				_, err := RunAllInspections(tt.args.ctx, tt.args.nodeURL)
				if (err != nil) != tt.wantErr {
					t.Errorf("RunAllInspections() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				// TODO: temporarily disabled
				// if !reflect.DeepEqual(got, tt.want) {
				// 	t.Errorf("RunAllInspections() got = %v, want %v", got, tt.want)
				// }
			},
		)
	}
}
