package subject

import "errors"

type FieldError struct {
	Field string
	Err   error
}

func (err *FieldError) Error() string {
	return err.Field + ": " + err.Err.Error()
}
func (err *FieldError) Unwrap() error {
	return err.Err
}
func WrapField(field string, err error) error {
	return &FieldError{Field: field, Err: err}
}
func HasCause(err, target error) bool {
	return errors.Is(err, target)
}
func FindField(err error) (*FieldError, bool) {
	var fieldErr *FieldError
	ok := errors.As(err, &fieldErr)
	return fieldErr, ok
}
