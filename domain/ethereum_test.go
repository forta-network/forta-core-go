package domain

import (
	"testing"

	"github.com/forta-network/forta-core-go/protocol"
	"github.com/stretchr/testify/assert"
)

// 2022-02-28T14:32:46Z
var testTSSeconds int64 = 1646058766
var testTSHex = "0x621cdd0e"

func TestBlock_GetTimestamp(t *testing.T) {
	b := Block{Timestamp: testTSHex}
	ts, err := b.GetTimestamp()
	assert.NoError(t, err)
	assert.Equal(t, testTSSeconds, ts.Unix())
}

func TestTrace_ToProto(t *testing.T) {
	trace := Trace{
		Action: TraceAction{
			CallType: safeToPointer("call"),
			To:       safeToPointer("0x1234567890"),
			Input:    safeToPointer("input data"),
			// Add other fields as necessary
		},
		BlockHash:   safeToPointer("block hash"),
		BlockNumber: safeToPointer(123),
		Result: &TraceResult{
			GasUsed: safeToPointer("gas used"),
			Address: safeToPointer("address"),
			Code:    safeToPointer("code"),
			Output:  safeToPointer("output"),
			// Add other fields as necessary
		},
		Subtraces:           2,
		TraceAddress:        []int{1, 2, 3},
		TransactionHash:     safeToPointer("transaction hash"),
		TransactionPosition: safeToPointer(456),
		Type:                "type",
		Error:               safeToPointer("error"),
	}

	expectedProto := &protocol.TransactionEvent_Trace{
		Action: &protocol.TransactionEvent_TraceAction{
			CallType: "call",
			To:       "0x1234567890",
			Input:    "input data",
			// Add other fields as necessary
		},
		BlockHash:   "block hash",
		BlockNumber: 123,
		Result: &protocol.TransactionEvent_TraceResult{
			GasUsed: "gas used",
			Address: "address",
			Code:    "code",
			Output:  "output",
			// Add other fields as necessary
		},
		Subtraces:           int64(2),
		TraceAddress:        []int64{1, 2, 3},
		TransactionHash:     "transaction hash",
		TransactionPosition: 456,
		Type:                "type",
		Error:               "error",
	}

	actualProto := trace.ToProto()

	assert.Equal(t, expectedProto, actualProto)
}

func TestTransaction_ToProto(t *testing.T) {
	transaction := Transaction{
		BlockHash:            "block hash",
		BlockNumber:          "block number",
		From:                 "from address",
		Gas:                  "gas",
		GasPrice:             "gas price",
		Hash:                 "transaction hash",
		Input:                safeToPointer("input data"),
		Nonce:                "nonce",
		To:                   safeToPointer("to address"),
		TransactionIndex:     "transaction index",
		Value:                safeToPointer("value"),
		V:                    "v",
		R:                    "r",
		S:                    "s",
		MaxFeePerGas:         safeToPointer("max fee per gas"),
		MaxPriorityFeePerGas: safeToPointer("max priority fee per gas"),
	}

	expectedProto := &protocol.TransactionEvent_EthTransaction{
		Type:                 "",
		Nonce:                "nonce",
		GasPrice:             "gas price",
		Gas:                  "gas",
		Value:                "value",
		Input:                "input data",
		V:                    "v",
		R:                    "r",
		S:                    "s",
		To:                   "to address",
		Hash:                 "transaction hash",
		From:                 "from address",
		MaxFeePerGas:         "max fee per gas",
		MaxPriorityFeePerGas: "max priority fee per gas",
	}

	actualProto := transaction.ToProto()

	assert.Equal(t, expectedProto, actualProto)
}

// Helper functions to create pointers to values
func safeToPointer[T any](t T) *T {
	return &t
}
