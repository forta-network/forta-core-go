package feeds

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/forta-network/forta-core-go/clients/graphql"
	"github.com/forta-network/forta-core-go/protocol"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/forta-network/forta-core-go/clients/health"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/utils"
)

type afHandler struct {
	Handler func(evt *domain.AlertEvent) error
	ErrCh   chan<- error
}

type alertFeed struct {
	offset    int
	ctx       context.Context
	cache     utils.Cache
	started   bool
	rateLimit *time.Ticker

	lastAlert health.MessageTracker
	workers   int

	handlers   []afHandler
	handlersMu sync.RWMutex
	alertCh    chan *domain.AlertEvent
	client     *graphql.Client

	botSubscriptions []string
	botsMu           sync.RWMutex
	// createdSince the earliest alert in milliseconds
	createdSince uint
}

func (af *alertFeed) Bots() []string {
	af.botsMu.RLock()
	defer af.botsMu.RUnlock()
	return af.botSubscriptions
}

func (af *alertFeed) SetBots(bots []string) {
	af.botsMu.Lock()
	defer af.botsMu.Unlock()
	af.botSubscriptions = bots
}

type AlertFeedConfig struct {
	Offset    int
	RateLimit *time.Ticker
	APIUrl    string
}

func (af *alertFeed) initialize() error {
	return nil
}

func (af *alertFeed) IsStarted() bool {
	return af.started
}

func (af *alertFeed) Start() {
	if !af.started {
		go af.loop()
	}
}

func (af *alertFeed) loop() {
	if err := af.initialize(); err != nil {
		log.WithError(err).Panic("failed to initialize")
	}

	af.started = true
	defer func() {
		af.started = false
	}()
	err := af.forEachBlock()
	if err == nil {
		return
	}
	if err != ErrEndBlockReached {
		log.WithError(err).Warn("failed while processing blocks")
	}
	for _, handler := range af.handlers {
		handler.ErrCh <- err
	}
}

func (af *alertFeed) Subscribe(handler func(evt *domain.AlertEvent) error) <-chan error {
	af.handlersMu.Lock()
	defer af.handlersMu.Unlock()

	errCh := make(chan error)
	af.handlers = append(
		af.handlers, afHandler{
			Handler: handler,
			ErrCh:   errCh,
		},
	)
	return errCh
}

func (af *alertFeed) forEachBlock() error {
	for {
		if af.ctx.Err() != nil {
			return af.ctx.Err()
		}
		if af.rateLimit != nil {
			<-af.rateLimit.C
		}

		af.botsMu.RLock()
		alerts, err := af.client.GetAlerts(
			af.ctx,
			&graphql.AlertsInput{Bots: af.botSubscriptions},
		)
		af.botsMu.RUnlock()
		if err != nil {
			return err
		}

		for _, alert := range alerts {
			evt := &domain.AlertEvent{
				Alert: &protocol.AlertEvent{
					Alert: alert,
				},
			}
			af.handlersMu.RLock()
			handlers := af.handlers
			af.handlersMu.RUnlock()
			for _, handler := range handlers {
				if err := handler.Handler(evt); err != nil {
					return err
				}
			}
		}
	}
}

func (af *alertFeed) ForEachAlert(alertHandler func(evt *domain.AlertEvent) error) error {
	grp, _ := errgroup.WithContext(af.ctx)

	// iterate over blocks
	grp.Go(
		func() error {
			errCh := af.Subscribe(
				func(evt *domain.AlertEvent) error {
					af.alertCh <- evt
					var blockHandlerErr error
					if alertHandler != nil {
						blockHandlerErr = alertHandler(evt)
					}
					return blockHandlerErr
				},
			)
			err := <-errCh
			close(af.alertCh)
			if err == ErrEndBlockReached {
				return nil
			}
			return err
		},
	)

	// block until above all finish (when context is cancelled or error returns)
	return grp.Wait()
}

// Name returns the name of this implementation.
func (af *alertFeed) Name() string {
	return "alert-feed"
}

// Health implements the health.Reporter interface.
func (af *alertFeed) Health() health.Reports {
	return health.Reports{
		af.lastAlert.GetReport("last-alert"),
	}
}

func NewAlertFeed(ctx context.Context, cfg AlertFeedConfig) (*alertFeed, error) {
	if cfg.Offset < 0 {
		return nil, fmt.Errorf("offset cannot be below zero: offset=%d", cfg.Offset)
	}
	ac := graphql.NewClient(cfg.APIUrl)

	bf := &alertFeed{
		offset:    cfg.Offset,
		ctx:       ctx,
		client:    ac,
		cache:     utils.NewCache(10000),
		rateLimit: cfg.RateLimit,
	}
	return bf, nil
}
