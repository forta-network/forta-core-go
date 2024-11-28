package slicering

import "sync"

// Ring defines the common interface for ring implementations in this package.
type Ring[T any] interface {
	Current() T
	Next() T
	Elements() []T
}

// ThreadSafeRing is a slice-backed ring protected with a mutex.
type ThreadSafeRing[T any] struct {
	elements []T
	curr     int
	mu       sync.RWMutex
}

// NewThreadSafeRing creates a new thread safe ring.
func NewThreadSafeRing[T any](elements ...T) *ThreadSafeRing[T] {
	if len(elements) == 0 {
		panic("zero elements provided to thread-safe ring")
	}
	return &ThreadSafeRing[T]{elements: elements}
}

// Current returns the current element being pointed at.
func (r *ThreadSafeRing[T]) Current() T {
	if len(r.elements) == 1 {
		return r.elements[0]
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.elements[r.curr]
}

// Next points at the next item and returns it.
func (r *ThreadSafeRing[T]) Next() T {
	if len(r.elements) == 1 {
		return r.elements[0]
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.curr += 1
	if r.curr == len(r.elements) {
		r.curr = 0
	}
	return r.elements[r.curr]
}

// Elements returns all elements.
func (r *ThreadSafeRing[T]) Elements() []T {
	return r.elements
}
