package subject

import (
	"fmt"
)

func Wrap(operation string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %v", operation, err)
}
