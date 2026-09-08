package subject

import (
	"fmt"
)

func Capture(work func() error) (err error, panicked bool) {
	defer func() {
		if value := recover(); value != nil {
			err = fmt.Errorf("%v", value)
			panicked = true
		}
	}()
	return work(), false
}
