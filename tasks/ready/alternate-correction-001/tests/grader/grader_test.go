package subject_test

import (
	"slices"
	"task.local/alternate-correction-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.CorrectAlternating([]int{7, 8, 9, 10, 11}); !slices.Equal(got, []int{8, 8, 10, 10, 12}) {
		t.Errorf("CorrectAlternating: got %v, want %v", got, []int{8, 8, 10, 10, 12})
	}
}
