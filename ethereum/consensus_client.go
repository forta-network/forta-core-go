package ethereum

import (
	"context"

	"github.com/prysmaticlabs/prysm/v3/api/client/beacon"
	"github.com/prysmaticlabs/prysm/v3/consensus-types/interfaces"
	"github.com/prysmaticlabs/prysm/v3/encoding/ssz/detect"
)

// ConsensusClient is an interface encompassing all ethereum consensus actions
type ConsensusClient interface {
	GetBlocksBySlot(ctx context.Context, slot string) (interfaces.SignedBeaconBlock, error)
}

// consensusEthClient wraps an ethereum client purpose-built for communicating with consensus layer.
type consensusEthClient struct {
	apiName string
	url     string

	client *beacon.Client
}

func (c *consensusEthClient) GetBlocksBySlot(ctx context.Context, slot string) (interfaces.SignedBeaconBlock, error) {
	// use state information to get correct unmarshaler
	sb, err := c.client.GetState(ctx, beacon.StateOrBlockId(slot))
	if err != nil {
		return nil, err
	}
	vu, err := detect.FromState(sb)
	if err != nil {
		return nil, err
	}

	// get beacon block data
	bb, err := c.client.GetBlock(ctx, beacon.StateOrBlockId(slot))
	if err != nil {
		return nil, err
	}

	// unmarshal data with appropriate unmarshaler
	b, err := vu.UnmarshalBeaconBlock(bb)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// NewConsensusETHClient creates a new ethereum client
func NewConsensusETHClient(apiName, url string) (*consensusEthClient, error) {
	beaconClient, err := beacon.NewClient(url)
	if err != nil {
		return nil, err
	}

	return &consensusEthClient{apiName: apiName, url: url, client: beaconClient}, nil
}
