package subject_test

import "testing"

func TestLongOwnedBatch(t *testing.T) {
	values := make([]int, 32)
	for i := range values {
		values[i] = i - 16
	}
	exercise(t, values, true)
}
