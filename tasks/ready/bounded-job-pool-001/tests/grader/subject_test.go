package subject_test

import (
	"errors"
	"testing"
)

func TestEveryAllowedPoolSize(t *testing.T) {
	marker := errors.New("bounded matrix failure")
	for workers := 1; workers <= 4; workers++ {
		for count := 0; count <= 12; count++ {
			values := make([]int, count)
			errs := make([]error, count)
			for i := range values {
				values[i] = i*7 - count
				if i%3 == 0 {
					errs[i] = marker
				}
			}
			heldPool(t, workers, values, errs)
		}
	}
}
