package subject_test

import "testing"

func TestMaximumFiniteStream(t *testing.T) {
	values := make([]int, 32)
	for i := range values {
		values[i] = i%7 - 3
	}
	checkOpen(t, values)
}
