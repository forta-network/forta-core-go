package registry

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/forta-protocol/forta-core-go/contracts/contract_agent_registry"
	"github.com/forta-protocol/forta-core-go/contracts/contract_dispatch"
	"github.com/forta-protocol/forta-core-go/contracts/contract_scanner_registry"
	"github.com/forta-protocol/forta-core-go/ens"
	"github.com/forta-protocol/forta-core-go/ethereum"
	"github.com/forta-protocol/forta-core-go/utils"
)

//Client retrieves agent, scanner, and assignment information from the registry contracts
type Client interface {
	//PegLatestBlock will set the opts so that every call uses same block
	PegLatestBlock() error

	//ResetOpts unsets the options for the store
	ResetOpts()

	//GetAssignmentHash returns a hash of all agents, helpful for knowing scanner's agents have changed
	GetAssignmentHash(scannerID string) (*AssignmentHash, error)

	//ForEachAssignedAgent invokes a handler for each agent assigned to the scanner
	ForEachAssignedAgent(scannerID string, handler func(a *Agent) error) error

	//IsEnabledScanner returns true if the scanner exists and is enabled
	IsEnabledScanner(scannerID string) (bool, error)

	//GetAgent returns the registry information for the agent
	GetAgent(agentID string) (*Agent, error)
}

var zero = big.NewInt(0)

type client struct {
	ctx context.Context
	cfg ClientConfig
	eth ethereum.Client

	// call PegLatestBlock to peg the context to the latest block
	opts *bind.CallOpts

	ar *contract_agent_registry.AgentRegistryCaller
	sr *contract_scanner_registry.ScannerRegistryCaller
	dp *contract_dispatch.DispatchCaller
}

type ClientConfig struct {
	//JsonRpcUrl is the url used for resolving the registry (polygon)
	JsonRpcUrl string `json:"jsonRpcUrl"`

	//ENSAddress is if there's not a default contract for ENS resolution
	ENSAddress string `json:"ensAddress"`

	//Name is used for logging
	Name string `json:"name"`
}

var defaultConfig = ClientConfig{
	JsonRpcUrl: "https://polygon-rpc.com",
	ENSAddress: "0x08f42fcc52a9C2F391bF507C4E8688D0b53e1bd7",
	Name:       "registry-client",
}

func NewDefaultClient(ctx context.Context) (*client, error) {
	return NewClient(ctx, defaultConfig)
}

func NewClient(ctx context.Context, cfg ClientConfig) (*client, error) {
	eth, err := ethereum.NewStreamEthClient(ctx, cfg.Name, cfg.JsonRpcUrl)
	if err != nil {
		return nil, err
	}

	ensStore, err := ens.DialENSStoreAt(cfg.JsonRpcUrl, cfg.ENSAddress)
	if err != nil {
		return nil, err
	}

	regContracts, err := ensStore.ResolveRegistryContracts()

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

	return &client{
		eth: eth,

		sr: sr,
		ar: ar,
		dp: dp,
	}, err
}

func scannerIDtoBigInt(scannerID string) *big.Int {
	return common.HexToHash(scannerID).Big()
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

//PegLatestBlock will set the opts so that every call uses same block
func (c *client) PegLatestBlock() error {
	opts, err := c.latestOpts()
	if err != nil {
		return err
	}
	c.opts = opts
	return nil
}

func (c *client) GetAssignmentHash(scannerID string) (*AssignmentHash, error) {
	sh, err := c.dp.ScannerHash(c.opts, scannerIDtoBigInt(scannerID))
	if err != nil {
		return nil, err
	}
	return &AssignmentHash{
		AgentLength: sh.Length.Int64(),
		Hash:        utils.Bytes32ToHex(sh.Manifest),
	}, nil
}

func (c *client) ForEachAssignedAgent(scannerID string, handler func(a *Agent) error) error {
	// if opts not set, peg for this loop
	opts := c.opts
	if opts == nil {
		currentOpts, err := c.latestOpts()
		if err != nil {
			return err
		}
		opts = currentOpts
	}

	sID := scannerIDtoBigInt(scannerID)
	length, err := c.dp.AgentsFor(opts, sID)
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
			AgentID:  utils.BigIntToHex(agt.AgentId),
			ChainIDs: utils.IntArray(agt.ChainIds),
			Enabled:  agt.Enabled,
			Manifest: agt.Metadata,
		}); err != nil {
			return err
		}
	}
	return nil
}

func isZeroAddress(address common.Address) bool {
	return 0 == common.HexToHash(address.Hex()).Big().Cmp(zero)
}

func (c *client) IsEnabledScanner(scannerID string) (bool, error) {
	sID := scannerIDtoBigInt(scannerID)
	owner, err := c.sr.OwnerOf(c.opts, sID)
	if err != nil || isZeroAddress(owner) {
		// owner returns an error when not existing
		return false, nil
	}
	return c.sr.IsEnabled(c.opts, sID)
}

func (c *client) GetAgent(agentID string) (*Agent, error) {
	aID, err := utils.HexToBigInt(agentID)
	if err != nil {
		return nil, err
	}
	agt, err := c.ar.GetAgent(c.opts, aID)
	if err != nil {
		return nil, err
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
	}, nil
}
