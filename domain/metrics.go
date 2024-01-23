package domain

func IsMetricAllowed(metric string) bool {
	_, ok := allMetrics[metric]
	return ok
}

// Metric names that are allowed for processing by Forta
const (
	MetricDockerEventPull                = "docker.image.pull"
	MetricDockerEventContainerCreate     = "docker.container.create"
	MetricDockerEventContainerDestroy    = "docker.container.destroy"
	MetricDockerEventContainerConnect    = "docker.container.connect"
	MetricDockerEventContainerDisconnect = "docker.container.disconnect"
	MetricDockerEventNetworkCreate       = "docker.network.create"
	MetricDockerEventNetworkDestroy      = "docker.network.destroy"
	MetricDockerEventNetworkConnect      = "docker.network.connect"
	MetricDockerEventNetworkDisconnect   = "docker.network.disconnect"
	MetricDockerResourcesCPU             = "docker.resources.cpu"
	MetricDockerResourcesMemory          = "docker.resources.memory"
	MetricDockerResourcesNetworkSent     = "docker.resources.network.sent"
	MetricDockerResourcesNetworkReceive  = "docker.resources.network.receive"
)

var (
	allMetrics = map[string]interface{}{
		MetricDockerEventPull:                nil,
		MetricDockerEventContainerCreate:     nil,
		MetricDockerEventContainerDestroy:    nil,
		MetricDockerEventContainerConnect:    nil,
		MetricDockerEventContainerDisconnect: nil,
		MetricDockerEventNetworkCreate:       nil,
		MetricDockerEventNetworkDestroy:      nil,
		MetricDockerEventNetworkConnect:      nil,
		MetricDockerEventNetworkDisconnect:   nil,
		MetricDockerResourcesCPU:             nil,
		MetricDockerResourcesMemory:          nil,
		MetricDockerResourcesNetworkSent:     nil,
		MetricDockerResourcesNetworkReceive:  nil,
	}
)
