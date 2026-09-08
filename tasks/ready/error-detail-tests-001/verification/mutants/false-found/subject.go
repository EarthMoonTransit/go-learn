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
	var detail *Detail
	ok := errors.As(err, &detail)
	_ = ok
	return detail, false
}
