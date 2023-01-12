package registry

import (
	"strings"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_dispatch"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/sirupsen/logrus"
)

const Upgrade = "Upgrade"

type UpgradeMessage struct {
	Message
	Proxy             string `json:"proxy"`
	NewImplementation string `json:"newImplementation"`
}

func (um *UpgradeMessage) LogFields() logrus.Fields {
	return logrus.Fields{"proxy": um.Proxy}
}

func NewUpgradeMessage(evt *contract_dispatch.DispatchUpgraded, l types.Log, blk *domain.Block) *UpgradeMessage {
	return &UpgradeMessage{
		Message:           MessageFrom(evt.Raw.TxHash.Hex(), blk, Upgrade),
		Proxy:             strings.ToLower(l.Address.Hex()),
		NewImplementation: strings.ToLower(evt.Implementation.Hex()),
	}
}
