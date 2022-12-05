package parallel

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func strArrs(count int) ([]string, []*string) {
	var strs []string
	var ptrs []*string
	for i := 0; i < count; i++ {
		s := fmt.Sprintf("%d", i)
		ptrs = append(ptrs, &s)
		strs = append(strs, s)
	}
	sort.Strings(strs)
	return strs, ptrs
}

func TestForEach(t *testing.T) {
	ctx := context.Background()
	mux := sync.Mutex{}
	var found []string
	expected, ptrs := strArrs(100)
	err := ForEach[string](ctx, 10, ptrs, func(ctx context.Context, item *string) error {
		mux.Lock()
		found = append(found, *item)
		mux.Unlock()
		return nil
	})
	assert.NoError(t, err)
	sort.Strings(found)
	assert.Equal(t, expected, found)
}

func TestForEachString(t *testing.T) {
	ctx := context.Background()
	var found []string
	mux := sync.Mutex{}
	items, _ := strArrs(100)
	err := ForEachString(ctx, 10, items, func(ctx context.Context, item string) error {
		mux.Lock()
		found = append(found, item)
		mux.Unlock()
		return nil
	})
	assert.NoError(t, err)
	sort.Strings(found)
	assert.Equal(t, items, found)
}
