package subject_test

import (
	"task.local/external-window-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	w := subject.NewWindow(2, 5)
	if !w.Contains(3) {
		t.Error("interior missing")
	}
}

// Add cases for the remaining requirements in TASK.md.
