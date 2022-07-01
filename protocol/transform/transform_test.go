package transform_test

import (
	"testing"

	"github.com/forta-network/forta-core-go/protocol"
	"github.com/forta-network/forta-core-go/protocol/transform"
	"github.com/stretchr/testify/require"
)

func TestBatchToAlertList(t *testing.T) {
	chainID := uint64(1234)
	blockHash := "block-hash"
	blockNumber := 1337
	blockTimestamp := "block-timestamp"
	alertID := "alert-id"
	alertHash := "alert-hash"
	blockAlertType := protocol.AlertType_BLOCK
	alertTimestamp := "alert-timestamp"
	protocolName := "protocol"
	severity := protocol.Finding_CRITICAL
	findingType := protocol.Finding_DEGRADED
	name := "some alert name"
	description := "some alert description"
	agentID := "agent-id"
	agentImage := "agent-image"
	agentManifest := "agent-manifest"
	transactionHash := "transaction-hash"

	batch := &protocol.AlertBatch{
		ChainId: chainID,
		Results: []*protocol.BlockResults{
			{
				Block: &protocol.Block{
					BlockHash:      blockHash,
					BlockNumber:    uint64(blockNumber),
					BlockTimestamp: blockTimestamp,
				},
				Results: []*protocol.AgentAlerts{
					{
						Alerts: []*protocol.SignedAlert{
							{
								Alert: &protocol.Alert{
									Id:        alertHash,
									Timestamp: alertTimestamp,
									Type:      blockAlertType,
									Finding: &protocol.Finding{
										Protocol:    protocolName,
										Severity:    severity,
										Type:        findingType,
										AlertId:     alertID,
										Name:        name,
										Description: description,
									},
									Agent: &protocol.AgentInfo{
										Id:       agentID,
										Image:    agentImage,
										Manifest: agentManifest,
									},
								},
							},
						},
					},
				},
				Transactions: []*protocol.TransactionResults{
					{
						Transaction: &protocol.TransactionEvent{
							Transaction: &protocol.TransactionEvent_EthTransaction{
								Hash: transactionHash,
							},
						},
						Results: []*protocol.AgentAlerts{
							{
								Alerts: []*protocol.SignedAlert{
									{
										Alert: &protocol.Alert{
											Id:        alertHash,
											Timestamp: alertTimestamp,
											Type:      blockAlertType,
											Finding: &protocol.Finding{
												Protocol:    protocolName,
												Severity:    severity,
												Type:        findingType,
												AlertId:     alertID,
												Name:        name,
												Description: description,
											},
											Agent: &protocol.AgentInfo{
												Id:       agentID,
												Image:    agentImage,
												Manifest: agentManifest,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	alertList := transform.ToWebhookAlertList(batch)
	r := require.New(t)

	r.Len(alertList, 2)
	blockAlert := alertList[0]
	blockAlertAgent := blockAlert.Source.Bot
	blockAlertBlock := blockAlert.Source.Block
	r.Equal(alertID, blockAlert.AlertID)
	r.Equal(alertTimestamp, blockAlert.CreatedAt)
	r.Equal(description, blockAlert.Description)
	r.Equal(findingType.String(), blockAlert.FindingType)
	r.Equal(alertHash, blockAlert.Hash)
	r.Equal(name, blockAlert.Name)
	r.Equal(protocolName, blockAlert.Protocol)
	r.Equal(severity.String(), blockAlert.Severity)
	r.Equal(agentID, blockAlertAgent.ID)
	r.Equal(agentImage, blockAlertAgent.Image)
	r.Equal(agentManifest, blockAlertAgent.Reference)
	r.Equal(blockHash, blockAlertBlock.Hash)
	r.Equal(chainID, uint64(blockAlertBlock.ChainID))
	r.Equal(uint64(blockNumber), blockAlertBlock.Number)
	r.Equal(blockTimestamp, blockAlertBlock.Timestamp)
	r.Empty(blockAlert.Source.TransactionHash)

	txAlert := alertList[1]
	txAlertAgent := txAlert.Source.Bot
	txAlertBlock := txAlert.Source.Block
	r.Equal(alertID, txAlert.AlertID)
	r.Equal(alertTimestamp, txAlert.CreatedAt)
	r.Equal(description, txAlert.Description)
	r.Equal(findingType.String(), txAlert.FindingType)
	r.Equal(alertHash, txAlert.Hash)
	r.Equal(name, txAlert.Name)
	r.Equal(protocolName, txAlert.Protocol)
	r.Equal(severity.String(), txAlert.Severity)
	r.Equal(agentID, txAlertAgent.ID)
	r.Equal(agentImage, txAlertAgent.Image)
	r.Equal(agentManifest, txAlertAgent.Reference)
	r.Equal(blockHash, txAlertBlock.Hash)
	r.Equal(chainID, uint64(txAlertBlock.ChainID))
	r.Equal(uint64(blockNumber), txAlertBlock.Number)
	r.Equal(blockTimestamp, txAlertBlock.Timestamp)
	r.Equal(transactionHash, txAlert.Source.TransactionHash)
}
