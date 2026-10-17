package subject_test

import (
	"slices"
	"task.local/prepend-values-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	// Individual prefix arguments follow the original values slice.
	if got := subject.PrependItems([]int{3, 2, 6}, 5, 8); !slices.Equal(got, []int{5, 8, 3, 2, 6}) {
		t.Errorf("PrependItems: got %v, want %v", got, []int{5, 8, 3, 2, 6})
	}
}

func TestPrependExpandedPrefix(t *testing.T) {
	prefix := []int{9, 4}
	// prefix... supplies each slice element as a separate variadic argument.
	if got := subject.PrependItems([]int{2}, prefix...); !slices.Equal(got, []int{9, 4, 2}) {
		t.Errorf("PrependItems expanded prefix: got %v, want %v", got, []int{9, 4, 2})
	}
}

// Add cases for the remaining requirements in TASK.md.
