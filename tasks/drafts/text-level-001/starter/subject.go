package subject

import (
	"errors"
)

type Level uint8

var ErrLevel = errors.New("invalid level")

func (l Level) MarshalText() ([]byte, error)     { return nil, nil }
func (l *Level) UnmarshalText(text []byte) error { return nil }
