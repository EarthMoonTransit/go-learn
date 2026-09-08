package subject_test

import (
	"slices"
	"task.local/variadic-tail-sums-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.SumTails([]int{1, 2}, []int{0, 9}); !slices.Equal(got, []int{2, 9}) {
		t.Errorf("SumTails: got %v, want %v", got, []int{2, 9})
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.SumTails([]int{}, []int{3, 4, 5}); !slices.Equal(got, []int{0, 9}) {
		t.Errorf("SumTails: got %v, want %v", got, []int{0, 9})
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.SumTails(nil, []int{7}); !slices.Equal(got, []int{0, 0}) {
		t.Errorf("SumTails: got %v, want %v", got, []int{0, 0})
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.SumTails(); !slices.Equal(got, []int{}) {
		t.Errorf("SumTails: got %v, want %v", got, []int{})
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.SumTails([]int{9, -2, 0, 3}); !slices.Equal(got, []int{1}) {
		t.Errorf("SumTails: got %v, want %v", got, []int{1})
	}
}

func TestPublicCase06(t *testing.T) {
	a := []int{1, 2, 3}
	subject.SumTails(a)
	if !slices.Equal(a, []int{1, 2, 3}) {
		t.Errorf("input changed: %v", a)
	}
}
