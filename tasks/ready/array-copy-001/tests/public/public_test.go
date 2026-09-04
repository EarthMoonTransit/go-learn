package subject_test

import (
	"task.local/array-copy-001/subject"
	"testing"
)

func TestReplacementExamples(t *testing.T) {
	input := [3]int{8, -2, 4}
	got := subject.Replaced(input, 1, 7)
	if got != ([3]int{8, 7, 4}) {
		t.Fatalf("replacement=%v, want [8 7 4]", got)
	}
	if input != ([3]int{8, -2, 4}) {
		t.Fatal("caller value changed")
	}
	if got := subject.Replaced(input, 3, 7); got != input {
		t.Fatal("out-of-range index must preserve input")
	}
}
