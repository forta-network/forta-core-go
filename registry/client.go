package registry

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/forta-network/forta-core-go/contracts/contract_forta_staking"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_pool_registry"
	"github.com/forta-network/forta-core-go/domain/registry"
	"github.com/forta-network/forta-core-go/security/eip712"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/forta-network/forta-core-go/contracts/contract_agent_registry"
	"github.com/forta-network/forta-core-go/contracts/contract_dispatch"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_node_version"
	"github.com/forta-network/forta-core-go/contracts/contract_scanner_registry"
	"github.com/forta-network/forta-core-go/ens"
	"github.com/forta-network/forta-core-go/ethereum"
	"github.com/forta-network/forta-core-go/utils"
)

const (
	scannerRegistryChainID     = 137
	scannerRegistryDeployBlock = 20187154
	defaultEnsAddress          = "0x08f42fcc52a9C2F391bF507C4E8688D0b53e1bd7"
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

	// RegisterScannerToPool registers a scanner to a pool by using scanner's signature and operator private key.
	RegisterScannerToPool(scannerAddress string, poolID *big.Int, chainID int64, metadata string, signature []byte) (txHash string, err error)

	// GenerateScannerRegistrationSignature generates a scanner registration signature from given data.
	GenerateScannerRegistrationSignature(reg *eip712.ScannerNodeRegistration) ([]byte, error)

	// RegistryContracts returns the ens-resolved registry contracts
	RegistryContracts() *registry.RegistryContracts

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

	// GetStakingThreshold returns the min/max/activated flag for a given address
	GetStakingThreshold(scannerID string) (*StakingThreshold, error)

	// EnableScanner enables a scanner.
	EnableScanner(ScannerPermission ScannerPermission, scannerAddress string) (txHash string, err error)

	// DisableScanner disables a scanner.
	DisableScanner(ScannerPermission ScannerPermission, scannerAddress string) (txHash string, err error)

	// GetScannerPoolOwner finds out the owner of the scanner pool.
	GetScannerPoolOwner(poolID *big.Int) (owner string, err error)
}

type client struct {
	ctx context.Context
	cfg ClientConfig
	eth ethereum.Client
	ec  *ethclient.Client

	// call PegLatestBlock to peg the context to the latest block
	opts       *bind.CallOpts
	privateKey *ecdsa.PrivateKey

	contracts         *registry.RegistryContracts
	agentReg          *contract_agent_registry.AgentRegistryCaller
	scannerReg        *contract_scanner_registry.ScannerRegistryCaller
	scannerPoolReg    *contract_scanner_pool_registry.ScannerPoolRegistryCaller
	dispatch          *contract_dispatch.DispatchCaller
	scannerVersion    *contract_scanner_node_version.ScannerNodeVersionCaller
	fortaStaking      *contract_forta_staking.FortaStakingCaller
	scannerRegFil     *contract_scanner_registry.ScannerRegistryFilterer
	scannerPoolRegFil *contract_scanner_pool_registry.ScannerPoolRegistryFilterer
	agentRegFil       *contract_agent_registry.AgentRegistryFilterer
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

	rpc, err := ethereum.NewRpcClient(cfg.JsonRpcUrl)
	if err != nil {
		return nil, err
	}
	ec := ethclient.NewClient(rpc)

	agentReg, err := contract_agent_registry.NewAgentRegistryCaller(regContracts.AgentRegistry, ec)
	if err != nil {
		return nil, err
	}

	scannerReg, err := contract_scanner_registry.NewScannerRegistryCaller(regContracts.ScannerRegistry, ec)
	if err != nil {
		return nil, err
	}

	scannerPoolReg, err := contract_scanner_pool_registry.NewScannerPoolRegistryCaller(regContracts.ScannerPoolRegistry, ec)
	if err != nil {
		return nil, err
	}

	dispatch, err := contract_dispatch.NewDispatchCaller(regContracts.Dispatch, ec)
	if err != nil {
		return nil, err
	}

	scannerVersion, err := contract_scanner_node_version.NewScannerNodeVersionCaller(regContracts.ScannerNodeVersion, ec)
	if err != nil {
		return nil, err
	}

	fortaStaking, err := contract_forta_staking.NewFortaStakingCaller(regContracts.FortaStaking, ec)
	if err != nil {
		return nil, err
	}

	scannerRegFil, err := contract_scanner_registry.NewScannerRegistryFilterer(regContracts.ScannerRegistry, ec)
	if err != nil {
		return nil, err
	}

	scannerPoolRegFil, err := contract_scanner_pool_registry.NewScannerPoolRegistryFilterer(regContracts.ScannerPoolRegistry, ec)
	if err != nil {
		return nil, err
	}

	agentRegFil, err := contract_agent_registry.NewAgentRegistryFilterer(regContracts.AgentRegistry, ec)
	if err != nil {
		return nil, err
	}

	return &client{
		ctx: ctx,
		cfg: cfg,
		eth: eth,
		ec:  ec,

		privateKey: cfg.PrivateKey,

		contracts:         regContracts,
		scannerReg:        scannerReg,
		scannerPoolReg:    scannerPoolReg,
		agentReg:          agentReg,
		dispatch:          dispatch,
		scannerVersion:    scannerVersion,
		fortaStaking:      fortaStaking,
		scannerRegFil:     scannerRegFil,
		scannerPoolRegFil: scannerPoolRegFil,
		agentRegFil:       agentRegFil,
	}, err
}

func NewClient(ctx context.Context, cfg ClientConfig) (*client, error) {
	// avoids need to provide ENS address outside of dev environment use cases
	ensAddr := cfg.ENSAddress
	if ensAddr == "" {
		ensAddr = defaultEnsAddress
	}
	ensStore, err := ens.DialENSStoreAt(cfg.JsonRpcUrl, ensAddr)
	if err != nil {
		return nil, err
	}
	return NewClientWithENSStore(ctx, cfg, ensStore)
}

func (c *client) RegistryContracts() *registry.RegistryContracts {
	return c.contracts
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
	res, err := c.scannerReg.GetStakeThreshold(c.opts, sID)
	if err != nil {
		return nil, err
	}
	return (*StakingThreshold)(&res), nil
}

func (c *client) GetScannerNodeVersion() (string, error) {
	scannerVersion, err := c.scannerVersion.ScannerNodeVersion(c.opts)
	if err != nil {
		return "", err
	}
	return scannerVersion, nil
}

func (c *client) GetScannerNodePrereleaseVersion() (string, error) {
	scannerVersion, err := c.scannerVersion.ScannerNodeBetaVersion(c.opts)
	if err != nil {
		return "", err
	}
	return scannerVersion, nil
}

func (c *client) GetAssignmentHash(scannerID string) (*AssignmentHash, error) {
	sh, err := c.dispatch.ScannerHash(c.opts, utils.ScannerIDHexToBigInt(scannerID))
	if err != nil {
		return nil, err
	}
	return &AssignmentHash{
		AgentLength: sh.Length.Int64(),
		Hash:        utils.Bytes32ToHex(sh.Manifest),
	}, nil
}

func (c *client) getOps() (*bind.CallOpts, error) {
	if c.opts != nil {
		return c.opts, nil
	}
	return c.latestOpts()
}

func (c *client) IsAssigned(scannerID string, agentID string) (bool, error) {
	agtID := utils.AgentHexToBigInt(agentID)
	scnID := utils.ScannerIDHexToBigInt(scannerID)
	return c.dispatch.AreTheyLinked(c.opts, agtID, scnID)
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
	opts, err := c.getOps()
	if err != nil {
		return err
	}

	it, err := c.scannerRegFil.FilterScannerUpdated(&bind.FilterOpts{
		Start:   block,
		Context: c.ctx,
	}, nil, nil)

	if err != nil {
		return err
	}

	for it.Next() {
		if it.Event != nil {
			scn, err := c.scannerReg.GetScannerState(opts, it.Event.ScannerId)
			if err != nil {
				return err
			}
			if err := handler(it.Event, &Scanner{
				ScannerID: utils.ScannerIDBigIntToHex(it.Event.ScannerId),
				ChainID:   scn.ChainId.Int64(),
				Enabled:   scn.Enabled,
				Manifest:  scn.Metadata,
				Owner:     scn.Owner.Hex(),
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *client) ForEachPoolScannerSinceBlock(
	block uint64, handler func(event *contract_scanner_pool_registry.ScannerPoolRegistryScannerUpdated, s *Scanner) error,
) error {
	opts, err := c.getOps()
	if err != nil {
		return err
	}

	it, err := c.scannerPoolRegFil.FilterScannerUpdated(&bind.FilterOpts{
		Start:   block,
		Context: c.ctx,
	}, nil, nil)

	if err != nil {
		return err
	}

	for it.Next() {
		if it.Event != nil {
			scannerAddr := common.HexToAddress(utils.ScannerIDBigIntToHex(it.Event.ScannerId))
			scn, err := c.scannerPoolReg.GetScannerState(opts, scannerAddr)
			if err != nil {
				return err
			}
			if err := handler(it.Event, &Scanner{
				ScannerID: utils.ScannerIDBigIntToHex(it.Event.ScannerId),
				ChainID:   scn.ChainId.Int64(),
				Enabled:   scn.Operational,
				Manifest:  scn.Metadata,
				Owner:     scn.Owner.Hex(),
				PoolID:    utils.PoolIDBigIntToHex(it.Event.ScannerPool),
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *client) ForEachChainAgent(chainID int64, handler func(a *Agent) error) error {
	opts, err := c.getOps()
	if err != nil {
		return err
	}

	cID := big.NewInt(chainID)
	length, err := c.agentReg.GetAgentCountByChain(opts, cID)
	if err != nil {
		return err
	}

	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		agtID, err := c.agentReg.GetAgentByChainAndIndex(opts, cID, idx)
		if err != nil {
			return err
		}
		agt, err := c.agentReg.GetAgentState(opts, agtID)
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
	opts, err := c.getOps()
	if err != nil {
		return err
	}

	length, err := c.agentReg.GetAgentCount(opts)
	if err != nil {
		return err
	}

	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		agtID, err := c.agentReg.GetAgentByIndex(opts, idx)
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
	opts, err := c.getOps()
	if err != nil {
		return err
	}

	return c.ForEachAgentID(func(agentID string) error {
		agt, err := c.agentReg.GetAgentState(opts, utils.AgentHexToBigInt(agentID))
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
	opts, err := c.getOps()
	if err != nil {
		return err
	}

	it, err := c.agentRegFil.FilterAgentUpdated(&bind.FilterOpts{
		Start:   block,
		Context: c.ctx,
	}, nil, nil)

	if err != nil {
		return err
	}

	for it.Next() {
		if it.Event != nil {
			agt, err := c.agentReg.GetAgentState(opts, it.Event.AgentId)
			if err != nil {
				return err
			}
			if err := handler(it.Event, &Agent{
				AgentID:  utils.AgentBigIntToHex(it.Event.AgentId),
				ChainIDs: utils.IntArray(agt.ChainIds),
				Enabled:  agt.Enabled,
				Manifest: agt.Metadata,
				Owner:    agt.Owner.Hex(),
			}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *client) ForEachAssignedScanner(agentID string, handler func(s *Scanner) error) error {
	// if opts not set, peg for this loop
	opts, err := c.getOps()
	if err != nil {
		return err
	}

	aID := utils.AgentHexToBigInt(agentID)
	length, err := c.dispatch.NumScannersFor(opts, aID)
	if err != nil {
		return err
	}
	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		scn, err := c.dispatch.ScannerRefAt(opts, aID, idx)
		if err != nil {
			return err
		}
		if err := handler(&Scanner{
			ScannerID: utils.ScannerIDBigIntToHex(scn.ScannerId),
			ChainID:   scn.ChainId.Int64(),
			Enabled:   scn.Operational,
			Manifest:  scn.Metadata,
			Owner:     scn.Owner.Hex(),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (c *client) ForEachAssignedAgent(scannerID string, handler func(a *Agent) error) error {
	// if opts not set, peg for this loop
	opts, err := c.getOps()
	if err != nil {
		return err
	}

	sID := utils.ScannerIDHexToBigInt(scannerID)
	length, err := c.dispatch.NumAgentsFor(opts, sID)
	if err != nil {
		return err
	}
	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		agt, err := c.dispatch.AgentRefAt(opts, sID, idx)
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
	sID := utils.ScannerIDHexToBigInt(scannerID)
	return c.scannerReg.IsEnabled(c.opts, sID)
}

func (c *client) IsOperationalScanner(scannerID string) (bool, error) {
	return c.scannerPoolReg.IsScannerOperational(c.opts, common.HexToAddress(scannerID))
}

func (c *client) GetActiveScannerStake(scannerID string) (*big.Int, error) {
	sID := utils.ScannerIDHexToBigInt(scannerID)
	return c.fortaStaking.ActiveStakeFor(c.opts, 0, sID)
}

func (c *client) GetScanner(scannerID string) (*Scanner, error) {
	sID := utils.ScannerIDHexToBigInt(scannerID)
	scn, err := c.scannerReg.GetScanner(c.opts, sID)

	if err != nil {
		return nil, err
	}

	if !scn.Registered {
		return nil, nil
	}

	enabled, err := c.scannerReg.IsEnabled(c.opts, sID)
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
	sID := common.HexToAddress(scannerID)
	scn, err := c.scannerPoolReg.GetScanner(c.opts, sID)

	if err != nil {
		return nil, err
	}

	if !scn.Registered {
		return nil, nil
	}

	enabled, err := c.scannerPoolReg.IsScannerOperational(c.opts, sID)
	if err != nil {
		return nil, err
	}

	owner, err := c.scannerPoolReg.OwnerOf(c.opts, scn.ScannerPoolId)
	if err != nil {
		return nil, err
	}

	return &Scanner{
		ScannerID: sID.Hex(),
		ChainID:   scn.ChainId.Int64(),
		Enabled:   enabled,
		Manifest:  scn.Metadata,
		Owner:     owner.Hex(),
		PoolID:    utils.PoolIDBigIntToHex(scn.ScannerPoolId),
	}, nil
}

func (c *client) GetAgent(agentID string) (*Agent, error) {
	aID := utils.AgentHexToBigInt(agentID)
	agt, err := c.agentReg.GetAgent(c.opts, aID)
	if err != nil {
		return nil, err
	}

	// if agt does not exist, return nil
	if !agt.Registered {
		return nil, nil
	}

	enabled, err := c.agentReg.IsEnabled(c.opts, aID)
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

func (c *client) RegisterScannerToPool(scannerAddress string, poolID *big.Int, chainID int64, metadata string, signature []byte) (txHash string, err error) {
	if c.privateKey == nil {
		return "", errors.New("no operator private key provided to the client")
	}
	reg, err := contract_scanner_pool_registry.NewScannerPoolRegistryTransactor(c.contracts.ScannerPoolRegistry, c.ec)
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
	tx, err := reg.RegisterScannerNode(opts, contract_scanner_pool_registry.ScannerPoolRegistryCoreScannerNodeRegistration{
		Scanner:       common.HexToAddress(scannerAddress),
		ScannerPoolId: poolID,
		ChainId:       big.NewInt(int64(chainID)),
		Metadata:      metadata,
		Timestamp:     big.NewInt(time.Now().Unix()),
	}, signature)
	if err != nil {
		return "", fmt.Errorf("failed to send the transaction: %v", err)
	}
	return tx.Hash().Hex(), nil
}

func (c *client) EnableScanner(permission ScannerPermission, scannerAddress string) (txHash string, err error) {
	if c.privateKey == nil {
		return "", errors.New("no private key provided to the client")
	}
	reg, err := contract_scanner_registry.NewScannerRegistryTransactor(c.contracts.ScannerRegistry, c.ec)
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
	reg, err := contract_scanner_registry.NewScannerRegistryTransactor(c.contracts.ScannerRegistry, c.ec)
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

func (c *client) GenerateScannerRegistrationSignature(reg *eip712.ScannerNodeRegistration) ([]byte, error) {
	return eip712.SignScannerRegistration(c.privateKey, c.contracts.ScannerPoolRegistry, reg)
}

func (c *client) GetScannerPoolOwner(poolID *big.Int) (owner string, err error) {
	addr, err := c.scannerPoolReg.OwnerOf(c.opts, poolID)
	if err != nil {
		return "", nil
	}
	return addr.Hex(), nil
}
