package domain

func IsMetricAllowed(metric string) bool {
	_, ok := allMetrics[metric]
	return ok
}

// Metric names that are allowed for processing by Forta
const (
	MetricClientDial  = "agent.client.dial"
	MetricClientClose = "agent.client.close"

	MetricStatusRunning     = "agent.status.running"
	MetricStatusAttached    = "agent.status.attached"
	MetricStatusInitialized = "agent.status.initialized"
	MetricStatusStopping    = "agent.status.stopping"
	MetricStatusActive      = "agent.status.active"
	MetricStatusInactive    = "agent.status.inactive"

	MetricActionUpdate      = "agent.action.update"
	MetricActionRestart     = "agent.action.restart"
	MetricActionSubscribe   = "agent.action.subscribe"
	MetricActionUnsubscribe = "agent.action.unsubscribe"

	MetricFailurePull               = "agent.failure.pull"
	MetricFailureLaunch             = "agent.failure.launch"
	MetricFailureStop               = "agent.failure.stop"
	MetricFailureDial               = "agent.failure.dial"
	MetricFailureInitialize         = "agent.failure.initialize"
	MetricFailureInitializeResponse = "agent.failure.initialize.response"
	MetricFailureInitializeValidate = "agent.failure.initialize.validate"
	MetricFailureTooManyErrs        = "agent.failure.too-many-errs"

	MetricHealthCheckAttempt = "agent.health.attempt"
	MetricHealthCheckSuccess = "agent.health.success"
	MetricHealthCheckError   = "agent.health.error"

	MetricFinding       = "finding"
	MetricTxRequest     = "tx.request"
	MetricTxLatency     = "tx.latency"
	MetricTxError       = "tx.error"
	MetricTxSuccess     = "tx.success"
	MetricTxDrop        = "tx.drop"
	MetricTxBlockAge    = "tx.block.age"
	MetricTxEventAge    = "tx.event.age"
	MetricBlockBlockAge = "block.block.age"
	MetricBlockEventAge = "block.event.age"
	MetricBlockRequest  = "block.request"
	MetricBlockLatency  = "block.latency"
	MetricBlockError    = "block.error"
	MetricBlockSuccess  = "block.success"
	MetricBlockDrop     = "block.drop"

	MetricJSONRPCLatency          = "jsonrpc.latency"
	MetricJSONRPCRequest          = "jsonrpc.request"
	MetricJSONRPCSuccess          = "jsonrpc.success"
	MetricJSONRPCThrottled        = "jsonrpc.throttled"
	MetricPublicAPIProxyLatency   = "publicapi.latency"
	MetricPublicAPIProxyRequest   = "publicapi.request"
	MetricPublicAPIProxySuccess   = "publicapi.success"
	MetricPublicAPIProxyThrottled = "publicapi.throttled"
	MetricFindingsDropped         = "findings.dropped"
	MetricCombinerRequest         = "combiner.request"
	MetricCombinerLatency         = "combiner.latency"
	MetricCombinerError           = "combiner.error"
	MetricCombinerSuccess         = "combiner.success"
	MetricCombinerDrop            = "combiner.drop"
)

var (
	allMetrics = map[string]interface{}{
		MetricClientDial:  nil,
		MetricClientClose: nil,

		MetricStatusRunning:     nil,
		MetricStatusAttached:    nil,
		MetricStatusInitialized: nil,
		MetricStatusStopping:    nil,
		MetricStatusActive:      nil,
		MetricStatusInactive:    nil,

		MetricActionUpdate:      nil,
		MetricActionRestart:     nil,
		MetricActionSubscribe:   nil,
		MetricActionUnsubscribe: nil,

		MetricFailurePull:               nil,
		MetricFailureLaunch:             nil,
		MetricFailureStop:               nil,
		MetricFailureDial:               nil,
		MetricFailureInitialize:         nil,
		MetricFailureInitializeResponse: nil,
		MetricFailureInitializeValidate: nil,
		MetricFailureTooManyErrs:        nil,

		MetricHealthCheckAttempt: nil,
		MetricHealthCheckSuccess: nil,
		MetricHealthCheckError:   nil,

		MetricFinding:       nil,
		MetricTxRequest:     nil,
		MetricTxLatency:     nil,
		MetricTxError:       nil,
		MetricTxSuccess:     nil,
		MetricTxDrop:        nil,
		MetricTxBlockAge:    nil,
		MetricTxEventAge:    nil,
		MetricBlockBlockAge: nil,
		MetricBlockEventAge: nil,
		MetricBlockRequest:  nil,
		MetricBlockLatency:  nil,
		MetricBlockError:    nil,
		MetricBlockSuccess:  nil,
		MetricBlockDrop:     nil,

		MetricJSONRPCLatency:          nil,
		MetricJSONRPCRequest:          nil,
		MetricJSONRPCSuccess:          nil,
		MetricJSONRPCThrottled:        nil,
		MetricPublicAPIProxyLatency:   nil,
		MetricPublicAPIProxyRequest:   nil,
		MetricPublicAPIProxySuccess:   nil,
		MetricPublicAPIProxyThrottled: nil,
		MetricFindingsDropped:         nil,
		MetricCombinerRequest:         nil,
		MetricCombinerLatency:         nil,
		MetricCombinerError:           nil,
		MetricCombinerSuccess:         nil,
		MetricCombinerDrop:            nil,
	}
)
