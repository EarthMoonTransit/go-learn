package subject_test

import (
	"task.local/decision-table-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	if got := subject.SelectAction("ace", "ace", "jack"); got != "P" {
		t.Errorf("pair got%q wantP", got)
	}
}

// Add cases for the remaining requirements in TASK.md.
