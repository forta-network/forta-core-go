package registry

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/forta-network/forta-core-go/contracts/generated/contract_rewards_distributor_0_1_0"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/forta-network/forta-core-go/contracts/generated/contract_agent_registry_0_1_6"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_dispatch_0_1_5"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_forta_staking_0_1_2"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_scanner_node_version_0_1_1"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_scanner_pool_registry_0_1_0"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_scanner_registry_0_1_4"
	"github.com/forta-network/forta-core-go/contracts/generated/contract_stake_allocator_0_1_0"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/domain/registry"
	"github.com/forta-network/forta-core-go/domain/registry/regmsg"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/forta-network/forta-core-go/feeds"
	"github.com/forta-network/forta-core-go/utils"
)

type listener struct {
	ctx    context.Context
	cfg    ListenerConfig
	logs   feeds.LogFeed
	client Client
	eth    ethereum.Client

	publisher  MessagePublisher
	handlerReg *HandlerRegistry
	handler    regmsg.HandlerFunc[regmsg.Interface]
}

// MessagePublisher sends messages to a remote consumer.
type MessagePublisher interface {
	Publish(ctx context.Context, logger *log.Entry, msg regmsg.Interface) error
}

// MessageSubscriber receives messages from a remote producer.
type MessageSubscriber interface {
	Subscribe(ctx context.Context, handler regmsg.HandlerFunc[regmsg.Interface]) error
}

type ContractFilter struct {
	AgentRegistry       bool
	ScannerRegistry     bool
	ScannerPoolRegistry bool
	DispatchRegistry    bool
	FortaStaking        bool
	ScannerVersion      bool
	StakeAllocator      bool
	RewardDistributor   bool
}

type ListenerConfig struct {
	Name           string
	JsonRpcURL     string
	ENSAddress     string
	StartBlock     *big.Int
	EndBlock       *big.Int
	BlockOffset    int
	Handlers       Handlers
	ContractFilter *ContractFilter
	Topics         []string
	Publisher      MessagePublisher
	NoRefresh      bool
}

type Listener interface {
	Listen() error
	ProcessLastBlocks(blocksAgo int64) error
	ProcessBlockRange(startBlock *big.Int, endBlock *big.Int) error
}

func (l *listener) handleScannerPoolRegistryEvent(contracts *Contracts, le types.Log, blk *domain.Block, logger *log.Entry) error {
	switch getTopic(le) {
	case contract_scanner_pool_registry_0_1_0.ScannerUpdatedTopic:
		su, err := contracts.ScannerPoolRegFil.ParseScannerUpdated(le)
		if err != nil {
			return err
		}
		scannerID := utils.ScannerIDBigIntToHex(su.ScannerId)
		enabled, err := l.client.IsEnabledScanner(scannerID)
		if err != nil {
			return err
		}
		return l.handler(l.ctx, logger, registry.NewScannerSaveMessageFromPool(su, enabled, blk))

	case contract_scanner_pool_registry_0_1_0.ManagedStakeThresholdChangedTopic:
		evt, err := contracts.ScannerPoolRegFil.ParseManagedStakeThresholdChanged(le)
		if err != nil {
			return err
		}
		return l.handler(l.ctx, logger, registry.NewScannerManagedStakeThresholdMessage(evt, le, blk))

	case contract_scanner_pool_registry_0_1_0.TransferTopic:
		evt, err := contracts.ScannerPoolRegFil.ParseTransfer(le)
		if err != nil {
			return err
		}
		chainID, err := l.client.Contracts().ScannerPoolReg.MonitoredChainId(nil, evt.TokenId)
		if err != nil {
			return fmt.Errorf("failed to get the monitored chain id for the pool: %v", err)
		}
		return l.handler(l.ctx, logger, registry.NewScannerPoolMessageFromTransfer(evt, chainID, blk))

	case contract_scanner_pool_registry_0_1_0.ScannerPoolRegisteredTopic:
		evt, err := contracts.ScannerPoolRegFil.ParseScannerPoolRegistered(le)
		if err != nil {
			return err
		}
		owner, err := l.client.GetScannerPoolOwner(evt.ScannerPoolId)
		if err != nil {
			return fmt.Errorf("failed to get scanner pool owner: %v", err)
		}
		return l.handler(l.ctx, logger, registry.NewScannerPoolMessageFromRegistration(evt, owner, blk))

	case contract_scanner_pool_registry_0_1_0.EnabledScannersChangedTopic:
		evt, err := contracts.ScannerPoolRegFil.ParseEnabledScannersChanged(le)
		if err != nil {
			return err
		}
		chainID, err := l.client.Contracts().ScannerPoolReg.MonitoredChainId(nil, evt.ScannerPoolId)
		if err != nil {
			return fmt.Errorf("failed to get the monitored chain id for the pool: %v", err)
		}
		owner, err := l.client.GetScannerPoolOwner(evt.ScannerPoolId)
		if err != nil {
			return fmt.Errorf("failed to get scanner pool owner: %v", err)
		}
		return l.handler(l.ctx, logger, registry.NewScannerPoolMessageFromEnablement(evt, owner, chainID, blk))
	}
	return nil
}

func (l *listener) handleScannerVersionEvent(contracts *Contracts, le types.Log, blk *domain.Block, logger *log.Entry) error {
	switch getTopic(le) {
	case contract_scanner_node_version_0_1_1.ScannerNodeVersionUpdatedTopic:
		su, err := contracts.ScannerVersionFil.ParseScannerNodeVersionUpdated(le)
		if err != nil {
			return err
		}
		return l.handler(l.ctx, logger, registry.NewScannerNodeVersionUpdated(su, blk))
	}
	return nil
}

func (l *listener) handleAgentRegistryEvent(contracts *Contracts, le types.Log, blk *domain.Block, logger *log.Entry) error {
	switch getTopic(le) {
	case contract_agent_registry_0_1_6.AgentUpdatedTopic: // same with prev version
		au, err := contracts.AgentRegFil.ParseAgentUpdated(le)
		if err != nil {
			return err
		}
		agt, err := l.client.GetAgent(utils.AgentBigIntToHex(au.AgentId))
		if err != nil {
			return err
		}
		return l.handler(l.ctx, logger, registry.NewAgentSaveMessage(au, agt.Enabled, blk))

	case contract_agent_registry_0_1_6.AgentEnabledTopic: // same with prev version
		ae, err := contracts.AgentRegFil.ParseAgentEnabled(le)
		if err != nil {
			return err
		}
		return l.handler(l.ctx, logger, registry.NewAgentMessage(ae, blk))

	case contract_agent_registry_0_1_6.StakeThresholdChangedTopic: // same with prev version
		stc, err := contracts.AgentRegFil.ParseStakeThresholdChanged(le)
		if err != nil {
			return err
		}
		return l.handler(l.ctx, logger, registry.NewAgentStakeThresholdMessage(stc, le, blk))
	}
	return nil
}

func (l *listener) handleFortaStakingEvent(contracts *Contracts, le types.Log, blk *domain.Block, logger *log.Entry) error {
	var subjectType uint8
	var subjectID *big.Int
	var changeType string
	var value *big.Int
	var acct common.Address

	switch getTopic(le) {
	case contract_forta_staking_0_1_2.StakeDepositedTopic: // same with prev version
		evt, err := contracts.FortaStakingFil.ParseStakeDeposited(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		value = evt.Amount
		acct = evt.Account
		changeType = registry.ChangeTypeDeposit

	case contract_forta_staking_0_1_2.WithdrawalInitiatedTopic: // same with prev version
		evt, err := contracts.FortaStakingFil.ParseWithdrawalInitiated(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		changeType = registry.ChangeTypeWithdrawal

	case contract_forta_staking_0_1_2.FrozeTopic:
		evt, err := contracts.FortaStakingFil.ParseFroze(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		if evt.IsFrozen {
			changeType = registry.ChangeTypeFroze
		} else {
			changeType = registry.ChangeTypeUnfroze
		}

	case contract_forta_staking_0_1_2.SlashedTopic: // same with prev version
		evt, err := contracts.FortaStakingFil.ParseSlashed(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		value = evt.Value
		changeType = registry.ChangeTypeSlash

	case contract_forta_staking_0_1_2.TransferSingleTopic: // same with prev version
		evt, err := contracts.FortaStakingFil.ParseTransferSingle(le)
		if err != nil {
			return err
		}
		m, err := registry.TransferSharesMessageFromSingle(le, evt, blk)
		if err != nil {
			return err
		}
		return l.handler(l.ctx, logger, m)

	case contract_forta_staking_0_1_2.TransferBatchTopic: // same with prev version
		evt, err := contracts.FortaStakingFil.ParseTransferBatch(le)
		if err != nil {
			return err
		}
		ms, err := registry.TransferSharesMessagesFromBatch(le, evt, blk)
		if err != nil {
			return err
		}
		for _, m := range ms {
			if err := l.handler(l.ctx, logger, m); err != nil {
				return err
			}
		}
		return nil

	default:
		logger.Debug("unhandled topic, ignoring")
		return nil
	}

	// parse ID for agent or scanner
	switch subjectType {
	case SubjectTypeScanner:
		scannerID := utils.ScannerIDBigIntToHex(subjectID)
		return l.handler(l.ctx, logger, registry.NewScannerStakeMessage(le, changeType, acct, scannerID, value, blk))

	case SubjectTypeAgent:
		agentID := utils.AgentBigIntToHex(subjectID)
		return l.handler(l.ctx, logger, registry.NewAgentStakeMessage(le, changeType, acct, agentID, value, blk))

	case SubjectTypeScannerPool:
		poolID := subjectID.String()
		return l.handler(l.ctx, logger, registry.NewScannerPoolStakeMessage(le, changeType, acct, poolID, value, blk))

	default:
		logger.WithField("subjectID", subjectType).Warn("unhandled subject ID, ignoring")
	}
	return nil
}

func (l *listener) handleStakeAllocatorEvent(contracts *Contracts, le types.Log, blk *domain.Block, logger *log.Entry) error {
	switch getTopic(le) {
	case contract_stake_allocator_0_1_0.AllocatedStakeTopic:
		evt, err := contracts.StakeAllocatorFil.ParseAllocatedStake(le)
		if err != nil {
			return err
		}
		stakePerManaged, err := l.client.GetAllocatedStakePerManaged(big.NewInt(0).SetUint64(evt.Raw.BlockNumber), evt.Subject)
		if err != nil {
			return fmt.Errorf("failed to get stake per managed: %v", err)
		}
		return l.handler(l.ctx, logger, registry.NewScannerPoolAllocationMessage(le, blk, evt, stakePerManaged))
	}
	return nil
}

func (l *listener) handleDispatchEvent(contracts *Contracts, le types.Log, blk *domain.Block, logger *log.Entry) error {
	switch getTopic(le) {
	case contract_dispatch_0_1_5.LinkTopic: // same with prev version
		link, err := contracts.DispatchFil.ParseLink(le)
		if err != nil {
			return err
		}
		return l.handler(l.ctx, logger, registry.NewDispatchMessage(link, blk))

	case contract_dispatch_0_1_5.AlreadyLinkedTopic: // same with prev version
		link, err := contracts.DispatchFil.ParseAlreadyLinked(le)
		if err != nil {
			return err
		}
		return l.handler(l.ctx, logger, registry.NewAlreadyLinkedDispatchMessage(link, blk))
	}
	return nil
}

func (l *listener) handleRewardDistributorEvent(contracts *Contracts, le types.Log, blk *domain.Block, logger *log.Entry) error {
	switch getTopic(le) {
	case contract_rewards_distributor_0_1_0.RewardedTopic:
		link, err := contracts.RewardsDistributorFil.ParseRewarded(le)
		if err != nil {
			return err
		}
		return l.handler(l.ctx, logger, registry.NewRewardedMessage(link, blk))

	case contract_rewards_distributor_0_1_0.ClaimedRewardsTopic:
		link, err := contracts.RewardsDistributorFil.ParseClaimedRewards(le)
		if err != nil {
			return err
		}
		return l.handler(l.ctx, logger, registry.NewClaimedRewardsMessage(link, blk))
	}
	return nil
}

func isUpgradeOrMigration(le types.Log) bool {
	switch getTopic(le) {
	case UpgradedTopic,
		contract_forta_staking_0_1_2.StakeHelpersConfiguredTopic,
		contract_scanner_registry_0_1_4.ConfiguredMigrationTopic:
		return true

	default:
		return false
	}
}

func (l *listener) handleUpgradeEvent(contracts *Contracts, le types.Log, blk *domain.Block, logger *log.Entry) error {
	if !isUpgradeOrMigration(le) {
		return nil
	}

	// refresh contracts and log filter addresses
	if err := l.client.RefreshContracts(); err != nil {
		return err
	}
	l.setLogFilterAddrs()

	var upgradeMsg *registry.UpgradeMessage

	switch getTopic(le) {
	case UpgradedTopic:
		// use any contract's filterer to unpack the event - pick dispatch
		upgraded, err := contracts.DispatchFil.ParseUpgraded(le)
		if err != nil {
			return err
		}
		upgradeMsg = registry.NewUpgradeMessageFromUpgrade(upgraded, le, blk)

	default:
		upgradeMsg = registry.NewUpgradeMessageFromConfigurationChange(le.TxHash, le, blk)
	}

	return l.handler(l.ctx, logger, upgradeMsg)
}

func (l *listener) handleLog(blk *domain.Block, le types.Log) error {
	if l.ctx.Err() != nil {
		return l.ctx.Err()
	}

	logger := getLoggerForLog(le)
	logger.Info("handling tx log")
	contracts := l.client.Contracts()

	// always try to handle the upgrade event first because we don't enforce address checks
	// the addresses are limited by the filtered addresses
	// this is no-op if the topic is irrelevant
	// we refresh all contracts if it's relevant to any of the contracts
	if err := l.handleUpgradeEvent(contracts, le, blk, logger); err != nil {
		return err
	}

	if equalsAddress(le.Address, contracts.Addresses.ScannerNodeVersion.Hex()) {
		return l.handleScannerVersionEvent(contracts, le, blk, logger)
	}
	if equalsAddress(le.Address, contracts.Addresses.AgentRegistry.Hex()) {
		return l.handleAgentRegistryEvent(contracts, le, blk, logger)
	}
	if equalsAddress(le.Address, contracts.Addresses.Dispatch.Hex()) {
		return l.handleDispatchEvent(contracts, le, blk, logger)
	}
	if equalsAddress(le.Address, contracts.Addresses.FortaStaking.Hex()) {
		return l.handleFortaStakingEvent(contracts, le, blk, logger)
	}
	if equalsAddress(le.Address, contracts.Addresses.ScannerPoolRegistry.Hex()) {
		return l.handleScannerPoolRegistryEvent(contracts, le, blk, logger)
	}
	if equalsAddress(le.Address, contracts.Addresses.StakeAllocator.Hex()) {
		return l.handleStakeAllocatorEvent(contracts, le, blk, logger)
	}
	if equalsAddress(le.Address, contracts.Addresses.Rewards.Hex()) {
		return l.handleRewardDistributorEvent(contracts, le, blk, logger)
	}
	return nil
}

func (l *listener) handleAfterBlock(blk *domain.Block) error {
	if l.ctx.Err() != nil {
		return l.ctx.Err()
	}
	if l.handlerReg.afterBlockHandler != nil {
		return l.handlerReg.afterBlockHandler(blk)
	}
	return nil
}

type page struct {
	Start int64
	End   int64
}

// ProcessBlockRange pages over a range of blocks, 10k blocks per page
func (l *listener) ProcessBlockRange(startBlock *big.Int, endBlock *big.Int) error {
	start := startBlock
	pageSize := big.NewInt(2000)
	if endBlock == nil {
		bn, err := l.eth.BlockNumber(context.Background())
		if err != nil {
			return err
		}
		endBlock = bn
	}
	end := math.BigMin(big.NewInt(0).Add(start, pageSize), endBlock)
	pages := make(chan page)
	grp, ctx := errgroup.WithContext(l.ctx)
	mux := sync.Mutex{}

	for i := 0; i < 25; i++ {
		grp.Go(func() error {
			for p := range pages {
				if ctx.Err() != nil {
					return ctx.Err()
				}
				logs, err := l.logs.GetLogsForRange(big.NewInt(p.Start), big.NewInt(p.End))
				if err != nil {
					return err
				}
				var block *domain.Block
				for _, lg := range logs {
					if block == nil || block.Number != utils.BigIntToHex(big.NewInt(int64(lg.BlockNumber))) {
						blk, err := l.eth.BlockByNumber(l.ctx, big.NewInt(int64(lg.BlockNumber)))
						if err != nil {
							return err
						}
						block = blk
					}
					// avoids concurrent handleLogs
					mux.Lock()
					err := l.handleLog(block, lg)
					mux.Unlock()
					if err != nil {
						return err
					}
				}
			}
			return nil
		})
	}

	grp.Go(func() error {
		defer close(pages)
		for end.Cmp(endBlock) <= 0 {
			pages <- page{
				Start: start.Int64(),
				End:   end.Int64(),
			}
			if end.Cmp(endBlock) == 0 {
				return nil
			}
			start = big.NewInt(0).Add(end, big.NewInt(1))
			end = math.BigMin(big.NewInt(0).Add(start, pageSize), endBlock)
		}
		return nil
	})

	return grp.Wait()
}

// ProcessLastBlocks fetches the logs in a single pass and calls handlers for them
func (l *listener) ProcessLastBlocks(blocksAgo int64) error {
	bn, err := l.eth.BlockNumber(context.Background())
	if err != nil {
		return err
	}
	if bn.Int64() == 0 {
		return errors.New("current block is unexpectedly 0")
	}
	start := bn
	end := big.NewInt(bn.Int64() - blocksAgo)
	return l.ProcessBlockRange(start, end)
}

func (l *listener) Listen() error {
	return l.logs.ForEachLog(l.handleLog, l.handleAfterBlock)
}

func NewDefaultListener(ctx context.Context, handlers Handlers) (*listener, error) {
	return NewListener(ctx, ListenerConfig{
		Name:       "listener",
		JsonRpcURL: defaultConfig.JsonRpcUrl,
		ENSAddress: defaultEnsAddress,
		Handlers:   handlers,
	})
}

func NewListener(ctx context.Context, cfg ListenerConfig) (*listener, error) {
	jsonRpc := cfg.JsonRpcURL
	if jsonRpc == "" {
		jsonRpc = defaultConfig.JsonRpcUrl
	}
	ensAddr := cfg.ENSAddress
	if ensAddr == "" {
		ensAddr = defaultEnsAddress
	}

	ethClient, err := ethereum.NewStreamEthClient(ctx, cfg.Name, jsonRpc)
	if err != nil {
		return nil, err
	}

	regClient, err := NewClient(ctx, ClientConfig{
		JsonRpcUrl: jsonRpc,
		ENSAddress: ensAddr,
		Name:       "registry-listener",
		NoRefresh:  cfg.NoRefresh,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create registry client: %v", err)
	}

	return NewListenerWithClients(ctx, cfg, ethClient, regClient, cfg.Publisher)
}

func (l *listener) setLogFilterAddrs() {
	filter := l.cfg.ContractFilter
	regContracts := l.client.Contracts().Addresses

	var addrs []string

	if filter != nil {
		if filter.AgentRegistry {
			addrs = append(addrs, regContracts.AgentRegistry.Hex())
		}
		if filter.ScannerRegistry {
			addrs = append(addrs, regContracts.ScannerRegistry.Hex())
		}
		if filter.FortaStaking {
			addrs = append(addrs, regContracts.FortaStaking.Hex())
		}
		if filter.DispatchRegistry {
			addrs = append(addrs, regContracts.Dispatch.Hex())
		}
		if filter.ScannerVersion {
			addrs = append(addrs, regContracts.ScannerNodeVersion.Hex())
		}
		if filter.ScannerPoolRegistry {
			addrs = append(addrs, regContracts.ScannerPoolRegistry.Hex())
		}
		if filter.StakeAllocator {
			addrs = append(addrs, regContracts.StakeAllocator.Hex())
		}
		if filter.RewardDistributor {
			addrs = append(addrs, regContracts.Rewards.Hex())
		}
	} else {
		// include all contracts
		addrs = getAllContractAddrs(regContracts)
	}

	if len(addrs) == 0 {
		panic("empty filter")
	}

	for _, addr := range addrs {
		l.logs.AddAddress(addr)
	}
}

func NewListenerWithClients(ctx context.Context, cfg ListenerConfig, ethClient ethereum.Client, regClient Client, publisher MessagePublisher) (*listener, error) {
	li := &listener{
		ctx:        ctx,
		client:     regClient,
		cfg:        cfg,
		eth:        ethClient,
		publisher:  publisher,
		handlerReg: NewHandlerRegistry(cfg.Handlers),
	}
	// if there is a publisher, publish only
	// otherwise, handle directly
	if publisher != nil {
		li.handler = li.publisher.Publish
	} else {
		li.handler = li.handlerReg.Handle
	}

	var topics [][]string
	if len(cfg.Topics) > 0 {
		topics = [][]string{cfg.Topics}
	}

	var err error
	li.logs, err = feeds.NewLogFeed(ctx, ethClient, feeds.LogFeedConfig{
		Topics:     topics,
		StartBlock: cfg.StartBlock,
		EndBlock:   cfg.EndBlock,
		Offset:     cfg.BlockOffset,
	})
	if err != nil {
		return nil, err
	}
	li.setLogFilterAddrs()

	return li, nil
}

// StartWorker handles incoming messages.
func StartWorker(ctx context.Context, subscriber MessageSubscriber, handlers Handlers) error {
	handlerReg := NewHandlerRegistry(handlers)
	return subscriber.Subscribe(ctx, handlerReg.Handle)
}

// ListenToUpgrades listens for contract upgrades and refreshes the contracts.
func ListenToUpgrades(ctx context.Context, client Client, blockFeed feeds.BlockFeed) <-chan error {
	return blockFeed.Subscribe(func(evt *domain.BlockEvent) error {
		regContracts := client.Contracts().Addresses
		addrs := getAllContractAddrs(regContracts)
		for _, evtLog := range evt.Logs {
			ethLog := evtLog.ToTypesLog()
			if !isAddrIn(addrs, ethLog.Address) {
				continue
			}
			if !isUpgradeOrMigration(ethLog) {
				continue
			}
			if err := client.RefreshContracts(); err != nil {
				log.WithError(err).Warn("failed to refresh contracts")
			}
			// we were able to match the topic and the address and attempted to handle
			// nothing else left to do for this block
			break
		}
		return nil
	})
}
