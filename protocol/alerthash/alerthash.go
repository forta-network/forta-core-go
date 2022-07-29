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

// Inputs hold all of the possible inputs for alert hash calculation.
type Inputs struct {
	Block       *protocol.BlockEvent
	Transaction *protocol.TransactionEvent
	Finding     *protocol.Finding
	BotInfo
}

// BotInfo contains bot-related info useful during hash calculation.
type BotInfo struct {
	BotImage string
	BotID    string
}

// ForBlockAlert calculates the hash for the block alert.
func ForBlockAlert(inputs *Inputs) string {
	idStr := strings.Join([]string{
		Version, "|",
		inputs.Block.Network.ChainId,
		inputs.Block.BlockHash,
		inputs.Finding.AlertId,
		inputs.Finding.Name,
		inputs.Finding.Description,
		inputs.Finding.Protocol,
		inputs.Finding.Type.String(),
		inputs.Finding.Severity.String(),
		inputs.BotImage,
		inputs.BotID,
		strings.Join(inputs.Finding.Addresses, "")}, "")
	return crypto.Keccak256Hash([]byte(idStr)).Hex()
}

// ForTransactionAlert calculates the hash for the transaction alert.
func ForTransactionAlert(inputs *Inputs) string {
	txAddrs := utils.MapKeys(inputs.Transaction.TxAddresses)
	sort.Strings(txAddrs)
	idStr := strings.Join([]string{
		Version, "|",
		inputs.Transaction.Network.ChainId,
		inputs.Transaction.Transaction.Hash,
		inputs.Finding.Name,
		inputs.Finding.Description,
		inputs.Finding.Protocol,
		inputs.Finding.Type.String(),
		inputs.Finding.AlertId,
		inputs.Finding.Severity.String(),
		inputs.BotImage,
		inputs.BotID,
		strings.Join(txAddrs, ""),
		strings.Join(inputs.Finding.Addresses, "")}, "")
	return crypto.Keccak256Hash([]byte(idStr)).Hex()
}
