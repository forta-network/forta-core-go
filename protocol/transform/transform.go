package transform

import (
	"github.com/forta-network/forta-core-go/clients/webhook/models"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/forta-network/forta-core-go/utils"
)

// ToWebhookAlertList transforms an alert batch to a webhook alert list.
func ToWebhookAlertList(batch *protocol.AlertBatch) *models.AlertList {
	var alertList models.AlertList
	for _, resultsForBlock := range batch.Results {
		for _, blockResult := range resultsForBlock.Results {
			for _, alert := range blockResult.Alerts {
				alertList.Alerts = append(alertList.Alerts, ToWebhookAlert(
					alert.Alert,
					batch.ChainId,
					resultsForBlock.Block,
					nil,
				))
			}
		}
		for _, resultsForTransaction := range resultsForBlock.Transactions {
			for _, transactionResult := range resultsForTransaction.Results {
				for _, alert := range transactionResult.Alerts {
					alertList.Alerts = append(alertList.Alerts, ToWebhookAlert(
						alert.Alert,
						batch.ChainId,
						resultsForBlock.Block,
						resultsForTransaction.Transaction,
					))
				}
			}
		}
	}
	return &alertList
}

// ToWebhookAlert converts given alert and extra data to webhook alert.
func ToWebhookAlert(alert *protocol.Alert, chainID uint64, block *protocol.Block, transaction *protocol.TransactionEvent) *models.Alert {
	webhookAlert := &models.Alert{
		AlertID:     alert.Finding.AlertId,
		CreatedAt:   alert.Timestamp,
		Description: alert.Finding.Description,
		FindingType: alert.Finding.Type.String(),
		Hash:        alert.Id,
		Metadata:    alert.Finding.Metadata,
		Name:        alert.Finding.Name,
		Protocol:    alert.Finding.Protocol,
		Severity:    alert.Finding.Severity.String(),
		Source: &models.AlertSource{
			Agent: &models.AlertAgent{
				ID:        alert.Agent.Id,
				Image:     alert.Agent.Image,
				Reference: alert.Agent.Manifest,
			},
		},
	}
	if block != nil {
		webhookAlert.Source.Block = &models.AlertBlock{
			ChainID:   chainID,
			Hash:      block.BlockHash,
			Number:    block.BlockNumber,
			Timestamp: block.BlockTimestamp,
		}
	}
	if transaction != nil {
		webhookAlert.Addresses = utils.MapKeys(transaction.Addresses)
		webhookAlert.Source.TransactionHash = transaction.Transaction.Hash
	}
	return webhookAlert
}
