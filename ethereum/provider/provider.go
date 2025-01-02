package provider

import (
	"github.com/forta-network/forta-core-go/utils/slicering"
)

// Element is the type of element provided by the provider.
type Element interface {
	Close()
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
	for _, el := range rp.Elements() {
		el.Close()
	}
}

// Ensuring type checks below.

var _ Provider[*dummyElement] = &RingProvider[*dummyElement]{}

type dummyElement struct{}

func (dc *dummyElement) Close() {}
