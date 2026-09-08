package subject

import (
	"errors"
)

type Level uint8

var ErrLevel = errors.New("invalid level")

func (l Level) MarshalText() ([]byte, error) {
	if l > 9 {
		return nil, ErrLevel
	}
	return []byte{byte(l) + '0'}, nil
}
func (l *Level) UnmarshalText(text []byte) error {
	if len(text) != 1 || text[0] < '0' || text[0] > '9' {
		return errors.New("invalid level")
	}
	*l = Level(text[0] - '0')
	return nil
}
