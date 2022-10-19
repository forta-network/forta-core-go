package feeds

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/forta-network/forta-core-go/clients/graphql"
	"github.com/forta-network/forta-core-go/clients/health"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/sirupsen/logrus"
)

type cfHandler struct {
	Handler func(evt *domain.AlertEvent) error
	ErrCh   chan<- error
}

type combinerFeed struct {
	offset    int
	ctx       context.Context
	cache     utils.Cache
	started   bool
	rateLimit *time.Ticker

	lastAlert health.MessageTracker
	workers   int

	handlers   []cfHandler
	handlersMu sync.RWMutex
	alertCh    chan *domain.AlertEvent
	client     *graphql.Client

	botSubscriptions map[string][]string
	botsMu           sync.RWMutex
	// createdSince the earliest alert in milliseconds
	createdSince uint
	alertCache   utils.Cache
}

func (cf *combinerFeed) Subscriptions() map[string][]string {
	cf.botsMu.RLock()
	defer cf.botsMu.RUnlock()
	return cf.botSubscriptions
}
func (cf *combinerFeed) SubscribedBots() (bots []string) {
	cf.botsMu.RLock()
	defer cf.botsMu.RUnlock()

	for sub := range cf.botSubscriptions {
		bots = append(bots, sub)
	}
	return
}
func (cf *combinerFeed) AddSubscription(subscription, subscriber string) {
	cf.botsMu.Lock()
	defer cf.botsMu.Unlock()
	for _, s := range cf.botSubscriptions[subscription] {
		if s == subscriber {
			return
		}
	}
	cf.botSubscriptions[subscription] = append(cf.botSubscriptions[subscription], subscriber)
}
func (cf *combinerFeed) RemoveSubscription(subscription, subscriber string) {
	cf.botsMu.Lock()
	defer cf.botsMu.Unlock()

	for i, s := range cf.botSubscriptions[subscription] {
		if s == subscriber {
			cf.botSubscriptions[subscription] = append(cf.botSubscriptions[subscription][:i], cf.botSubscriptions[subscription][i+1:]...)
		}
	}
	if len(cf.botSubscriptions[subscription]) == 0 {
		delete(cf.botSubscriptions, subscription)
	}
}

type CombinerFeedConfig struct {
	Offset    int
	RateLimit *time.Ticker
	APIUrl    string
}

func (cf *combinerFeed) initialize() error {
	return nil
}

func (cf *combinerFeed) IsStarted() bool {
	return cf.started
}

func (cf *combinerFeed) Start() {

}

func (cf *combinerFeed) Subscribe(handler func(evt *domain.AlertEvent) error) <-chan error {
	cf.handlersMu.Lock()
	defer cf.handlersMu.Unlock()

	errCh := make(chan error)
	cf.handlers = append(
		cf.handlers, cfHandler{
			Handler: handler,
			ErrCh:   errCh,
		},
	)
	return errCh
}

func (cf *combinerFeed) ForEachAlert(alertHandler func(evt *domain.AlertEvent) error) error {
	for {
		if cf.ctx.Err() != nil {
			return cf.ctx.Err()
		}
		if cf.rateLimit != nil {
			<-cf.rateLimit.C
		}

		// skip query if there are no alert subscriptions
		if len(cf.SubscribedBots()) == 0 {
			continue
		}

		var alerts []*protocol.AlertEvent
		bo := backoff.NewExponentialBackOff()
		err := backoff.Retry(
			func() error {
				var cErr error
				alerts, cErr = cf.client.GetAlerts(
					cf.ctx,
					&graphql.AlertsInput{Bots: cf.SubscribedBots()},
				)
				if cErr != nil {
					logrus.WithError(cErr).Warn("error retrieving alerts")
					return cErr
				}

				return nil
			}, bo,
		)
		if err != nil {
			return err
		}

		for _, alert := range alerts {
			if cf.alertCache.Exists(alert.Alert.Hash) {
				continue
			}

			alertCA, err := time.Parse(time.RFC3339, alert.Alert.CreatedAt)
			if err != nil {
				return err
			}
			evt := &domain.AlertEvent{
				Event: alert,
				Timestamps: &domain.TrackingTimestamps{
					Block: alertCA,
					Feed:  time.Now().UTC(),
				},
			}

			if err := alertHandler(evt); err != nil {
				return err
			}

			cf.alertCache.Add(alert.Alert.Hash)
		}
	}
}

// Name returns the name of this implementation.
func (cf *combinerFeed) Name() string {
	return "alert-feed"
}

// Health implements the health.Reporter interface.
func (cf *combinerFeed) Health() health.Reports {
	return health.Reports{
		cf.lastAlert.GetReport("last-alert"),
	}
}

func NewCombinerFeed(ctx context.Context, cfg CombinerFeedConfig) (AlertFeed, error) {
	if cfg.Offset < 0 {
		return nil, fmt.Errorf("offset cannot be below zero: offset=%d", cfg.Offset)
	}
	ac := graphql.NewClient(cfg.APIUrl)
	alerts := make(chan *domain.AlertEvent, 10)

	bf := &combinerFeed{
		offset:           cfg.Offset,
		ctx:              ctx,
		client:           ac,
		cache:            utils.NewCache(10000),
		rateLimit:        cfg.RateLimit,
		alertCh:          alerts,
		botSubscriptions: map[string][]string{},
		alertCache:       utils.NewCache(1e6),
	}
	return bf, nil
}
