package subject_test

import (
	"fmt"
	"testing"
)

func TestGraderScopeMatrix(t *testing.T) {
	for count := 0; count <= 4; count++ {
		for errorAt := -1; errorAt < count; errorAt++ {
			for _, canceled := range []bool{false, true} {
				name := fmt.Sprintf("n%d-error%d-canceled%v", count, errorAt, canceled)
				t.Run(name, func(t *testing.T) {
					values := make([]int, count)
					for index := range values {
						values[index] = (index - 2) * 37
					}
					checkScopeSequence(t, values, errorAt, canceled, true)
				})
			}
		}
	}
}
