package subject_test

import (
	"testing"
)

func TestMaximumInput(t *testing.T) {
	values := make([]int, 16)
	for i := range values {
		values[i] = i - 8
	}
	checkStream(t, values, 5)
}
