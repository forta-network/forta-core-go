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

var (
	ErrCombinerStopReached   = fmt.Errorf("combiner stop reached")
	DefaultRatelimitDuration = time.Minute
	ErrAccessControlDenied   = fmt.Errorf("query was not granted by the access control service")
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

	botSubscriptions []*domain.CombinerBotSubscription
	botsMu           sync.RWMutex
	// bloom filter could be a better choice but the miss rate was too high
	// and we don't know the expected item count
	alertCache *cache.Cache

	handlers    []cfHandler
	handlersMu  sync.Mutex
	cfg         CombinerFeedConfig
	maxAlertAge time.Duration
}

func (cf *combinerFeed) Subscriptions() []*domain.CombinerBotSubscription {
	cf.botsMu.RLock()
	defer cf.botsMu.RUnlock()

	return cf.botSubscriptions
}

func (cf *combinerFeed) AddSubscription(subscription *domain.CombinerBotSubscription) error {
	if subscription == nil || subscription.Subscription == nil || subscription.Subscriber == nil {
		return fmt.Errorf("nil subscription data")
	}
	
	// subscriptions should be bot <-> bot
	if subscription.Subscription.BotId == "" {
		return fmt.Errorf("subscription must have valid bot id")
	}

	cf.botsMu.Lock()
	defer cf.botsMu.Unlock()

	for _, s := range cf.botSubscriptions {
		// existing subscriptions must not be added to prevent duplicate graphql queries, the error however is usually safe to ignore.
		if isSameSubscription(s, subscription) {
			return fmt.Errorf("incoming subscription already exists")
		}
	}

	cf.botSubscriptions = append(cf.botSubscriptions, subscription)
	return nil
}

func (cf *combinerFeed) RemoveSubscription(subscription *domain.CombinerBotSubscription) {
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
	logger := log.WithField("component", "combinerFeed")
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
			logger = logger.WithFields(
				log.Fields{
					"subscriberBotId": subscription.Subscriber.BotID,
					"subscribedBotId": subscription.Subscription.BotId,
				},
			)
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
	ctx context.Context, alertHandlers []cfHandler, subscription *domain.CombinerBotSubscription, createdSince int64,
	createdBefore int64,
) error {
	lg := log.WithFields(
		log.Fields{
			"subscriberBotId":    subscription.Subscriber.BotID,
			"subscriberBotOwner": subscription.Subscriber.BotOwner,
			"subscribedTo":       subscription.Subscription.BotId,
		},
	)

	var alerts []*protocol.AlertEvent

	bo := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)

	err := backoff.Retry(
		func() error {
			var cErr error
			// subscriber information must be passed as headers to the public-api, to comply with fee subscriptions.
			authHeaders := subscriberInfoToHeaders(subscription.Subscriber)

			// convert subscription to graphql input and fetch alerts
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
				authHeaders,
			)
			if cErr != nil {
				lg.WithError(cErr).Warn("error retrieving alerts")

				// don't retry unauthorized errors
				if strings.Contains(cErr.Error(), ErrAccessControlDenied.Error()) {
					return backoff.Permanent(cErr)
				}

				// it is safe to return nil on context deadlines, no need for error handling.
				if !errors.Is(cErr, context.DeadlineExceeded) {
					return cErr
				}
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

		// to prevent duplicate results, processed alerts must be checked/cached.
		if _, exists := cf.alertCache.Get(alertCacheEncode(subscription.Subscriber.BotID, alert.Alert.Hash)); exists {
			continue
		}

		// set incoming alert/subscriber pair to cache
		cf.alertCache.Set(
			alertCacheEncode(subscription.Subscriber.BotID, alert.Alert.Hash), struct{}{}, cache.DefaultExpiration,
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
			Subscriber: subscription.Subscriber,
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

// subscriberInfoToHeaders converts subscriber information to map[string]string
func subscriberInfoToHeaders(subscriber *domain.Subscriber) map[string]string {
	return map[string]string{
		"bot-id":    subscriber.BotID,
		"bot-owner": subscriber.BotOwner,
	}
}

// isSameSubscription checks for field-for-field equality for subscriptions.
func isSameSubscription(a, b *domain.CombinerBotSubscription) bool {
	if a == nil || b == nil {
		return false
	}

	if a.Subscriber == nil || b.Subscriber == nil {
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

	if !strings.EqualFold(strings.Join(a.Subscription.AlertIds, ","), strings.Join(b.Subscription.AlertIds, ",")) {
		return false
	}

	// Since the protocol enforces subscription fees, subscriptions from 2 different bots or bot owners can not
	// be treated as same, because one can fail while the other succeeds.
	if a.Subscriber.BotID != b.Subscriber.BotID {
		return false
	}

	if a.Subscriber.BotOwner != b.Subscriber.BotOwner {
		return false
	}

	return true
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
	url := fmt.Sprintf("%s/graphql", cfg.APIUrl)
	ac := graphql.NewClient(url)
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
			_ = os.RemoveAll(cfg.CombinerCachePath)
			return nil, fmt.Errorf("malformed combiner cache: %v", err)
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
		botSubscriptions: []*domain.CombinerBotSubscription{},
		cfg:              cfg,
		alertCache:       alertCache,
	}

	return bf, nil
}

// alertCacheEncode must encode alerts to prevent missing subscriptions to the same target bot
// from several deployed bots
func alertCacheEncode(subscriberBotID, alertHash string) string {
	return fmt.Sprintf("%s|%s", subscriberBotID, alertHash)
}