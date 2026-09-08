package subject_test

import (
	"task.local/lexical-choice-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	if got := subject.ChooseLabel("Beta", "Delta"); got != "selected: Beta" {
		t.Errorf("result got %q want %q", got, "selected: Beta")
	}
}

// Add cases for the remaining requirements in TASK.md.
