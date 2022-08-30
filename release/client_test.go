package release

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetReleaseManifest(t *testing.T) {
	c, err := NewClient("https://ipfs.forta.network")
	assert.NoError(t, err)

	rm, err := c.GetReleaseManifest(context.Background(), "QmPU5yx6Puapj7o79zY4n1LkyhnZowLjoDaZv7TRhytYpt")
	assert.NoError(t, err)

	assert.Equal(t, "7ef0aca978a3d7ed9d10170442e80a4e12a9a7b5", rm.Release.Commit)
}
