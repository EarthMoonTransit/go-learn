package subject

import (
	"errors"
)

type Dictionary map[string]string

var ErrMissing = errors.New("missing key")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok || value == "" {
		return "", ErrMissing
	}
	return value, nil
}
