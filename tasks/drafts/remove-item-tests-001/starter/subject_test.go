package subject_test

import (
	"slices"
	"task.local/remove-item-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	if got := subject.RemoveItem([]int{3, 2, 6, 4, 8}, 2); !slices.Equal(got, []int{3, 2, 4, 8}) {
		t.Errorf("RemoveItem: got %v, want %v", got, []int{3, 2, 4, 8})
	}
}

// Add cases for the remaining requirements in TASK.md.
