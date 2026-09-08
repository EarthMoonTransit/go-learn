package subject

import (
	"errors"
	"time"
)

var ErrDate = errors.New("invalid date")

func IsAfternoon(text string) (bool, error) {
	value, err := time.Parse("Monday, January 2, 2006 15:04:05", text)
	if err != nil {
		return false, nil
	}
	return value.Hour() >= 12 && value.Hour() < 18, nil
}
