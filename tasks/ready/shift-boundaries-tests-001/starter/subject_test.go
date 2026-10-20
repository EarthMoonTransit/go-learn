package subject_test

import (
	subject "task.local/shift-boundaries-tests-001/subject"
	"testing"
)

func TestZeroWord(t *testing.T) {
	if subject.Window(0, 0, 0) != 0 {
		t.Fatal("zero word: want zero")
	}
}
