package subject_test

import (
	"task.local/counter-read-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	n := 2
	if got := subject.CounterValue(&n); got != 2 {
		t.Errorf("value got %d want 2", got)
	}
}

// Add cases for the remaining requirements in TASK.md.
