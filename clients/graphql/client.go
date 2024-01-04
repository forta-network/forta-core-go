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
	"github.com/vektah/gqlparser/v2/gqlerror"
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

		response := responses[defaultInputAlias]
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

func (ac *client) GetAlertsBatch(ctx context.Context, inputs []*AlertsInput,
	headers map[string]string) ([]*protocol.AlertEvent, error) {
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

	// iterate until there are no more alerts to retrieve
	for {
		response, err := fetchAlertsBatch(ctx, ac.url, inputs, headers)
		// TODO: ignore partial failures for now
		if len(err) > 0 {
			return nil, fmt.Errorf("failed to fetch alerts: %v", err)
		}

		// iterator to overwrite inputs
		i := -1
		shouldBreak := true
		for _, responseItem := range response {
			i++
			alerts = append(alerts, responseItem.ToAlertEvents()...)

			// check if there are more alerts
			if !responseItem.PageInfo.HasNextPage {
				continue
			}

			// check if response is ok
			if responseItem.PageInfo.EndCursor == nil {
				continue
			}

			inputs[i].After = &AlertEndCursorInput{
				AlertId:     responseItem.PageInfo.EndCursor.AlertId,
				BlockNumber: responseItem.PageInfo.EndCursor.BlockNumber,
			}
			shouldBreak = false
		}

		if shouldBreak {
			break
		}
	}

	return alerts, nil
}

func fetchAlertsBatch(ctx context.Context, client string, inputs []*AlertsInput,
	headers map[string]string) (BatchGetAlertsResponse, gqlerror.List) {
	query, variables := createGetAlertsQuery(inputs)
	req := &graphql.Request{
		OpName:    "getAlerts",
		Query:     query,
		Variables: variables,
	}

	respBody, err := makeRequest(ctx, client, req, headers)
	if err != nil {
		return nil, gqlerror.List{gqlerror.Errorf(err.Error())}
	}

	resp := parseBatchResponse(respBody)
	if err != nil {
		return nil, gqlerror.List{gqlerror.Errorf(err.Error())}
	}

	return *resp.Data.(*BatchGetAlertsResponse), resp.Errors
}

func makeRequest(ctx context.Context, client string, req *graphql.Request, headers map[string]string) ([]byte, error) {
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
	return respBody, err
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
