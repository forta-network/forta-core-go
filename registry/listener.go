package registry

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/forta-network/forta-core-go/contracts/contract_agent_registry"
	"github.com/forta-network/forta-core-go/contracts/contract_dispatch"
	"github.com/forta-network/forta-core-go/contracts/contract_forta_staking"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_node_version"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_pool_registry"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_registry"
	"github.com/forta-network/forta-core-go/contracts/contract_stake_allocator"
	"github.com/forta-network/forta-core-go/domain"
	"github.com/forta-network/forta-core-go/domain/registry"
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

	scannerAddr        string
	scannerVersionAddr string
	agentAddr          string
	dispatchAddr       string
	fortaStakingAddr   string
	scannerPoolAddr    *string
	stakeAllocatorAddr *string

	fortaStakingFilterer   *contract_forta_staking.FortaStakingFilterer
	scannerFilterer        *contract_scanner_registry.ScannerRegistryFilterer
	scannerPoolFilterer    *contract_scanner_pool_registry.ScannerPoolRegistryFilterer
	agentsFilterer         *contract_agent_registry.AgentRegistryFilterer
	dispatchFilterer       *contract_dispatch.DispatchFilterer
	scannerVersionFilterer *contract_scanner_node_version.ScannerNodeVersionFilterer
	stakeAllocatorFilterer *contract_stake_allocator.StakeAllocatorFilterer
}

type Handlers struct {
	AfterBlockHandler func(blk *domain.Block) error

	// registration
	SaveAgentHandler         func(logger *log.Entry, msg *registry.AgentSaveMessage) error
	AgentActionHandler       func(logger *log.Entry, msg *registry.AgentMessage) error
	SaveScannerHandler       func(logger *log.Entry, msg *registry.ScannerSaveMessage) error
	ScannerActionHandler     func(logger *log.Entry, msg *registry.ScannerMessage) error
	UpdateScannerPoolHandler func(logger *log.Entry, msg *registry.UpdateScannerPoolMessage) error

	// scanner node versions
	ScannerNodeVersionHandler func(logger *log.Entry, msg *registry.ScannerNodeVersionMessage) error

	// assignment
	DispatchHandler func(logger *log.Entry, msg *registry.DispatchMessage) error

	// staking
	AgentStakeHandler       func(logger *log.Entry, msg *registry.AgentStakeMessage) error
	ScannerStakeHandler     func(logger *log.Entry, msg *registry.ScannerStakeMessage) error
	TransferSharesHandler   func(logger *log.Entry, msg *registry.TransferSharesMessage) error
	ScannerPoolStakeHandler func(logger *log.Entry, msg *registry.ScannerPoolStakeMessage) error

	AgentStakeThresholdHandler   func(logger *log.Entry, msg *registry.AgentStakeThresholdMessage) error
	ScannerStakeThresholdHandler func(logger *log.Entry, msg *registry.ScannerStakeThresholdMessage) error

	ScannerPoolAllocationHandler func(logger *log.Entry, msg *registry.ScannerPoolAllocationMessage) error
}

type ContractFilter struct {
	AgentRegistry       bool
	ScannerRegistry     bool
	ScannerPoolRegistry bool
	DispatchRegistry    bool
	FortaStaking        bool
	ScannerVersion      bool
	StakeAllocator      bool
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
}

type Listener interface {
	Listen() error
	ProcessLastBlocks(blocksAgo int64) error
	ProcessBlockRange(startBlock *big.Int, endBlock *big.Int) error
}

func (l *listener) handleScannerRegistryEvent(le types.Log, blk *domain.Block, logger *log.Entry) error {
	switch getTopic(le) {
	case contract_scanner_registry.ScannerUpdatedTopic:
		su, err := l.scannerFilterer.ParseScannerUpdated(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.SaveScannerHandler != nil {
			scannerID := utils.ScannerIDBigIntToHex(su.ScannerId)
			enabled, err := l.client.IsEnabledScanner(scannerID)
			if err != nil {
				return err
			}
			return l.cfg.Handlers.SaveScannerHandler(logger, registry.NewScannerSaveMessage(su, enabled, blk))
		}

	case contract_scanner_registry.ScannerEnabledTopic:
		se, err := l.scannerFilterer.ParseScannerEnabled(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.ScannerActionHandler != nil {
			return l.cfg.Handlers.ScannerActionHandler(logger, registry.NewScannerMessage(se, blk))
		}

	case contract_scanner_registry.StakeThresholdChangedTopic:
		evt, err := l.scannerFilterer.ParseStakeThresholdChanged(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.ScannerStakeThresholdHandler != nil {
			return l.cfg.Handlers.ScannerStakeThresholdHandler(logger, registry.NewScannerStakeThresholdMessage(evt, le, blk))
		}
	}
	return nil
}

func (l *listener) handleScannerPoolRegistryEvent(le types.Log, blk *domain.Block, logger *log.Entry) error {
	switch getTopic(le) {
	case contract_scanner_pool_registry.ScannerUpdatedTopic:
		su, err := l.scannerPoolFilterer.ParseScannerUpdated(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.SaveScannerHandler != nil {
			scannerID := utils.ScannerIDBigIntToHex(su.ScannerId)
			enabled, err := l.client.IsEnabledScanner(scannerID)
			if err != nil {
				return err
			}
			return l.cfg.Handlers.SaveScannerHandler(logger, registry.NewScannerSaveMessageFromPool(su, enabled, blk))
		}

	case contract_scanner_pool_registry.ScannerEnabledTopic:
		se, err := l.scannerPoolFilterer.ParseScannerEnabled(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.ScannerActionHandler != nil {
			return l.cfg.Handlers.ScannerActionHandler(logger, registry.NewScannerMessageFromPool(se, blk))
		}

	case contract_scanner_pool_registry.ManagedStakeThresholdChangedTopic:
		evt, err := l.scannerPoolFilterer.ParseManagedStakeThresholdChanged(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.ScannerStakeThresholdHandler != nil {
			return l.cfg.Handlers.ScannerStakeThresholdHandler(logger, registry.NewScannerManagedStakeThresholdMessage(evt, le, blk))
		}

	case contract_scanner_pool_registry.TransferTopic:
		evt, err := l.scannerPoolFilterer.ParseTransfer(le)
		if err != nil {
			return err
		}
		// detect the new owner only if not minted
		if evt.From.Hex() == utils.ZeroAddress {
			return nil
		}
		if l.cfg.Handlers.UpdateScannerPoolHandler != nil {
			return l.cfg.Handlers.UpdateScannerPoolHandler(logger, registry.NewScannerPoolMessageFromTransfer(evt, blk))
		}

	case contract_scanner_pool_registry.ScannerPoolRegisteredTopic:
		evt, err := l.scannerPoolFilterer.ParseScannerPoolRegistered(le)
		if err != nil {
			return err
		}
		owner, err := l.client.GetScannerPoolOwner(evt.ScannerPoolId)
		if err != nil {
			return fmt.Errorf("failed to get scanner pool owner: %v", err)
		}
		if l.cfg.Handlers.UpdateScannerPoolHandler != nil {
			return l.cfg.Handlers.UpdateScannerPoolHandler(logger, registry.NewScannerPoolMessageFromRegistration(evt, owner, blk))
		}
	}
	return nil
}

func (l *listener) handleScannerVersionEvent(le types.Log, blk *domain.Block, logger *log.Entry) error {
	switch getTopic(le) {
	case contract_scanner_node_version.ScannerNodeVersionUpdatedTopic:
		su, err := l.scannerVersionFilterer.ParseScannerNodeVersionUpdated(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.ScannerNodeVersionHandler != nil {
			return l.cfg.Handlers.ScannerNodeVersionHandler(logger, registry.NewScannerNodeVersionUpdated(su, blk))
		}
	}
	return nil
}

func (l *listener) handleAgentRegistryEvent(le types.Log, blk *domain.Block, logger *log.Entry) error {
	switch getTopic(le) {
	case contract_agent_registry.AgentUpdatedTopic:
		au, err := l.agentsFilterer.ParseAgentUpdated(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.SaveAgentHandler != nil {
			agt, err := l.client.GetAgent(utils.AgentBigIntToHex(au.AgentId))
			if err != nil {
				return err
			}
			return l.cfg.Handlers.SaveAgentHandler(logger, registry.NewAgentSaveMessage(au, agt.Enabled, blk))
		}

	case contract_agent_registry.AgentEnabledTopic:
		ae, err := l.agentsFilterer.ParseAgentEnabled(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.AgentActionHandler != nil {
			return l.cfg.Handlers.AgentActionHandler(logger, registry.NewAgentMessage(ae, blk))
		}

	case contract_agent_registry.StakeThresholdChangedTopic:
		stc, err := l.agentsFilterer.ParseStakeThresholdChanged(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.AgentStakeThresholdHandler != nil {
			return l.cfg.Handlers.AgentStakeThresholdHandler(logger, registry.NewAgentStakeThresholdMessage(stc, le, blk))
		}
	}
	return nil
}

func (l *listener) handleFortaStakingEvent(le types.Log, blk *domain.Block, logger *log.Entry) error {
	var subjectType uint8
	var subjectID *big.Int
	var changeType string
	var value *big.Int
	var acct common.Address

	switch getTopic(le) {
	case contract_forta_staking.StakeDepositedTopic:
		evt, err := l.fortaStakingFilterer.ParseStakeDeposited(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		value = evt.Amount
		acct = evt.Account
		changeType = registry.ChangeTypeDeposit

	case contract_forta_staking.WithdrawalInitiatedTopic:
		evt, err := l.fortaStakingFilterer.ParseWithdrawalInitiated(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		changeType = registry.ChangeTypeWithdrawal

	case contract_forta_staking.SlashedTopic:
		evt, err := l.fortaStakingFilterer.ParseSlashed(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		value = evt.Value
		changeType = registry.ChangeTypeSlash

	case contract_forta_staking.TransferSingleTopic:
		evt, err := l.fortaStakingFilterer.ParseTransferSingle(le)
		if err != nil {
			return err
		}
		m, err := registry.TransferSharesMessageFromSingle(le, evt, blk)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.TransferSharesHandler != nil {
			return l.cfg.Handlers.TransferSharesHandler(logger, m)
		}

	case contract_forta_staking.TransferBatchTopic:
		evt, err := l.fortaStakingFilterer.ParseTransferBatch(le)
		if err != nil {
			return err
		}
		ms, err := registry.TransferSharesMessagesFromBatch(le, evt, blk)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.TransferSharesHandler != nil {
			for _, m := range ms {
				if err := l.cfg.Handlers.TransferSharesHandler(logger, m); err != nil {
					return err
				}
			}
			return nil
		}

	default:
		logger.Debug("unhandled topic, ignoring")
		return nil
	}

	// parse ID for agent or scanner
	switch subjectType {
	case SubjectTypeScanner:
		scannerID := utils.ScannerIDBigIntToHex(subjectID)
		if l.cfg.Handlers.ScannerStakeHandler != nil {
			return l.cfg.Handlers.ScannerStakeHandler(logger, registry.NewScannerStakeMessage(le, changeType, acct, scannerID, value, blk))
		}

	case SubjectTypeAgent:
		agentID := utils.AgentBigIntToHex(subjectID)
		if l.cfg.Handlers.AgentStakeHandler != nil {
			return l.cfg.Handlers.AgentStakeHandler(logger, registry.NewAgentStakeMessage(le, changeType, acct, agentID, value, blk))
		}

	case SubjectTypeScannerPool:
		poolID := subjectID.String()
		if l.cfg.Handlers.ScannerPoolStakeHandler != nil {
			return l.cfg.Handlers.ScannerPoolStakeHandler(logger, registry.NewScannerPoolStakeMessage(le, changeType, acct, poolID, value, blk))
		}

	default:
		logger.WithField("subjectID", subjectType).Warn("unhandled subject ID, ignoring")
	}
	return nil
}

func (l *listener) handleStakeAllocatorEvent(le types.Log, blk *domain.Block, logger *log.Entry) error {
	switch getTopic(le) {
	case contract_stake_allocator.AllocatedStakeTopic:
		evt, err := l.stakeAllocatorFilterer.ParseAllocatedStake(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.ScannerPoolAllocationHandler != nil {
			return l.cfg.Handlers.ScannerPoolAllocationHandler(logger, registry.NewScannerPoolAllocationMessage(le, blk, evt))
		}
	}
	return nil
}

func (l *listener) handleDispatchEvent(le types.Log, blk *domain.Block, logger *log.Entry) error {
	switch getTopic(le) {
	case contract_dispatch.LinkTopic:
		link, err := l.dispatchFilterer.ParseLink(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.DispatchHandler != nil {
			return l.cfg.Handlers.DispatchHandler(logger, registry.NewDispatchMessage(link, blk))
		}

	case contract_dispatch.AlreadyLinkedTopic:
		link, err := l.dispatchFilterer.ParseAlreadyLinked(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.DispatchHandler != nil {
			return l.cfg.Handlers.DispatchHandler(logger, registry.NewAlreadyLinkedDispatchMessage(link, blk))
		}
	}
	return nil
}

func (l *listener) handleLog(blk *domain.Block, le types.Log) error {
	if l.ctx.Err() != nil {
		return l.ctx.Err()
	}
	logger := getLoggerForLog(le)
	if equalsAddress(le.Address, l.scannerAddr) {
		return l.handleScannerRegistryEvent(le, blk, logger)
	}
	if equalsAddress(le.Address, l.scannerVersionAddr) {
		return l.handleScannerVersionEvent(le, blk, logger)
	}
	if equalsAddress(le.Address, l.agentAddr) {
		return l.handleAgentRegistryEvent(le, blk, logger)
	}
	if equalsAddress(le.Address, l.dispatchAddr) {
		return l.handleDispatchEvent(le, blk, logger)
	}
	if equalsAddress(le.Address, l.fortaStakingAddr) {
		return l.handleFortaStakingEvent(le, blk, logger)
	}
	if l.scannerPoolAddr != nil && equalsAddress(le.Address, *l.scannerPoolAddr) {
		return l.handleScannerPoolRegistryEvent(le, blk, logger)
	}
	if l.stakeAllocatorAddr != nil && equalsAddress(le.Address, *l.stakeAllocatorAddr) {
		return l.handleStakeAllocatorEvent(le, blk, logger)
	}
	return nil
}

func (l *listener) handleAfterBlock(blk *domain.Block) error {
	if l.ctx.Err() != nil {
		return l.ctx.Err()
	}
	if l.cfg.Handlers.AfterBlockHandler != nil {
		return l.cfg.Handlers.AfterBlockHandler(blk)
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
	pageSize := big.NewInt(10000)
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
	})
	if err != nil {
		return nil, err
	}

	return NewListenerWithClients(ctx, cfg, ethClient, regClient)
}

func NewListenerWithClients(ctx context.Context, cfg ListenerConfig, ethClient ethereum.Client, regClient Client) (*listener, error) {
	regContracts := regClient.RegistryContracts()

	sf, err := contract_scanner_registry.NewScannerRegistryFilterer(regContracts.ScannerRegistry, nil)
	if err != nil {
		return nil, err
	}

	snv, err := contract_scanner_node_version.NewScannerNodeVersionFilterer(regContracts.ScannerNodeVersion, nil)
	if err != nil {
		return nil, err
	}

	af, err := contract_agent_registry.NewAgentRegistryFilterer(regContracts.AgentRegistry, nil)
	if err != nil {
		return nil, err
	}

	df, err := contract_dispatch.NewDispatchFilterer(regContracts.Dispatch, nil)
	if err != nil {
		return nil, err
	}

	stkf, err := contract_forta_staking.NewFortaStakingFilterer(regContracts.FortaStaking, nil)
	if err != nil {
		return nil, err
	}

	var spf *contract_scanner_pool_registry.ScannerPoolRegistryFilterer
	if regContracts.ScannerPoolRegistry != nil {
		spf, err = contract_scanner_pool_registry.NewScannerPoolRegistryFilterer(*regContracts.ScannerPoolRegistry, nil)
		if err != nil {
			return nil, err
		}
	}

	var saf *contract_stake_allocator.StakeAllocatorFilterer
	if regContracts.StakeAllocator != nil {
		saf, err = contract_stake_allocator.NewStakeAllocatorFilterer(*regContracts.StakeAllocator, nil)
		if err != nil {
			return nil, err
		}
	}

	var addrs []string
	if cfg.ContractFilter != nil {
		if cfg.ContractFilter.AgentRegistry {
			addrs = append(addrs, regContracts.AgentRegistry.Hex())
		}
		if cfg.ContractFilter.ScannerRegistry {
			addrs = append(addrs, regContracts.ScannerRegistry.Hex())
		}
		if cfg.ContractFilter.FortaStaking {
			addrs = append(addrs, regContracts.FortaStaking.Hex())
		}
		if cfg.ContractFilter.DispatchRegistry {
			addrs = append(addrs, regContracts.Dispatch.Hex())
		}
		if cfg.ContractFilter.ScannerVersion {
			addrs = append(addrs, regContracts.ScannerNodeVersion.Hex())
		}
		if cfg.ContractFilter.ScannerPoolRegistry && regContracts.ScannerPoolRegistry != nil {
			addrs = append(addrs, regContracts.ScannerPoolRegistry.Hex())
		}
		if cfg.ContractFilter.StakeAllocator && regContracts.StakeAllocator != nil {
			addrs = append(addrs, regContracts.StakeAllocator.Hex())
		}

	} else {
		addrs = []string{
			regContracts.AgentRegistry.Hex(),
			regContracts.ScannerRegistry.Hex(),
			regContracts.Dispatch.Hex(),
			regContracts.FortaStaking.Hex(),
			regContracts.ScannerNodeVersion.Hex(),
		}
		if regContracts.ScannerPoolRegistry != nil {
			addrs = append(addrs, regContracts.ScannerPoolRegistry.Hex())
		}
		if regContracts.StakeAllocator != nil {
			addrs = append(addrs, regContracts.StakeAllocator.Hex())
		}
	}

	var topics [][]string
	if len(cfg.Topics) > 0 {
		topics = [][]string{cfg.Topics}
	}

	logFeed, err := feeds.NewLogFeed(ctx, ethClient, feeds.LogFeedConfig{
		Addresses:  addrs,
		Topics:     topics,
		StartBlock: cfg.StartBlock,
		EndBlock:   cfg.EndBlock,
		Offset:     cfg.BlockOffset,
	})

	if err != nil {
		return nil, err
	}

	var scannerPoolAddr string
	if regContracts.ScannerPoolRegistry != nil {
		scannerPoolAddr = regContracts.ScannerPoolRegistry.Hex()
	}
	var stakeAllocatorAddr string
	if regContracts.StakeAllocator != nil {
		stakeAllocatorAddr = regContracts.StakeAllocator.Hex()
	}

	return &listener{
		ctx:    ctx,
		client: regClient,
		cfg:    cfg,
		logs:   logFeed,
		eth:    ethClient,

		scannerAddr:        regContracts.ScannerRegistry.Hex(),
		scannerVersionAddr: regContracts.ScannerNodeVersion.Hex(),
		agentAddr:          regContracts.AgentRegistry.Hex(),
		dispatchAddr:       regContracts.Dispatch.Hex(),
		fortaStakingAddr:   regContracts.FortaStaking.Hex(),
		scannerPoolAddr:    &scannerPoolAddr,
		stakeAllocatorAddr: &stakeAllocatorAddr,

		scannerFilterer:        sf,
		scannerVersionFilterer: snv,
		agentsFilterer:         af,
		dispatchFilterer:       df,
		fortaStakingFilterer:   stkf,
		scannerPoolFilterer:    spf,
		stakeAllocatorFilterer: saf,
	}, nil
}
