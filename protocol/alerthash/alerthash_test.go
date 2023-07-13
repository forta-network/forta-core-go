package alerthash

import (
	"github.com/forta-network/forta-core-go/protocol"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestForBlockAlert(t *testing.T) {
	scenarios := []struct {
		name     string
		inputs   *Inputs
		expected string
	}{
		{
			name: "basic scenario",
			inputs: &Inputs{
				BlockEvent: &protocol.BlockEvent{
					Network: &protocol.BlockEvent_Network{
						ChainId: "1",
					},
					Block: &protocol.BlockEvent_EthBlock{
						Hash: "0x123",
					},
				},
				Finding: &protocol.Finding{
					AlertId:     "alert1",
					Name:        "Alert 1",
					Description: "This is Alert 1",
					Protocol:    "Protocol 1",
					Type:        protocol.Finding_EXPLOIT,
					Severity:    protocol.Finding_HIGH,
				},
			},
			expected: "0x30d8edcfb73e206876e737218c10c78b108a8790a3d1b4b9944b2958ef5a62bf",
		}, {
			name: "unique key",
			inputs: &Inputs{
				Finding: &protocol.Finding{
					AlertId:     "alert1",
					Name:        "Alert 1",
					Description: "This is Alert 1",
					Protocol:    "Protocol 1",
					Type:        protocol.Finding_EXPLOIT,
					Severity:    protocol.Finding_HIGH,
					UniqueKey:   "unique",
				},
			},
			expected: "0xe518cb690acf4dca0950069e234975b17358b549906206ab3065d26e6379a810",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			hash := ForBlockAlert(scenario.inputs)
			assert.Equal(t, scenario.expected, hash)
		})
	}
}

func TestForTransactionAlert(t *testing.T) {
	scenarios := []struct {
		name     string
		inputs   *Inputs
		expected string
	}{
		{
			name: "basic scenario",
			inputs: &Inputs{
				TransactionEvent: &protocol.TransactionEvent{
					Network: &protocol.TransactionEvent_Network{
						ChainId: "1",
					},
					Transaction: &protocol.TransactionEvent_EthTransaction{
						Hash: "0x123",
					},
				},
				Finding: &protocol.Finding{
					AlertId:     "alert1",
					Name:        "Alert 1",
					Description: "This is Alert 1",
					Protocol:    "Protocol 1",
					Type:        protocol.Finding_EXPLOIT,
					Severity:    protocol.Finding_HIGH,
				},
			},
			expected: "0x4b817e325c8e6a8f30844d7a850876a8620370c31fd3cc7c4e128bff1e7ac80f",
		}, {
			name: "unique key",
			inputs: &Inputs{
				Finding: &protocol.Finding{
					AlertId:     "alert1",
					Name:        "Alert 1",
					Description: "This is Alert 1",
					Protocol:    "Protocol 1",
					Type:        protocol.Finding_EXPLOIT,
					Severity:    protocol.Finding_HIGH,
					UniqueKey:   "unique",
				},
			},
			expected: "0xe518cb690acf4dca0950069e234975b17358b549906206ab3065d26e6379a810",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			hash := ForTransactionAlert(scenario.inputs)
			assert.Equal(t, scenario.expected, hash)
		})
	}
}

func TestForCombinationAlert(t *testing.T) {
	scenarios := []struct {
		name     string
		inputs   *Inputs
		expected string
	}{
		{
			name: "basic scenario",
			inputs: &Inputs{
				AlertEvent: &protocol.AlertEvent{
					Alert: &protocol.AlertEvent_Alert{AlertId: "0xID"},
				},
				Finding: &protocol.Finding{
					AlertId:     "alert1",
					Name:        "Alert 1",
					Description: "This is Alert 1",
					Protocol:    "Protocol 1",
					Type:        protocol.Finding_EXPLOIT,
					Severity:    protocol.Finding_HIGH,
				},
			},
			expected: "0xea92e7d01ecc36d3f3c3316dd8d9b1f45e88035eca9c031963df67a28ee22a78",
		}, {
			name: "unique key",
			inputs: &Inputs{
				Finding: &protocol.Finding{
					AlertId:     "alert1",
					Name:        "Alert 1",
					Description: "This is Alert 1",
					Protocol:    "Protocol 1",
					Type:        protocol.Finding_EXPLOIT,
					Severity:    protocol.Finding_HIGH,
					UniqueKey:   "unique",
				},
			},
			expected: "0xe518cb690acf4dca0950069e234975b17358b549906206ab3065d26e6379a810",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			hash := ForCombinationAlert(scenario.inputs)
			assert.Equal(t, scenario.expected, hash)
		})
	}
}
