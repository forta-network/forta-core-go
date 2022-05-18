package registry

import (
	"github.com/forta-network/forta-core-go/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEthConvert(t *testing.T) {
	res, _ := utils.HexToBigInt("0x468cd99899c5da6c")
	v := toEthFloat(res)
	assert.Equal(t, 5.083677328832584, v)
}
