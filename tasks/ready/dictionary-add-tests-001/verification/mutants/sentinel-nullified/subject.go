package subject

import (
	"errors"
)

type Dictionary map[string]string

var ErrExists = errors.New("key exists")

func (d Dictionary) Add(key, value string) error {
	ErrExists = nil
	if _, ok := d[key]; ok {
		return ErrExists
	}
	d[key] = value
	return nil
}
