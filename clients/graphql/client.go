package graphql

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/forta-network/forta-core-go/protocol"
)

const (
	DefaultLastNMinutes = time.Minute * 10
	DefaultPageSize     = 1e3
)

type client struct {
	url    string
	client graphql.Client
}

type Client interface {
	GetAlerts(ctx context.Context, input *AlertsInput, headers map[string]string) ([]*protocol.AlertEvent, error)
	GetAlertsBatch(ctx context.Context, input []*AlertsInput, headers map[string]string) ([]*protocol.AlertEvent, error)
}

func NewClient(url string) Client {
	c := graphql.NewClient(url, nil)

	return &client{url: url, client: c}
}

// Deprecated: Use GetAlertsBatch instead.
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

	// iterate until there are no more alerts to retrieve
	for {
		response, err := fetchAlerts(ctx, ac.url, input, headers)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch alerts: %v", err)
		}

		alerts = append(alerts, response.Alerts.ToAlertEvents()...)

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
) (*GetAlertsResponse, error) {
	req := &graphql.Request{
		OpName:    "getAlerts",
		Query:     getAlertsOperation,
		Variables: __getAlertsInput{Input: input},
	}
	var err error

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

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Accept-Encoding", "gzip")

	queryTime := time.Now().Truncate(time.Minute).UnixMilli()
	httpReq.Header.Set("Forta-Query-Timestamp", fmt.Sprintf("%d", queryTime))

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

	// Check if the response is compressed with gzip
	var respBodyReader = httpResp.Body
	if strings.Contains(httpResp.Header.Get("Content-Encoding"), "gzip") {
		respBodyReader, err = gzip.NewReader(httpResp.Body)
		if err != nil {
			return nil, err
		}
		defer respBodyReader.Close()
	}

	// Parse response
	respBody, err := io.ReadAll(respBodyReader)
	if err != nil {
		return nil, err
	}

	resp, data, err := parseResponse(respBody)
	if err != nil {
		return nil, err
	}

	if len(resp.Errors) > 0 {
		return nil, resp.Errors
	}

	return data, err
}

func parseResponse(responseBody []byte) (*graphql.Response, *GetAlertsResponse, error) {
	var data GetAlertsResponse
	resp := &graphql.Response{Data: &data}

	err := json.Unmarshal(responseBody, resp)
	if err != nil {
		return nil, nil, err
	}

	return resp, &data, err
}

// GetAlertsBatch is a method that retrieves alerts in batches using pagination.
// It takes a context, a slice of AlertsInput, and a map of headers as parameters.
// It returns a slice of AlertEvent and an error.
//
// The method pre-processes the inputs by assigning default values to some fields if they are not provided.
//
// It then iterates until there are no more pagination inputs to query. In each iteration, it calls the fetchAlertsBatch function to fetch alerts based on the inputs and headers.
// If an error occurs during fetching, the method returns nil and the error.
//
// After fetching alerts, the method calls the paginateBatch function to paginate the inputs and the response. It assigns the new inputs and the alert page to variables.
// If an error occurs during pagination, the method returns nil and the error.
//
// Finally, the method appends the alert page to the alerts slice and repeats the iteration until there are no more pagination inputs to query.
// It then returns the alerts slice and nil.
func (ac *client) GetAlertsBatch(ctx context.Context, inputs []*AlertsInput, headers map[string]string) ([]*protocol.AlertEvent, error) {
	// pre-process inputs
	for _, input := range inputs {
		if input.BlockSortDirection == "" {
			input.BlockSortDirection = SortDesc
		}

		// have a default of 10m
		if input.CreatedSince == 0 {
			input.CreatedSince = uint(DefaultLastNMinutes.Milliseconds())
		}

		if input.First == 0 {
			input.First = DefaultPageSize
		}
	}

	var alerts []*protocol.AlertEvent

	// iterate until there are no more pagination inputs to query
	for len(inputs) > 0 {
		response, err := fetchAlertsBatch(ctx, ac.url, inputs, headers)
		if err != nil {
			return nil, err
		}

		var alertPage []*protocol.AlertEvent
		inputs, alertPage, err = paginateBatch(inputs, response)
		if err != nil {
			return nil, err
		}

		alerts = append(alerts, alertPage...)
	}

	return alerts, nil
}

func parseBatchResponse(responseBody []byte) *graphql.Response {
	var data BatchGetAlertsResponse
	resp := &graphql.Response{Data: &data}

	err := json.Unmarshal(responseBody, resp)
	if err != nil {
		return nil
	}

	return resp
}
