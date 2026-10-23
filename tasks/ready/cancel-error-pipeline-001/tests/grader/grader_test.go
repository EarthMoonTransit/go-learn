package subject_test

import "testing"

func TestPipelineFiniteExtensions(t *testing.T) {
	for length := 0; length <= 12; length++ {
		values := make([]int, length)
		for i := range values {
			values[i] = (length+5*i)%11 - 5
		}
		pipelineCase(t, values, -1)
		if length > 0 {
			pipelineCase(t, values, length/2)
		}
	}
}
