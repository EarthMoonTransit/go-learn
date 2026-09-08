package subject

import (
	"errors"
)

func Matches(err, target error) bool { return errors.Is(err, target) }
