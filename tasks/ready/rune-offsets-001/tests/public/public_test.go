package subject_test

import (
	"slices"
	"task.local/rune-offsets-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.RuneOffsets(""); !slices.Equal(got, []int{}) {
		t.Errorf("RuneOffsets: got %v, want %v", got, []int{})
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.RuneOffsets("abc"); !slices.Equal(got, []int{0, 1, 2}) {
		t.Errorf("RuneOffsets: got %v, want %v", got, []int{0, 1, 2})
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.RuneOffsets("Aé🔍中"); !slices.Equal(got, []int{0, 1, 3, 7}) {
		t.Errorf("RuneOffsets: got %v, want %v", got, []int{0, 1, 3, 7})
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.RuneOffsets("e\u0301"); !slices.Equal(got, []int{0, 1}) {
		t.Errorf("RuneOffsets: got %v, want %v", got, []int{0, 1})
	}
}
