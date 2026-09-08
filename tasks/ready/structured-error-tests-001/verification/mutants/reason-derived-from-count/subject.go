package subject

import (
	"fmt"
)

type InvalidCountError struct {
	Count  int
	Reason string
}

func (e *InvalidCountError) Error() string {
	reason := "negative count"
	if e.Count == 0 {
		reason = "empty count"
	}
	return fmt.Sprintf("%d items are invalid: %s", e.Count, reason)
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
