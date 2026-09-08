package subject

import (
	"errors"
)

type Dictionary map[string]string

var ErrExists = errors.New("key exists")

func (d Dictionary) Add(key, value string) error { d[key] = value; return nil }
