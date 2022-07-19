package inspect

import (
	"context"
	"testing"
)

func TestRunNetworkInspection(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]float64
		wantErr bool
	}{
		{
			name:    "can pass inspection",
			args:    args{ctx: context.TODO()},
			want:    map[string]float64{MetricContainerNetworkOutboundAccess: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := RunNetworkInspection(tt.args.ctx)
				if (err != nil) != tt.wantErr {
					t.Errorf("RunNetworkInspection() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr {
					if got[MetricContainerNetworkDownloadSpeed] <= 0 {
						t.Errorf("RunNetworkInspection() download speed test failed")
					}
					if got[MetricContainerNetworkUploadSpeed] <= 0 {
						t.Errorf("RunNetworkInspection() upload speed test failed")
					}
				}
			},
		)
	}
}
