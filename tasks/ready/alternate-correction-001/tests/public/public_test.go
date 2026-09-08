package subject_test

import (
	"slices"
	"task.local/alternate-correction-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.CorrectAlternating([]int{2, 5, 0, 7, 4, 1}); !slices.Equal(got, []int{3, 5, 1, 7, 5, 1}) {
		t.Errorf("CorrectAlternating: got %v, want %v", got, []int{3, 5, 1, 7, 5, 1})
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.CorrectAlternating([]int{4}); !slices.Equal(got, []int{5}) {
		t.Errorf("CorrectAlternating: got %v, want %v", got, []int{5})
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.CorrectAlternating([]int{0, 0, 0}); !slices.Equal(got, []int{1, 0, 1}) {
		t.Errorf("CorrectAlternating: got %v, want %v", got, []int{1, 0, 1})
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.CorrectAlternating(nil); !slices.Equal(got, []int{}) {
		t.Errorf("CorrectAlternating: got %v, want %v", got, []int{})
	}
}

func TestPublicCase05(t *testing.T) {
	input := []int{2, 3, 4}
	subject.CorrectAlternating(input)
	if !slices.Equal(input, []int{3, 3, 5}) {
		t.Errorf("input got %v want [3 3 5]", input)
	}
}
