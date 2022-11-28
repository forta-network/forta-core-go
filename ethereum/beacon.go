package ethereum

import (
	"context"

	"github.com/prysmaticlabs/prysm/v3/api/client/beacon"
	"github.com/prysmaticlabs/prysm/v3/consensus-types/interfaces"
	"github.com/prysmaticlabs/prysm/v3/encoding/ssz/detect"
)

// BeaconClient is an interface encompassing all ethereum consensus actions
type BeaconClient interface {
	GetBlock(ctx context.Context, slot beacon.StateOrBlockId) (interfaces.SignedBeaconBlock, error)
}

// beaconClient wraps an ethereum client purpose-built for communicating with consensus layer.
type beaconClient struct {
	apiName string
	url     string

	client *beacon.Client
}

func (c *beaconClient) GetBlock(ctx context.Context, slot beacon.StateOrBlockId) (interfaces.SignedBeaconBlock, error) {
	// get beacon block data
	bb, err := c.client.GetBlock(ctx, beacon.StateOrBlockId(slot))
	if err != nil {
		return nil, err
	}

	// unmarshal data with appropriate unmarshaler
	// this depends on the slot. Pre-merge and Post-merge blocks are incompatible to use the same unmarshaler.
	unmarshaler, err := c.getVersionedUnmarshaler(ctx, slot)
	if err != nil {
		return nil, err
	}

	b, err := unmarshaler.UnmarshalBeaconBlock(bb)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *beaconClient) getVersionedUnmarshaler(ctx context.Context, state beacon.StateOrBlockId) (*detect.VersionedUnmarshaler, error) {
	f, err := c.client.GetFork(ctx, state)
	if err != nil {
		return nil, err
	}

	var version [4]byte
	copy(version[:], f.CurrentVersion[:4])

	return detect.FromForkVersion(version)
}

// NewBeaconClient creates a new ethereum client
func NewBeaconClient(apiName, url string) (BeaconClient, error) {
	bClient, err := beacon.NewClient(url)
	if err != nil {
		return nil, err
	}

	return &beaconClient{apiName: apiName, url: url, client: bClient}, nil
}

