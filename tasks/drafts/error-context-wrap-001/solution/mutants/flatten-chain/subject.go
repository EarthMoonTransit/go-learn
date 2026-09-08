package subject

import (
	"errors"
	"fmt"
)

func Wrap(operation string, err error) error {
	if err == nil {
		return nil
	}
	if previous := errors.Unwrap(err); previous != nil {
		err = previous
	}
	return fmt.Errorf("%s: %w", operation, err)
}
