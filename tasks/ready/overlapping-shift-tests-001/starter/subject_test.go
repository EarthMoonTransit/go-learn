package subject_test

import (
	"slices"
	"task.local/overlapping-shift-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	v := []int{1, 2, 3, 4}
	alias := v
	n := subject.ShiftRight(v, 1)
	if n != 4 || !slices.Equal(v, []int{1, 0, 2, 3}) || !slices.Equal(alias, v) {
		t.Errorf("shift got%v removed%d", v, n)
	}
}

// Add cases for the remaining requirements in TASK.md.
