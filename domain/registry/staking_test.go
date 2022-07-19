package registry

import (
	"testing"

	"github.com/forta-network/forta-core-go/utils"
	"github.com/stretchr/testify/assert"
)

func TestExtractions(t *testing.T) {
	type typeTest struct {
		tokenID   string
		stakeType string
		active    bool
	}

	tests := []typeTest{
		{stakeType: ScannerStake, active: true, tokenID: "0xa6b9aec07b2ded8ea06a28b559757b4c826b43aae8c8823379637f5c933cb900"},
		{stakeType: ScannerStake, active: false, tokenID: "0x5994d5db1538c3334c15787cffb7e5dd4ff4fcabd6cd0723ca7da483edf02000"},
	}

	for _, test := range tests {
		tokenID, _ := utils.HexToBigInt(test.tokenID)
		stakeType, err := extractStakeType(tokenID)
		assert.NoError(t, err, test.tokenID)
		active := isActive(tokenID)
		assert.Equal(t, test.stakeType, stakeType, test.tokenID)
		assert.Equal(t, test.active, active, test.tokenID)
	}
}
