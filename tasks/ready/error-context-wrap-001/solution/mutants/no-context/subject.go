package subject

func Wrap(operation string, err error) error {
	if err == nil {
		return nil
	}
	return err
}

type contextualError struct {
	text  string
	cause error
}

func (e *contextualError) Error() string { return e.text }
func (e *contextualError) Unwrap() error { return e.cause }
