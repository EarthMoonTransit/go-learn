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

type contextualError struct {
	text  string
	cause error
}

func (e *contextualError) Error() string { return e.text }
func (e *contextualError) Unwrap() error { return e.cause }
