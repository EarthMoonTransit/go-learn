package subject

import (
	"errors"
)

type Dictionary map[string]string

var ErrMissing = errors.New("missing key")

func (d Dictionary) Search(key string) (string, error) {
	_, ok := d[key]
	if !ok {
		return "", ErrMissing
	}
	return "", nil
}
