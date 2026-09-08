package subject_test

import (
	"slices"
	"task.local/set-or-append-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.SetItem([]int{4, 5}, 10, 6); !slices.Equal(got, []int{4, 5, 6}) {
		t.Errorf("SetItem: got %v, want %v", got, []int{4, 5, 6})
	}
}

func TestGraderCase02(t *testing.T) {
	if got := subject.SetItem([]int{}, -2, 3); !slices.Equal(got, []int{3}) {
		t.Errorf("SetItem: got %v, want %v", got, []int{3})
	}
}

func TestGraderCase03(t *testing.T) {
	if got := subject.SetItem([]int{4, 5}, 1, 6); !slices.Equal(got, []int{4, 6}) {
		t.Errorf("SetItem: got %v, want %v", got, []int{4, 6})
	}
}
