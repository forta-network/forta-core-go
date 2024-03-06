package inspect

import (
	"context"
	"fmt"
	"net/http"

	"github.com/forta-network/forta-core-go/utils/httpclient"
	"github.com/hashicorp/go-multierror"
	"github.com/showwin/speedtest-go/speedtest"
	log "github.com/sirupsen/logrus"
)

const (
	IndicatorNetworkOutboundAccess = "network.access.outbound"
	IndicatorNetworkDownloadSpeed  = "network.speed.download"
	IndicatorNetworkUploadSpeed    = "network.speed.upload"
)

// Global flags
var (
	networkIndicators = []string{
		IndicatorNetworkOutboundAccess, IndicatorNetworkDownloadSpeed, IndicatorNetworkUploadSpeed,
	}
)

// NetworkInspector is an inspector implementation.
type NetworkInspector struct{}

// compile time check: it should implement the interface
var _ Inspector = &NetworkInspector{}

// Name returns the name of the inspector.
func (ni *NetworkInspector) Name() string {
	return "network"
}

// Inspect inspects network access and quality.
func (ni *NetworkInspector) Inspect(ctx context.Context, inspectionCfg InspectionConfig) (results *InspectionResults, resultErr error) {
	results = NewInspectionResults()
	results.Indicators = defaultIndicators(networkIndicators)

	err := sendOutboundRequest(ctx)
	if err != nil {
		resultErr = multierror.Append(resultErr, err)
		results.Indicators[IndicatorNetworkOutboundAccess] = ResultFailure
		results.Indicators[IndicatorNetworkDownloadSpeed] = ResultFailure
		results.Indicators[IndicatorNetworkUploadSpeed] = ResultFailure
	} else {
		results.Indicators[IndicatorNetworkOutboundAccess] = ResultSuccess
	}

	// TODO: Disabling until we have a better use case for this.
	// downloadSpeed, uploadSpeed, err := speedTest()
	// if err != nil {
	// 	resultErr = multierror.Append(resultErr, fmt.Errorf("could not test speed: %w", err))
	// 	results.Indicators[IndicatorNetworkDownloadSpeed] = ResultInternalProblem
	// 	results.Indicators[IndicatorNetworkUploadSpeed] = ResultInternalProblem
	// } else {
	// 	results.Indicators[IndicatorNetworkDownloadSpeed] = downloadSpeed
	// 	results.Indicators[IndicatorNetworkUploadSpeed] = uploadSpeed
	// }

	return
}

// SpeedTest returns download speed, upload speed and error.
func SpeedTest() (float64, float64, error) {
	speedTest := speedtest.New()

	servers, err := speedTest.FetchServers()
	if err != nil {
		return 0, 0, err
	}

	if len(servers) == 0 {
		return 0, 0, fmt.Errorf("no available servers")
	}

	// just test with a single server
	server := servers[0]

	err = server.PingTest(nil)
	if err != nil {
		return 0, 0, err
	}

	logger := log.WithField("inspector", "network")

	logger.Info("started download test")
	err = server.DownloadTest()
	if err != nil {
		return 0, 0, err
	}
	logger.Info("finished download test")

	logger.Info("started upload test")
	err = server.UploadTest()
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

	resp, err := httpclient.Default.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	return nil
}
