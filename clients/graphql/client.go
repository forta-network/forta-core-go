package graphql

import (
	"context"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/forta-network/forta-core-go/protocol"
)

const (
	DefaultLastNMinutes = time.Minute * 10
	DefaultQueryLimit   = 1e3
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
) ([]*protocol.AlertEvent, error) {
	if input.BlockSortDirection == "" {
		input.BlockSortDirection = SortDesc
	}

	// have a default of 1m
	if input.CreatedSince == 0 {
		input.CreatedSince = uint(DefaultLastNMinutes.Milliseconds())
	}

	if input.First == 0 {
		input.First = DefaultQueryLimit
	}

	var alerts []*protocol.AlertEvent

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

func (v getAlertsResponse) ToProto() []*protocol.AlertEvent {
	return ToProto(v)
}

func ToProto(response getAlertsResponse) []*protocol.AlertEvent {
	var resp []*protocol.AlertEvent
	for _, alert := range response.Alerts.Alerts {
		contracts := make([]*protocol.AlertEvent_Alert_Contract, len(alert.Contracts))
		projects := make([]*protocol.AlertEvent_Alert_Project, len(alert.Projects))
		for i, contract := range alert.Contracts {
			contracts[i] = &protocol.AlertEvent_Alert_Contract{
				Name:      contract.Name,
				ProjectId: contract.ProjectId,
			}
		}
		for i, project := range alert.Projects {
			projects[i] = &protocol.AlertEvent_Alert_Project{
				Id: project.Id,
			}
		}
		resp = append(
			resp, &protocol.AlertEvent{
				Alert: &protocol.AlertEvent_Alert{
					AlertId:       alert.AlertId,
					Addresses:     alert.Addresses,
					Contracts:     contracts,
					CreatedAt:     alert.CreatedAt,
					Description:   alert.Description,
					Hash:          alert.Hash,
					Metadata:      alert.Metadata,
					Name:          alert.Name,
					Projects:      projects,
					ScanNodeCount: int32(alert.ScanNodeCount),
					Severity:      alert.Severity,
					Source: &protocol.AlertEvent_Alert_Source{
						TransactionHash: alert.Source.TransactionHash,
						Bot: &protocol.AlertEvent_Alert_Bot{
							ChainIds:     alert.Source.Bot.ChainIds,
							CreatedAt:    alert.Source.Bot.CreatedAt,
							Description:  alert.Source.Bot.Description,
							Developer:    alert.Source.Bot.Developer,
							DocReference: alert.Source.Bot.DocReference,
							Enabled:      alert.Source.Bot.Enabled,
							Id:           alert.Source.Bot.Id,
							Image:        alert.Source.Bot.Image,
							Name:         alert.Source.Bot.Name,
							Reference:    alert.Source.Bot.Reference,
							Repository:   alert.Source.Bot.Repository,
							Projects:     alert.Source.Bot.Projects,
							ScanNodes:    alert.Source.Bot.ScanNodes,
							Version:      alert.Source.Bot.Version,
						},
						Block: &protocol.AlertEvent_Alert_Block{
							Number:    uint64(alert.Source.Block.Number),
							Hash:      alert.Source.Block.Hash,
							Timestamp: alert.Source.Block.Timestamp,
							ChainId:   uint64(alert.Source.Block.ChainId),
						},
					},
					FindingType: alert.FindingType,
				},
			},
		)
	}

	return resp
}
