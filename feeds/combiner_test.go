package feeds

import (
	"context"
	"testing"
	"time"

	"github.com/forta-network/forta-core-go/protocol"
)

func Test_combinerFeed_StartRange(t *testing.T) {
	type args struct {
		start uint64
		end   uint64
		rate  int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "successfully feeds range",
			args: args{
				start: uint64(time.Now().Add(-1 * time.Hour).UnixMilli()),
				end:   uint64(time.Now().UnixMilli()),
				rate:  0,
			},
		}, {
			name: "no range",
			args: args{
				start: 0,
				end:   0,
				rate:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				ctx, cancel := context.WithTimeout(
					context.Background(),
					time.Second*555,
				)
				defer cancel()
				cf, err := NewCombinerFeed(
					ctx, CombinerFeedConfig{
						APIUrl: "https://api.forta.network/graphql",
					},
				)
				if err != nil {
					t.Fatal(err)
				}
				cf.AddSubscription(
					&protocol.CombinerBotSubscription{
						BotId:   "0x5e13c2f3a97c292695b598090056ba5d52f9dcc7790bcdaa8b6cd87c1a1ebc0f",
						AlertId: "FORTA-6",
					},
				)
				cf.StartRange(tt.args.start, tt.args.end, tt.args.rate)
				<-ctx.Done()
			},
		)
	}
}
