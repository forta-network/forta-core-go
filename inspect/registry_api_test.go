package inspect

import (
	"context"
	"testing"

	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
)

var testRegistryEnv struct {
	RegistryAPI string `envconfig:"registry_api" default:"https://rpc.ankr.com/polygon"`
}

func init() {
	envconfig.MustProcess("test", &testRegistryEnv)
}

func TestRegistryAPIInspection(t *testing.T) {
	r := require.New(t)

	inspector := &RegistryAPIInspector{}
	results, err := inspector.Inspect(
		context.Background(), InspectionConfig{
			RegistryAPIURL:     testRegistryEnv.RegistryAPI,
			ENSContractAddress: "0x08f42fcc52a9C2F391bF507C4E8688D0b53e1bd7",
			ScannerAddress:     "0x3DC45b47B7559Ca3b231E5384D825F9B461A0398",
		},
	)
	r.NoError(err)

	r.Equal(
		map[string]float64{
			IndicatorRegistryAPIAccessible:  ResultSuccess,
			IndicatorRegistryAPIENS:         ResultSuccess,
			IndicatorRegistryAPIAssignments: ResultSuccess,
		}, results.Indicators,
	)
}
