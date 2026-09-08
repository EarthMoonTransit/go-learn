package subject_test

import (
	"task.local/constant-value-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	if subject.Rate != 0.05 {
		t.Errorf("Rate got%v want.05", subject.Rate)
	}
}

// Add cases for the remaining requirements in TASK.md.
