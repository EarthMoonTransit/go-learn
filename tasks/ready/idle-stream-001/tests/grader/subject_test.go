package subject_test

import (
	"testing"
)

func TestFullBoundAndNegativeSequence(t *testing.T) {
	checkClosedInput(t, []int{-4, 2, 0, -4, 7, 1, 8, 1, 3, 0, 2, 6, -1, 5, 9, 4})
}
