package feeds

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/stretchr/testify/assert"
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
						RateLimit: rate,
						APIUrl:    "https://api-dev.forta.network/graphql",
					},
				)
				r.NoError(err)
				cf.AddSubscription(
					&CombinerBotSubscription{
						Subscription: &protocol.CombinerBotSubscription{
							BotId:   "0x2bee737433c0c8cdbd924bbb68306cfd8abcf0e46a6ce8994fa7d474361bb186",
							AlertId: "FORTA_1",
						},
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

func Test_combinerFeed_AddSubscription(t *testing.T) {
	type fields struct {
		botSubscriptions []*CombinerBotSubscription
		botsMu           sync.RWMutex
	}
	type args struct {
		subscription *CombinerBotSubscription
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "can handle nil subscriptions",
			fields:  fields{botSubscriptions: []*CombinerBotSubscription{}},
			args:    args{subscription: nil},
			wantErr: true,
		},
		{
			name: "can handle bad subscriptions",
			fields: fields{
				botSubscriptions: []*CombinerBotSubscription{
				},
			},
			args: args{
				subscription: &CombinerBotSubscription{
					Subscription: &protocol.CombinerBotSubscription{BotId: ""},
				},
			},
			wantErr: true,
		},
		{
			name: "can handle adding same subscription",
			fields: fields{
				botSubscriptions: []*CombinerBotSubscription{
					{
						Subscription: &protocol.CombinerBotSubscription{BotId: "0xexists"},
					},
				},
			},
			args: args{
				subscription: &CombinerBotSubscription{
					Subscription: &protocol.CombinerBotSubscription{BotId: "0xexists"},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				cf := &combinerFeed{
					botSubscriptions: tt.fields.botSubscriptions,
					botsMu:           tt.fields.botsMu,
				}
				err := cf.AddSubscription(tt.args.subscription)
				assert.Equal(t, tt.wantErr, err != nil)
			},
		)
	}
}