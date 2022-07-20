package registry

import (
	"context"
	"errors"
	"math/big"
	"os"
	"testing"

	"github.com/forta-network/forta-core-go/contracts/contract_agent_registry"
	"github.com/forta-network/forta-core-go/contracts/contract_dispatch"
	"github.com/forta-network/forta-core-go/contracts/contract_forta_staking"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_registry"
	"github.com/forta-network/forta-core-go/domain/registry"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
)

// This actually calls out to polygon to get some known blocks and parse the actions
// by default it gets the data at https://polygon-rpc.com, otherwise override with ENV var POLYGON_JSON_RPC
func testListener(ctx context.Context, filter *ContractFilter, topic string, handlers Handlers) Listener {
	jrpc := os.Getenv("POLYGON_JSON_RPC")
	if jrpc == "" {
		jrpc = defaultConfig.JsonRpcUrl
	}
	l, err := NewListener(ctx, ListenerConfig{
		Name:           "listener",
		JsonRpcURL:     jrpc,
		ENSAddress:     defaultEnsAddress,
		ContractFilter: filter,
		Topics:         []string{topic},
		Handlers:       handlers,
	})
	if err != nil {
		panic(err)
	}
	return l
}

type listenerTest struct {
	name     string
	listener Listener
	block    int64
}

func TestListener_Listen(t *testing.T) {
	ctx := context.Background()
	found := errors.New("found")
	tests := []listenerTest{
		{
			name: "agent-enable",
			listener: testListener(ctx, &ContractFilter{AgentRegistry: true},
				contract_agent_registry.AgentEnabledTopic,
				Handlers{
					AgentActionHandler: func(logger *log.Entry, msg *registry.AgentMessage) error {
						assert.Equal(t, int64(26030393), msg.Source.BlockNumberDecimal)
						assert.Equal(t, registry.EnableAgent, msg.Action)
						assert.Equal(t, "0x4bc8273b69f070c209c4866907d5def4c6f899af4cb47dee5e263aba6defad69", msg.AgentID)
						return found
					},
				}),
			block: 26030393,
		},
		{
			name: "agent-disable",
			listener: testListener(ctx, &ContractFilter{AgentRegistry: true},
				contract_agent_registry.AgentEnabledTopic,
				Handlers{
					AgentActionHandler: func(logger *log.Entry, msg *registry.AgentMessage) error {
						assert.Equal(t, int64(26029118), msg.Source.BlockNumberDecimal)
						assert.Equal(t, registry.DisableAgent, msg.Action)
						assert.Equal(t, "0x4bc8273b69f070c209c4866907d5def4c6f899af4cb47dee5e263aba6defad69", msg.AgentID)
						return found
					},
				}),
			block: 26029118,
		},
		{
			name: "scanner-enable",
			listener: testListener(ctx, &ContractFilter{ScannerRegistry: true},
				contract_scanner_registry.ScannerEnabledTopic,
				Handlers{
					ScannerActionHandler: func(logger *log.Entry, msg *registry.ScannerMessage) error {
						assert.Equal(t, int64(27984372), msg.Source.BlockNumberDecimal)
						assert.Equal(t, registry.EnableScanner, msg.Action)
						assert.Equal(t, "0xec542ecf3d371bb95c45af6776e20e7e5ff6cbc0", msg.ScannerID)
						return found
					},
				}),
			block: 27984372,
		},
		{
			name: "scanner-disable",
			listener: testListener(ctx, &ContractFilter{ScannerRegistry: true},
				contract_scanner_registry.ScannerEnabledTopic,
				Handlers{
					ScannerActionHandler: func(logger *log.Entry, msg *registry.ScannerMessage) error {
						assert.Equal(t, int64(28005870), msg.Source.BlockNumberDecimal)
						assert.Equal(t, registry.DisableScanner, msg.Action)
						assert.Equal(t, "0x7c4cd89c996bdc1589a50e7b2116b43fdca7e6d2", msg.ScannerID)
						return found
					},
				}),
			block: 28005870,
		},
		{
			name: "scanner-save",
			listener: testListener(ctx, &ContractFilter{ScannerRegistry: true},
				contract_scanner_registry.ScannerUpdatedTopic,
				Handlers{
					SaveScannerHandler: func(logger *log.Entry, msg *registry.ScannerSaveMessage) error {
						assert.Equal(t, int64(25809030), msg.Source.BlockNumberDecimal)
						assert.Equal(t, registry.SaveScanner, msg.Action)
						assert.Equal(t, "0xdec088fea5feab7dc17789a92bffc10393a769de", msg.ScannerID)
						return found
					},
				}),
			block: 25809030,
		},
		{
			name: "agent-save",
			listener: testListener(ctx, &ContractFilter{AgentRegistry: true},
				contract_agent_registry.AgentUpdatedTopic,
				Handlers{
					SaveAgentHandler: func(logger *log.Entry, msg *registry.AgentSaveMessage) error {
						assert.Equal(t, int64(25730681), msg.Source.BlockNumberDecimal)
						assert.Equal(t, registry.SaveAgent, msg.Action)
						assert.Equal(t, "0xdc75bb779e1cbe73a21b8d3810867411f6d71eb64f79cec7e7d4fbcaa40de990", msg.AgentID)
						return found
					},
				}),
			block: 25730681,
		},
		{
			name: "dispatch-link",
			listener: testListener(ctx, &ContractFilter{DispatchRegistry: true},
				contract_dispatch.LinkTopic,
				Handlers{
					DispatchHandler: func(logger *log.Entry, msg *registry.DispatchMessage) error {
						assert.Equal(t, int64(25730716), msg.Source.BlockNumberDecimal)
						assert.Equal(t, registry.Link, msg.Action)
						assert.Equal(t, "0xdc75bb779e1cbe73a21b8d3810867411f6d71eb64f79cec7e7d4fbcaa40de990", msg.AgentID)
						assert.Equal(t, "0xb0697be4e0ee8f18d741b5b1c940c4bcac9c7eb4", msg.ScannerID)
						return found
					},
				}),
			block: 25730716,
		},
		{
			name: "dispatch-unlink",
			listener: testListener(ctx, &ContractFilter{DispatchRegistry: true},
				contract_dispatch.LinkTopic,
				Handlers{
					DispatchHandler: func(logger *log.Entry, msg *registry.DispatchMessage) error {
						assert.Equal(t, int64(25772370), msg.Source.BlockNumberDecimal)
						assert.Equal(t, registry.Unlink, msg.Action)
						assert.Equal(t, "0x841f771742ce7d9904d061db29dcf31ef500979d73e1859b4b81c7c739499f2a", msg.AgentID)
						assert.Equal(t, "0xaa64b99aed1b93e705b35f4a5ff2822871888b9d", msg.ScannerID)
						return found
					},
				}),
			block: 25772370,
		},
		{
			name: "scanner-stake-threshold",
			listener: testListener(ctx, &ContractFilter{ScannerRegistry: true},
				contract_scanner_registry.StakeThresholdChangedTopic,
				Handlers{
					ScannerStakeThresholdHandler: func(logger *log.Entry, msg *registry.ScannerStakeThresholdMessage) error {
						assert.Equal(t, int64(26465762), msg.Source.BlockNumberDecimal)
						assert.Equal(t, registry.ScannerStakeThreshold, msg.Action)
						assert.Equal(t, "0", msg.Min)
						assert.Equal(t, "750000000000000000000", msg.Max)
						return found
					},
				}),
			block: 26465762,
		},
		{
			name: "scanner-stake",
			listener: testListener(ctx, &ContractFilter{FortaStaking: true},
				contract_forta_staking.StakeDepositedTopic,
				Handlers{
					ScannerStakeHandler: func(logger *log.Entry, msg *registry.ScannerStakeMessage) error {
						assert.Equal(t, int64(26498238), msg.Source.BlockNumberDecimal)
						assert.Equal(t, registry.ScannerStake, msg.Action)
						assert.Equal(t, registry.ChangeTypeDeposit, msg.ChangeType)
						assert.Equal(t, "0x58ee631aaef6882a392da1c25486ee181ff1b7d5", msg.ScannerID)
						assert.Equal(t, "500000000000000000000", msg.Amount)
						return found
					},
				}),
			block: 26498238,
		},
		{
			name: "transfer-shares",
			listener: testListener(ctx, &ContractFilter{FortaStaking: true},
				contract_forta_staking.TransferSingleTopic,
				Handlers{
					TransferSharesHandler: func(logger *log.Entry, msg *registry.TransferSharesMessage) error {
						assert.Equal(t, int64(26498238), msg.Source.BlockNumberDecimal)
						assert.Equal(t, registry.TransferShares, msg.Action)
						assert.Equal(t, "0x9b0a8a8e6b2c23d572d7145f3da14438fed35374", msg.To)
						assert.Equal(t, "500000000000000000000", msg.Amount)
						return found
					},
				}),
			block: 26498238,
		},
	}

	grp, _ := errgroup.WithContext(ctx)

	for _, lt := range tests {
		tst := lt
		grp.Go(func() error {
			err := tst.listener.ProcessBlockRange(big.NewInt(tst.block), big.NewInt(tst.block))
			assert.Equal(t, found, err, tst.name)
			return nil
		})
	}
	assert.NoError(t, grp.Wait())
}
