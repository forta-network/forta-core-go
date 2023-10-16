package ethutils

import (
	"errors"
	"reflect"

	"github.com/forta-network/forta-core-go/utils"
)

// EventIterator iterates over events of given type.
type EventIterator[S any] interface {
	Value() (*S, bool)
	abigenIterator
}

type abigenIterator interface {
	Next() bool
	Error() error
	Close() error
}

// NewEventIterator creates a new event iterator by finding and wrapping the concrete
// abigen-generated iterator found in given struct.
func NewEventIterator[S any](v any) (EventIterator[S], error) {
	it, ok := utils.GetImplementation[abigenIterator](v)
	if !ok {
		return nil, errors.New("iterator not found")
	}
	return &eventIterator[S]{abigenIterator: it}, nil
}

type eventIterator[S any] struct {
	abigenIterator
}

func (it *eventIterator[S]) Value() (*S, bool) {
	itVal := reflect.ValueOf(it.abigenIterator)
	event := itVal.Elem().FieldByName("Event").Elem()

	var ret S
	retVal := reflect.ValueOf(&ret).Elem()
	for i := 0; i < event.NumField(); i++ {
		field := event.Field(i)
		fieldName := event.Type().Field(i).Name
		if _, ok := retVal.Type().FieldByName(fieldName); !ok {
			continue
		}
		retVal.FieldByName(fieldName).Set(field)
	}

	return &ret, true
}
