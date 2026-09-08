package subject_test

import (
	"slices"
	"task.local/slice-sequence-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	n, _, _ := subject.Sequences()
	if !slices.Equal(n, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("numbers got%v", n)
	}
}

// Add cases for the remaining requirements in TASK.md.
