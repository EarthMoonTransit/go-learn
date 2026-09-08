package subject

func Wrap(operation string, err error) error {
	if err == nil {
		return nil
	}
	return err
}
