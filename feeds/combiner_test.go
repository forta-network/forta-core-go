package feeds

import (
	"context"
	"testing"
	"time"

	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/stretchr/testify/require"
)

func Test_combinerFeed_Start(t *testing.T) {
	type args struct {
		rate                int64
		stopAfterFirstAlert bool
		expectErr           error
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "successfully feeds alerts",
			args: args{
				rate:                int64(time.Nanosecond),
				stopAfterFirstAlert: false,
				expectErr:           context.DeadlineExceeded,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				r := require.New(t)

				ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
				defer cancel()

				rate := time.NewTicker(time.Duration(tt.args.rate))
				cf, err := NewCombinerFeed(
					ctx, CombinerFeedConfig{
						APIUrl:    "https://api.forta.network/graphql",
						RateLimit: rate,
					},
				)
				r.NoError(err)
				cf.AddSubscription(
					&protocol.CombinerBotSubscription{
						BotId:   "0x5e13c2f3a97c292695b598090056ba5d52f9dcc7790bcdaa8b6cd87c1a1ebc0f",
						AlertId: "FORTA-6",
					},
				)
				errCh := cf.RegisterHandler(
					func(evt *domain.AlertEvent) error {
						t.Logf("got alert: %s, created at: %s", evt.Event.Alert.Hash, evt.Event.Alert.CreatedAt)
						if tt.args.stopAfterFirstAlert {
							cancel()
						}
						return nil
					},
				)
				cf.Start()
				r.Equal(tt.args.expectErr, <-errCh)
			},
		)
	}
}
