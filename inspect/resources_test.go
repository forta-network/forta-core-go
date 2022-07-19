package inspect

import (
	"context"
	"testing"
)

func TestGetResourceMetrics(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"can read resources", false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				_, err := GetResources(context.Background())
				if (err != nil) != tt.wantErr {
					t.Errorf("GetResources() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}

func Test_cpuBenchmark(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "can benchmark cpu",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				cpuStressTestBenchmark()
			},
		)
	}
}
