package inspect

import (
	"context"
	"testing"
	"time"
)

func TestOffsetInspector_Inspect(t *testing.T) {
	type args struct {
		ctx           context.Context
		inspectionCfg InspectionConfig
	}

	tests := []struct {
		name        string
		args        args
	}{
		{
			name: "different apis",
			args: args{
				ctx: context.Background(),
				inspectionCfg: InspectionConfig{
					ScanAPIURL:  "http://172.31.84.3:8545",
					ProxyAPIURL: "http://172.31.71.62:8545",
				},
			},
		}, {
			name: "dns",
			args: args{
				ctx: context.Background(),
				inspectionCfg: InspectionConfig{
					ScanAPIURL:  "http://erigon-eth-prod.forta.internal:8545",
					ProxyAPIURL: "http://erigon-eth-prod.forta.internal:8545",
				},
			},
		}, {
			name: "same node",
			args: args{
				ctx: context.Background(),
				inspectionCfg: InspectionConfig{
					ScanAPIURL:  "http://172.31.84.3:8545",
					ProxyAPIURL: "http://172.31.84.3:8545",
				},
			},
		}, {
			name: "free and full combo",
			args: args{
				ctx: context.Background(),
				inspectionCfg: InspectionConfig{
					ScanAPIURL:  "https://cloudflare-eth.com",
					ProxyAPIURL: "http://erigon-eth-prod.forta.internal:8545",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				pai := &OffsetInspector{}
				results, err := pai.Inspect(tt.args.ctx, tt.args.inspectionCfg)
				if err != nil {
					t.Log(err)
				}

				if results.Indicators[IndicatorOffsetScanProxySynced] == ResultFailure {
					t.Logf(
						"there is an offset!: %s \t %f \n", time.Now().String(),
						results.Indicators[IndicatorOffsetScanBlockNumber],
					)
				} else {
					t.Logf(
						"no offset: %s \t %f \n", time.Now().String(),
						results.Indicators[IndicatorOffsetScanBlockNumber],
					)
				}

				time.Sleep(time.Millisecond * 500)
			},
		)
	}
}
