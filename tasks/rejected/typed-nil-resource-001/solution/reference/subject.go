package subject

import "reflect"

type Resource interface {
	Close() error
}

func IsNilResource(value Resource) bool {
	if value == nil {
		return true
	}
	dynamic := reflect.ValueOf(value)
	switch dynamic.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return dynamic.IsNil()
	default:
		return false
	}
}
