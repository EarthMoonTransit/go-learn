package subject_test

import (
	"task.local/total-counts-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	if got := subject.TotalCounts([]int{9, 0, 8, 4, 5, 1, 3}); got != 30 {
		t.Errorf("TotalCounts: got %v, want %v", got, 30)
	}
}

// Add cases for the remaining requirements in TASK.md.
