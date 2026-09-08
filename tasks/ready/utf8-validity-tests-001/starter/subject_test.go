package subject_test

import (
	"task.local/utf8-validity-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	if got := subject.ValidUTF8("ASCII"); got != true {
		t.Errorf("ValidUTF8: got %v, want %v", got, true)
	}
}

// Add cases for the remaining requirements in TASK.md.
