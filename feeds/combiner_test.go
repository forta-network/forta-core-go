package feeds

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/forta-network/forta-core-go/clients/graphql"
	mock_graphql "github.com/forta-network/forta-core-go/clients/mocks"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_combinerFeed_Start(t *testing.T) {
	type args struct {
		rate                uint64
		stopAfterFirstAlert bool
		expectErr           error
	}

	subscriberBot := "0xsubscriber"
	subscribeeBot := "0xsubscribee"

	ctrl := gomock.NewController(t)
	successfulMockClient := mock_graphql.NewMockClient(ctrl)
	successfulMockClient.EXPECT().GetAlerts(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		[]*protocol.AlertEvent{
			{
				Alert: &protocol.AlertEvent_Alert{
					Hash:      "0xaaaaa",
					CreatedAt: time.Now().Format(time.RFC3339),
					Source: &protocol.AlertEvent_Alert_Source{
						Bot: &protocol.AlertEvent_Alert_Bot{Id: subscribeeBot},
					},
				},
			},
		}, nil,
	)

	tests := []struct {
		name   string
		args   args
		client graphql.Client
	}{
		{
			name: "successfully feeds alerts",
			args: args{
				rate:                uint64(time.Second.Milliseconds()),
				stopAfterFirstAlert: true,
				expectErr:           context.Canceled,
			},
			client: successfulMockClient,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				r := require.New(t)

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
				defer cancel()

				cf, err := NewCombinerFeedWithClient(
					ctx, CombinerFeedConfig{
						QueryInterval: tt.args.rate,
					}, tt.client,
				)

				r.NoError(err)

				err = cf.AddSubscription(
					&domain.CombinerBotSubscription{
						Subscription: &protocol.CombinerBotSubscription{
							BotId: subscribeeBot,
						},
						Subscriber: &domain.Subscriber{BotID: subscriberBot, BotOwner: "0x"},
					},
				)
				r.NoError(err)

				errCh := cf.RegisterHandler(
					func(evt *domain.AlertEvent) error {
						if tt.args.stopAfterFirstAlert {
							cancel()
						}
						return nil
					},
				)
				cf.Start()
				r.Error(tt.args.expectErr, <-errCh)
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
