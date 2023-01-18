package registry

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_stake_allocator"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/domain/registry/regmsg"
	"github.com/forta-network/forta-core-go/utils"
	"github.com/sirupsen/logrus"
)

const (
	ScannerPoolAllocatedStake = "ScannerPoolAllocatedStake"
)

type ScannerPoolAllocationMessage struct {
	regmsg.Message
	PoolID            string `json:"poolId"`
	ChangeAmount      string `json:"changeAmount"`
	TotalAmount       string `json:"totalAmount"`
	Increase          bool   `json:"increase"`
	CurrentPerScanner string `json:"currentPerScanner"`
}

func (spam *ScannerPoolAllocationMessage) LogFields() logrus.Fields {
	return logrus.Fields{}
}

func NewScannerPoolAllocationMessage(l types.Log, blk *domain.Block, evt *contract_stake_allocator.StakeAllocatorAllocatedStake, currentPerScanner *big.Int) *ScannerPoolAllocationMessage {
	return &ScannerPoolAllocationMessage{
		Message:           regmsg.From(l.TxHash.Hex(), blk, ScannerPoolAllocatedStake),
		PoolID:            utils.PoolIDToString(evt.Subject),
		ChangeAmount:      evt.Amount.String(),
		TotalAmount:       evt.TotalAllocated.String(),
		Increase:          evt.Increase,
		CurrentPerScanner: currentPerScanner.String(),
	}
}
