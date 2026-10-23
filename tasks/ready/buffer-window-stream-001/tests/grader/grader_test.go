package subject_test

import "testing"

func TestWindowFiniteExtensions(t *testing.T) {
	for length := 0; length <= 12; length++ {
		values := make([]int, length)
		for i := range values {
			values[i] = (i*7+length)%9 - 4
		}
		for limit := 1; limit <= 4; limit++ {
			windowCase(t, values, limit)
		}
	}
}
