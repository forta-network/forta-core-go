package registry

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/forta-network/forta-core-go/contracts/contract_forta_staking"
	"github.com/forta-network/forta-core-go/domain/registry"

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
	zeroAddress                = "0x0000000000000000000000000000000000000000"
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

	// RegisterScanner registers a scanner using private key.
	RegisterScanner(ownerAddress string, chainID int64, metadata string) (txHash string, err error)

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

	// ForEachScanner gets all scanners
	ForEachScanner(handler func(s *Scanner) error) error

	// ForEachAssignedScanner loops over scanners by agent
	ForEachAssignedScanner(agentID string, handler func(s *Scanner) error) error

	// GetStakingThreshold returns the min/max/activated flag for a given address
	GetStakingThreshold(scannerID string) (*StakingThreshold, error)

	// EnableScanner enables a scanner.
	EnableScanner(ScannerPermission ScannerPermission, scannerAddress string) (txHash string, err error)

	// DisableScanner disables a scanner.
	DisableScanner(ScannerPermission ScannerPermission, scannerAddress string) (txHash string, err error)
}

type client struct {
	ctx context.Context
	cfg ClientConfig
	eth ethereum.Client
	ec  *ethclient.Client

	// call PegLatestBlock to peg the context to the latest block
	opts       *bind.CallOpts
	privateKey *ecdsa.PrivateKey

	contracts *registry.RegistryContracts
	ar        *contract_agent_registry.AgentRegistryCaller
	sr        *contract_scanner_registry.ScannerRegistryCaller
	dp        *contract_dispatch.DispatchCaller
	sv        *contract_scanner_node_version.ScannerNodeVersionCaller
	fs        *contract_forta_staking.FortaStakingCaller
	srf       *contract_scanner_registry.ScannerRegistryFilterer
}

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

	ar, err := contract_agent_registry.NewAgentRegistryCaller(regContracts.AgentRegistry, ec)
	if err != nil {
		return nil, err
	}

	sr, err := contract_scanner_registry.NewScannerRegistryCaller(regContracts.ScannerRegistry, ec)
	if err != nil {
		return nil, err
	}

	dp, err := contract_dispatch.NewDispatchCaller(regContracts.Dispatch, ec)
	if err != nil {
		return nil, err
	}

	sv, err := contract_scanner_node_version.NewScannerNodeVersionCaller(regContracts.ScannerNodeVersion, ec)
	if err != nil {
		return nil, err
	}

	fs, err := contract_forta_staking.NewFortaStakingCaller(regContracts.FortaStaking, ec)
	if err != nil {
		return nil, err
	}

	srf, err := contract_scanner_registry.NewScannerRegistryFilterer(regContracts.ScannerRegistry, ec)
	if err != nil {
		return nil, err
	}

	return &client{
		ctx: ctx,
		cfg: cfg,
		eth: eth,
		ec:  ec,

		privateKey: cfg.PrivateKey,

		contracts: regContracts,
		sr:        sr,
		ar:        ar,
		dp:        dp,
		sv:        sv,
		fs:        fs,
		srf:       srf,
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

//ResetOpts unsets the options for the store
func (c *client) ResetOpts() {
	c.opts = nil
}

//latestOpts returns the callopts for the latest block so that calls can use a same block
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

//PegLatestBlock will set the opts so that every call uses same block
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
	res, err := c.sr.GetStakeThreshold(c.opts, sID)
	if err != nil {
		return nil, err
	}
	return (*StakingThreshold)(&res), nil
}

func (c *client) GetScannerNodeVersion() (string, error) {
	sv, err := c.sv.ScannerNodeVersion(c.opts)
	if err != nil {
		return "", err
	}
	return sv, nil
}

func (c *client) GetScannerNodePrereleaseVersion() (string, error) {
	sv, err := c.sv.ScannerNodeBetaVersion(c.opts)
	if err != nil {
		return "", err
	}
	return sv, nil
}

func (c *client) GetAssignmentHash(scannerID string) (*AssignmentHash, error) {
	sh, err := c.dp.ScannerHash(c.opts, utils.ScannerIDHexToBigInt(scannerID))
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
	return c.dp.AreTheyLinked(c.opts, agtID, scnID)
}

func (c *client) ForEachScanner(handler func(s *Scanner) error) error {
	opts, err := c.getOps()
	if err != nil {
		return err
	}

	it, err := c.srf.FilterTransfer(&bind.FilterOpts{
		Start:   scannerRegistryDeployBlock,
		Context: c.ctx,
	},
		nil, nil, nil)

	if err != nil {
		return err
	}

	for it.Next() {
		if it.Event != nil && it.Event.From.Hex() == zeroAddress {
			scn, err := c.sr.GetScannerState(opts, it.Event.TokenId)
			if err != nil {
				return err
			}
			if err := handler(&Scanner{
				ScannerID: utils.ScannerIDBigIntToHex(it.Event.TokenId),
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

func (c *client) ForEachChainAgent(chainID int64, handler func(a *Agent) error) error {
	opts, err := c.getOps()
	if err != nil {
		return err
	}

	cID := big.NewInt(chainID)
	length, err := c.ar.GetAgentCountByChain(opts, cID)
	if err != nil {
		return err
	}

	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		agtID, err := c.ar.GetAgentByChainAndIndex(opts, cID, idx)
		if err != nil {
			return err
		}
		agt, err := c.ar.GetAgentState(opts, agtID)
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

	length, err := c.ar.GetAgentCount(opts)
	if err != nil {
		return err
	}

	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		agtID, err := c.ar.GetAgentByIndex(opts, idx)
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
		agt, err := c.ar.GetAgentState(opts, utils.AgentHexToBigInt(agentID))
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

func (c *client) ForEachAssignedScanner(agentID string, handler func(s *Scanner) error) error {
	// if opts not set, peg for this loop
	opts, err := c.getOps()
	if err != nil {
		return err
	}

	aID := utils.AgentHexToBigInt(agentID)
	length, err := c.dp.NumScannersFor(opts, aID)
	if err != nil {
		return err
	}
	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		scn, err := c.dp.ScannerRefAt(opts, aID, idx)
		if err != nil {
			return err
		}
		if err := handler(&Scanner{
			ScannerID: utils.ScannerIDBigIntToHex(scn.ScannerId),
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

func (c *client) ForEachAssignedAgent(scannerID string, handler func(a *Agent) error) error {
	// if opts not set, peg for this loop
	opts, err := c.getOps()
	if err != nil {
		return err
	}

	sID := utils.ScannerIDHexToBigInt(scannerID)
	length, err := c.dp.NumAgentsFor(opts, sID)
	if err != nil {
		return err
	}
	for i := int64(0); i < length.Int64(); i++ {
		idx := big.NewInt(i)
		agt, err := c.dp.AgentRefAt(opts, sID, idx)
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
	return c.sr.IsEnabled(c.opts, sID)
}

func (c *client) GetActiveScannerStake(scannerID string) (*big.Int, error) {
	sID := utils.ScannerIDHexToBigInt(scannerID)
	return c.fs.ActiveStakeFor(c.opts, 0, sID)
}

func (c *client) GetScanner(scannerID string) (*Scanner, error) {
	sID := utils.ScannerIDHexToBigInt(scannerID)
	scn, err := c.sr.GetScanner(c.opts, sID)

	if err != nil {
		return nil, err
	}

	if !scn.Registered {
		return nil, nil
	}

	enabled, err := c.sr.IsEnabled(c.opts, sID)
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

func (c *client) GetAgent(agentID string) (*Agent, error) {
	aID := utils.AgentHexToBigInt(agentID)
	agt, err := c.ar.GetAgent(c.opts, aID)
	if err != nil {
		return nil, err
	}

	// if agt does not exist, return nil
	if !agt.Created {
		return nil, nil
	}

	enabled, err := c.ar.IsEnabled(c.opts, aID)
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

func (c *client) RegisterScanner(ownerAddress string, chainID int64, metadata string) (txHash string, err error) {
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
	tx, err := reg.Register(opts, common.HexToAddress(ownerAddress), big.NewInt(int64(chainID)), metadata)
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
