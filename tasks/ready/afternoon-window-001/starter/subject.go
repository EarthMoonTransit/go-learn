package subject

import (
	"errors"
)

var ErrDate = errors.New("invalid date")

func IsAfternoon(text string) (bool, error) { return false, nil }
