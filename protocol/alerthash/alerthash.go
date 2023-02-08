package alerthash

import (
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/forta-network/forta-core-go/utils"
)

// Version represents alert hash version.
const Version = "1"

// Inputs hold all the possible inputs for alert hash calculation.
type Inputs struct {
	BlockEvent       *protocol.BlockEvent
	TransactionEvent *protocol.TransactionEvent
	AlertEvent       *protocol.AlertEvent
	Finding          *protocol.Finding
	BotInfo
}

// BotInfo contains bot-related info useful during hash calculation.
type BotInfo struct {
	BotImage string
	BotID    string
}

// ForBlockAlert calculates the hash for the block alert.
func ForBlockAlert(inputs *Inputs) string {
	sort.Strings(inputs.Finding.Addresses)
	idStr := strings.Join(
		[]string{
			Version, "|",
			inputs.BlockEvent.Network.ChainId,
			inputs.BlockEvent.BlockHash,
			inputs.Finding.AlertId,
			inputs.Finding.Name,
			inputs.Finding.Description,
			inputs.Finding.Protocol,
			inputs.Finding.Type.String(),
			inputs.Finding.Severity.String(),
			inputs.BotImage,
			inputs.BotID,
			strings.Join(inputs.Finding.Addresses, ""),
		}, "",
	)
	return crypto.Keccak256Hash([]byte(idStr)).Hex()
}

// ForTransactionAlert calculates the hash for the transaction alert.
func ForTransactionAlert(inputs *Inputs) string {
	sort.Strings(inputs.Finding.Addresses)
	txAddrs := utils.MapKeys(inputs.TransactionEvent.TxAddresses)
	sort.Strings(txAddrs)
	idStr := strings.Join(
		[]string{
			Version, "|",
			inputs.TransactionEvent.Network.ChainId,
			inputs.TransactionEvent.Transaction.Hash,
			inputs.Finding.Name,
			inputs.Finding.Description,
			inputs.Finding.Protocol,
			inputs.Finding.Type.String(),
			inputs.Finding.AlertId,
			inputs.Finding.Severity.String(),
			inputs.BotImage,
			inputs.BotID,
			strings.Join(txAddrs, ""),
			strings.Join(inputs.Finding.Addresses, ""),
		}, "",
	)
	return crypto.Keccak256Hash([]byte(idStr)).Hex()
}

// ForCombinationAlert calculates the hash for the alert handler alert.
func ForCombinationAlert(inputs *Inputs) string {
	sort.Strings(inputs.Finding.Addresses)
	sort.Strings(inputs.Finding.RelatedAlerts)

	idStr := strings.Join(
		[]string{
			Version, "|",
			inputs.AlertEvent.Alert.Hash,
			inputs.Finding.Name,
			inputs.Finding.Description,
			inputs.Finding.Protocol,
			inputs.Finding.Type.String(),
			inputs.Finding.AlertId,
			inputs.Finding.Severity.String(),
			inputs.BotImage,
			inputs.BotID,
			// source alert hash
			strings.Join(inputs.Finding.Addresses, ""),
			strings.Join(inputs.Finding.RelatedAlerts, ""),
		}, "",
	)
	return crypto.Keccak256Hash([]byte(idStr)).Hex()
}
