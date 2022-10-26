package feeds

import (
	"context"
	"testing"
	"time"
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
				start: 1666766210406,
				end:   uint64(1666766210406 + time.Minute.Milliseconds()),
				rate:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
				defer cancel()
				cf, err := NewCombinerFeed(
					ctx, CombinerFeedConfig{
						APIUrl: "https://api.forta.network/graphql",
					},
				)
				if err != nil {
					t.Fatal(err)
				}
				cf.AddSubscription("0x5e13c2f3a97c292695b598090056ba5d52f9dcc7790bcdaa8b6cd87c1a1ebc0f", "0xlocalbot")
				cf.StartRange(tt.args.start, tt.args.end, tt.args.rate)
			},
		)
	}
}
