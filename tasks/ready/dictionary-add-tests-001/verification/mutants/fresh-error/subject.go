package subject

import (
	"errors"
)

type Dictionary map[string]string

var ErrExists = errors.New("key exists")

func (d Dictionary) Add(key, value string) error {
	if _, ok := d[key]; ok {
		return errors.New("key exists")
	}
	d[key] = value
	return nil
}
