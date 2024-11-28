package provider

import (
	"github.com/forta-network/forta-core-go/utils/slicering"
)

// Element is the type of element provided by the provider.
type Element interface {
	Close()
	comparable
}

// Provider keeps one or many elements provides when requested.
type Provider[T Element] interface {
	Provide() T
	Next() T
	Close()
}

// RingProvider provides elements from a slice-backed thread-safe ring.
type RingProvider[T Element] struct {
	*slicering.ThreadSafeRing[T]
}

// NewRingProvider creates a new ring provider.
func NewRingProvider[T Element](elements ...T) Provider[T] {
	return &RingProvider[T]{ThreadSafeRing: slicering.NewThreadSafeRing(elements...)}
}

// Provide provides the currently pointed element.
func (rp *RingProvider[T]) Provide() T {
	return rp.Current()
}

// Close closes all elements in the ring.
func (rp *RingProvider[T]) Close() {
	first := rp.Current()
	first.Close()
	for {
		curr := rp.Next()
		if curr == first {
			break
		}
		curr.Close()
	}
}

// StaticProvider keeps a single element and provides it when requested.
// It is a thread-unsafe alternative to use whenever there is a single element.
type StaticProvider[T Element] struct {
	element T
}

// NewStaticProvider creates a provider for a single element.
func NewStaticProvider[T Element](element T) Provider[T] {
	return &StaticProvider[T]{element: element}
}

// Provide provides the element.
func (sp *StaticProvider[T]) Provide() T {
	return sp.element
}

// Next implements the Provider interface.
func (sp *StaticProvider[T]) Next() T {
	return sp.element
}

// Close closes the element.
func (sp *StaticProvider[T]) Close() {
	sp.element.Close()
}

// Ensuring type checks below.

var _ Provider[*dummyElement] = &RingProvider[*dummyElement]{}

var _ Provider[*dummyElement] = &StaticProvider[*dummyElement]{}

type dummyElement struct{}

func (dc *dummyElement) Close() {}
