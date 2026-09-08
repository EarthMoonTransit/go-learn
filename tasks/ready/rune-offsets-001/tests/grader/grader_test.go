package subject_test

import (
	"slices"
	"task.local/rune-offsets-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.RuneOffsets("🔍aé"); !slices.Equal(got, []int{0, 4, 5}) {
		t.Errorf("RuneOffsets: got %v, want %v", got, []int{0, 4, 5})
	}
}

func TestGraderCase02(t *testing.T) {
	if got := subject.RuneOffsets("\x00中!"); !slices.Equal(got, []int{0, 1, 4}) {
		t.Errorf("RuneOffsets: got %v, want %v", got, []int{0, 1, 4})
	}
}
