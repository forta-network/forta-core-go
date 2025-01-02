package slicering_test

import (
	"testing"

	"github.com/forta-network/forta-core-go/utils/slicering"
	"github.com/stretchr/testify/require"
)

func TestThreadSafeRing(t *testing.T) {
	r := require.New(t)

	ring := slicering.NewThreadSafeRing(1, 2)
	r.Equal(1, ring.Current())
	r.Equal(1, ring.Current())
	r.Equal(2, ring.Next())
	r.Equal(2, ring.Current())
	r.Equal(2, ring.Current())
	r.Equal(1, ring.Next())
	r.Equal(1, ring.Current())

	ring = slicering.NewThreadSafeRing(1)
	r.Equal(1, ring.Current())
	r.Equal(1, ring.Current())
	r.Equal(1, ring.Next())
	r.Equal(1, ring.Current())

	r.Equal([]int{1}, ring.Elements())
}
