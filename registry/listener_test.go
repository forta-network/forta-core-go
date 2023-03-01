package registry

import (
	"context"
	"errors"
	"math/big"
	"os"
	"strings"
	"testing"

	"github.com/forta-network/forta-core-go/contracts/generated/contract_agent_registry_0_1_4"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_dispatch_0_1_4"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_forta_staking_0_1_1"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_forta_staking_0_1_2"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_scanner_node_version_0_1_0"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_scanner_registry_0_1_3"

	"github.com/forta-network/forta-core-go/domain/registry"
	"github.com/forta-network/forta-core-go/domain/registry/regmsg"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		NoRefresh:      true,
	})
	if err != nil {
		panic(err)
	}
	return l
}

func testMumbaiListener(ctx context.Context, filter *ContractFilter, topic string, handlers Handlers) Listener {
	l, err := NewListener(ctx, ListenerConfig{
		Name:           "listener",
		JsonRpcURL:     "https://rpc.ankr.com/polygon_mumbai",
		ENSAddress:     devConfig.ENSAddress,
		ContractFilter: filter,
		Topics:         []string{topic},
		Handlers:       handlers,
		NoRefresh:      true,
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
			name: "upgrade",
			listener: testListener(ctx, &ContractFilter{ScannerVersion: true},
				UpgradedTopic,
				Handlers{
					UpgradeHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.UpgradeMessage) error {
							assert.Equal(t, int64(27061991), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.Upgrade, msg.Action)
							assert.Equal(t, strings.ToLower("0x4720c872425876B6f4b4E9130CDef667aDE553b2"), msg.Proxy)
							assert.Equal(t, strings.ToLower("0x68608f260ba3be8808b83fd686eccfe6a9f468f7"), msg.NewImplementation)
							return found
						},
					),
				},
			),
			block: 27061991,
		},
		{
			name: "configuration-change",
			listener: testMumbaiListener(ctx, &ContractFilter{FortaStaking: true},
				contract_forta_staking_0_1_2.StakeHelpersConfiguredTopic,
				Handlers{
					UpgradeHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.UpgradeMessage) error {
							assert.Equal(t, int64(31808525), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.ConfigurationChange, msg.Action)
							assert.Equal(t, strings.ToLower("0x64d5192f03bd98db1de2aa8b4abac5419eac32ce"), msg.Proxy)
							return found
						},
					),
				},
			),
			block: 31808525,
		},
		{
			name: "agent-enable",
			listener: testListener(ctx, &ContractFilter{AgentRegistry: true},
				contract_agent_registry_0_1_4.AgentEnabledTopic,
				Handlers{
					AgentActionHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.AgentMessage) error {
							assert.Equal(t, int64(26030393), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.EnableAgent, msg.Action)
							assert.Equal(t, "0x4bc8273b69f070c209c4866907d5def4c6f899af4cb47dee5e263aba6defad69", msg.AgentID)
							return found
						},
					),
				},
			),
			block: 26030393,
		},
		{
			name: "agent-disable",
			listener: testListener(ctx, &ContractFilter{AgentRegistry: true},
				contract_agent_registry_0_1_4.AgentEnabledTopic,
				Handlers{
					AgentActionHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.AgentMessage) error {
							assert.Equal(t, int64(26029118), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.DisableAgent, msg.Action)
							assert.Equal(t, "0x4bc8273b69f070c209c4866907d5def4c6f899af4cb47dee5e263aba6defad69", msg.AgentID)
							return found
						},
					),
				},
			),
			block: 26029118,
		},
		{
			name: "scanner-save",
			listener: testListener(ctx, &ContractFilter{ScannerRegistry: true},
				contract_scanner_registry_0_1_3.ScannerUpdatedTopic,
				Handlers{
					SaveScannerHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.ScannerSaveMessage) error {
							assert.Equal(t, int64(25809030), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.SaveScanner, msg.Action)
							assert.Equal(t, "0xdec088fea5feab7dc17789a92bffc10393a769de", msg.ScannerID)
							return found
						},
					),
				},
			),
			block: 25809030,
		},
		{
			name: "agent-save",
			listener: testListener(ctx, &ContractFilter{AgentRegistry: true},
				contract_agent_registry_0_1_4.AgentUpdatedTopic,
				Handlers{
					SaveAgentHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.AgentSaveMessage) error {
							assert.Equal(t, int64(25730681), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.SaveAgent, msg.Action)
							assert.Equal(t, "0xdc75bb779e1cbe73a21b8d3810867411f6d71eb64f79cec7e7d4fbcaa40de990", msg.AgentID)
							return found
						},
					),
				},
			),
			block: 25730681,
		},
		{
			name: "dispatch-link",
			listener: testListener(ctx, &ContractFilter{DispatchRegistry: true},
				contract_dispatch_0_1_4.LinkTopic,
				Handlers{
					DispatchHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.DispatchMessage) error {
							assert.Equal(t, int64(25730716), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.Link, msg.Action)
							assert.Equal(t, "0xdc75bb779e1cbe73a21b8d3810867411f6d71eb64f79cec7e7d4fbcaa40de990", msg.AgentID)
							assert.Equal(t, "0xb0697be4e0ee8f18d741b5b1c940c4bcac9c7eb4", msg.ScannerID)
							return found
						},
					),
				},
			),
			block: 25730716,
		},
		{
			name: "dispatch-unlink",
			listener: testListener(ctx, &ContractFilter{DispatchRegistry: true},
				contract_dispatch_0_1_4.LinkTopic,
				Handlers{
					DispatchHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.DispatchMessage) error {
							assert.Equal(t, int64(25772370), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.Unlink, msg.Action)
							assert.Equal(t, "0x841f771742ce7d9904d061db29dcf31ef500979d73e1859b4b81c7c739499f2a", msg.AgentID)
							assert.Equal(t, "0xaa64b99aed1b93e705b35f4a5ff2822871888b9d", msg.ScannerID)
							return found
						},
					),
				},
			),
			block: 25772370,
		},
		{
			name: "scanner-stake-threshold",
			listener: testListener(ctx, &ContractFilter{ScannerRegistry: true},
				contract_scanner_registry_0_1_3.StakeThresholdChangedTopic,
				Handlers{
					ScannerStakeThresholdHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.ScannerStakeThresholdMessage) error {
							assert.Equal(t, int64(26465762), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.ScannerStakeThreshold, msg.Action)
							assert.Equal(t, "0", msg.Min)
							assert.Equal(t, "750000000000000000000", msg.Max)
							return found
						},
					),
				},
			),
			block: 26465762,
		},
		{
			name: "scanner-stake",
			listener: testListener(ctx, &ContractFilter{FortaStaking: true},
				contract_forta_staking_0_1_1.StakeDepositedTopic,
				Handlers{
					ScannerStakeHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.ScannerStakeMessage) error {
							assert.Equal(t, int64(26498238), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.ScannerStake, msg.Action)
							assert.Equal(t, registry.ChangeTypeDeposit, msg.ChangeType)
							assert.Equal(t, "0x58ee631aaef6882a392da1c25486ee181ff1b7d5", msg.ScannerID)
							assert.Equal(t, "500000000000000000000", msg.Amount)
							return found
						},
					),
				},
			),
			block: 26498238,
		},
		{
			name: "transfer-shares",
			listener: testListener(ctx, &ContractFilter{FortaStaking: true},
				contract_forta_staking_0_1_1.TransferSingleTopic,
				Handlers{
					TransferSharesHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.TransferSharesMessage) error {
							assert.Equal(t, int64(26498238), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.TransferShares, msg.Action)
							assert.Equal(t, "0x9b0a8a8e6b2c23d572d7145f3da14438fed35374", msg.To)
							assert.Equal(t, "500000000000000000000", msg.Amount)
							return found
						},
					),
				},
			),
			block: 26498238,
		}, {
			name: "scanner-version",
			listener: testListener(ctx, &ContractFilter{ScannerVersion: true},
				contract_scanner_node_version_0_1_0.ScannerNodeVersionUpdatedTopic,
				Handlers{
					ScannerNodeVersionHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.ScannerNodeVersionMessage) error {
							assert.Equal(t, int64(31739343), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.ScannerNodeVersionUpdated, msg.Action)
							assert.Equal(t, "QmPU5yx6Puapj7o79zY4n1LkyhnZowLjoDaZv7TRhytYpt", msg.NewVersion)
							assert.Equal(t, "QmQX4FELAScb8n1cjAr6wK4WF73b2jGv9FSRtUWhysnKu1", msg.OldVersion)
							return found
						},
					),
				},
			),
			block: 31739343,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			r := require.New(t)
			err := testCase.listener.ProcessBlockRange(big.NewInt(testCase.block), big.NewInt(testCase.block))
			r.Equal(found, err)
		})
	}
}
