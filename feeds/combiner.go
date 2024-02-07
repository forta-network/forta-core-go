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
	log "github.com/sirupsen/logrus"
)

var (
	ErrCombinerStopReached   = fmt.Errorf("combiner stop reached")
	DefaultRatelimitDuration = time.Minute * 2
	DefaultLookbackPeriod    = time.Minute * 5
	ErrUnauthorized          = fmt.Errorf("query was not granted")
	ErrBadRequest            = fmt.Errorf("bad public api request")
)

const (
	DefaultBatchSize = 10
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
	client  graphql.Client

	botSubscriptions []*domain.CombinerBotSubscription
	botsMu           sync.RWMutex

	combinerCache *combinerCache

	handlers    []cfHandler
	handlersMu  sync.Mutex
	cfg         CombinerFeedConfig
	maxAlertAge time.Duration
	batchSize   int
}

func (cf *combinerFeed) Subscriptions() []*domain.CombinerBotSubscription {
	cf.botsMu.RLock()
	defer cf.botsMu.RUnlock()

	return cf.botSubscriptions
}

// AddSubscription adds the given CombinerBotSubscription to the list of subscriptions for the combiner feed.
// Returns an error if the subscription data is nil or if the subscription doesn't have a valid bot id.
// Also checks if the subscription already exists in the list of subscriptions to prevent duplicate graphql queries.
// This method is thread-safe, as it acquires a lock on the botSubscriptions mutex before accessing or modifying the botSubscriptions slice.
// The error returned when an existing subscription is found is usually safe to ignore, as it indicates that the subscription was already present and therefore not added again.
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
		if s.Equal(subscription) {
			return fmt.Errorf("incoming subscription already exists")
		}
	}

	cf.botSubscriptions = append(cf.botSubscriptions, subscription)
	return nil
}

// RemoveSubscription removes the given CombinerBotSubscription from the list of subscriptions for the combiner feed.
// This method is thread-safe, as it acquires a lock on the botSubscriptions mutex before accessing or modifying the botSubscriptions slice.
// If the subscription is not found in the list of subscriptions, this method does nothing.
func (cf *combinerFeed) RemoveSubscription(subscription *domain.CombinerBotSubscription) {
	cf.botsMu.Lock()
	defer cf.botsMu.Unlock()

	for i, s := range cf.botSubscriptions {
		if s.Equal(subscription) {
			cf.botSubscriptions = append(
				cf.botSubscriptions[:i], cf.botSubscriptions[i+1:]...,
			)
		}
	}
}

// RegisterHandler registers the given alert handler function to receive alert events from the combiner feed.
// The alertHandler function takes an AlertEvent pointer as input and returns an error.
// Returns a channel that the caller can use to receive errors from the alert handler function.
// This method is thread-safe, as it acquires a lock on the handlers mutex before accessing or modifying the handlers slice.
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
	QueryInterval     uint64 // query interval in milliseconds
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

// forEachAlert retrieves alerts for each subscription, and processes them by calling the alert handlers passed in as an argument.
// It waits for the rate limit, if any, and saves the alert cache to a persistent file, if configured.
// This method is thread-safe, as it acquires a lock on the subscriptions mutex before accessing or modifying them,
// and on the combinerCache mutex before accessing or modifying it.
func (cf *combinerFeed) forEachAlert(alertHandlers []cfHandler) error {
	// Set up logger and firstRun flag
	logger := log.WithField("component", "combinerFeed")
	firstRun := true

	// Loop until context is done or an error occurs
	for {
		// Check if context is done and return error if so
		if cf.ctx.Err() != nil {
			return cf.ctx.Err()
		}

		// Wait for the rate limit, if any
		if cf.rateLimit != nil {
			if !firstRun {
				<-cf.rateLimit.C
			}
			firstRun = false
		}

		// Skip query if there are no alert subscriptions
		if len(cf.Subscriptions()) == 0 {
			continue
		}

		// Set lower and upper bounds for alert creation date
		lowerBound := DefaultLookbackPeriod
		upperBound := int64(0)

		// Query all subscriptions and process alerts
		cf.handleSubscriptions(alertHandlers, cf.Subscriptions(), lowerBound, upperBound, logger)

		// Save alert cache to persistent file, if configured
		if cf.cfg.CombinerCachePath != "" {
			if err := cf.combinerCache.DumpToFile(cf.cfg.CombinerCachePath); err != nil {
				log.Panic(err)
			}
		}
	}
}

func (cf *combinerFeed) handleSubscriptions(alertHandlers []cfHandler, subscriptions []*domain.CombinerBotSubscription, lowerBound time.Duration, upperBound int64, logger *log.Entry) {
	// create a lookup map to batch per subscriber (ie. bot)
	subscriberBatchMap := make(map[domain.Subscriber][]*protocol.CombinerBotSubscription)
	for _, subscription := range subscriptions {
		subscriber := *subscription.Subscriber
		subscriberBatchMap[subscriber] = append(subscriberBatchMap[subscriber], subscription.Subscription)
	}

	// handle subscriptions in batches
	for subscriber, botSubscriptions := range subscriberBatchMap {
		logger = logger.WithFields(
			log.Fields{
				"subscriberBotId":    subscriber.BotID,
				"subscriberBotOwner": subscriber.BotOwner,
				"subscriberBotImage": subscriber.BotImage,
			})

		var subscriberAlerts []*protocol.AlertEvent
		// iterate over batches and handle
		for i := 0; i < len(botSubscriptions); {
			currentBatchSize := cf.batchSize
			for {
				// Determine the end of the current batch
				end := i + currentBatchSize
				if end > len(botSubscriptions) {
					end = len(botSubscriptions)
				}

				// Create a batch
				batch := botSubscriptions[i:end]

				alerts, err := cf.fetchAlertsBatch(cf.ctx, logger, &subscriber, batch, lowerBound.Milliseconds(), upperBound)
				if err != nil {
					if currentBatchSize > 1 {
						// Reduce batch size and retry
						currentBatchSize /= 2
						logger.WithError(err).Warnf("error querying batch, reducing size to %d and retrying", currentBatchSize)
						continue
					} else {
						// batch size at minimum, do not retry
						logger.WithError(err).Warn("failed to fetch alerts")
					}
				}

				subscriberAlerts = append(subscriberAlerts, alerts...)
				i += currentBatchSize
				break
			}
		}
		cf.processAlerts(cf.ctx, logger, subscriber, subscriberAlerts, alertHandlers)
	}
}

func (cf *combinerFeed) fetchAlertsBatch(ctx context.Context, logger *log.Entry, subscriber *domain.Subscriber,
	subscriptions []*protocol.CombinerBotSubscription, createdSince int64, createdBefore int64) ([]*protocol.AlertEvent,
	error) {
	var alerts []*protocol.AlertEvent

	// construct auth headers for the subscriber
	authHeaders := subscriberInfoToHeaders(subscriber)

	inputs := subscriptionsToAlertInputs(subscriptions, createdSince, createdBefore)

	// call the GraphQL client's GetAlerts method with retries
	err := cf.retryWithBackoff(
		ctx, func() error {
			var cErr error
			alerts, cErr = cf.client.GetAlertsBatch(cf.ctx, inputs, authHeaders)
			if cErr != nil {
				logger.WithError(cErr).Warn("error retrieving alerts")

				// make any error is non-retryable
				return backoff.Permanent(cErr)
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return alerts, nil
}

func subscriptionsToAlertInputs(subscriptions []*protocol.CombinerBotSubscription, createdSince int64, createdBefore int64) []*graphql.AlertsInput {
	inputs := make([]*graphql.AlertsInput, len(subscriptions))
	for i, subscription := range subscriptions {
		// construct the graphql.AlertsInput object based on the subscription data
		inputs[i] = &graphql.AlertsInput{
			Bots:          []string{subscription.BotId},
			CreatedSince:  uint(createdSince),
			CreatedBefore: uint(createdBefore),
			AlertIds:      subscription.AlertIds,
			AlertId:       subscription.AlertId,
			ChainId:       uint(subscription.ChainId),
		}
	}
	return inputs
}

// processAlerts processes a slice of alerts by filtering out those that are too old or have already been processed and then passing
// the remaining alerts to the alert handlers passed in as an argument.
// It uses a cache to prevent duplicate processing of alerts and creates an AlertEvent object to pass to each alert handler.
// It is thread-safe as it acquires a lock on the combinerCache mutex before accessing or modifying it.
func (cf *combinerFeed) processAlerts(_ context.Context, logger *log.Entry, subscriber domain.Subscriber,
	alerts []*protocol.AlertEvent,
	alertHandlers []cfHandler) {
	for _, alert := range alerts {
		if tooOld, age := alertIsTooOld(alert, cf.maxAlertAge); tooOld {
			logger.WithFields(log.Fields{"age": age, "alert": alert.Alert.Hash}).Warnf(
				"alert is older than %v - setting current alert iterator head to now", cf.maxAlertAge,
			)
			continue
		}

		// check if the alert has already been processed
		if cf.combinerCache.Exists(subscriber, alert) {
			continue
		}

		// add the alert to the cache to prevent duplicate processing
		cf.combinerCache.Set(subscriber, alert)

		// create an AlertEvent object to pass to each alert handler
		alertCA, err := time.Parse(time.RFC3339, alert.Alert.CreatedAt)
		if err != nil {
			// safe to continue processing rest of alerts - alert specific problem
			logger.WithError(err).Warn("failed to process alert")
			continue
		}

		evt := &domain.AlertEvent{
			Event: alert,
			Timestamps: &domain.TrackingTimestamps{
				Feed:        time.Now().UTC(),
				SourceAlert: alertCA,
			},
			Subscriber: subscriber,
		}

		// call each alert handler with the AlertEvent object
		for _, alertHandler := range alertHandlers {
			if err := alertHandler.Handler(evt); err != nil {
				// safe to continue processing rest of alerts - alert specific problem
				logger.WithError(err).Warn("error executing alert handler")
				continue
			}
		}
	}
}

func (cf *combinerFeed) retryWithBackoff(ctx context.Context, f func() error) error {
	bo := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)

	return backoff.Retry(
		func() error {
			err := f()
			if err != nil {
				log.WithError(err).Warn("error executing function, will retry...")
			}
			return err
		}, bo,
	)
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

	return NewCombinerFeedWithClient(ctx, cfg, ac)
}

func NewCombinerFeedWithClient(ctx context.Context, cfg CombinerFeedConfig, client graphql.Client) (AlertFeed, error) {
	alerts := make(chan *domain.AlertEvent, 10)

	c, err := newCombinerCache(cfg.CombinerCachePath)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize combiner cache: %v", err)
	}

	rateLimit := time.NewTicker(DefaultRatelimitDuration)
	// Use configured query interval if exists, max interval is set to the default interval to prevent protocol-wide delays.
	if cfg.QueryInterval > 0 && cfg.QueryInterval < uint64(DefaultRatelimitDuration.Milliseconds()) {
		rateLimit = time.NewTicker(time.Millisecond * time.Duration(cfg.QueryInterval))
	}

	bf := &combinerFeed{
		maxAlertAge:      time.Minute * 20,
		ctx:              ctx,
		client:           client,
		rateLimit:        rateLimit,
		alertCh:          alerts,
		botSubscriptions: []*domain.CombinerBotSubscription{},
		cfg:              cfg,
		combinerCache:    c,
		batchSize:        DefaultBatchSize,
	}

	return bf, nil
}
