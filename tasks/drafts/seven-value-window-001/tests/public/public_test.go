package subject_test

import (
	"slices"
	"task.local/seven-value-window-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.WeekSum([]int{3, 0, 5, 1, 0, 4, 1, 0, 3, 4, 3, 0, 8, 0}, 1); got != 14 {
		t.Errorf("WeekSum: got %v, want %v", got, 14)
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.WeekSum([]int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}, 2); got != 12 {
		t.Errorf("WeekSum: got %v, want %v", got, 12)
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.WeekSum([]int{3, 0, 3, 3, 2, 1, 0}, 1); got != 12 {
		t.Errorf("WeekSum: got %v, want %v", got, 12)
	}
}

func TestPublicCase04(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	subject.WeekSum(input, 1)
	if !slices.Equal(input, []int{1, 2, 3, 4, 5, 6, 7}) {
		t.Errorf("input changed: %v", input)
	}
}
