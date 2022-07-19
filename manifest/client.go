package manifest

import (
	"context"

	"github.com/forta-network/forta-core-go/ipfs"
)

type Client interface {
	GetAgentManifest(ctx context.Context, reference string) (*SignedAgentManifest, error)
}

type client struct {
	ic ipfs.Client
}

func (c *client) GetAgentManifest(ctx context.Context, reference string) (*SignedAgentManifest, error) {
	var m SignedAgentManifest
	if err := c.ic.UnmarshalJson(ctx, reference, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func NewClient(ipfsGateway string) (*client, error) {
	ic, err := ipfs.NewClient(ipfsGateway)
	if err != nil {
		return nil, err
	}
	return &client{ic: ic}, nil
}
