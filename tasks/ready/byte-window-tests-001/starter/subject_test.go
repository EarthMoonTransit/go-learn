package subject_test

import (
	"task.local/byte-window-tests-001/subject"
	"testing"
)

func TestByteWindowInterior(t *testing.T) {
	got, ok := subject.ByteWindow("abcd", 1, 3)
	if !ok || got != "bc" {
		t.Fatalf("ByteWindow=(%q,%v), want (bc,true)", got, ok)
	}
}

// Add independent cases for the requirements in TASK.md.
