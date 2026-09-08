package subject_test

import (
	"slices"
	"task.local/guarded-item-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.GetItem([]int{5, 2, 10, 6, 8, 7, 0, 9}, 4); got != 8 {
		t.Errorf("GetItem: got %v, want %v", got, 8)
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.GetItem([]int{5, 2, 9}, 0); got != 5 {
		t.Errorf("GetItem: got %v, want %v", got, 5)
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.GetItem([]int{5, 2, 9}, 2); got != 9 {
		t.Errorf("GetItem: got %v, want %v", got, 9)
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.GetItem([]int{5, 2, 9}, -1); got != -1 {
		t.Errorf("GetItem: got %v, want %v", got, -1)
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.GetItem([]int{5, 2, 9}, 3); got != -1 {
		t.Errorf("GetItem: got %v, want %v", got, -1)
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.GetItem(nil, 0); got != -1 {
		t.Errorf("GetItem: got %v, want %v", got, -1)
	}
}

func TestPublicCase07(t *testing.T) {
	input := []int{4, 7}
	subject.GetItem(input, 1)
	if !slices.Equal(input, []int{4, 7}) {
		t.Errorf("input changed: %v", input)
	}
}
