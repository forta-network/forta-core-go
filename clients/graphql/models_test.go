package graphql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createGetAlertsQuery(t *testing.T) {
	testInput := AlertsInput{
		AlertId: "0xabc",
	}
	resp, variables := createGetAlertsQuery([]*AlertsInput{&testInput})
	assert.Equal(t, resp, mockExpectedQuery)
	variable, ok := variables["input0"].(*AlertsInput)
	assert.True(t, ok)
	assert.Equal(t, variable.AlertId, testInput.AlertId)
}

const mockExpectedQuery = `query getAlerts($input0: AlertsInput) {alerts0: alerts(input: $input0) {
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
}}`
