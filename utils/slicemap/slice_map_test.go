package slicemap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceMap(t *testing.T) {
	r := require.New(t)

	var m SliceMap[string, int]

	inserted := m.Set("a", 2)
	r.True(inserted)
	inserted = m.Set("a", 3)
	r.False(inserted)
	inserted = m.Set("abc", 4)
	r.True(inserted)

	r.Len(m, 2)

	v, ok := m.Get("a")
	r.True(ok)
	r.Equal(2, v)

	v, ok = m.Get("abc")
	r.True(ok)
	r.Equal(4, v)

	deleted := m.Delete("abc")
	r.True(deleted)

	r.Len(m, 1)

	v, ok = m.Get("a")
	r.True(ok)
	r.Equal(2, v)

	deleted = m.Delete("a")
	r.True(deleted)

	r.Len(m, 0)
}
