package settings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateChainSettings(t *testing.T) {
	for _, chainSettings := range allChainSettings {
		t.Run(chainSettings.Name, func(t *testing.T) {
			r := require.New(t)

			r.True(ValidateChainSettings(chainSettings.ChainID))
		})
	}
}
