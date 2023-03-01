package ens

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/forta-network/forta-core-go/domain/registry"
	"github.com/wealdtech/go-ens/v3"
)

const (
	DispatchContract            = "dispatch.forta.eth"
	AgentRegistryContract       = "agents.registries.forta.eth"
	ScannerRegistryContract     = "scanners.registries.forta.eth"
	ScannerNodeVersionContract  = "scanner-node-version.forta.eth"
	StakingContract             = "staking.forta.eth"
	FortaContract               = "forta.eth"
	RewardsDistributorContract  = "rewards.forta.eth"
	StakeAllocatorContract      = "allocator.forta.eth"
	ScannerPoolRegistryContract = "pool.registries.forta.eth"

	ENSResolverDev  = "0x5f7c5bbBa72e1e1fae689120D76D2f334A390Ae9"
	ENSResolverProd = "0x08f42fcc52a9C2F391bF507C4E8688D0b53e1bd7"
)

// ENS resolves inputs.
type ENS interface {
	Resolver
	ResolveRegistryContracts() (*registry.RegistryContracts, error)
}

// ENSStore wraps the ENS client which interacts with namespace contract(s).
type ENSStore struct {
	Resolver
}

// NewENSStore creates a new store.
func NewENSStore(backend bind.ContractBackend) *ENSStore {
	return &ENSStore{Resolver: &ENSResolver{backend: backend}}
}

// DialENSStore dials an Ethereum API and creates a new store.
func DialENSStore(rpcUrl string) (*ENSStore, error) {
	client, err := rpc.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	return &ENSStore{Resolver: &ENSResolver{backend: ethclient.NewClient(client)}}, nil
}

// DialENSStoreAt dials an Ethereum API and creates a new store that works with a resolver at given address.
func DialENSStoreAt(rpcUrl, resolverAddr string) (*ENSStore, error) {
	client, err := rpc.Dial(rpcUrl)
	if err != nil {
		return nil, err
	}
	return &ENSStore{Resolver: &ENSResolver{backend: ethclient.NewClient(client), resolverAddr: resolverAddr}}, nil
}

// NewENStoreWithResolver creates a new store with custom resolver.
func NewENStoreWithResolver(resolver Resolver) *ENSStore {
	return &ENSStore{Resolver: resolver}
}

// Resolver resolves inputs.
type Resolver interface {
	Resolve(input string) (common.Address, error)
}

// ENSResolver resolves names from an ENS contract.
type ENSResolver struct {
	backend      bind.ContractBackend
	resolverAddr string
}

// Resolve resolves an input to an address.
func (ensResolver *ENSResolver) Resolve(input string) (common.Address, error) {
	if len(ensResolver.resolverAddr) == 0 {
		return ens.Resolve(ensResolver.backend, input)
	}
	resolver, err := ens.NewResolverAt(ensResolver.backend, input, common.HexToAddress(ensResolver.resolverAddr))
	if err != nil {
		return common.Address{}, err
	}
	// Resolve the domain
	address, err := resolver.Address()
	if err != nil {
		return ens.UnknownAddress, err
	}
	if bytes.Equal(address.Bytes(), ens.UnknownAddress.Bytes()) {
		return ens.UnknownAddress, fmt.Errorf("no address for %s", input)
	}
	return address, nil
}

// ResolverFunc helps implementing a custom resolver with a function.
type ResolverFunc func(input string) (common.Address, error)

// Resolve implements the Resolver interface.
func (rf ResolverFunc) Resolve(input string) (common.Address, error) {
	return rf(input)
}

func (ensstore *ENSStore) ResolveRegistryContracts() (*registry.RegistryContracts, error) {
	agentReg, err := ensstore.Resolve(AgentRegistryContract)
	if err != nil {
		return nil, err
	}

	scannerReg, err := ensstore.Resolve(ScannerRegistryContract)
	if err != nil {
		return nil, err
	}

	dispatch, err := ensstore.Resolve(DispatchContract)
	if err != nil {
		return nil, err
	}

	scannerNodeVersion, err := ensstore.Resolve(ScannerNodeVersionContract)
	if err != nil {
		return nil, err
	}

	fortaStaking, err := ensstore.Resolve(StakingContract)
	if err != nil {
		return nil, err
	}

	forta, err := ensstore.Resolve(FortaContract)
	if err != nil {
		return nil, err
	}

	rewards, err := ensstore.Resolve(RewardsDistributorContract)
	if err != nil {
		return nil, err
	}

	regContracts := &registry.RegistryContracts{
		AgentRegistry:      agentReg,
		ScannerRegistry:    scannerReg,
		Dispatch:           dispatch,
		ScannerNodeVersion: scannerNodeVersion,
		FortaStaking:       fortaStaking,
		Forta:              forta,
		RewardsDistributor: rewards,
	}

	return regContracts, nil

}
