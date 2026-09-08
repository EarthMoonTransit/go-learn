package subject

import (
	"errors"
)

func Matches(err, target error) bool {
	if err == nil || target == nil {
		return false
	}
	return errors.Is(err, target)
}
