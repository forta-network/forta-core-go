package graphql

import (
	"context"
	"strconv"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/forta-network/forta-core-go/protocol"
)

type GetAlertsResponse getAlertsResponse

type Client struct {
	url    string
	client graphql.Client
}

func NewClient(url string) *Client {
	client := graphql.NewClient(url, nil)

	return &Client{url: url, client: client}
}

func (ac *Client) GetAlerts(
	ctx context.Context, input *AlertsInput,
) ([]*protocol.Alert, error) {
	if input.BlockSortDirection == "" {
		input.BlockSortDirection = SortDesc
	}

	// have a default of 1m
	if input.CreatedSince == 0 {
		input.CreatedSince = uint(time.Minute.Milliseconds())
	}

	if input.First == 0 {
		input.First = 500
	}
	var alerts []*protocol.Alert

	for {
		response, err := getAlerts(ctx, ac.client, input)
		if err != nil {
			return nil, err
		}

		alerts = append(alerts, response.ToProto()...)

		// check if there are more alerts
		if !response.Alerts.PageInfo.HasNextPage {
			break
		}
		// check if response is ok
		if response.Alerts.PageInfo.EndCursor == nil {
			break
		}
		input.After = &AlertEndCursorInput{
			AlertId:     response.Alerts.PageInfo.EndCursor.AlertId,
			BlockNumber: response.Alerts.PageInfo.EndCursor.BlockNumber,
		}
	}

	return alerts, nil
}

func (v getAlertsResponse) ToProto() []*protocol.Alert {
	return ToProto(v)
}
func ToProto(response getAlertsResponse) []*protocol.Alert {
	var resp []*protocol.Alert
	for _, alert := range response.Alerts.Alerts {
		blockNum := strconv.FormatUint(uint64(alert.Source.Block.Number), 10)

		resp = append(
			resp, &protocol.Alert{
				Id: alert.AlertId,
				Finding: &protocol.Finding{
					Protocol:    alert.Protocol,
					Severity:    protocol.Finding_Severity(protocol.Finding_Severity_value[alert.Severity]),
					Metadata:    alert.Metadata,
					Type:        protocol.Finding_FindingType(protocol.Finding_FindingType_value[alert.FindingType]),
					AlertId:     alert.AlertId,
					Name:        alert.Name,
					Description: alert.Description,
					EverestId:   "",
					Private:     false,
					Addresses:   alert.Addresses,
					Indicators:  nil,
				},
				Timestamp: alert.CreatedAt,
				Metadata:  alert.Metadata,
				Agent: &protocol.AgentInfo{
					Image:     alert.Source.Bot.Image,
					ImageHash: "",
					Id:        alert.Source.Bot.Image,
					IsTest:    false,
					Manifest:  "",
				},
				Tags: nil,
				Scanner: &protocol.ScannerInfo{
					Address: "",
				},
				Timestamps: &protocol.TrackingTimestamps{
					Block:       blockNum,
					Feed:        "",
					BotRequest:  "",
					BotResponse: "",
				},
			},
		)
	}

	return resp
}
