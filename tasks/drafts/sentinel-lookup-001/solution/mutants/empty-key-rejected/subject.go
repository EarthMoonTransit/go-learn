package subject

import (
	"errors"
)

type Dictionary map[string]string

var ErrMissing = errors.New("missing key")

func (d Dictionary) Search(key string) (string, error) {
	if key == "" {
		return "", ErrMissing
	}
	value, ok := d[key]
	if !ok {
		return "", ErrMissing
	}
	return value, nil
}
