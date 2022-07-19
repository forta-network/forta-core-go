package inspect

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-multierror"
	"github.com/showwin/speedtest-go/speedtest"
)

const (
	MetricContainerNetworkOutboundAccess = "container.network.outboundaccess"
	MetricContainerNetworkDownloadSpeed  = "container.network.downloadspeed"
	MetricContainerNetworkUploadSpeed    = "container.network.uploadspeed"
)

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

// RunNetworkInspection tests network access and quality of the node.
func RunNetworkInspection(ctx context.Context) (result map[string]float64, resultError error) {
	result = make(map[string]float64)

	err := sendOutboundRequest(ctx)
	if err != nil {
		resultError = multierror.Append(resultError, err)
		result[MetricContainerNetworkOutboundAccess] = StateError
	} else {
		result[MetricContainerNetworkOutboundAccess] = StateSuccess
	}

	downloadSpeed, uploadSpeed, err := speedTest()
	if err != nil {
		resultError = multierror.Append(resultError, fmt.Errorf("could not test speed: %w", err))
	}

	result[MetricContainerNetworkDownloadSpeed] = downloadSpeed
	result[MetricContainerNetworkUploadSpeed] = uploadSpeed

	return result, nil
}

// speedTest returns download speed, upload speed and error.
func speedTest() (float64, float64, error) {
	st := speedtest.New()
	users, err := st.FetchUserInfo()
	if err != nil {
		return 0, 0, err
	}

	svs, err := st.FetchServers(users)
	if err != nil {
		return 0, 0, err
	}

	if len(svs) == 0 {
		return 0, 0, fmt.Errorf("no available servers")
	}

	// just test with a single server
	sv := svs[0]

	err = sv.PingTest()
	if err != nil {
		return 0, 0, err
	}

	err = sv.DownloadTest(false)
	if err != nil {
		return 0, 0, err
	}

	err = sv.UploadTest(false)
	if err != nil {
		return 0, 0, err
	}

	return sv.DLSpeed, sv.ULSpeed, nil
}
