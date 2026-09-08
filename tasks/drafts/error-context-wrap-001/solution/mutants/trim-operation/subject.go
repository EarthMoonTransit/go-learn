package subject

import (
	"fmt"
	"strings"
)

func Wrap(operation string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", strings.TrimSpace(operation), err)
}
