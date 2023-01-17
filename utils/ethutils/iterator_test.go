package ethutils

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

type testKnownEvent struct {
	A int
	B *big.Int
	D float64
}

type testActualEvent struct {
	A int
	B *big.Int
	C string
}

type testAbigenIterator[E any] struct {
	Event *E

	returnEvents   []*E
	currEventIndex int
}

func (it *testAbigenIterator[E]) Next() bool {
	if it.currEventIndex >= len(it.returnEvents) {
		return false
	}
	it.Event = it.returnEvents[it.currEventIndex]
	it.currEventIndex++
	return true
}

func (it *testAbigenIterator[E]) Error() error {
	return nil
}

func (it *testAbigenIterator[E]) Close() error {
	return nil
}

type testIteratorResult[E any] struct {
	Iterator *testAbigenIterator[E]
}

func TestEventIterator(t *testing.T) {
	r := require.New(t)

	itResult := testIteratorResult[testActualEvent]{
		Iterator: &testAbigenIterator[testActualEvent]{
			returnEvents: []*testActualEvent{
				{
					A: 123,
					B: big.NewInt(123),
					C: "123",
				},
				{
					A: 456,
					B: big.NewInt(456),
					C: "456",
				},
			},
		},
	}

	it, err := NewEventIterator[testKnownEvent](itResult)
	r.NoError(err)

	expectedEvents := []*testKnownEvent{
		{
			A: 123,
			B: big.NewInt(123),
			D: 0,
		},
		{
			A: 456,
			B: big.NewInt(456),
			D: 0,
		},
	}
	var receivedEvents []*testKnownEvent
	for it.Next() {
		v, ok := it.Value()
		if !ok {
			break
		}
		receivedEvents = append(receivedEvents, v)
	}

	r.Len(receivedEvents, len(expectedEvents))
	for i, receivedEv := range receivedEvents {
		expectedEv := expectedEvents[i]
		r.Equal(expectedEv.A, receivedEv.A)
		r.Equal(expectedEv.D, receivedEv.D)
		r.Equal(0, expectedEv.B.Cmp(receivedEv.B)) // equal *big.Int
	}
}
