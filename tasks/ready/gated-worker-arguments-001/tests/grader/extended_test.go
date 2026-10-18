package subject_test

import (
	"testing"
	"testing/synctest"
)

func TestLongRepeatedBatch(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		values := make([]int, 23)
		for i := range values {
			values[i] = i%7 - 3
		}
		p19Exercise(t, values)
	})
}
func TestExtremeArguments(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		max := int(^uint(0) >> 1)
		p19Exercise(t, []int{max, -max - 1, 0, max, 1})
	})
}
