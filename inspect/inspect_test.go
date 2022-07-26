package inspect

import (
	"context"
	"testing"
)

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
			name: "test eth rpc with no trace",
			args: args{
				nodeURL: "https://cloudflare-eth.com",
				ctx:     context.Background(),
			},
			wantErr: false,
			want: map[string]float64{
				IndicatorProxyAPIAccessible: ResultSuccess,
				IndicatorScanAPIChainID:     1,
				IndicatorScanAPIModuleEth:   ResultSuccess,
				IndicatorScanAPIModuleNet:   ResultSuccess,
				IndicatorTraceSupported:     ResultSuccess,
			},
		},
		{
			name: "detects on bad node access",
			args: args{
				nodeURL: "https://polygon-invalid__url__.co",
				ctx:     context.Background(),
			},
			want: map[string]float64{
				IndicatorScanAPIAccessible: ResultFailure,
				IndicatorScanAPIChainID:    ResultFailure,
				IndicatorScanAPIModuleEth:  ResultFailure,
				IndicatorScanAPIModuleNet:  ResultFailure,
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
				IndicatorScanAPIAccessible: ResultSuccess,
				IndicatorScanAPIChainID:    137,
				IndicatorScanAPIModuleEth:  ResultSuccess,
				IndicatorScanAPIModuleNet:  ResultSuccess,
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
				IndicatorScanAPIAccessible: ResultFailure,
				IndicatorScanAPIChainID:    ResultFailure,
				IndicatorScanAPIModuleEth:  ResultFailure,
				IndicatorScanAPIModuleNet:  ResultFailure,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				_, err := Inspect(tt.args.ctx, InspectionConfig{ScanAPIURL: tt.args.nodeURL, ProxyAPIURL: tt.args.nodeURL, TraceAPIURL: tt.args.nodeURL, CheckTrace: false})
				if (err != nil) != tt.wantErr {
					t.Errorf("Inspect() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}
