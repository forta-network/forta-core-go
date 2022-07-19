package inspect

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
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
				r := require.New(t)

				got, err := RunNetworkInspection(tt.args.ctx)
				if tt.wantErr {
					r.Error(err)
					return
				}

				r.NoError(err)
				r.Greater(got[MetricContainerNetworkDownloadSpeed], float64(0), "download speed test failed")
				r.Greater(got[MetricContainerNetworkUploadSpeed], float64(0), "upload speed test failed")
			},
		)
	}
}
