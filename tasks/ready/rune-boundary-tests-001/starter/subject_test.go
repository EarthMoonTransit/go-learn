package subject_test

import (
	"task.local/rune-boundary-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	if a, b := subject.RuneBoundary("abc", 1); a != 1 || b != true {
		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 1, true)
	}
}

// Add cases for the remaining requirements in TASK.md.
