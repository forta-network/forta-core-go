package feeds

import (
	"context"
	"testing"
	"time"

	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/stretchr/testify/require"
)

func Test_combinerFeed_StartRange(t *testing.T) {
	type args struct {
		start               uint64
		end                 uint64
		rate                int64
		stopAfterFirstAlert bool
		expectErr           error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "successfully feeds range",
			args: args{
				start:               uint64(time.Now().Add(time.Second * -15).UnixMilli()),
				end:                 uint64(time.Now().UnixMilli()),
				rate:                int64(time.Nanosecond),
				stopAfterFirstAlert: false,
				expectErr:           ErrCombinerStopReached,
			},
		}, {
			name: "no range",
			args: args{
				start:               0,
				end:                 0,
				rate:                int64(time.Nanosecond),
				stopAfterFirstAlert: true,
				expectErr:           context.Canceled,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				r := require.New(t)

				ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
				defer cancel()
				cf, err := NewCombinerFeed(
					ctx, CombinerFeedConfig{
						APIUrl: "https://api.forta.network/graphql",
					},
				)
				r.NoError(err)
				cf.AddSubscription(
					&protocol.CombinerBotSubscription{
						BotId:   "0x5e13c2f3a97c292695b598090056ba5d52f9dcc7790bcdaa8b6cd87c1a1ebc0f",
						AlertId: "FORTA-6",
					},
				)
				errCh := cf.RegisterHandler(func(evt *domain.AlertEvent) error {
					t.Logf("got alert: %s", evt.Event.Alert.Hash)
					if tt.args.stopAfterFirstAlert {
						cancel()
					}
					return nil
				})
				cf.StartRange(tt.args.start, tt.args.end, tt.args.rate)
				r.Equal(tt.args.expectErr, <-errCh)
			},
		)
	}
}
