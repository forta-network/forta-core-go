package graphql

import (
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/forta-network/forta-core-go/protocol"
)

type GetAlertsResponse struct {
	Alerts struct {
		PageInfo *PageInfo                    `json:"pageInfo"`
		Alerts   []*protocol.AlertEvent_Alert `json:"alerts"`
	}
}

// AlertsInput Alert list input
type AlertsInput struct {
	// Indicate a list of addresses.
	// Alerts returned will have those addresses involved.
	Addresses []string `json:"addresses"`
	// Search results after the specified cursor
	After *AlertEndCursorInput `json:"after,omitempty"`
	// Get a specific alert by alert hash.
	AlertHash string `json:"alertHash"`
	// Filter alerts by alert-id.
	AlertId string `json:"alertId"`
	// Filter alerts by multiple alert-ids.
	AlertIds []string `json:"alertIds"`
	// Filter alerts by alert name.
	AlertName string `json:"alertName"`
	// Block Date range
	// Alerts returned will be between the specified start and end block timestamp dates when the threats were detected.
	// By default, start and end date will be set to local query execution date.
	BlockDateRange *DateRange `json:"blockDateRange,omitempty"`
	// Block number range
	// Alerts for the block number range will be returned.
	BlockNumberRange *BlockRange `json:"blockNumberRange,omitempty"`
	// Indicate sorting order by block number,
	// 'desc' or 'asc'.
	// Default is 'desc'.
	BlockSortDirection Sort `json:"blockSortDirection"`
	// Block Timestamp range
	// Alerts returned will be between the specified start and end block timestamp when the threats were detected.
	BlockTimestampRange *TimestampRange `json:"blockTimestampRange,omitempty"`
	// Indicate a list of bot hashes.
	// Alerts returned will only be from any of those bots.
	Bots []string `json:"bots"`
	// Indicate a chain Id: EIP155 identifier of the chain
	// Alerts returned will only be from the specific chain Id
	// Default is 1 = Ethereum Mainnet.
	ChainId uint `json:"chainId"`
	// Indicate number of milliseconds
	// Alerts returned will be alerts created before the number of milliseconds indicated ago.
	CreatedBefore uint `json:"createdBefore"`
	// Indicate number of milliseconds
	// Alerts returned will be alerts created since the number of milliseconds indicated ago.
	CreatedSince uint `json:"createdSince"`
	// Indicate max number of results.
	First uint `json:"first"`
	// Indicate a project Id.
	// Alerts returned will only be from that project.
	ProjectId string `json:"projectId"`
	// Filter alerts by number of scan nodes confirming the alert.
	ScanNodeConfirmations *ScanNodeFilters `json:"scanNodeConfirmations,omitempty"`
	// Filter alerts by severity levels.
	Severities []string `json:"severities"`
	// Indicate a transaction hash
	// Alerts returned will only be from that transaction.
	TransactionHash string `json:"transactionHash"`
}

// AlertEndCursorInput Search after specified block number and alertId
type AlertEndCursorInput struct {
	AlertId     string `json:"alertId"`
	BlockNumber uint   `json:"blockNumber"`
}

// DateRange Date format: YYYY-MM-DD
type DateRange struct {
	EndDate   string `json:"endDate"`
	StartDate string `json:"startDate"`
}

type Sort string

const (
	SortAsc  Sort = "asc"
	SortDesc Sort = "desc"
)

type BlockRange struct {
	EndBlockNumber   uint `json:"endBlockNumber"`
	StartBlockNumber uint `json:"startBlockNumber"`
}

// ScanNodeFilters Filter by number of scan nodes confirming the alert.
type ScanNodeFilters struct {
	Gte uint `json:"gte"`
	Lte uint `json:"lte"`
}

// TimestampRange Timestamp format: number of milliseconds from start of UNIX epoch
type TimestampRange struct {
	EndTimestamp   uint `json:"endTimestamp"`
	StartTimestamp uint `json:"startTimestamp"`
}

type PageInfo struct {
	HasNextPage bool       `json:"hasNextPage"`
	EndCursor   *EndCursor `json:"endCursor"`
}

type EndCursor struct {
	AlertId     string `json:"alertId"`
	BlockNumber uint   `json:"blockNumber"`
}

type __getAlertsInput struct {
	Input *AlertsInput `json:"input,omitempty"`
}

func (g *GetAlertsResponse) ToAlertEvents() []*protocol.AlertEvent {
	resp := make([]*protocol.AlertEvent, len(g.Alerts.Alerts))
	for i, alert := range g.Alerts.Alerts {
		// alert will be the source alert in consumer's perspective,
		// so it should be exposed as the SourceAlert tracking timestamp
		t, _ := time.Parse(time.RFC3339, alert.CreatedAt)
		alertTimestamp := hexutil.EncodeUint64(uint64(t.Unix()))

		resp[i] = &protocol.AlertEvent{
			Alert: alert,
			Timestamps: &protocol.TrackingTimestamps{
				SourceAlert: alertTimestamp,
			},
		}
	}

	return resp
}

// The query or mutation executed by getAlerts.
const getAlertsOperation = `
query getAlerts ($input: AlertsInput) {
	alerts(input: $input) {
		pageInfo {
			hasNextPage
			endCursor {
				alertId
				blockNumber
			}
		}
		alerts {
			alertId
			addresses
			contracts {
				name
				projectId
			}
			createdAt
			description
			hash
			metadata
			name
			projects {
				id
			}
			protocol
			scanNodeCount
			severity
			source {
				transactionHash
				bot {
					chainIds
					createdAt
					description
					developer
					docReference
					enabled
					id
					image
					name
					reference
					repository
					projects
					scanNodes
					version
				}
				block {
					number
					hash
					timestamp
					chainId
				}
				sourceAlert {
					hash
					botId
					timestamp
					chainId
				}
			}
			alertDocumentType
			findingType
			relatedAlerts
			chainId
			labels {
				label
				confidence
				entity
				entityType
				remove
				metadata
				uniqueKey
				embedding
			}
			addressBloomFilter {
				bitset
				itemCount
				k
				m
			}
		}
	}
}
`
