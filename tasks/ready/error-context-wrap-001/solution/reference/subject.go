package subject

type contextualError struct {
	text  string
	cause error
}

func (e *contextualError) Error() string { return e.text }
func (e *contextualError) Unwrap() error { return e.cause }
func Wrap(operation string, err error) error {
	if err == nil {
		return nil
	}
	return &contextualError{text: operation + ": " + err.Error(), cause: err}
}
