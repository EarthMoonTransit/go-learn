package subject

import "reflect"

type Resource interface {
	Close() error
}

func IsNilResource(value Resource) bool {
	if value == nil {
		return true
	}
	return reflect.ValueOf(value).IsNil()
}
