package subject

import (
	"fmt"
)

func Wrap(operation string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %v", operation, err)
}

type contextualError struct {
	text  string
	cause error
}

func (e *contextualError) Error() string { return e.text }
func (e *contextualError) Unwrap() error { return e.cause }
