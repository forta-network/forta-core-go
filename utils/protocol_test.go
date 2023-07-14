package utils

import (
	"testing"

	"github.com/forta-network/forta-core-go/protocol"
	"github.com/stretchr/testify/assert"
)

func TestGetChainIDsForAlert(t *testing.T) {
	scenarios := []struct {
		name     string
		alert    *protocol.SignedAlert
		expected []uint64
		wantErr  bool
	}{
		{
			name: "valid alert with all valid chainIDs",
			alert: &protocol.SignedAlert{
				Alert: &protocol.Alert{
					Finding: &protocol.Finding{
						Source: &protocol.Source{
							Chains:       []*protocol.Source_ChainSource{{ChainId: 1}, {ChainId: 2}},
							Transactions: []*protocol.Source_TransactionSource{{ChainId: 3}},
							Blocks:       []*protocol.Source_BlockSource{{ChainId: 4}},
						},
					},
				},
				ChainId: "0x5",
			},
			expected: []uint64{1, 2, 3, 4, 5},
			wantErr:  false,
		},
		{
			name: "valid alert with zeros that are ignored",
			alert: &protocol.SignedAlert{
				Alert: &protocol.Alert{
					Finding: &protocol.Finding{
						Source: &protocol.Source{
							Chains:       []*protocol.Source_ChainSource{{ChainId: 1}, {ChainId: 2}, {ChainId: 0}},
							Transactions: []*protocol.Source_TransactionSource{{ChainId: 3}, {ChainId: 0}},
							Blocks:       []*protocol.Source_BlockSource{{ChainId: 4}, {ChainId: 0}},
						},
					},
				},
				ChainId: "0x5",
			},
			expected: []uint64{1, 2, 3, 4, 5},
			wantErr:  false,
		},
		{
			name: "invalid alert with no scanner chainId",
			alert: &protocol.SignedAlert{
				Alert: &protocol.Alert{
					Finding: &protocol.Finding{
						Source: &protocol.Source{
							Chains:       []*protocol.Source_ChainSource{{ChainId: 1}, {ChainId: 2}},
							Transactions: []*protocol.Source_TransactionSource{{ChainId: 3}},
							Blocks:       []*protocol.Source_BlockSource{{ChainId: 4}},
						},
					},
				},
			},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result, err := GetChainIDsForAlert(scenario.alert)
			if scenario.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.ElementsMatch(t, scenario.expected, result)
		})
	}
}
