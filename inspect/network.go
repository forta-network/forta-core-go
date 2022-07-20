package inspect

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-multierror"
	"github.com/showwin/speedtest-go/speedtest"
	log "github.com/sirupsen/logrus"
)

const (
	MetricNetworkOutboundAccess = "network.access.outbound"
	MetricNetworkDownloadSpeed  = "network.speed.download"
	MetricNetworkUploadSpeed    = "network.speed.upload"
)

// NetworkInspector is an inspector implementation.
type NetworkInspector struct{}

var _ Inspector = &NetworkInspector{}

// Name returns the name of the inspector.
func (ni *NetworkInspector) Name() string {
	return "network"
}

// Inspect inspects network access and quality.
func (ni *NetworkInspector) Inspect(ctx context.Context, inspectionCfg InspectionConfig) (results *InspectionResults, resultErr error) {
	results = NewInspectionResults()

	err := sendOutboundRequest(ctx)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
		results.Metrics[MetricNetworkOutboundAccess] = ResultFailure
	} else {
		results.Metrics[MetricNetworkOutboundAccess] = ResultSuccess
	}

	downloadSpeed, uploadSpeed, err := speedTest()
	if err != nil {
		resultErr = multierror.Append(resultErr, fmt.Errorf("could not test speed: %w", err))
	}

	results.Metrics[MetricNetworkDownloadSpeed] = downloadSpeed
	results.Metrics[MetricNetworkUploadSpeed] = uploadSpeed

	return
}

// speedTest returns download speed, upload speed and error.
func speedTest() (float64, float64, error) {
	speedTest := speedtest.New()
	users, err := speedTest.FetchUserInfo()
	if err != nil {
		return 0, 0, err
	}

	servers, err := speedTest.FetchServers(users)
	if err != nil {
		return 0, 0, err
	}

	if len(servers) == 0 {
		return 0, 0, fmt.Errorf("no available servers")
	}

	// just test with a single server
	server := servers[0]

	err = server.PingTest()
	if err != nil {
		return 0, 0, err
	}

	logger := log.WithField("inspector", "network")

	logger.Info("started download test")
	err = server.DownloadTest(false)
	if err != nil {
		return 0, 0, err
	}
	logger.Info("finished download test")

	logger.Info("started upload test")
	err = server.UploadTest(false)
	if err != nil {
		return 0, 0, err
	}
	logger.Info("finished upload test")

	return server.DLSpeed, server.ULSpeed, nil
}

func sendOutboundRequest(ctx context.Context) error {
	const url = "https://api.forta.network"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	return nil
}
