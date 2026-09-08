package subject

import (
	"errors"
)

func Matches(err, target error) bool { return errors.Is(target, err) }
