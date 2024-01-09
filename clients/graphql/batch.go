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
	log "github.com/sirupsen/logrus"
)

var (
	ErrResponseSizeTooBig = fmt.Errorf("response size too big")
)

// paginateBatch processes the response received from the server and extracts the relevant data.
// It takes an array of AlertsInput objects and a graphql.Response object as input.
// It then iterates over the inputs and extracts the response item based on the alias.
// If there is an error for the input, it logs a warning and continues to the next input.
// If there are more alerts for the input, it adds the input to the list of pagination inputs.
// The function returns the pagination inputs, the extracted alert events, and any encountered error.
func paginateBatch(inputs []*AlertsInput, response *graphql.Response) ([]*AlertsInput,
	[]*protocol.AlertEvent, error) {
	// type-checking response
	if response == nil {
		return nil, nil, fmt.Errorf("nil graphql response")
	}

	batchAlertsResponseUnsafe, ok := response.Data.(*BatchGetAlertsResponse)
	if !ok {
		return nil, nil, fmt.Errorf("invalid pagination response")
	}
	if batchAlertsResponseUnsafe == nil {
		return nil, nil, fmt.Errorf("nil pagination response")
	}

	batchAlertsResponse := *batchAlertsResponseUnsafe

	var pagination []*AlertsInput
	var alerts []*protocol.AlertEvent
	for inputIdx := range inputs {
		alias := idxToResponseAlias(inputIdx)

		logger := log.WithFields(log.Fields{
			"input": idxToInputAlias(inputIdx),
		})

		responseItem, ok := batchAlertsResponse[alias]
		if !ok {
			logger.Warn("no response for input")
			continue
		}

		// check if there is an error for the input
		err := HasError(response.Errors, inputIdx)
		if err != nil {
			logger.WithError(err).Warn("error response for input")
			continue
		}

		alerts = append(alerts, responseItem.ToAlertEvents()...)

		// check if there are more alerts for the input
		if !responseItem.PageInfo.HasNextPage {
			continue
		}

		// check if there are more alerts for the input
		if responseItem.PageInfo.EndCursor == nil {
			continue
		}

		// add input to the list of pagination inputs
		inputs[inputIdx].After = &AlertEndCursorInput{
			AlertId:     responseItem.PageInfo.EndCursor.AlertId,
			BlockNumber: responseItem.PageInfo.EndCursor.BlockNumber,
		}

		pagination = append(pagination, inputs[inputIdx])
	}

	return pagination, alerts, nil
}

// fetchAlertsBatch retrieves the alerts in batches from the server using the provided client, inputs, and headers.
func fetchAlertsBatch(ctx context.Context, client string, inputs []*AlertsInput, headers map[string]string) (*graphql.Response, error) {
	query, variables := createGetAlertsQuery(inputs)
	req := &graphql.Request{
		OpName:    "getAlerts",
		Query:     query,
		Variables: variables,
	}

	respBody, err := makeRequest(ctx, client, req, headers)
	if err != nil {
		return nil, err
	}

	resp := parseBatchResponse(respBody)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// makeRequest sends a GraphQL request to the specified client and returns the response body.
// It takes the context, client URL, request, and headers as input parameters.
// It marshals the request into JSON and creates an HTTP request.
// It sets the custom headers and executes the query with the default HTTP client.
// If the response status code is not OK, it returns an error with the status and response body.
// If the response is gzip compressed, it decompresses the body before parsing.
// It reads the response body and returns it along with any encountered error.
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

	if httpResp.StatusCode == http.StatusInternalServerError {
		return nil, ErrResponseSizeTooBig
	}

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
