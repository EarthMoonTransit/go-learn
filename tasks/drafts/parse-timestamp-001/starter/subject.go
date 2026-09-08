package subject

import (
	"errors"
	"time"
)

var ErrDate = errors.New("invalid date")

func ParseTimestamp(text string) (time.Time, error) { return time.Time{}, nil }
