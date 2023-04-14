package graphql

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/sirupsen/logrus"
)

const (
	DefaultLastNMinutes = time.Minute * 10
	DefaultPageSize     = 1e3
)

type GetAlertsResponse getAlertsResponse

type client struct {
	url    string
	client graphql.Client
}

type Client interface {
	GetAlerts(ctx context.Context, input *AlertsInput, headers map[string]string) ([]*protocol.AlertEvent, error)
}

func NewClient(url string) Client {
	c := graphql.NewClient(url, nil)

	return &client{url: url, client: c}
}

func (ac *client) GetAlerts(
	ctx context.Context, input *AlertsInput, headers map[string]string,
) ([]*protocol.AlertEvent, error) {
	if input.BlockSortDirection == "" {
		input.BlockSortDirection = SortDesc
	}

	// have a default of 1m
	if input.CreatedSince == 0 {
		input.CreatedSince = uint(DefaultLastNMinutes.Milliseconds())
	}

	if input.First == 0 {
		input.First = DefaultPageSize
	}

	var alerts []*protocol.AlertEvent

	for {
		response, err := fetchAlerts(ctx, ac.url, input, headers)
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

func fetchAlerts(
	ctx context.Context,
	client string,
	input *AlertsInput,
	headers map[string]string,
) (*getAlertsResponse, error) {
	req := &graphql.Request{
		OpName: "getAlerts",
		Query:  getAlerts_Operation,
		Variables: &__getAlertsInput{
			Input: input,
		},
	}
	var err error

	var data getAlertsResponse
	resp := &graphql.Response{Data: &data}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequest(
		http.MethodPost,
		client,
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}

	httpReq = httpReq.WithContext(ctx)

	// set custom headers
	for key, val := range headers {
		httpReq.Header.Set(key, val)
	}

	// execute query
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		var respBody []byte
		respBody, err = io.ReadAll(httpResp.Body)
		if err != nil {
			respBody = []byte(fmt.Sprintf("<unreadable: %v>", err))
		}
		return nil, fmt.Errorf("returned error %v: %s", httpResp.Status, respBody)
	}

	// parse response
	err = json.NewDecoder(httpResp.Body).Decode(resp)
	if err != nil {
		return nil, err
	}
	if len(resp.Errors) > 0 {
		return nil, resp.Errors
	}

	return &data, err
}

func (v *getAlertsResponse) ToProto() []*protocol.AlertEvent {
	return ToProto(v)
}

func ToProto(response *getAlertsResponse) []*protocol.AlertEvent {
	var resp []*protocol.AlertEvent
	for _, alert := range response.Alerts.Alerts {
		resp = append(resp, AlertToProto(alert))
	}

	return resp
}

func AlertToProto(alert *getAlertsAlertsAlertsResponseAlertsAlert) *protocol.AlertEvent {
	contracts := make([]*protocol.AlertEvent_Alert_Contract, len(alert.Contracts))
	projects := make([]*protocol.AlertEvent_Alert_Project, len(alert.Projects))
	labels := make([]*protocol.AlertEvent_Alert_Label, len(alert.Labels))

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

	for i, label := range alert.Labels {
		if label == nil {
			logrus.WithField("alert-id", alert.AlertId).Warn("malformed alert label")
			continue
		}

		labels[i] = &protocol.AlertEvent_Alert_Label{
			Label:      label.Label,
			Confidence: float32(label.Confidence),
			Entity:     label.Entity,
			EntityType: label.EntityType,
			Remove:     label.Remove,
			Metadata:   label.Metadata,
		}
	}

	t, _ := time.Parse(time.RFC3339, alert.Source.Block.Timestamp)
	blockTimestamp := hexutil.EncodeUint64(uint64(t.Unix()))
	t, _ = time.Parse(time.RFC3339, alert.CreatedAt)
	alertTimestamp := hexutil.EncodeUint64(uint64(t.Unix()))

	a := &protocol.AlertEvent{
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
			FindingType:   alert.FindingType,
			RelatedAlerts: alert.RelatedAlerts,
			ChainId:       uint64(alert.ChainId),
			Labels:        labels,
		},
		Timestamps: &protocol.TrackingTimestamps{SourceAlert: alertTimestamp},
	}

	a.Alert.Source = &protocol.AlertEvent_Alert_Source{
		TransactionHash: alert.Source.TransactionHash,
	}

	if alert.Source == nil {
		return a
	}

	// fill source bot
	if alert.Source.Bot != nil {
		a.Alert.Source.Bot = &protocol.AlertEvent_Alert_Bot{
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
		}
	}

	// fill source block
	if alert.Source.Block != nil {
		a.Alert.Source.Block = &protocol.AlertEvent_Alert_Block{
			Number:    uint64(alert.Source.Block.Number),
			Hash:      alert.Source.Block.Hash,
			Timestamp: blockTimestamp,
			ChainId:   uint64(alert.Source.Block.ChainId),
		}
	}

	// fill source alert
	if alert.Source.SourceAlert != nil {
		a.Alert.Source.SourceEvent = &protocol.AlertEvent_Alert_SourceAlertEvent{
			BotId:     alert.Source.SourceAlert.BotId,
			AlertHash: alert.Source.SourceAlert.Hash,
			Timestamp: alert.Source.SourceAlert.Timestamp,
			ChainId:   fmt.Sprintf("%d", alert.Source.SourceAlert.ChainId),
		}
	}

	return a
}
