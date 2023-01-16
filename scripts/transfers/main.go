package main

import (
	"context"
	"math/big"
	"os"

	"github.com/forta-network/forta-core-go/contracts/generated/contract_forta_staking_0_1_1"
	rd "github.com/forta-network/forta-core-go/domain/registry"
	"github.com/forta-network/forta-core-go/domain/registry/regmsg"
	"github.com/forta-network/forta-core-go/registry"
	log "github.com/sirupsen/logrus"
)

// this script prints any stake transfers between two block ranges+
func main() {
	ctx := context.Background()
	l, err := registry.NewListener(ctx, registry.ListenerConfig{
		JsonRpcURL: os.Getenv("POLYGON_JSON_RPC"),
		Handlers: registry.Handlers{
			TransferSharesHandlers: regmsg.Handlers(
				func(ctx context.Context, logger *log.Entry, msg *rd.TransferSharesMessage) error {
					log.WithFields(log.Fields{
						"to":     msg.To,
						"from":   msg.From,
						"amount": msg.Amount,
						"type":   msg.StakeType,
						"burn":   msg.IsBurn(),
						"mint":   msg.IsMint(),
					}).Info("event")
					return nil
				},
			),
		},
		ContractFilter: &registry.ContractFilter{
			FortaStaking: true,
		},
		Topics: []string{contract_forta_staking_0_1_1.TransferSingleTopic, contract_forta_staking_0_1_1.TransferBatchTopic},
	})
	if err != nil {
		panic(err)
	}

	start := big.NewInt(30172379)
	end := big.NewInt(30189948)

	if err := l.ProcessBlockRange(start, end); err != nil {
		panic(err)
	}
}
