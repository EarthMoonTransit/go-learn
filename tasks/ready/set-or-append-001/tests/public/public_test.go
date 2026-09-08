package subject_test

import (
	"slices"
	"task.local/set-or-append-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.SetItem([]int{1, 2, 4, 1}, 2, 6); !slices.Equal(got, []int{1, 2, 6, 1}) {
		t.Errorf("SetItem: got %v, want %v", got, []int{1, 2, 6, 1})
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.SetItem([]int{1, 2, 4, 1}, -1, 6); !slices.Equal(got, []int{1, 2, 4, 1, 6}) {
		t.Errorf("SetItem: got %v, want %v", got, []int{1, 2, 4, 1, 6})
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.SetItem([]int{1, 2}, 2, 7); !slices.Equal(got, []int{1, 2, 7}) {
		t.Errorf("SetItem: got %v, want %v", got, []int{1, 2, 7})
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.SetItem(nil, 0, 7); !slices.Equal(got, []int{7}) {
		t.Errorf("SetItem: got %v, want %v", got, []int{7})
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.SetItem([]int{1}, 0, 9); !slices.Equal(got, []int{9}) {
		t.Errorf("SetItem: got %v, want %v", got, []int{9})
	}
}

func TestPublicCase06(t *testing.T) {
	input := []int{1, 2, 3}
	view := input[1:]
	subject.SetItem(input, 1, 9)
	if view[0] != 9 {
		t.Errorf("shared element got %d want 9", view[0])
	}
}
