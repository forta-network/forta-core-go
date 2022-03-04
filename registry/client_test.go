package registry

import (
	"context"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestClient_GetScanner(t *testing.T) {
	c, err := NewDevClient(context.Background())
	assert.NoError(t, err)

	scannerID := "0x3f88c2b3e267e6b8e9de017cdb47a59ac9ecb284"
	depositedAt := int64(25365106)
	withdrawAt := int64(25365157)

	c.PegBlock(big.NewInt(depositedAt - 1))
	beforeDeposit, err := c.GetActiveScannerStake(scannerID)
	assert.NoError(t, err)

	c.PegBlock(big.NewInt(depositedAt + 1))
	afterDeposit, err := c.GetActiveScannerStake(scannerID)
	assert.NoError(t, err)

	c.PegBlock(big.NewInt(withdrawAt + 1))
	afterWithdraw, err := c.GetActiveScannerStake(scannerID)
	assert.NoError(t, err)

	t.Logf("before deposit: %d", beforeDeposit.Int64())
	t.Logf("after deposit: %d", afterDeposit.Int64())
	t.Logf("after withdraw: %d", afterWithdraw.Int64())
}
