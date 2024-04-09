package registry

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/forta-network/forta-core-go/contracts/generated/contract_dispatch_0_1_5"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_rewards_distributor"
	"github.com/forta-network/go-multicall"

	"github.com/forta-network/forta-core-go/contracts/merged/contract_agent_registry"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_dispatch"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_forta_staking"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_scanner_node_version"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_scanner_pool_registry"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_scanner_registry"
	"github.com/forta-network/forta-core-go/contracts/merged/contract_stake_allocator"
	"github.com/forta-network/forta-core-go/utils/ethutils"

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
	defaultScannerRegistryChainID = uint64(137)
	scannerRegistryDeployBlock    = 20187154
	defaultEnsAddress             = "0x08f42fcc52a9C2F391bF507C4E8688D0b53e1bd7"
)

// Registry errors
var (
	ErrContractNotReady        = errors.New("contract not ready")
	ErrUnknownAssignmentLength = errors.New("unknown assignment length")
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
	// Close closes the client.
	Close()

	// RefreshContracts refreshes all contract implementations.
	RefreshContracts() error

	// Contracts returns the ens-resolved registry contracts
	Contracts() *Contracts

	//PegLatestBlock will set the opts so that every call uses same block
	PegLatestBlock() error
	PegBlock(blockNum *big.Int)

	// ResetOpts unsets the options for the store
	ResetOpts()

	// SetRegistryChainID sets the registry chain ID in the client.
	SetRegistryChainID(chainID uint64)

	// GetTransactionOpts returns the default transaction opts used by the client whenever a transaction is sent.
	GetTransactionOpts() (*bind.TransactOpts, error)

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

	// GenerateScannerRegistrationSignature generates a scanner registration signature from given data.
	GenerateScannerRegistrationSignature(reg *eip712.ScannerNodeRegistration) (*ScannerRegistrationInfo, error)

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

	// NumScannersForByChain gets total number of assignments for bot for a given chain.
	NumScannersForByChain(agentID string, chainID *big.Int) (*big.Int, error)

	// GetActiveAgentStake returns the active stake for an agent
	GetActiveAgentStake(blockNumber *big.Int, botID string) (*big.Int, error)

	// GetActiveScannerStake returns the active stake for a scanner
	GetActiveScannerStake(blockNumber *big.Int, scannerID string) (*big.Int, error)

	// GetActivePoolStake returns the total active stake for a pool at a specific block
	GetActivePoolStake(blockNumber, poolID *big.Int) (*big.Int, error)

	// GetAllocatedStakePerManaged Returns allocatedManagedStake (own + delegator's) in DELEGATED / total managed subjects, or 0 if not DELEGATED
	GetAllocatedStakePerManaged(blockNumber, poolID *big.Int) (*big.Int, error)

	// GetAssignmentList returns the indexed bot assignment list for a given node.
	GetAssignmentList(blockNumber, chainId *big.Int, scannerAddress string) ([]*Assignment, error)

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
	AgentRegTx  *contract_agent_registry.AgentRegistryTransactor

	ScannerReg    *contract_scanner_registry.ScannerRegistryCaller
	ScannerRegFil *contract_scanner_registry.ScannerRegistryFilterer

	Dispatch      *contract_dispatch.DispatchCaller
	DispatchFil   *contract_dispatch.DispatchFilterer
	DispatchMulti *multicall.Contract

	ScannerVersion    *contract_scanner_node_version.ScannerNodeVersionCaller
	ScannerVersionFil *contract_scanner_node_version.ScannerNodeVersionFilterer

	FortaStaking    *contract_forta_staking.FortaStakingCaller
	FortaStakingFil *contract_forta_staking.FortaStakingFilterer

	// delegated staking contracts

	ScannerPoolReg    *contract_scanner_pool_registry.ScannerPoolRegistryCaller
	ScannerPoolRegFil *contract_scanner_pool_registry.ScannerPoolRegistryFilterer

	StakeAllocator    *contract_stake_allocator.StakeAllocatorCaller
	StakeAllocatorFil *contract_stake_allocator.StakeAllocatorFilterer

	RewardsDistributor    *contract_rewards_distributor.RewardsDistributorCaller
	RewardsDistributorFil *contract_rewards_distributor.RewardsDistributorFilterer
}

type client struct {
	ctx         context.Context
	cfg         ClientConfig
	eth         ethereum.Client
	ec          *ethclient.Client
	multiCaller *multicall.Caller
	chainID     *big.Int

	// call PegLatestBlock to peg the context to the latest block
	opts       *bind.CallOpts
	privateKey *ecdsa.PrivateKey

	contracts Contracts

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

	// MulticallAddress is the contract address used for the multicalls
	MulticallAddress string
}

var defaultConfig = ClientConfig{
	JsonRpcUrl: "https://polygon-rpc.com",
	ENSAddress: defaultEnsAddress,
	Name:       "registry-client",
}

var devConfig = ClientConfig{
	JsonRpcUrl: "https://sepolia.base.org",
	ENSAddress: "0x650AFCA8545964064b60ad040F9a09F788F714ed",
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
		ctx:     ctx,
		cfg:     cfg,
		eth:     eth,
		ec:      ec,
		chainID: big.NewInt(0).SetUint64(defaultScannerRegistryChainID),

		privateKey: cfg.PrivateKey,

		versionManager: &VersionManager{},
	}
	cl.contracts.Addresses = *regContracts

	if len(cfg.MulticallAddress) > 0 {
		cl.multiCaller, err = multicall.New(cl.ec, cfg.MulticallAddress)
	} else {
		cl.multiCaller, err = multicall.New(cl.ec)
	}
	if err != nil {
		return nil, err
	}

	cl.contracts.AgentReg, err = contract_agent_registry.NewAgentRegistryCaller(regContracts.AgentRegistry, ec)
	if err != nil {
		return nil, err
	}
	cl.contracts.AgentRegFil, err = contract_agent_registry.NewAgentRegistryFilterer(regContracts.AgentRegistry, ec)
	if err != nil {
		return nil, err
	}
	cl.contracts.AgentRegTx, err = contract_agent_registry.NewAgentRegistryTransactor(regContracts.AgentRegistry, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("AgentRegistry", cl.contracts.AgentReg, cl.contracts.AgentReg, cl.contracts.AgentRegFil, cl.contracts.AgentRegTx)

	cl.contracts.ScannerReg, err = contract_scanner_registry.NewScannerRegistryCaller(regContracts.ScannerRegistry, ec)
	if err != nil {
		return nil, err
	}
	cl.contracts.ScannerRegFil, err = contract_scanner_registry.NewScannerRegistryFilterer(regContracts.ScannerRegistry, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("ScannerRegistry", cl.contracts.ScannerReg, cl.contracts.ScannerReg, cl.contracts.ScannerRegFil)

	cl.contracts.ScannerPoolReg, err = contract_scanner_pool_registry.NewScannerPoolRegistryCaller(regContracts.ScannerPoolRegistry, ec)
	if err != nil {
		return nil, err
	}
	cl.contracts.ScannerPoolRegFil, err = contract_scanner_pool_registry.NewScannerPoolRegistryFilterer(regContracts.ScannerPoolRegistry, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("ScannerPoolRegistry", cl.contracts.ScannerPoolReg, cl.contracts.ScannerPoolReg, cl.contracts.ScannerPoolRegFil)

	cl.contracts.Dispatch, err = contract_dispatch.NewDispatchCaller(regContracts.Dispatch, ec)
	if err != nil {
		return nil, err
	}
	cl.contracts.DispatchFil, err = contract_dispatch.NewDispatchFilterer(regContracts.Dispatch, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("Dispatch", cl.contracts.Dispatch, cl.contracts.Dispatch, cl.contracts.DispatchFil)
	cl.contracts.DispatchMulti, err = multicall.NewContract(contract_dispatch_0_1_5.DispatchMetaData.ABI, regContracts.Dispatch.Hex())
	if err != nil {
		return nil, err
	}

	cl.contracts.ScannerVersion, err = contract_scanner_node_version.NewScannerNodeVersionCaller(regContracts.ScannerNodeVersion, ec)
	if err != nil {
		return nil, err
	}
	cl.contracts.ScannerVersionFil, err = contract_scanner_node_version.NewScannerNodeVersionFilterer(regContracts.ScannerNodeVersion, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("ScannerNodeVersion", cl.contracts.ScannerVersion, cl.contracts.ScannerVersion, cl.contracts.ScannerVersionFil)

	cl.contracts.FortaStaking, err = contract_forta_staking.NewFortaStakingCaller(regContracts.FortaStaking, ec)
	if err != nil {
		return nil, err
	}
	cl.contracts.FortaStakingFil, err = contract_forta_staking.NewFortaStakingFilterer(regContracts.FortaStaking, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("FortaStaking", cl.contracts.FortaStaking, cl.contracts.FortaStaking, cl.contracts.FortaStakingFil)

	cl.contracts.StakeAllocator, err = contract_stake_allocator.NewStakeAllocatorCaller(regContracts.StakeAllocator, ec)
	if err != nil {
		return nil, err
	}
	cl.contracts.StakeAllocatorFil, err = contract_stake_allocator.NewStakeAllocatorFilterer(regContracts.StakeAllocator, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("StakeAllocator", cl.contracts.StakeAllocator, cl.contracts.StakeAllocator, cl.contracts.StakeAllocatorFil)

	cl.contracts.RewardsDistributor, err = contract_rewards_distributor.NewRewardsDistributorCaller(regContracts.Rewards, ec)
	if err != nil {
		return nil, err
	}
	cl.contracts.RewardsDistributorFil, err = contract_rewards_distributor.NewRewardsDistributorFilterer(regContracts.Rewards, ec)
	if err != nil {
		return nil, err
	}
	cl.versionManager.SetUpdateRule("RewardsDistributor", cl.contracts.RewardsDistributor, cl.contracts.RewardsDistributor, cl.contracts.RewardsDistributorFil)

	if cfg.NoRefresh {
		return cl, nil
	}

	if err := cl.RefreshContracts(); err != nil {
		return nil, err
	}

	return cl, nil
}

func (c *client) RefreshContracts() error {
	return c.versionManager.Refresh()
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

func (c *client) Close() {
	c.ec.Close()
	c.eth.Close()
}

func (c *client) Contracts() *Contracts {
	// return reference to a copy
	contracts := c.contracts
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

// SetRegistryChainID sets the registry chain ID in the client.
func (c *client) SetRegistryChainID(chainID uint64) {
	c.chainID = big.NewInt(0).SetUint64(chainID)
}

// GetTransactionOpts returns the default transaction opts used by the client whenever a transaction is sent.
func (c *client) GetTransactionOpts() (*bind.TransactOpts, error) {
	if c.privateKey == nil {
		return nil, errors.New("registry client does not have a private key - please provide it using the config")
	}
	if c.chainID != nil {
		return bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	}
	return bind.NewKeyedTransactor(c.privateKey), nil
}

func (c *client) GetScannerNodeVersion() (string, error) {
	scannerVersion, err := c.contracts.ScannerVersion.ScannerNodeVersion(c.opts)
	if err != nil {
		return "", err
	}
	return scannerVersion, nil
}

func (c *client) GetScannerNodePrereleaseVersion() (string, error) {
	scannerVersion, err := c.contracts.ScannerVersion.ScannerNodeBetaVersion(c.opts)
	if err != nil {
		return "", err
	}
	return scannerVersion, nil
}

func (c *client) GetAssignmentHash(scannerID string) (*AssignmentHash, error) {
	sh, err := c.contracts.Dispatch.ScannerHash(c.opts, utils.ScannerIDHexToBigInt(scannerID))
	if err != nil {
		return nil, err
	}
	// this can be nil sometimes, if the JSON-RPC API messes up the response
	if sh.Length == nil {
		return nil, ErrUnknownAssignmentLength
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
	linked, err := c.contracts.Dispatch.AreTheyLinked(c.opts, agtID, scnID)
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

	contracts := c.contracts

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
	contracts := c.contracts
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

	contracts := c.contracts

	cID := big.NewInt(chainID)
	length, err := contracts.AgentReg.GetAgentCountByChain(opts, cID)
	if err != nil {
		return err
	}

	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		agtID, err := c.contracts.AgentReg.GetAgentByChainAndIndex(opts, cID, idx)
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

	contracts := c.contracts

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
		agt, err := c.contracts.AgentReg.GetAgentState(opts, utils.AgentHexToBigInt(agentID))
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

	contracts := c.contracts

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

	contracts := c.contracts

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
	length, err := c.contracts.Dispatch.NumScannersFor(opts, aID)
	if err != nil {
		return nil, err
	}

	contracts := c.contracts
	var idxByChain int64
	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		scn, err := contracts.Dispatch.ScannerRefAt(opts, aID, idx)
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

func (c *client) NumScannersForByChain(agentID string, chainID *big.Int) (*big.Int, error) {
	opts, err := c.getOpts()
	if err != nil {
		return nil, err
	}
	aID := utils.AgentHexToBigInt(agentID)
	length, err := c.contracts.Dispatch.NumScannersFor(opts, aID)
	if err != nil {
		return nil, err
	}

	contracts := c.contracts
	var assigns int64
	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		scn, err := contracts.Dispatch.ScannerRefAt(opts, aID, idx)
		if err != nil {
			return nil, err
		}

		// if filtered by chain, ignore.
		if chainID != nil && scn.ChainId.Cmp(chainID) != 0 {
			continue
		}

		assigns++
	}

	return big.NewInt(assigns), nil
}

func (c *client) NumScannersFor(agentID string) (*big.Int, error) {
	opts, err := c.getOpts()
	if err != nil {
		return nil, err
	}
	aID := utils.AgentHexToBigInt(agentID)

	return c.contracts.Dispatch.NumScannersFor(opts, aID)
}

func (c *client) ForEachAssignedAgent(scannerID string, handler func(a *Agent) error) error {
	// if opts not set, peg for this loop
	opts, err := c.getOpts()
	if err != nil {
		return err
	}

	contracts := c.contracts

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
	contracts := c.contracts
	return contracts.ScannerPoolReg.IsScannerOperational(c.opts, common.HexToAddress(scannerID))
}

func (c *client) IsOperationalScanner(scannerID string) (bool, error) {
	contracts := c.contracts
	if contracts.ScannerPoolReg == nil || contracts.ScannerPoolRegFil == nil {
		return false, ErrContractNotReady
	}
	return contracts.ScannerPoolReg.IsScannerOperational(c.opts, common.HexToAddress(scannerID))
}

func (c *client) getBlockOpts(blockNumber *big.Int) *bind.CallOpts {
	if c.opts != nil {
		opts := *c.opts
		opts.BlockNumber = blockNumber
		return &opts
	}
	return &bind.CallOpts{
		BlockNumber: blockNumber,
	}
}

func (c *client) GetActiveAgentStake(blockNumber *big.Int, botID string) (*big.Int, error) {
	opts := c.getBlockOpts(blockNumber)
	bID := utils.AgentHexToBigInt(botID)
	return c.contracts.FortaStaking.ActiveStakeFor(opts, SubjectTypeAgent, bID)
}

func (c *client) GetActiveScannerStake(blockNumber *big.Int, scannerID string) (*big.Int, error) {

	opts := c.getBlockOpts(blockNumber)

	sID := utils.ScannerIDHexToBigInt(scannerID)
	return c.contracts.FortaStaking.ActiveStakeFor(opts, SubjectTypeScanner, sID)
}

func (c *client) GetScanner(scannerID string) (*Scanner, error) {
	return c.GetPoolScanner(scannerID)
}

func (c *client) GetPoolScanner(scannerID string) (*Scanner, error) {
	contracts := c.contracts
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
	contracts := c.contracts

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

func (c *client) GenerateScannerRegistrationSignature(reg *eip712.ScannerNodeRegistration) (*ScannerRegistrationInfo, error) {
	_, sig, err := eip712.SignScannerRegistration(c.privateKey, c.contracts.Addresses.ScannerPoolRegistry, c.chainID, reg)
	if err != nil {
		return nil, fmt.Errorf("failed to sign the registration data: %v", err)
	}
	info, err := MakeScannerRegistrationInfo(reg, sig)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare the scanner registration info: %v", err)
	}
	return info, nil
}

func (c *client) GetScannerPoolOwner(poolID *big.Int) (owner string, err error) {
	contracts := c.contracts
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
	return c.contracts.ScannerPoolReg.WillNewScannerShutdownPool(c.opts, poolID)
}

func (c *client) GetActivePoolStake(blockNumber, poolID *big.Int) (*big.Int, error) {
	contracts := c.contracts
	if contracts.ScannerPoolReg == nil {
		return nil, ErrContractNotReady
	}

	opts := c.getBlockOpts(blockNumber)
	poolStake, err := contracts.FortaStaking.ActiveStakeFor(opts, SubjectTypeScannerPool, poolID)
	if err != nil {
		return nil, err
	}
	delegatorPoolStake, err := contracts.FortaStaking.ActiveStakeFor(opts, SubjectTypeDelegatorScannerPool, poolID)
	if err != nil {
		return nil, err
	}

	return big.NewInt(0).Add(poolStake, delegatorPoolStake), nil
}

func (c *client) GetAllocatedStakePerManaged(blockNumber, poolID *big.Int) (*big.Int, error) {
	contracts := c.contracts
	if contracts.StakeAllocator == nil {
		return nil, ErrContractNotReady
	}

	opts := c.getBlockOpts(blockNumber)
	return contracts.StakeAllocator.AllocatedStakePerManaged(opts, SubjectTypeScannerPool, poolID)
}
