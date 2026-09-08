package subject

import (
	"errors"
)

func Matches(err, target error) bool { return err == target || errors.Unwrap(err) == target }
