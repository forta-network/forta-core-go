package feeds

import (
	"time"

	"github.com/forta-network/forta-core-go/clients/graphql"
)

type Iterator interface {
	HasNext() bool
	LowerBound() int64
	UpperBound() int64
	Next() int64
	Head() int64
	SetHead(head int64)
	End() int64
}

var combinerInterval = graphql.DefaultLastNMinutes.Milliseconds()

type combinerIterator struct {
	start    int64
	end      int64
	head     int64
	stepSize int64
}

func (i *combinerIterator) SetHead(head int64) {
	i.head = head
}

func (i *combinerIterator) End() int64 {
	return i.end
}

func (i *combinerIterator) Head() int64 {
	return i.head
}

func (i *combinerIterator) HasNext() bool {
	return i.head <= i.end || i.end == 0
}

func (i *combinerIterator) LowerBound() (lowerBound int64) {
	lowerBound = time.Now().UnixMilli() - i.head
	if lowerBound < 0 {
		lowerBound = 0
	}

	return
}

func (i *combinerIterator) UpperBound() (upperBound int64) {
	now := time.Now().UnixMilli()

	// if the head is too close to end, get the biggest possible upper bound:
	// head: 12:10
	// end: 12:17
	// upper bound should be 7 minutes instead of 10 minutes (default interval)
	//
	// if there is no end specified, no need to account the end value
	if i.end == 0 || i.end > i.head+combinerInterval {
		upperBound = now - i.head - interval
	} else {
		upperBound = now - i.end
	}

	if upperBound < 0 {
		upperBound = 0
	}

	return
}

func (i *combinerIterator) Next() int64 {
	i.head += i.stepSize
	return i.head
}

func NewCombinerIterator(start, end, stepSize int64) Iterator {
	return &combinerIterator{start: start, end: end, stepSize: stepSize, head: start}
}
