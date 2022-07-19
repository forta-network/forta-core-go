package domain

import (
	"testing"

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
