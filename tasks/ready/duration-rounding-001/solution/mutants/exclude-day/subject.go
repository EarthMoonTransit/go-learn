package subject

import (
	"errors"
	"time"
)

var ErrDuration = errors.New("invalid duration")

func WholeMinutes(value time.Duration) (int64, error) {
	if value < 0 || value >= 24*time.Hour {
		return 0, ErrDuration
	}
	return int64(value / time.Minute), nil
}
