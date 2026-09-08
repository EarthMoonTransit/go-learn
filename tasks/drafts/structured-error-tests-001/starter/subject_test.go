package subject_test

import (
	"task.local/structured-error-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	if e := subject.ValidateCount(2); e != nil {
		t.Errorf("positive got %v wantnil", e)
	}
}

// Add cases for the remaining requirements in TASK.md.
