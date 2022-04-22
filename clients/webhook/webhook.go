package webhook

import (
	"fmt"
	"net/url"

	"github.com/forta-network/forta-core-go/clients/webhook/client"
	"github.com/forta-network/forta-core-go/clients/webhook/client/operations"
	"github.com/go-openapi/runtime"
)

// AlertWebhookClient makes webhook requests.
type AlertWebhookClient interface {
	SendAlerts(params *operations.SendAlertsParams, opts ...operations.ClientOption) (*operations.SendAlertsOK, error)
}

type alertWebhookClient struct {
	operations.ClientService
}

func (awc *alertWebhookClient) SendAlerts(params *operations.SendAlertsParams, opts ...operations.ClientOption) (*operations.SendAlertsOK, error) {
	return awc.ClientService.SendAlerts(params, operations.ClientOption(func(co *runtime.ClientOperation) {
		co.PathPattern = "/" // override /alerts
	}))
}

// NewAlertWebhookClient creates a new webhook client to make requests to '/'.
func NewAlertWebhookClient(dest string) (AlertWebhookClient, error) {
	u, err := url.Parse(dest)
	if err != nil {
		return nil, fmt.Errorf("invalid webhook url: %v", err)
	}
	client := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host:     u.Host,
		BasePath: u.Path,
		Schemes:  []string{u.Scheme},
	}).Operations
	return &alertWebhookClient{
		ClientService: client,
	}, nil
}
