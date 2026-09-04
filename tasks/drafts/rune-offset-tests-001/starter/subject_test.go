package subject_test

import (
	"slices"
	"testing"

	"task.local/rune-offset-tests-001/subject"
)

func TestRuneOffsetsASCII(t *testing.T) {
	if got, want := subject.RuneOffsets("go"), []int{0, 1, 2}; !slices.Equal(got, want) {
		t.Fatalf("RuneOffsets returned %v; want %v", got, want)
	}
}
