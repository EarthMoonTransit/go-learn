package subject

import (
	"errors"
)

type Dictionary map[string]string

var ErrExists = errors.New("key exists")

func (d Dictionary) Add(key, value string) error {
	if v, ok := d[key]; ok && v != "" {
		return ErrExists
	}
	d[key] = value
	return nil
}
