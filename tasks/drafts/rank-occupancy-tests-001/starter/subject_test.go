package subject_test

import (
	"task.local/rank-occupancy-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	if got := subject.CountRank(map[string][]bool{"A": {true, false, true, false, false, false, false, true}, "B": {false, false, false, false, true, false, false, false}}, 1); got != 1 {
		t.Errorf("CountRank: got %v, want %v", got, 1)
	}
}

// Add cases for the remaining requirements in TASK.md.
