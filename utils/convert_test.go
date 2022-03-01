package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
