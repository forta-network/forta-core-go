package registry

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/forta-network/forta-core-go/contracts/generated/contract_scanner_registry_0_1_3"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_agent_registry"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_dispatch"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_forta_staking"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_scanner_node_version"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_scanner_pool_registry"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_scanner_registry"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_stake_allocator"
	"github.com/forta-network/forta-core-go/utils/ethutils"
	log "github.com/sirupsen/logrus"

	"github.com/forta-network/forta-core-go/domain/registry"
	"github.com/forta-network/forta-core-go/security/eip712"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/forta-network/forta-core-go/ens"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/forta-network/forta-core-go/utils"
)

const (
	scannerRegistryChainID     = 137
	scannerRegistryDeployBlock = 20187154
	defaultEnsAddress          = "0x08f42fcc52a9C2F391bF507C4E8688D0b53e1bd7"
)

// Registry errors
var (
	ErrContractNotReady = errors.New("contract not ready")
)

// ScannerPermission is an alias to use for enums.
type ScannerPermission uint8

// ScannerPermission values align with an enum in the ScannerRegistry smart contract
const (
	ScannerPermissionAdmin ScannerPermission = iota
	ScannerPermissionSelf
	ScannerPermissionOwner
	ScannerPermissionManager
)

// Client retrieves agent, scanner, and assignment information from the registry contracts
type Client interface {
	// RefreshContracts refreshes all contract implementations.
	RefreshContracts() error

	// Contracts returns the ens-resolved registry contracts
	Contracts() *Contracts

	//PegLatestBlock will set the opts so that every call uses same block
	PegLatestBlock() error
	PegBlock(blockNum *big.Int)

	// ResetOpts unsets the options for the store
	ResetOpts()

	// GetAssignmentHash returns a hash of all agents, helpful for knowing scanner's agents have changed
	GetAssignmentHash(scannerID string) (*AssignmentHash, error)

	// IsEnabledScanner returns true if the scanner exists and is enabled
	IsEnabledScanner(scannerID string) (bool, error)

	// IsOperationalScanner returns true if the scanner exists and is enabled
	IsOperationalScanner(scannerID string) (bool, error)

	// GetScannerNodeVersion returns the current ipfs reference for the latest scanner node release
	GetScannerNodeVersion() (string, error)

	// GetScannerNodePrereleaseVersion returns the current ipfs reference for the latest scanner node prerelease
	GetScannerNodePrereleaseVersion() (string, error)

	// GetAgent returns the registry information for the agent
	GetAgent(agentID string) (*Agent, error)

	// IsAssigned returns true if the scanner is assigned to the agent
	IsAssigned(scannerID string, agentID string) (bool, error)

	// GetScanner returns a scanner
	GetScanner(scannerID string) (*Scanner, error)

	// GetPoolScanner returns a scanner
	GetPoolScanner(scannerID string) (*Scanner, error)

	// TotalScannersRegistered returns the count of scanners registered in a pool
	TotalScannersRegistered(poolID *big.Int) (*big.Int, error)

	// RegisterScannerOld executes the pre-delegated-staking registration.
	RegisterScannerOld(ownerAddress string, chainID int64, metadata string) (txHash string, err error)

	// GenerateScannerRegistrationSignature generates a scanner registration signature from given data.
	GenerateScannerRegistrationSignature(reg *eip712.ScannerNodeRegistration) (encodedData []byte, sig []byte, err error)

	// ForEachAssignedAgent invokes a handler for each agent assigned to the scanner
	ForEachAssignedAgent(scannerID string, handler func(a *Agent) error) error

	// ForEachChainAgent loops over all agents for a chainID
	ForEachChainAgent(chainID int64, handler func(a *Agent) error) error

	// ForEachAgent loops over all agents
	ForEachAgent(handler func(a *Agent) error) error

	// ForEachAgentID loops over all agents only returning the ID
	ForEachAgentID(handler func(agentID string) error) error

	// ForEachAgentSinceBlock loops over all agents since a provided block number
	ForEachAgentSinceBlock(block uint64, handler func(event *contract_agent_registry.AgentRegistryAgentUpdated, a *Agent) error) error

	// ForEachScanner gets all scanners
	ForEachScanner(handler func(s *Scanner) error) error

	// ForEachScannerSinceBlock loops over all scanners since a provided block number
	ForEachScannerSinceBlock(block uint64, handler func(event *contract_scanner_registry.ScannerRegistryScannerUpdated, s *Scanner) error) error

	// ForEachPoolScannerSinceBlock loops over all scanners since a provided block number
	ForEachPoolScannerSinceBlock(block uint64, handler func(event *contract_scanner_pool_registry.ScannerPoolRegistryScannerUpdated, s *Scanner) error) error

	// ForEachAssignedScanner loops over scanners by agent
	ForEachAssignedScanner(agentID string, handler func(s *Scanner) error) error

	// IndexOfAssignedScannerByChain gets index of assigned scanner inside a chain.
	IndexOfAssignedScannerByChain(agentID, scannerID string, chainID *big.Int) (*big.Int, error)

	// NumScannersFor gets total number of assignments for bot.
	NumScannersFor(agentID string) (*big.Int, error)

	// GetStakingThreshold returns the min/max/activated flag for a given address
	GetStakingThreshold(scannerID string) (*StakingThreshold, error)

	// GetActiveScannerStake returns the active stake for a scanner
	GetActiveScannerStake(scannerID string) (*big.Int, error)

	// GetActivePoolStake returns the total active stake for a pool
	GetActivePoolStake(poolID *big.Int) (*big.Int, error)

	// EnableScanner enables a scanner. Deprecated.
	EnableScanner(ScannerPermission ScannerPermission, scannerAddress string) (txHash string, err error)

	// DisableScanner disables a scanner. Deprecated.
	DisableScanner(ScannerPermission ScannerPermission, scannerAddress string) (txHash string, err error)

	// GetScannerPoolOwner finds out the owner of the scanner pool.
	GetScannerPoolOwner(poolID *big.Int) (owner string, err error)

	// WillNewScannerShutdownPool tells if registering a new scanner could shutdown a pool.
	WillNewScannerShutdownPool(poolID *big.Int) (bool, error)
}

// Contracts contains the latest state of the contracts.
type Contracts struct {
	Addresses registry.RegistryContracts

	AgentReg    *contract_agent_registry.AgentRegistryCaller
	AgentRegFil *contract_agent_registry.AgentRegistryFilterer

	ScannerReg    *contract_scanner_registry.ScannerRegistryCaller
	ScannerRegFil *contract_scanner_registry.ScannerRegistryFilterer

	Dispatch    *contract_dispatch.DispatchCaller
	DispatchFil *contract_dispatch.DispatchFilterer

	ScannerVersion    *contract_scanner_node_version.ScannerNodeVersionCaller
	ScannerVersionFil *contract_scanner_node_version.ScannerNodeVersionFilterer

	FortaStaking    *contract_forta_staking.FortaStakingCaller
	FortaStakingFil *contract_forta_staking.FortaStakingFilterer

	// post-migration contracts

	ScannerPoolReg    *contract_scanner_pool_registry.ScannerPoolRegistryCaller
	ScannerPoolRegFil *contract_scanner_pool_registry.ScannerPoolRegistryFilterer

	StakeAllocator    *contract_stake_allocator.StakeAllocatorCaller
	StakeAllocatorFil *contract_stake_allocator.StakeAllocatorFilterer
}

type client struct {
	ctx context.Context
	cfg ClientConfig
	eth ethereum.Client
	ec  *ethclient.Client

	// call PegLatestBlock to peg the context to the latest block
	opts       *bind.CallOpts
	privateKey *ecdsa.PrivateKey

	contractsUnsafe Contracts
	contractsMu     sync.RWMutex

	versionManager *VersionManager
}

var _ Client = &client{}

type ClientConfig struct {
	// JsonRpcUrl is the url used for resolving the registry (polygon)
	JsonRpcUrl string `json:"jsonRpcUrl"`

	// ENSAddress is if there's not a default contract for ENS resolution
	ENSAddress string `json:"ensAddress"`

	// Name is used for logging
	Name string `json:"name"`

	// PrivateKey is used for sending transactions
	PrivateKey *ecdsa.PrivateKey

	// NoRefresh tells if the contracts should not be refreshed and multiplexed to different versions.
	NoRefresh bool
}

var defaultConfig = ClientConfig{
	JsonRpcUrl: "https://polygon-rpc.com",
	ENSAddress: defaultEnsAddress,
	Name:       "registry-client",
}

var devConfig = ClientConfig{
	JsonRpcUrl: "https://rpc-mumbai.matic.today",
	ENSAddress: "0x5f7c5bbBa72e1e1fae689120D76D2f334A390Ae9",
	Name:       "registry-client",
}

func NewDevClient(ctx context.Context) (*client, error) {
	return NewClient(ctx, devConfig)
}

func NewDefaultClient(ctx context.Context) (*client, error) {
	return NewClient(ctx, defaultConfig)
}

func NewClientWithENSStore(ctx context.Context, cfg ClientConfig, ensStore ens.ENS) (*client, error) {
	regContracts, err := ensStore.ResolveRegistryContracts()
	if err != nil {
		return nil, err
	}

	eth, err := ethereum.NewStreamEthClient(ctx, cfg.Name, cfg.JsonRpcUrl)
	if err != nil {
		return nil, err
	}

	rpc, err := ethereum.NewRpcClient(ctx, cfg.JsonRpcUrl)
	if err != nil {
		return nil, err
	}
	ec := ethclient.NewClient(rpc)

	cl := &client{
		ctx: ctx,
		cfg: cfg,
		eth: eth,
		ec:  ec,

		privateKey: cfg.PrivateKey,

		versionManager: &VersionManager{},
	}
	cl.contractsUnsafe.Addresses = *regContracts

	cl.contractsUnsafe.AgentReg, err = contract_agent_registry.NewAgentRegistryCaller(regContracts.AgentRegistry, ec)
	if err != nil {
		return nil, err
	}
	cl.contractsUnsafe.AgentRegFil, err = contract_agent_registry.NewAgentRegistryFilterer(regContracts.AgentRegistry, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("AgentRegistry", cl.contractsUnsafe.AgentReg, cl.contractsUnsafe.AgentReg, cl.contractsUnsafe.AgentRegFil)

	cl.contractsUnsafe.ScannerReg, err = contract_scanner_registry.NewScannerRegistryCaller(regContracts.ScannerRegistry, ec)
	if err != nil {
		return nil, err
	}
	cl.contractsUnsafe.ScannerRegFil, err = contract_scanner_registry.NewScannerRegistryFilterer(regContracts.ScannerRegistry, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("ScannerRegistry", cl.contractsUnsafe.ScannerReg, cl.contractsUnsafe.ScannerReg, cl.contractsUnsafe.ScannerRegFil)

	cl.contractsUnsafe.Dispatch, err = contract_dispatch.NewDispatchCaller(regContracts.Dispatch, ec)
	if err != nil {
		return nil, err
	}
	cl.contractsUnsafe.DispatchFil, err = contract_dispatch.NewDispatchFilterer(regContracts.Dispatch, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("Dispatch", cl.contractsUnsafe.Dispatch, cl.contractsUnsafe.Dispatch, cl.contractsUnsafe.DispatchFil)

	cl.contractsUnsafe.ScannerVersion, err = contract_scanner_node_version.NewScannerNodeVersionCaller(regContracts.ScannerNodeVersion, ec)
	if err != nil {
		return nil, err
	}
	cl.contractsUnsafe.ScannerVersionFil, err = contract_scanner_node_version.NewScannerNodeVersionFilterer(regContracts.ScannerNodeVersion, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("ScannerNodeVersion", cl.contractsUnsafe.ScannerVersion, cl.contractsUnsafe.ScannerVersion, cl.contractsUnsafe.ScannerVersionFil)

	cl.contractsUnsafe.FortaStaking, err = contract_forta_staking.NewFortaStakingCaller(regContracts.FortaStaking, ec)
	if err != nil {
		return nil, err
	}
	cl.contractsUnsafe.FortaStakingFil, err = contract_forta_staking.NewFortaStakingFilterer(regContracts.FortaStaking, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("FortaStaking", cl.contractsUnsafe.FortaStaking, cl.contractsUnsafe.FortaStaking, cl.contractsUnsafe.FortaStakingFil)

	if cfg.NoRefresh {
		return cl, nil
	}

	if err := cl.RefreshContracts(); err != nil {
		return nil, err
	}

	return cl, nil
}

func (c *client) RefreshContracts() error {
	if err := c.versionManager.Refresh(); err != nil {
		return err
	}

	c.contractsMu.Lock()
	defer c.contractsMu.Unlock()

	// find out post-delegated-staking-migration contracts from known contracts

	var detectedNew bool

	scannerPoolRegAddr, err := c.contractsUnsafe.ScannerReg.ScannerPoolRegistry(c.opts)
	if err == nil && scannerPoolRegAddr.Hex() != utils.ZeroAddress {
		c.contractsUnsafe.Addresses.ScannerPoolRegistry = &scannerPoolRegAddr
		c.contractsUnsafe.ScannerPoolReg, err = contract_scanner_pool_registry.NewScannerPoolRegistryCaller(scannerPoolRegAddr, c.ec)
		if err != nil {
			return err
		}
		c.contractsUnsafe.ScannerPoolRegFil, err = contract_scanner_pool_registry.NewScannerPoolRegistryFilterer(scannerPoolRegAddr, c.ec)
		if err != nil {
			return err
		}
		c.versionManager.SetUpdateRule("ScannerPoolRegistry", c.contractsUnsafe.ScannerPoolReg, c.contractsUnsafe.ScannerPoolReg, c.contractsUnsafe.ScannerPoolRegFil)
		detectedNew = true
		version, err := c.contractsUnsafe.ScannerPoolReg.Version(c.opts)
		if err != nil {
			return err
		}
		log.WithFields(log.Fields{
			"name":    "ScannerPoolRegistry",
			"version": version,
		}).Info("detected post-migration contract")
	}

	stakeAllocatorAddr, err := c.contractsUnsafe.FortaStaking.Allocator(c.opts)
	if err == nil && stakeAllocatorAddr.Hex() != utils.ZeroAddress {
		c.contractsUnsafe.Addresses.StakeAllocator = &stakeAllocatorAddr
		c.contractsUnsafe.StakeAllocator, err = contract_stake_allocator.NewStakeAllocatorCaller(stakeAllocatorAddr, c.ec)
		if err != nil {
			return err
		}
		c.contractsUnsafe.StakeAllocatorFil, err = contract_stake_allocator.NewStakeAllocatorFilterer(stakeAllocatorAddr, c.ec)
		if err != nil {
			return err
		}
		c.versionManager.SetUpdateRule("StakeAllocator", c.contractsUnsafe.StakeAllocator, c.contractsUnsafe.StakeAllocator, c.contractsUnsafe.StakeAllocatorFil)
		detectedNew = true
		version, err := c.contractsUnsafe.StakeAllocator.Version(c.opts)
		if err != nil {
			return err
		}
		log.WithFields(log.Fields{
			"name":    "StakeAllocator",
			"version": version,
		}).Info("detected post-migration contract")
	}

	if detectedNew {
		return c.versionManager.Refresh()
	}

	return nil
}

func NewClient(ctx context.Context, cfg ClientConfig) (*client, error) {
	// avoids need to provide ENS address outside of dev environment use cases
	ensAddr := cfg.ENSAddress
	if ensAddr == "" {
		ensAddr = defaultEnsAddress
	}
	ensStore, err := ens.DialENSStoreAt(cfg.JsonRpcUrl, ensAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to dial ens store: %v", err)
	}
	return NewClientWithENSStore(ctx, cfg, ensStore)
}

func (c *client) Contracts() *Contracts {
	c.contractsMu.RLock()
	defer c.contractsMu.RUnlock()

	// return reference to a copy
	contracts := c.contractsUnsafe
	return &contracts
}

func (c *client) SetPrivateKey(privateKey *ecdsa.PrivateKey) {
	c.privateKey = privateKey
}

// ResetOpts unsets the options for the store
func (c *client) ResetOpts() {
	c.opts = nil
}

// latestOpts returns the callopts for the latest block so that calls can use a same block
func (c *client) latestOpts() (*bind.CallOpts, error) {
	blk, err := c.eth.BlockByNumber(c.ctx, nil)
	if err != nil {
		return nil, err
	}
	num, err := utils.HexToBigInt(blk.Number)
	if err != nil {
		return nil, err
	}
	return &bind.CallOpts{
		BlockNumber: num,
	}, nil
}

func (c *client) PegBlock(blockNum *big.Int) {
	c.opts = &bind.CallOpts{
		BlockNumber: blockNum,
	}
}

// PegLatestBlock will set the opts so that every call uses same block
func (c *client) PegLatestBlock() error {
	opts, err := c.latestOpts()
	if err != nil {
		return err
	}
	c.opts = opts
	return nil
}

func (c *client) GetStakingThreshold(scannerID string) (*StakingThreshold, error) {
	sID := utils.ScannerIDHexToBigInt(scannerID)
	res, err := c.Contracts().ScannerReg.GetStakeThreshold(c.opts, sID)
	if err != nil {
		return nil, err
	}
	return (*StakingThreshold)(res), nil
}

func (c *client) GetScannerNodeVersion() (string, error) {
	scannerVersion, err := c.Contracts().ScannerVersion.ScannerNodeVersion(c.opts)
	if err != nil {
		return "", err
	}
	return scannerVersion, nil
}

func (c *client) GetScannerNodePrereleaseVersion() (string, error) {
	scannerVersion, err := c.Contracts().ScannerVersion.ScannerNodeBetaVersion(c.opts)
	if err != nil {
		return "", err
	}
	return scannerVersion, nil
}

func (c *client) GetAssignmentHash(scannerID string) (*AssignmentHash, error) {
	sh, err := c.Contracts().Dispatch.ScannerHash(c.opts, utils.ScannerIDHexToBigInt(scannerID))
	if err != nil {
		return nil, err
	}
	return &AssignmentHash{
		AgentLength: sh.Length.Int64(),
		Hash:        utils.Bytes32ToHex(sh.Manifest),
	}, nil
}

func (c *client) getOpts() (*bind.CallOpts, error) {
	if c.opts != nil {
		return c.opts, nil
	}
	return c.latestOpts()
}

func (c *client) IsAssigned(scannerID string, agentID string) (bool, error) {
	agtID := utils.AgentHexToBigInt(agentID)
	scnID := utils.ScannerIDHexToBigInt(scannerID)
	linked, err := c.Contracts().Dispatch.AreTheyLinked(c.opts, agtID, scnID)
	if err != nil {
		return false, err
	}
	return linked, nil
}

func (c *client) ForEachScanner(handler func(s *Scanner) error) error {
	return c.ForEachScannerSinceBlock(
		scannerRegistryDeployBlock,
		func(_ *contract_scanner_registry.ScannerRegistryScannerUpdated, s *Scanner) error {
			return handler(s)
		})
}

func (c *client) ForEachScannerSinceBlock(
	block uint64, handler func(event *contract_scanner_registry.ScannerRegistryScannerUpdated, s *Scanner) error,
) error {
	opts, err := c.getOpts()
	if err != nil {
		return err
	}

	contracts := c.Contracts()

	iterators, err := contracts.ScannerRegFil.FilterScannerUpdated(&bind.FilterOpts{
		Start:   block,
		Context: c.ctx,
	}, nil, nil)

	if err != nil {
		return err
	}

	it, err := ethutils.NewEventIterator[contract_scanner_registry.ScannerRegistryScannerUpdated](iterators)
	if err != nil {
		return err
	}

	for it.Next() {
		event, ok := it.Value()
		if !ok {
			break
		}

		scn, err := contracts.ScannerReg.GetScannerState(opts, event.ScannerId)
		if err != nil {
			return err
		}
		if err := handler(event, &Scanner{
			ScannerID: utils.ScannerIDBigIntToHex(event.ScannerId),
			ChainID:   scn.ChainId.Int64(),
			Enabled:   scn.Enabled,
			Manifest:  scn.Metadata,
			Owner:     scn.Owner.Hex(),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (c *client) ForEachPoolScannerSinceBlock(
	block uint64, handler func(event *contract_scanner_pool_registry.ScannerPoolRegistryScannerUpdated, s *Scanner) error,
) error {
	contracts := c.Contracts()
	if contracts.ScannerPoolReg == nil || contracts.ScannerPoolRegFil == nil {
		return ErrContractNotReady
	}

	opts, err := c.getOpts()
	if err != nil {
		return err
	}

	it, err := contracts.ScannerPoolRegFil.FilterScannerUpdated(&bind.FilterOpts{
		Start:   block,
		Context: c.ctx,
	}, nil, nil)

	if err != nil {
		return err
	}

	for it.Next() {
		if it.Event != nil {
			scannerAddr := common.HexToAddress(utils.ScannerIDBigIntToHex(it.Event.ScannerId))
			scn, err := contracts.ScannerPoolReg.GetScannerState(opts, scannerAddr)
			if err != nil {
				return err
			}
			if err := handler((*contract_scanner_pool_registry.ScannerPoolRegistryScannerUpdated)(it.Event), &Scanner{
				ScannerID: utils.ScannerIDBigIntToHex(it.Event.ScannerId),
				ChainID:   scn.ChainId.Int64(),
				Enabled:   scn.Operational,
				Manifest:  scn.Metadata,
				Owner:     scn.Owner.Hex(),
				PoolID:    it.Event.ScannerPool.String(),
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *client) ForEachChainAgent(chainID int64, handler func(a *Agent) error) error {
	opts, err := c.getOpts()
	if err != nil {
		return err
	}

	contracts := c.Contracts()

	cID := big.NewInt(chainID)
	length, err := contracts.AgentReg.GetAgentCountByChain(opts, cID)
	if err != nil {
		return err
	}

	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		agtID, err := c.Contracts().AgentReg.GetAgentByChainAndIndex(opts, cID, idx)
		if err != nil {
			return err
		}
		agt, err := contracts.AgentReg.GetAgentState(opts, agtID)
		if err != nil {
			return err
		}
		if err := handler(&Agent{
			AgentID:  utils.AgentBigIntToHex(agtID),
			ChainIDs: utils.IntArray(agt.ChainIds),
			Enabled:  agt.Enabled,
			Manifest: agt.Metadata,
			Owner:    agt.Owner.Hex(),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (c *client) ForEachAgentID(handler func(agentID string) error) error {
	opts, err := c.getOpts()
	if err != nil {
		return err
	}

	contracts := c.Contracts()

	length, err := contracts.AgentReg.GetAgentCount(opts)
	if err != nil {
		return err
	}

	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		agtID, err := contracts.AgentReg.GetAgentByIndex(opts, idx)
		if err != nil {
			return err
		}
		if err := handler(utils.AgentBigIntToHex(agtID)); err != nil {
			return err
		}
	}
	return nil
}

func (c *client) ForEachAgent(handler func(a *Agent) error) error {
	opts, err := c.getOpts()
	if err != nil {
		return err
	}

	return c.ForEachAgentID(func(agentID string) error {
		agt, err := c.Contracts().AgentReg.GetAgentState(opts, utils.AgentHexToBigInt(agentID))
		if err != nil {
			return err
		}
		return handler(&Agent{
			AgentID:  agentID,
			ChainIDs: utils.IntArray(agt.ChainIds),
			Enabled:  agt.Enabled,
			Manifest: agt.Metadata,
			Owner:    agt.Owner.Hex(),
		})
	})
}

func (c *client) ForEachAgentSinceBlock(
	block uint64, handler func(event *contract_agent_registry.AgentRegistryAgentUpdated, a *Agent) error,
) error {
	opts, err := c.getOpts()
	if err != nil {
		return err
	}

	contracts := c.Contracts()

	iterators, err := contracts.AgentRegFil.FilterAgentUpdated(&bind.FilterOpts{
		Start:   block,
		Context: c.ctx,
	}, nil, nil)

	if err != nil {
		return err
	}

	it, err := ethutils.NewEventIterator[contract_agent_registry.AgentRegistryAgentUpdated](iterators)
	if err != nil {
		return err
	}

	for it.Next() {
		event, ok := it.Value()
		if !ok {
			break
		}

		agt, err := contracts.AgentReg.GetAgentState(opts, event.AgentId)
		if err != nil {
			return err
		}
		if err := handler(event, &Agent{
			AgentID:  utils.AgentBigIntToHex(event.AgentId),
			ChainIDs: utils.IntArray(agt.ChainIds),
			Enabled:  agt.Enabled,
			Manifest: agt.Metadata,
			Owner:    agt.Owner.Hex(),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (c *client) ForEachAssignedScanner(agentID string, handler func(s *Scanner) error) error {
	// if opts not set, peg for this loop
	opts, err := c.getOpts()
	if err != nil {
		return err
	}

	contracts := c.Contracts()

	aID := utils.AgentHexToBigInt(agentID)
	length, err := contracts.Dispatch.NumScannersFor(opts, aID)
	if err != nil {
		return err
	}
	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		scn, err := contracts.Dispatch.ScannerRefAt(opts, aID, idx)
		if err != nil {
			return err
		}
		if err := handler(&Scanner{
			ScannerID: utils.ScannerIDBigIntToHex(scn.ScannerId),
			ChainID:   scn.ChainId.Int64(),
			Enabled:   scn.Operational || scn.Enabled,
			Manifest:  scn.Metadata,
			Owner:     scn.Owner.Hex(),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (c *client) IndexOfAssignedScannerByChain(agentID, scannerID string, chainID *big.Int) (*big.Int, error) {
	opts, err := c.getOpts()
	if err != nil {
		return nil, err
	}
	aID := utils.AgentHexToBigInt(agentID)
	sID := utils.ScannerIDHexToBigInt(scannerID)
	length, err := c.Contracts().Dispatch.NumScannersFor(opts, aID)
	if err != nil {
		return nil, err
	}

	var idxByChain int64
	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		scn, err := c.Contracts().Dispatch.ScannerRefAt(opts, aID, idx)
		if err != nil {
			return nil, err
		}

		// if filtered by chain, ignore.
		if chainID != nil && scn.ChainId.Cmp(chainID) != 0 {
			continue
		}

		if scn.ScannerId.Cmp(sID) == 0 {
			return big.NewInt(idxByChain), nil
		}

		idxByChain++
	}

	return nil, fmt.Errorf("can't find shard index for %s", scannerID)
}

func (c *client) NumScannersFor(agentID string) (*big.Int, error) {
	opts, err := c.getOpts()
	if err != nil {
		return nil, err
	}
	aID := utils.AgentHexToBigInt(agentID)

	return c.contractsUnsafe.Dispatch.NumScannersFor(opts, aID)
}

func (c *client) ForEachAssignedAgent(scannerID string, handler func(a *Agent) error) error {
	// if opts not set, peg for this loop
	opts, err := c.getOpts()
	if err != nil {
		return err
	}

	contracts := c.Contracts()

	sID := utils.ScannerIDHexToBigInt(scannerID)
	length, err := contracts.Dispatch.NumAgentsFor(opts, sID)
	if err != nil {
		return err
	}
	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		agt, err := contracts.Dispatch.AgentRefAt(opts, sID, idx)
		if err != nil {
			return err
		}
		if err := handler(&Agent{
			AgentID:  utils.AgentBigIntToHex(agt.AgentId),
			ChainIDs: utils.IntArray(agt.ChainIds),
			Enabled:  agt.Enabled,
			Manifest: agt.Metadata,
			Owner:    agt.Owner.Hex(),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (c *client) IsEnabledScanner(scannerID string) (bool, error) {
	contracts := c.Contracts()
	if c.migrationEnded() && contracts.ScannerPoolReg != nil {
		return contracts.ScannerPoolReg.IsScannerOperational(c.opts, common.HexToAddress(scannerID))
	}
	return contracts.ScannerReg.IsEnabled(c.opts, utils.ScannerIDHexToBigInt(scannerID))
}

func (c *client) migrationEnded() bool {
	// ignoring error here because the old contract doesn't have this method in its ABI
	// only successful request and a "true" in response is true
	migrationEnded, _ := c.Contracts().ScannerReg.HasMigrationEnded(c.opts)
	return migrationEnded
}

func (c *client) IsOperationalScanner(scannerID string) (bool, error) {
	contracts := c.Contracts()
	if contracts.ScannerPoolReg == nil || contracts.ScannerPoolRegFil == nil {
		return false, ErrContractNotReady
	}
	return contracts.ScannerPoolReg.IsScannerOperational(c.opts, common.HexToAddress(scannerID))
}

func (c *client) GetActiveScannerStake(scannerID string) (*big.Int, error) {
	sID := utils.ScannerIDHexToBigInt(scannerID)
	return c.Contracts().FortaStaking.ActiveStakeFor(c.opts, 0, sID)
}

func (c *client) GetScanner(scannerID string) (*Scanner, error) {
	if c.migrationEnded() {
		return c.GetPoolScanner(scannerID)
	}

	contracts := c.Contracts()

	sID := utils.ScannerIDHexToBigInt(scannerID)
	scn, err := contracts.ScannerReg.GetScanner(c.opts, sID)

	if err != nil {
		return nil, err
	}

	if !scn.Registered {
		return nil, nil
	}

	enabled, err := contracts.ScannerReg.IsEnabled(c.opts, sID)
	if err != nil {
		return nil, err
	}

	return &Scanner{
		ScannerID: utils.ScannerIDBigIntToHex(sID),
		ChainID:   scn.ChainId.Int64(),
		Enabled:   enabled,
		Manifest:  scn.Metadata,
		Owner:     scn.Owner.Hex(),
	}, nil
}

func (c *client) GetPoolScanner(scannerID string) (*Scanner, error) {
	contracts := c.Contracts()
	if contracts.ScannerPoolReg == nil || contracts.ScannerPoolRegFil == nil {
		return nil, ErrContractNotReady
	}

	sID := common.HexToAddress(scannerID)
	scn, err := contracts.ScannerPoolReg.GetScanner(c.opts, sID)

	if err != nil {
		return nil, err
	}

	if !scn.Registered {
		return nil, nil
	}

	enabled, err := contracts.ScannerPoolReg.IsScannerOperational(c.opts, sID)
	if err != nil {
		return nil, err
	}

	owner, err := contracts.ScannerPoolReg.OwnerOf(c.opts, scn.ScannerPoolId)
	if err != nil {
		return nil, err
	}

	return &Scanner{
		ScannerID: sID.Hex(),
		ChainID:   scn.ChainId.Int64(),
		Enabled:   enabled,
		Manifest:  scn.Metadata,
		Owner:     owner.Hex(),
		PoolID:    scn.ScannerPoolId.String(),
	}, nil
}

func (c *client) GetAgent(agentID string) (*Agent, error) {
	contracts := c.Contracts()

	aID := utils.AgentHexToBigInt(agentID)
	agt, err := contracts.AgentReg.GetAgent(c.opts, aID)
	if err != nil {
		return nil, err
	}

	// if agt does not exist, return nil
	if !agt.Registered {
		return nil, nil
	}

	enabled, err := contracts.AgentReg.IsEnabled(c.opts, aID)
	if err != nil {
		return nil, err
	}

	return &Agent{
		AgentID:  agentID,
		ChainIDs: utils.IntArray(agt.ChainIds),
		Enabled:  enabled,
		Manifest: agt.Metadata,
		Owner:    agt.Owner.Hex(),
	}, nil
}

func (c *client) RegisterScannerOld(ownerAddress string, chainID int64, metadata string) (txHash string, err error) {
	if c.privateKey == nil {
		return "", errors.New("no private key provided to the client")
	}
	contracts := c.Contracts()
	// point directly to 0.1.3 because the registration method does not exist in newer versions
	reg, err := contract_scanner_registry_0_1_3.NewScannerRegistryTransactor(contracts.Addresses.ScannerRegistry, c.ec)
	if err != nil {
		return "", fmt.Errorf("failed to create contract transactor: %v", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(c.privateKey, big.NewInt(scannerRegistryChainID))
	if err != nil {
		return "", fmt.Errorf("failed to create transaction opts: %v", err)
	}
	opts.GasPrice, err = c.ec.SuggestGasPrice(c.ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get gas price suggestion: %v", err)
	}
	tx, err := reg.Register(opts, common.HexToAddress(ownerAddress), big.NewInt(int64(chainID)), metadata)
	if err != nil {
		return "", fmt.Errorf("failed to send the transaction: %v", err)
	}
	return tx.Hash().Hex(), nil
}

func (c *client) EnableScanner(permission ScannerPermission, scannerAddress string) (txHash string, err error) {
	if c.migrationEnded() {
		return "", errors.New("migration ended")
	}
	if c.privateKey == nil {
		return "", errors.New("no private key provided to the client")
	}
	contracts := c.Contracts()
	// point directly to 0.1.3 because the enable method does not exist in newer versions
	reg, err := contract_scanner_registry_0_1_3.NewScannerRegistryTransactor(contracts.Addresses.ScannerRegistry, c.ec)
	if err != nil {
		return "", fmt.Errorf("failed to create contract transactor: %v", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(c.privateKey, big.NewInt(scannerRegistryChainID))
	if err != nil {
		return "", fmt.Errorf("failed to create transaction opts: %v", err)
	}
	opts.GasPrice, err = c.ec.SuggestGasPrice(c.ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get gas price suggestion: %v", err)
	}
	tx, err := reg.EnableScanner(opts, utils.ScannerIDHexToBigInt(scannerAddress), uint8(permission))
	if err != nil {
		return "", fmt.Errorf("failed to send the transaction: %v", err)
	}
	return tx.Hash().Hex(), nil
}

func (c *client) DisableScanner(permission ScannerPermission, scannerAddress string) (txHash string, err error) {
	if c.privateKey == nil {
		return "", errors.New("no private key provided to the client")
	}
	contracts := c.Contracts()
	// point directly to 0.1.3 because the disable method does not exist in newer versions
	reg, err := contract_scanner_registry_0_1_3.NewScannerRegistryTransactor(contracts.Addresses.ScannerRegistry, c.ec)
	if err != nil {
		return "", fmt.Errorf("failed to create contract transactor: %v", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(c.privateKey, big.NewInt(scannerRegistryChainID))
	if err != nil {
		return "", fmt.Errorf("failed to create transaction opts: %v", err)
	}
	opts.GasPrice, err = c.ec.SuggestGasPrice(c.ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get gas price suggestion: %v", err)
	}
	tx, err := reg.DisableScanner(opts, utils.ScannerIDHexToBigInt(scannerAddress), uint8(permission))
	if err != nil {
		return "", fmt.Errorf("failed to send the transaction: %v", err)
	}
	return tx.Hash().Hex(), nil
}

func (c *client) GenerateScannerRegistrationSignature(reg *eip712.ScannerNodeRegistration) ([]byte, []byte, error) {
	contracts := c.Contracts()
	if contracts.Addresses.ScannerPoolRegistry == nil {
		return nil, nil, ErrContractNotReady
	}
	chainID, err := c.eth.ChainID(c.ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to determine registry chain id: %v", chainID)
	}
	return eip712.SignScannerRegistration(c.privateKey, *contracts.Addresses.ScannerPoolRegistry, chainID, reg)
}

func (c *client) GetScannerPoolOwner(poolID *big.Int) (owner string, err error) {
	contracts := c.Contracts()
	if contracts.ScannerPoolReg == nil {
		return "", ErrContractNotReady
	}
	addr, err := contracts.ScannerPoolReg.OwnerOf(c.opts, poolID)
	if err != nil {
		return "", nil
	}
	return addr.Hex(), nil
}

func (c *client) WillNewScannerShutdownPool(poolID *big.Int) (bool, error) {
	return c.Contracts().ScannerPoolReg.WillNewScannerShutdownPool(c.opts, poolID)
}

func (c *client) GetActivePoolStake(poolID *big.Int) (*big.Int, error) {
	contracts := c.Contracts()
	if contracts.ScannerPoolReg == nil {
		return nil, ErrContractNotReady
	}

	poolStake, err := contracts.FortaStaking.ActiveStakeFor(c.opts, SubjectTypeScannerPool, poolID)
	if err != nil {
		return nil, err
	}
	delegatorPoolStake, err := contracts.FortaStaking.ActiveStakeFor(c.opts, SubjectTypeDelegatorScannerPool, poolID)
	if err != nil {
		return nil, err
	}

	return big.NewInt(0).Add(poolStake, delegatorPoolStake), nil
}

func (c *client) TotalScannersRegistered(poolID *big.Int) (*big.Int, error) {
	contracts := c.Contracts()
	if contracts.ScannerPoolReg == nil {
		return nil, ErrContractNotReady
	}

	return contracts.ScannerPoolReg.TotalScannersRegistered(c.opts, poolID)
}
