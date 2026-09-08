package subject

import (
	"errors"
	"time"
)

var ErrDate = errors.New("invalid date")

func ParseTimestamp(text string) (time.Time, error) {
	value, err := time.Parse("1/2/2006 15:04:05", text)
	if err != nil {
		return time.Time{}, ErrDate
	}
	return value, nil
}
