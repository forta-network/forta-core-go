package utils

import "reflect"

// GetImplementation returns the first non-nil implementation contained as a field in a struct.
func GetImplementation[I interface{}](v interface{}) (it I, ok bool) {
	if v == nil {
		return
	}

	if it, ok := v.(I); ok {
		return it, true
	}

	t := reflect.ValueOf(v)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return
	}
	fieldCount := t.NumField()
	for i := 0; i < fieldCount; i++ {
		field := t.Field(i)
		if field.IsNil() {
			continue
		}
		it, ok = field.Interface().(I)
		if ok {
			return
		}
	}
	return
}
