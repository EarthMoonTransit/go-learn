package subject

type FieldError struct {
	Field string
	Err   error
}

func (err *FieldError) Error() string {
	return err.Field + ": " + err.Err.Error()
}

func (err *FieldError) Unwrap() error {
	return nil
}

func WrapField(field string, err error) error {
	return &FieldError{Field: field, Err: err}
}

func HasCause(err, target error) bool {
	return err != nil && target != nil && err.Error() == target.Error()
}

func FindField(err error) (*FieldError, bool) {
	fieldErr, ok := err.(*FieldError)
	return fieldErr, ok
}
