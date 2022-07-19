package release

import (
	"context"

	"github.com/forta-network/forta-core-go/ipfs"
)

type Client interface {
	GetReleaseManifest(ctx context.Context, reference string) (*ReleaseManifest, error)
}

type client struct {
	ic ipfs.Client
}

func (c *client) GetReleaseManifest(ctx context.Context, reference string) (*ReleaseManifest, error) {
	var rm ReleaseManifest
	if err := c.ic.UnmarshalJson(ctx, reference, &rm); err != nil {
		return nil, err
	}
	return &rm, nil
}

func NewClient(ipfsGateway string) (*client, error) {
	ic, err := ipfs.NewClient(ipfsGateway)
	if err != nil {
		return nil, err
	}
	return &client{ic: ic}, nil
}
