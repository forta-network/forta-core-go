package registry

import (
	"context"
	"errors"
	"math/big"
	"strings"
	"sync"
	"testing"

	"github.com/forta-network/forta-core-go/contracts/generated/contract_rewards_distributor_0_1_0"

	"github.com/forta-network/forta-core-go/contracts/generated/contract_agent_registry_0_1_4"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_dispatch_0_1_4"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_forta_staking_0_1_1"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_forta_staking_0_1_2"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_scanner_node_version_0_1_0"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_scanner_pool_registry_0_1_0"

	"github.com/forta-network/forta-core-go/domain/registry"
	"github.com/forta-network/forta-core-go/domain/registry/regmsg"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This actually calls out to polygon to get some known blocks and parse the actions
// by default it gets the data at https://polygon-rpc.com, otherwise override with ENV var POLYGON_JSON_RPC
func testListener(ctx context.Context, filter *ContractFilter, topic string, handlers Handlers) Listener {
	l, err := NewListener(ctx, ListenerConfig{
		Name:           "listener",
		JsonRpcURL:     "https://rpc.ankr.com/polygon",
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
	l, err := NewListener(
		ctx, ListenerConfig{
			Name:       "listener",
			JsonRpcURL: "https://rpc.ankr.com/polygon_mumbai",
			// ENSAddress:     devConfig.ENSAddress,
			ENSAddress:     "0x5f7c5bbBa72e1e1fae689120D76D2f334A390Ae9",
			ContractFilter: filter,
			Topics:         []string{topic},
			Handlers:       handlers,
			NoRefresh:      true,
		},
	)
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
	handledEvent := errors.New("handled event")
	mux := sync.Mutex{}
	found := make(map[string]bool)
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
							return handledEvent
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
							return handledEvent
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
							return handledEvent
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
							return handledEvent
						},
					),
				},
			),
			block: 26029118,
		},
		{
			name: "scanner-enable",
			listener: testListener(ctx, &ContractFilter{ScannerPoolRegistry: true},
				contract_scanner_pool_registry_0_1_0.EnabledScannersChangedTopic,
				Handlers{
					UpdateScannerPoolHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.UpdateScannerPoolMessage) error {
							assert.Equal(t, int64(39834786), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.UpdateScannerPool, msg.Action)
							assert.Equal(t, "70", msg.PoolID)
							return handledEvent
						},
					),
				},
			),
			block: 39834786,
		},
		{
			name: "scanner-save",
			listener: testMumbaiListener(ctx, &ContractFilter{ScannerPoolRegistry: true},
				contract_scanner_pool_registry_0_1_0.ScannerUpdatedTopic,
				Handlers{
					SaveScannerHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.ScannerSaveMessage) error {
							assert.Equal(t, int64(32910234), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.SaveScanner, msg.Action)
							assert.Equal(t, strings.ToLower("0x3f88c2b3e267e6b8e9de017cdb47a59ac9ecb284"), msg.ScannerID)
							return handledEvent
						},
					),
				},
			),
			block: 32910234,
		},
		{
			name: "agent-save-from-update",
			listener: testListener(ctx, &ContractFilter{AgentRegistry: true},
				contract_agent_registry_0_1_4.AgentUpdatedTopic,
				Handlers{
					SaveAgentHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.AgentSaveMessage) error {
							assert.Equal(t, int64(25730681), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.SaveAgent, msg.Action)
							assert.Equal(t, "0xdc75bb779e1cbe73a21b8d3810867411f6d71eb64f79cec7e7d4fbcaa40de990", msg.AgentID)
							return handledEvent
						},
					),
				},
			),
			block: 25730681,
		},
		{
			name: "agent-save-from-owner-transfer",
			listener: testListener(ctx, &ContractFilter{AgentRegistry: true},
				contract_agent_registry_0_1_4.TransferTopic,
				Handlers{
					SaveAgentHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.AgentSaveMessage) error {
							assert.Equal(t, int64(39058880), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.SaveAgent, msg.Action)
							assert.Equal(t, "0x6f022d4a65f397dffd059e269e1c2b5004d822f905674dbf518d968f744c2ede", msg.AgentID)
							return handledEvent
						},
					),
				},
			),
			block: 39058880,
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
							return handledEvent
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
							return handledEvent
						},
					),
				},
			),
			block: 25772370,
		},
		{
			name: "scanner-stake-threshold",
			listener: testMumbaiListener(ctx, &ContractFilter{ScannerPoolRegistry: true},
				contract_scanner_pool_registry_0_1_0.ManagedStakeThresholdChangedTopic,
				Handlers{
					ScannerStakeThresholdHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.ScannerStakeThresholdMessage) error {
							assert.Equal(t, int64(32901311), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.ScannerStakeThreshold, msg.Action)
							assert.Equal(t, "2500000000000000000000", msg.Min)
							assert.Equal(t, "15000000000000000000000", msg.Max)
							return handledEvent
						},
					),
				},
			),
			block: 32901311,
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
							return handledEvent
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
							return handledEvent
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
							return handledEvent
						},
					),
				},
			),
			block: 31739343,
		}, {
			name: "rewarded",
			listener: testListener(ctx, &ContractFilter{RewardDistributor: true},
				contract_rewards_distributor_0_1_0.RewardedTopic,
				Handlers{
					RewardedHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.RewardedMessage) error {
							mux.Lock()
							defer mux.Unlock()
							// if matching event already found, skip
							if _, ok := found["rewarded"]; ok {
								// should not fire twice for same pool
								assert.NotEqual(t, "820", msg.Subject)
								return handledEvent
							}
							// if not the right payment, skip it
							if msg.Subject != "820" {
								return nil
							}
							assert.Equal(t, int64(41447648), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.Rewarded, msg.Action)
							assert.Equal(t, "1765787550499999989760", msg.Amount)
							assert.Equal(t, "820", msg.Subject)
							found["rewarded"] = true
							return handledEvent
						},
					),
				},
			),
			block: 41447648,
		}, {
			name: "claimed-rewards",
			listener: testListener(ctx, &ContractFilter{RewardDistributor: true},
				contract_rewards_distributor_0_1_0.ClaimedRewardsTopic,
				Handlers{
					ClaimedRewardsHandlers: regmsg.Handlers(
						func(ctx context.Context, logger *log.Entry, msg *registry.ClaimedRewardsMessage) error {
							mux.Lock()
							defer mux.Unlock()

							// if matching found, just skip the next items
							if _, ok := found["claimed-rewards"]; ok {
								// should not fire twice for same pool/payee
								assert.False(t, msg.To == "0x7ecd70a3b139499c531ad6c84b772483fddff8f7" && msg.Subject == "229")
								return handledEvent
							}
							// if not the right claim, skip it
							if !(msg.To == "0x7ecd70a3b139499c531ad6c84b772483fddff8f7" && msg.Subject == "229") {
								return nil
							}
							assert.Equal(t, int64(41112688), msg.Source.BlockNumberDecimal)
							assert.Equal(t, registry.ClaimedRewards, msg.Action)
							assert.Equal(t, "140913039600000000000", msg.Amount)
							assert.Equal(t, "0x7ecd70a3b139499c531ad6c84b772483fddff8f7", msg.To)
							assert.Equal(t, "229", msg.Subject)
							assert.Equal(t, 2, msg.SubjectType)
							found["claimed-rewards"] = true
							return handledEvent
						},
					),
				},
			),
			block: 41112688,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			r := require.New(t)
			err := testCase.listener.ProcessBlockRange(big.NewInt(testCase.block), big.NewInt(testCase.block))
			r.Equal(handledEvent, err)
		})
	}
}
