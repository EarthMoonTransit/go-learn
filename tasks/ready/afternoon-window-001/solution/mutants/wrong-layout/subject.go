package subject

import (
	"errors"
	"time"
)

var ErrDate = errors.New("invalid date")

func IsAfternoon(text string) (bool, error) {
	value, err := time.Parse("1/2/2006 15:04:05", text)
	if err != nil {
		return false, ErrDate
	}
	return value.Hour() >= 12 && value.Hour() < 18, nil
}
