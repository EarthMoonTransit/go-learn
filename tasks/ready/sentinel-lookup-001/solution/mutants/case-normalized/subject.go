package subject

import (
	"errors"
	"strings"
)

type Dictionary map[string]string

var ErrMissing = errors.New("missing key")

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[strings.ToUpper(key)]
	if !ok {
		return "", ErrMissing
	}
	return value, nil
}
