package slicemap

type element[K comparable, V any] struct {
	key K
	val V
}

// SliceMap is a simple map implementation using slice. It is an option for cases
// when persistence and collision issues can have critical results and the size
// is not a concern.
type SliceMap[K comparable, V any] []*element[K, V]

// Get gets an element.
func (sm SliceMap[K, V]) Get(key K) (val V, ok bool) {
	for _, el := range sm {
		if el.key == key {
			return el.val, true
		}
	}
	return
}

// Set sets an element.
func (sm *SliceMap[K, V]) Set(key K, val V) bool {
	for _, el := range *sm {
		if el.key == key {
			return false
		}
	}
	*sm = append(*sm, &element[K, V]{key: key, val: val})
	return true
}

// Delete deletes an element.
func (sm *SliceMap[K, V]) Delete(key K) bool {
	for i, el := range *sm {
		if el.key == key {
			*sm = append((*sm)[:i], (*sm)[i+1:]...)
			return true
		}
	}
	return false
}
