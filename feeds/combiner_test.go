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

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
				defer cancel()

				rate := time.NewTicker(time.Duration(tt.args.rate))
				cf, err := NewCombinerFeed(
					ctx, CombinerFeedConfig{
						RateLimit: rate,
						APIUrl:    "https://api.forta.network",
					},
				)
				r.NoError(err)
				_ = cf.AddSubscription(
					&domain.CombinerBotSubscription{
						Subscription: &protocol.CombinerBotSubscription{
							BotId:   "0x77281ae942ee1fe141d0652e9dad7d001761552f906fb1684b2812603de31049",
							ChainId: 250,
						},
					},
				)
				errCh := cf.RegisterHandler(
					func(evt *domain.AlertEvent) error {
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
		botSubscriptions []*domain.CombinerBotSubscription
	}
	type args struct {
		subscription *domain.CombinerBotSubscription
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "can handle nil subscriptions",
			fields:  fields{botSubscriptions: []*domain.CombinerBotSubscription{}},
			args:    args{subscription: nil},
			wantErr: true,
		},
		{
			name: "can handle bad subscriptions",
			fields: fields{
				botSubscriptions: []*domain.CombinerBotSubscription{},
			},
			args: args{
				subscription: &domain.CombinerBotSubscription{
					Subscription: &protocol.CombinerBotSubscription{BotId: ""},
				},
			},
			wantErr: true,
		},
		{
			name: "can handle adding same subscription",
			fields: fields{
				botSubscriptions: []*domain.CombinerBotSubscription{
					{
						Subscription: &protocol.CombinerBotSubscription{BotId: "0xexists"},
					},
				},
			},
			args: args{
				subscription: &domain.CombinerBotSubscription{
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
					botsMu:           sync.RWMutex{},
				}
				err := cf.AddSubscription(tt.args.subscription)
				assert.Equal(t, tt.wantErr, err != nil)
			},
		)
	}
}
