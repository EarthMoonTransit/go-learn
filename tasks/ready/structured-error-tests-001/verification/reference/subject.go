package subject

import (
	"fmt"
)

type InvalidCountError struct {
	Count  int
	Reason string
}

func (e *InvalidCountError) Error() string {
	return fmt.Sprintf("%d items are invalid: %s", e.Count, e.Reason)
}
func ValidateCount(n int) error {
	if n < 0 {
		return &InvalidCountError{Count: n, Reason: "negative count"}
	}
	if n == 0 {
		return &InvalidCountError{Count: n, Reason: "empty count"}
	}
	return nil
}
