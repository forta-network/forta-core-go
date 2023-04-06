package feeds

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/forta-network/forta-core-go/clients/graphql"
	"github.com/forta-network/forta-core-go/clients/health"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
)

type Subscriber struct {
	BotID    string `json:"bot_id"`
	BotOwner string `json:"bot_owner"`
}

type CombinerBotSubscription struct {
	Subscription *protocol.CombinerBotSubscription `json:"subscription"`
	Subscriber   Subscriber                        `json:"subscriber"`
}

var (
	ErrCombinerStopReached   = fmt.Errorf("combiner stop reached")
	DefaultRatelimitDuration = time.Minute
)

type cfHandler struct {
	Handler func(evt *domain.AlertEvent) error
	ErrCh   chan<- error
}

type combinerFeed struct {
	ctx       context.Context
	started   bool
	rateLimit *time.Ticker

	lastAlert health.MessageTracker

	alertCh chan *domain.AlertEvent
	client  *graphql.Client

	botSubscriptions []*CombinerBotSubscription
	botsMu           sync.RWMutex
	// bloom filter could be a better choice but the miss rate was too high
	// and we don't know the expected item count
	alertCache *cache.Cache

	handlers    []cfHandler
	handlersMu  sync.Mutex
	cfg         CombinerFeedConfig
	maxAlertAge time.Duration
}

func (cf *combinerFeed) Subscriptions() []*CombinerBotSubscription {
	cf.botsMu.RLock()
	defer cf.botsMu.RUnlock()

	return cf.botSubscriptions
}

func (cf *combinerFeed) AddSubscription(subscription *CombinerBotSubscription) {
	cf.botsMu.Lock()
	defer cf.botsMu.Unlock()

	for _, s := range cf.botSubscriptions {
		if isSameSubscription(s, subscription) {
			return
		}
	}

	cf.botSubscriptions = append(cf.botSubscriptions, subscription)
}

func (cf *combinerFeed) RemoveSubscription(subscription *CombinerBotSubscription) {
	cf.botsMu.Lock()
	defer cf.botsMu.Unlock()

	for i, s := range cf.botSubscriptions {
		if isSameSubscription(s, subscription) {
			cf.botSubscriptions = append(
				cf.botSubscriptions[:i], cf.botSubscriptions[i+1:]...,
			)
		}
	}
}

func (cf *combinerFeed) RegisterHandler(alertHandler func(evt *domain.AlertEvent) error) <-chan error {
	cf.handlersMu.Lock()
	defer cf.handlersMu.Unlock()

	errCh := make(chan error)
	cf.handlers = append(
		cf.handlers, cfHandler{
			Handler: alertHandler,
			ErrCh:   errCh,
		},
	)
	return errCh
}

type CombinerFeedConfig struct {
	RateLimit         *time.Ticker
	APIUrl            string
	Start             uint64
	End               uint64
	CombinerCachePath string
}

func (cf *combinerFeed) Start() {
	if !cf.started {
		go cf.loop()
	}
}

func (cf *combinerFeed) initialize() error {
	if cf.rateLimit == nil {
		cf.rateLimit = time.NewTicker(DefaultRatelimitDuration)
	}
	return nil
}

func (cf *combinerFeed) forEachAlert(alertHandlers []cfHandler) error {
	firstRun := true
	for {
		if cf.ctx.Err() != nil {
			return cf.ctx.Err()
		}

		if cf.rateLimit != nil {
			if !firstRun {
				// wait for the ratelimit
				<-cf.rateLimit.C
			}
			firstRun = false
		}

		// skip query if there are no alert subscriptions
		if len(cf.Subscriptions()) == 0 {
			continue
		}

		lowerBound := time.Minute * 10
		upperBound := int64(0)

		// query all subscriptions and push
		for _, subscription := range cf.Subscriptions() {
			logger := log.WithField("bot-id", subscription.Subscriber.BotID)
			err := cf.fetchAlertsAndHandle(
				cf.ctx,
				alertHandlers, subscription, lowerBound.Milliseconds(), upperBound,
			)
			if err != nil {
				logger.WithError(err).Warn("failed to fetch alerts")
				continue
			}
		}

		// dump cache to persistent file
		if cf.cfg.CombinerCachePath != "" {
			d, err := json.Marshal(cf.alertCache.Items())
			if err != nil {
				log.Panic(err)
			}
			err = os.WriteFile(cf.cfg.CombinerCachePath, d, 0644)
			if err != nil {
				log.Panic(err)
			}
		}
	}
}

func (cf *combinerFeed) fetchAlertsAndHandle(
	ctx context.Context, alertHandlers []cfHandler, subscription *CombinerBotSubscription, createdSince int64,
	createdBefore int64,
) error {
	var alerts []*protocol.AlertEvent

	bo := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)

	err := backoff.Retry(
		func() error {
			var cErr error
			alerts, cErr = cf.client.GetAlerts(
				cf.ctx,
				&graphql.AlertsInput{
					Bots:          []string{subscription.Subscription.BotId},
					CreatedSince:  uint(createdSince),
					CreatedBefore: uint(createdBefore),
					AlertIds:      subscription.Subscription.AlertIds,
					AlertId:       subscription.Subscription.AlertId,
					ChainId:       uint(subscription.Subscription.ChainId),
				},
				subscriberInfoToHeaders(subscription.Subscriber),
			)
			if cErr != nil && errors.Is(cErr, context.DeadlineExceeded) {
				log.WithError(cErr).Warn("error retrieving alerts")
				return cErr
			}

			return nil
		}, bo,
	)
	if err != nil {
		return err
	}

	// alert post processing
	for _, alert := range alerts {
		if tooOld, age := alertIsTooOld(alert, cf.maxAlertAge); tooOld {
			log.WithField("age", age).Warnf(
				"alert is older than %v - setting current alert iterator head to now", cf.maxAlertAge,
			)

			continue
		}

		if _, exists := cf.alertCache.Get(alert.Alert.Hash); exists {
			continue
		}

		cf.alertCache.Set(
			alert.Alert.Hash, struct{}{}, cache.DefaultExpiration,
		)

		alertCA, err := time.Parse(time.RFC3339, alert.Alert.CreatedAt)
		if err != nil {
			return err
		}

		evt := &domain.AlertEvent{
			Event: alert,
			Timestamps: &domain.TrackingTimestamps{
				Feed:        time.Now().UTC(),
				SourceAlert: alertCA,
			},
		}

		for _, alertHandler := range alertHandlers {
			if err := alertHandler.Handler(evt); err != nil {
				return err
			}
		}
	}

	return nil
}

func (cf *combinerFeed) loop() {
	if err := cf.initialize(); err != nil {
		log.WithError(err).Panic("failed to initialize")
	}
	cf.started = true
	defer func() {
		cf.started = false
	}()

	err := cf.forEachAlert(cf.handlers)
	if err == nil {
		return
	}
	if err != ErrCombinerStopReached {
		log.WithError(err).Warn("failed while processing alerts")
	}
	for _, handler := range cf.handlers {
		if handler.ErrCh != nil {
			handler.ErrCh <- err
		}
	}
}

//
// Helper Methods
//

// subscriberInfoToHeaders converts subscriber information to map[string]string
func subscriberInfoToHeaders(subscriber Subscriber) map[string]string {
	return map[string]string{
		"bot-id":    subscriber.BotID,
		"bot-owner": subscriber.BotOwner,
	}
}

// isSameSubscription checks for field-for-field equality for subscriptions.
func isSameSubscription(a, b *CombinerBotSubscription) bool {
	if a == nil || b == nil {
		return false
	}

	if a.Subscription.BotId != b.Subscription.BotId {
		return false
	}

	if a.Subscription.AlertId != b.Subscription.AlertId {
		return false
	}

	if a.Subscription.ChainId != b.Subscription.ChainId {
		return false
	}

	sort.Strings(a.Subscription.AlertIds)
	sort.Strings(b.Subscription.AlertIds)

	return strings.EqualFold(strings.Join(a.Subscription.AlertIds, ","), strings.Join(b.Subscription.AlertIds, ","))
}
func alertIsTooOld(alert *protocol.AlertEvent, maxAge time.Duration) (bool, *time.Duration) {
	if maxAge == 0 {
		return false, nil
	}

	createdAt, err := time.Parse(time.RFC3339, alert.Alert.CreatedAt)
	age := time.Since(createdAt)
	if err != nil {
		log.WithFields(
			log.Fields{
				"alertHash": alert.Alert.Hash,
			},
		).WithError(err).Errorf("error getting age of block")
		return false, &age
	}

	return age > maxAge, &age
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
	ac := graphql.NewClient(cfg.APIUrl)
	alerts := make(chan *domain.AlertEvent, 10)

	var alertCache *cache.Cache
	if cfg.CombinerCachePath != "" {
		d, err := os.ReadFile(cfg.CombinerCachePath)
		if err != nil {
			return nil, err
		}

		var m map[string]cache.Item

		err = json.Unmarshal(d, &m)
		if err != nil {
			return nil, err
		}

		alertCache = cache.NewFrom(graphql.DefaultLastNMinutes*2, time.Minute, m)
	} else {
		alertCache = cache.New(graphql.DefaultLastNMinutes*2, time.Minute)
	}

	bf := &combinerFeed{
		maxAlertAge:      time.Minute * 20,
		ctx:              ctx,
		client:           ac,
		rateLimit:        cfg.RateLimit,
		alertCh:          alerts,
		botSubscriptions: []*CombinerBotSubscription{},
		cfg:              cfg,
		alertCache:       alertCache,
	}

	return bf, nil
}
