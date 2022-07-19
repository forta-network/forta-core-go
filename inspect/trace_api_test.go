package inspect

import (
	"context"
	"reflect"
	"testing"
)

func TestRunTraceAPIInspection(t *testing.T) {
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
		/*		{
				name: "can support trace",
				args: args{
					ctx:     context.TODO(),
					nodeURL: "http://erigon-eth-prod.forta.internal:8545",
				},
				wantErr: false,
				want: map[string]float64{
					MetricContainerTraceSupported:  StateSuccess,
					MetricContainerTraceAccessible: StateSuccess,
				},
			},*/
		{
			name: "detects eth-mainnet without tracing",
			args: args{
				ctx:     context.TODO(),
				nodeURL: "https://cloudflare-eth.com",
			},
			wantErr: true,
			want: map[string]float64{
				MetricContainerTraceSupported:  StateError,
				MetricContainerTraceAccessible: StateSuccess,
			},
		},
		{
			name: "no trace needed except eth-mainnet",
			args: args{
				ctx:     context.TODO(),
				nodeURL: "https://rpc.poa.psdk.io:8545",
			},
			wantErr: false,
			want: map[string]float64{
				MetricContainerTraceSupported:  StateError,
				MetricContainerTraceAccessible: StateSuccess,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := RunTraceAPIInspection(tt.args.ctx, tt.args.nodeURL)
				if (err != nil) != tt.wantErr {
					t.Errorf("RunTraceAPIInspection() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("RunTraceAPIInspection() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
