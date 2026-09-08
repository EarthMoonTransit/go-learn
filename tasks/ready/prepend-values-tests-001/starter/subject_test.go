package subject_test

import (
	"slices"
	"task.local/prepend-values-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	if got := subject.PrependItems([]int{3, 2, 6}, 5); !slices.Equal(got, []int{5, 3, 2, 6}) {
		t.Errorf("PrependItems: got %v, want %v", got, []int{5, 3, 2, 6})
	}
}

// Add cases for the remaining requirements in TASK.md.
