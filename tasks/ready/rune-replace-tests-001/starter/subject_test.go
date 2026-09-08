package subject_test

import (
	"task.local/rune-replace-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	if got := subject.Replace("❗ recommended product", '❗', '?'); got != "? recommended product" {
		t.Errorf("Replace: got %q, want %q", got, "? recommended product")
	}
}

// Add cases for the remaining requirements in TASK.md.
