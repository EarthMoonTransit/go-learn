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
	return dynamic.Kind() == reflect.Pointer && dynamic.IsNil()
}
