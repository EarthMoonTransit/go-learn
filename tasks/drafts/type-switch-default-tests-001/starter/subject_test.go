package subject_test

import (
	"task.local/type-switch-default-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	if got := subject.TypeLabel(3); got != "int" {
		t.Errorf("int got%q", got)
	}
}

// Add cases for the remaining requirements in TASK.md.
