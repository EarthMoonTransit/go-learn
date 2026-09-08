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
	if tree, ok := err.(interface{ Unwrap() []error }); ok {
		children := tree.Unwrap()
		for i := len(children) - 1; i >= 0; i-- {
			var d *Detail
			if errors.As(children[i], &d) {
				return d, true
			}
		}
		return nil, false
	}
	var d *Detail
	ok := errors.As(err, &d)
	return d, ok
}
