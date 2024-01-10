package feeds

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/forta-network/forta-core-go/clients/graphql"
	mock_graphql "github.com/forta-network/forta-core-go/clients/mocks"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_combinerFeed_Start(t *testing.T) {
	subscriberBot := "0xsubscriber"
	subscribeeBot := "0xsubscribee"
	rate := uint64(20 * time.Second.Milliseconds())

	ctrl := gomock.NewController(t)
	successfulAlertResponse := []*protocol.AlertEvent{
		{
			Alert: &protocol.AlertEvent_Alert{
				Hash:      "0xaaaaa",
				CreatedAt: time.Now().Format(time.RFC3339),
				Source: &protocol.AlertEvent_Alert_Source{
					Bot: &protocol.AlertEvent_Alert_Bot{Id: subscribeeBot},
				},
			},
		},
	}

	var subscriptions []*domain.CombinerBotSubscription
	for i := 0; i < 20; i++ {
		subscriptions = append(subscriptions, &domain.CombinerBotSubscription{
			Subscription: &protocol.CombinerBotSubscription{
				BotId: fmt.Sprintf("0xbot%d", i),
			},
			Subscriber: &domain.Subscriber{BotID: subscriberBot, BotOwner: "0x", BotImage: "0x123"},
		})
	}

	//
	// Test Case 1: can fetch subscriptions without any issues
	// Setup: 1 subscriber bot, has 20 subscriptions
	// Batch size is 10, meaning there should be 2 GetAlertsBatch requests
	successfulMockClient := mock_graphql.NewMockClient(ctrl)
	successfulMockClient.EXPECT().GetAlertsBatch(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(successfulAlertResponse, nil).
		Times(2)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	cf, err := NewCombinerFeedWithClient(
		ctx, CombinerFeedConfig{
			QueryInterval: rate,
		}, successfulMockClient,
	)

	assert.NoError(t, err)
	for _, subscription := range subscriptions {
		err = cf.AddSubscription(subscription)
		assert.NoError(t, err)
	}

	errCh := cf.RegisterHandler(
		func(evt *domain.AlertEvent) error {
			cancel()
			return nil
		},
	)
	cf.Start()
	assert.Error(t, context.Canceled, <-errCh)

	//
	//
	// Test Case 2: Retries in smaller chunks if there is a response size error
	// Test Setup: 1 bot with 20 subscriptions
	// Batch size is 10, there should be 2 requests. However, first request fails due to response size
	// Resulting in 4 GetAlertsBatch calls.
	//
	//
	ctrl = gomock.NewController(t)
	responseTooBigClient := mock_graphql.NewMockClient(ctrl)
	responseTooBigClient.EXPECT().GetAlertsBatch(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		nil, graphql.ErrResponseSizeTooBig,
	)
	responseTooBigClient.EXPECT().GetAlertsBatch(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		successfulAlertResponse, nil,
	).Times(3)

	ctxTooBig, cancelTooBig := context.WithTimeout(context.Background(), time.Second*50)
	defer cancelTooBig()

	cfTooBig, err := NewCombinerFeedWithClient(
		ctxTooBig, CombinerFeedConfig{
			QueryInterval: rate,
		}, responseTooBigClient,
	)
	assert.NoError(t, err)

	for i := 0; i < 20; i++ {
		err = cfTooBig.AddSubscription(&domain.CombinerBotSubscription{
			Subscription: &protocol.CombinerBotSubscription{
				BotId: fmt.Sprintf("0xbot%d", i),
			},
			Subscriber: &domain.Subscriber{BotID: subscriberBot, BotOwner: "0x", BotImage: "0x123"},
		})
		assert.NoError(t, err)
	}

	errChTooBig := cfTooBig.RegisterHandler(
		func(evt *domain.AlertEvent) error {
			cancelTooBig()
			return nil
		},
	)
	cfTooBig.Start()
	assert.Error(t, context.Canceled, <-errChTooBig)
	//
	//
	// Test Case 3: Retries in smaller chunks if there is a response size error, stops at some point
	// Test Setup: 1 bot with 20 subscriptions
	// Batch size is 10, there should be 2 requests. However, all requests fail.
	// Resulting in 80 GetAlertsBatch calls.
	//
	//
	ctrl = gomock.NewController(t)
	unauthClient := mock_graphql.NewMockClient(ctrl)
	unauthClient.EXPECT().GetAlertsBatch(gomock.Any(), gomock.Any(), gomock.Any()).Return(
		nil, fmt.Errorf("unauthorized"),
	).Times(80)

	ctxUnauth, cancelUnauth := context.WithTimeout(context.Background(), time.Second*50)
	defer cancelUnauth()

	cfUnauth, err := NewCombinerFeedWithClient(
		ctxUnauth, CombinerFeedConfig{
			QueryInterval: rate,
		}, unauthClient,
	)
	assert.NoError(t, err)

	for i := 0; i < 20; i++ {
		err = cfUnauth.AddSubscription(&domain.CombinerBotSubscription{
			Subscription: &protocol.CombinerBotSubscription{
				BotId: fmt.Sprintf("0xbot%d", i),
			},
			Subscriber: &domain.Subscriber{BotID: subscriberBot, BotOwner: "0x", BotImage: "0x123"},
		})
		assert.NoError(t, err)
	}

	// ensure no alerts are returned
	_ = cfTooBig.RegisterHandler(
		func(evt *domain.AlertEvent) error {
			t.Error("unexpected alert received")
			return nil
		},
	)

	cfUnauth.Start()
	time.Sleep(time.Second * 5)
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
