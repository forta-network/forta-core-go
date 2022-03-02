package registry

import (
	"context"
	"github.com/forta-protocol/forta-core-go/contracts/contract_forta_staking"
	"github.com/forta-protocol/forta-core-go/utils"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"

	"github.com/forta-protocol/forta-core-go/contracts/contract_agent_registry"
	"github.com/forta-protocol/forta-core-go/contracts/contract_dispatch"
	"github.com/forta-protocol/forta-core-go/contracts/contract_scanner_registry"
	"github.com/forta-protocol/forta-core-go/domain"
	"github.com/forta-protocol/forta-core-go/domain/registry"
	"github.com/forta-protocol/forta-core-go/ens"
	"github.com/forta-protocol/forta-core-go/ethereum"
	"github.com/forta-protocol/forta-core-go/feeds"
)

type listener struct {
	ctx  context.Context
	cfg  ListenerConfig
	logs feeds.LogFeed

	scannerAddr      string
	agentAddr        string
	dispatchAddr     string
	fortaStakingAddr string

	fortaStakingFilterer *contract_forta_staking.FortaStakingFilterer
	scannerFilterer      *contract_scanner_registry.ScannerRegistryFilterer
	agentsFilterer       *contract_agent_registry.AgentRegistryFilterer
	dispatchFilterer     *contract_dispatch.DispatchFilterer
}

type Handlers struct {
	AfterBlockHandler func(blk *domain.Block) error

	// registration
	SaveAgentHandler     func(logger *log.Entry, msg *registry.AgentSaveMessage) error
	AgentActionHandler   func(logger *log.Entry, msg *registry.AgentMessage) error
	SaveScannerHandler   func(logger *log.Entry, msg *registry.ScannerSaveMessage) error
	ScannerActionHandler func(logger *log.Entry, msg *registry.ScannerMessage) error

	// assignment
	DispatchHandler func(logger *log.Entry, msg *registry.DispatchMessage) error

	// staking
	AgentStakeHandler            func(logger *log.Entry, msg *registry.AgentStakeMessage) error
	ScannerStakeHandler          func(logger *log.Entry, msg *registry.ScannerStakeMessage) error
	AgentStakeThresholdHandler   func(logger *log.Entry, msg *registry.AgentStakeThresholdMessage) error
	ScannerStakeThresholdHandler func(logger *log.Entry, msg *registry.ScannerStakeThresholdMessage) error
}

type ListenerConfig struct {
	Name        string
	JsonRpcURL  string
	ENSAddress  string
	StartBlock  *big.Int
	BlockOffset int
	Handlers    Handlers
}

type Listener interface {
	Listen() error
	ProcessLastBlocks(blocksAgo int64) error
}

func (l *listener) handleScannerRegistryEvent(le types.Log, logger *log.Entry) error {
	if isEvent(le, contract_scanner_registry.ScannerUpdatedTopic) {
		su, err := l.scannerFilterer.ParseScannerUpdated(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.SaveScannerHandler != nil {
			return l.cfg.Handlers.SaveScannerHandler(logger, registry.NewScannerSaveMessage(su))
		}
	} else if isEvent(le, contract_scanner_registry.ScannerEnabledTopic) {
		se, err := l.scannerFilterer.ParseScannerEnabled(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.ScannerActionHandler != nil {
			return l.cfg.Handlers.ScannerActionHandler(logger, registry.NewScannerMessage(se))
		}
	} else if isEvent(le, contract_scanner_registry.StakeThresholdChangedTopic) {
		evt, err := l.scannerFilterer.ParseStakeThresholdChanged(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.ScannerStakeThresholdHandler != nil {
			return l.cfg.Handlers.ScannerStakeThresholdHandler(logger, registry.NewScannerStakeThresholdMessage(evt.ChainId.Int64()))
		}
	}
	return nil
}

func (l *listener) handleAgentRegistryEvent(le types.Log, logger *log.Entry) error {
	if isEvent(le, contract_agent_registry.AgentUpdatedTopic) {
		au, err := l.agentsFilterer.ParseAgentUpdated(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.SaveAgentHandler != nil {
			return l.cfg.Handlers.SaveAgentHandler(logger, registry.NewAgentSaveMessage(au))
		}
	} else if isEvent(le, contract_agent_registry.AgentEnabledTopic) {
		ae, err := l.agentsFilterer.ParseAgentEnabled(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.AgentActionHandler != nil {
			return l.cfg.Handlers.AgentActionHandler(logger, registry.NewAgentMessage(ae))
		}
	} else if isEvent(le, contract_agent_registry.StakeThresholdChangedTopic) {
		if l.cfg.Handlers.AgentStakeThresholdHandler != nil {
			return l.cfg.Handlers.AgentStakeThresholdHandler(logger, registry.NewAgentStakeThresholdMessage())
		}
	}
	return nil
}

func (l *listener) handleFortaStakingEvent(le types.Log, logger *log.Entry) error {
	var subjectType uint8
	var subjectID *big.Int
	var changeType string

	if isEvent(le, contract_forta_staking.StakeDepositedTopic) {
		evt, err := l.fortaStakingFilterer.ParseStakeDeposited(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		changeType = registry.ChangeTypeDeposit
	} else if isEvent(le, contract_forta_staking.WithdrawalInitiatedTopic) {
		evt, err := l.fortaStakingFilterer.ParseWithdrawalInitiated(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		changeType = registry.ChangeTypeWithdrawal
	} else if isEvent(le, contract_forta_staking.SlashedTopic) {
		evt, err := l.fortaStakingFilterer.ParseSlashed(le)
		if err != nil {
			return err
		}
		subjectType = evt.SubjectType
		subjectID = evt.Subject
		changeType = registry.ChangeTypeSlash
	} else {
		logger.Warn("unhandled topic, ignoring")
		return nil
	}

	// parse ID for agent or scanner
	if subjectType == SubjectTypeScanner {
		scannerID := utils.ScannerIDBigIntToHex(subjectID)
		if l.cfg.Handlers.ScannerStakeHandler != nil {
			return l.cfg.Handlers.ScannerStakeHandler(logger, registry.NewScannerStakeMessage(changeType, scannerID))
		}
	} else if subjectType == SubjectTypeAgent {
		agentID := utils.AgentBigIntToHex(subjectID)
		if l.cfg.Handlers.AgentStakeHandler != nil {
			return l.cfg.Handlers.AgentStakeHandler(logger, registry.NewAgentStakeMessage(changeType, agentID))
		}
	} else {
		logger.WithField("subjectID", subjectType).Warn("unhandled subject ID, ignoring")
	}

	return nil
}

func (l *listener) handleDispatcherEvent(le types.Log, logger *log.Entry) error {
	if isEvent(le, contract_dispatch.LinkTopic) {
		link, err := l.dispatchFilterer.ParseLink(le)
		if err != nil {
			return err
		}
		if l.cfg.Handlers.DispatchHandler != nil {
			return l.cfg.Handlers.DispatchHandler(logger, registry.NewDispatchMessage(link))
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
		return l.handleScannerRegistryEvent(le, logger)
	}
	if equalsAddress(le.Address, l.agentAddr) {
		return l.handleAgentRegistryEvent(le, logger)
	}
	if equalsAddress(le.Address, l.dispatchAddr) {
		return l.handleDispatcherEvent(le, logger)
	}
	if equalsAddress(le.Address, l.fortaStakingAddr) {
		return l.handleFortaStakingEvent(le, logger)
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

// ProcessLogs fetches the logs in a single pass and calls handlers for them
func (l *listener) ProcessLastBlocks(blocksAgo int64) error {
	logs, err := l.logs.GetLogsForLastBlocks(blocksAgo)
	if err != nil {
		return err
	}
	for _, lg := range logs {
		if err := l.handleLog(nil, lg); err != nil {
			return err
		}
	}
	return nil
}

func (l *listener) Listen() error {
	return l.logs.ForEachLog(l.handleLog, l.handleAfterBlock)
}

func NewListener(ctx context.Context, cfg ListenerConfig) (*listener, error) {
	client, err := ethereum.NewStreamEthClient(ctx, cfg.Name, cfg.JsonRpcURL)
	if err != nil {
		return nil, err
	}
	ensStore, err := ens.DialENSStoreAt(cfg.JsonRpcURL, cfg.ENSAddress)
	if err != nil {
		return nil, err
	}

	regContracts, err := ensStore.ResolveRegistryContracts()

	sf, err := contract_scanner_registry.NewScannerRegistryFilterer(regContracts.ScannerRegistry, nil)
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

	logFeed, err := feeds.NewLogFeed(ctx, client, feeds.LogFeedConfig{
		Addresses:  []string{regContracts.AgentRegistry.Hex(), regContracts.ScannerRegistry.Hex(), regContracts.Dispatch.Hex(), regContracts.FortaStaking.Hex()},
		StartBlock: cfg.StartBlock,
		Offset:     cfg.BlockOffset,
	})

	if err != nil {
		return nil, err
	}

	return &listener{
		ctx:                  ctx,
		cfg:                  cfg,
		logs:                 logFeed,
		scannerAddr:          regContracts.ScannerRegistry.Hex(),
		agentAddr:            regContracts.AgentRegistry.Hex(),
		dispatchAddr:         regContracts.Dispatch.Hex(),
		fortaStakingAddr:     regContracts.FortaStaking.Hex(),
		scannerFilterer:      sf,
		agentsFilterer:       af,
		dispatchFilterer:     df,
		fortaStakingFilterer: stkf,
	}, nil
}
