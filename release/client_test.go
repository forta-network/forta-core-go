package release

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetReleaseManifest(t *testing.T) {
	c, err := NewClient("https://ipfs.forta.network")
	assert.NoError(t, err)

	rm, err := c.GetReleaseManifest(context.Background(), "QmfFajdVEMZWNTkaK4Z49xGhK3VWDhx8tKwfKCmu7VApDx")
	assert.NoError(t, err)

	assert.Equal(t, "9470f9e7a71279dee141757f54c07d8e74b0dcb3", rm.Release.Commit)
}
