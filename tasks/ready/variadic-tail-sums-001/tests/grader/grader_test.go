package subject_test

import (
	"slices"
	"task.local/variadic-tail-sums-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.SumTails([]int{3, 2, 1}, []int{5, 6}, []int{7, 8}); !slices.Equal(got, []int{3, 6, 8}) {
		t.Errorf("SumTails: got %v, want %v", got, []int{3, 6, 8})
	}
}
