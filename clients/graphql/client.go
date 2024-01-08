package graphql

import (
	"context"
	"encoding/json"
	"fmt"
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
		responses, err := fetchAlertsBatch(ctx, ac.url, []*AlertsInput{input}, headers)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch alerts: %v", err)
		}

		resp := *responses.Data.(*BatchGetAlertsResponse)

		response := resp[defaultInputAlias]
		alerts = append(alerts, response.ToAlertEvents()...)

		// check if there are more alerts
		if !response.PageInfo.HasNextPage {
			break
		}

		// check if response is ok
		if response.PageInfo.EndCursor == nil {
			break
		}
		input.After = &AlertEndCursorInput{
			AlertId:     response.PageInfo.EndCursor.AlertId,
			BlockNumber: response.PageInfo.EndCursor.BlockNumber,
		}
	}

	return alerts, nil
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
