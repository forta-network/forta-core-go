package provider_test

import (
	"testing"

	"github.com/forta-network/forta-core-go/ethereum/provider"
	"github.com/stretchr/testify/require"
)

type testElement struct {
	closed bool
}

func (te *testElement) Close() {
	te.closed = true
}

func TestRingProvider(t *testing.T) {
	r := require.New(t)

	el1 := &testElement{}
	el2 := &testElement{}

	p := provider.NewRingProvider(el1, el2)
	r.Equal(el1, p.Provide())
	r.Equal(el1, p.Provide())
	r.Equal(el2, p.Next())
	r.Equal(el2, p.Provide())
	r.Equal(el1, p.Next())

	r.False(el1.closed)
	r.False(el2.closed)

	p.Close()

	r.True(el1.closed)
	r.True(el2.closed)
}
