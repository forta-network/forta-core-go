package graphql

import (
	"fmt"
	"testing"

	"github.com/Khan/genqlient/graphql"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/stretchr/testify/assert"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func TestPaginateBatch(t *testing.T) {
	var alerts0 ast.PathName = "alerts0"
	tests := []struct {
		name               string
		inputs             []*AlertsInput
		response           *graphql.Response
		expectedPagination []*AlertsInput
		expectedAlerts     []*protocol.AlertEvent
		expectedErr        error
	}{
		{
			name:        "Invalid Pagination Response",
			response:    &graphql.Response{},
			expectedErr: fmt.Errorf("invalid pagination response"),
		},
		{
			name:   "Test with two inputs and one error in response",
			inputs: []*AlertsInput{{}, {}},
			response: &graphql.Response{
				Data: &BatchGetAlertsResponse{
					"alerts0": {
						PageInfo: nil,
					},
					"alerts1": {
						PageInfo: &PageInfo{
							HasNextPage: true,
							EndCursor:   &EndCursor{AlertId: "0xaaa"},
						},
					},
				},
				Errors: gqlerror.List{{
					Path: ast.Path{
						alerts0,
					},
					Message: "test error",
				}},
			},
			expectedPagination: []*AlertsInput{
				{
					After: &AlertEndCursorInput{
						AlertId: "0xaaa",
					},
				},
			},
			expectedAlerts: nil,
			expectedErr:    nil,
		},
		// Add more test cases here for other scenarios.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pagination, alerts, err := paginateBatch(tt.inputs, tt.response)
			assert.Equal(t, tt.expectedPagination, pagination)
			assert.Equal(t, tt.expectedAlerts, alerts)
			if tt.expectedErr != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
