package release

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetReleaseManifestFromIPFS(t *testing.T) {
	c, err := NewClient("https://ipfs.forta.network", nil)
	assert.NoError(t, err)

	rm, err := c.GetReleaseManifest("QmPU5yx6Puapj7o79zY4n1LkyhnZowLjoDaZv7TRhytYpt")
	assert.NoError(t, err)

	assert.Equal(t, "7ef0aca978a3d7ed9d10170442e80a4e12a9a7b5", rm.Release.Commit)
}

func TestClient_GetReleaseManifestFromDist(t *testing.T) {
	c, err := NewClient("https://invalid.forta.network", []string{"https://dist-dev.forta.network/manifests/releases"})
	assert.NoError(t, err)

	rm, err := c.GetReleaseManifest("QmVDqk2pFVJ6joc8sNsNjsT1vzkjnYgz7X1dVw8eT1EgnR")
	assert.NoError(t, err)

	assert.Equal(t, "6c63eeee87cc99819899008ffc5db3165636bfe3", rm.Release.Commit)
}
