package subject

import (
	"errors"
)

type Detail struct {
	Code    int
	Message string
}

func (e *Detail) Error() string { return e.Message }
func FindDetail(err error) (*Detail, bool) {
	if d, ok := err.(*Detail); ok {
		return d, true
	}
	if tree, ok := err.(interface{ Unwrap() []error }); ok {
		for _, child := range tree.Unwrap() {
			if d, ok := FindDetail(child); ok {
				return d, true
			}
		}
		return nil, false
	}
	if child := errors.Unwrap(err); child != nil {
		return FindDetail(child)
	}
	return nil, false
}
