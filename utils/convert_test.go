package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHexToBigInt(t *testing.T) {
	hex := "0xabc"
	res, err := HexToBigInt(hex)
	assert.NoError(t, err)
	assert.Equal(t, int64(2748), res.Int64())
}

func TestBytesToHex(t *testing.T) {
	value := "ABCDEFG12382837581235uASDFASDFASDFzxcvzxcvzxcv"
	assert.Equal(t, "0x37353831323335754153444641534446415344467a7863767a7863767a786376", BytesToHex([]byte(value)))
}

func TestAgentBigIntToHex(t *testing.T) {
	agentID := "0x023dfceebe145aac511e103d9665163f87fa432bbc73d55d64d1b15ddd9184c9"
	res := AgentHexToBigInt(agentID)
	str := AgentBigIntToHex(res)
	assert.Equal(t, agentID, str)
}

func TestScannerBigIntToHex(t *testing.T) {
	scannerID := "0x3f88c2b3e267e6b8e9de017cdb47a59ac9ecb284"
	res := ScannerIDHexToBigInt(scannerID)
	str := ScannerIDBigIntToHex(res)
	assert.Equal(t, scannerID, str)
}

func TestIsValidBotID(t *testing.T) {
	testCases := []struct {
		name    string
		id      string
		success bool
	}{
		{
			name:    "valid",
			id:      "0x023dfceebe145aac511e103d9665163f87fa432bbc73d55d64d1b15ddd9184c9",
			success: true,
		},
		{
			name:    "too long",
			id:      "0x023dfceebe145aac511e103d9665163f87fa432bbc73d55d64d1b15ddd9184c99",
			success: false,
		},
		{
			name:    "too short",
			id:      "0x023dfceebe145aac511e103d9665163f87fa432bbc73d55d64d1b15ddd918",
			success: false,
		},
		{
			name:    "invalid prefix",
			id:      "023dfceebe145aac511e103d9665163f87fa432bbc73d55d64d1b15ddd9184c9",
			success: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			r := require.New(t)

			r.Equal(testCase.success, IsValidBotID(testCase.id))
		})
	}
}

func TestIsValidBotID1(t *testing.T) {
	type args struct {
		botID string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid bot id",
			args: args{"0xbe1872858e63b6ed4ef7b84fc453970dc8d89968715797662a4f43c01d598aab"},
			want: true,
		}, {
			name: "invalid bot id (empty)",
			args: args{""},
			want: false,
		}, {
			name: "invalid bot id (bad length)",
			args: args{"0xbe1872858e63b6ed4ef7b84fc453970dc8d8996871579766"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.Equalf(t, tt.want, IsValidBotID(tt.args.botID), "IsValidBotID(%v)", tt.args.botID)
			},
		)
	}
}
