package subject_test

import (
	"slices"
	"task.local/total-counts-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	if got := subject.TotalCounts([]int{9, 0, 8, 4, 5, 1, 3}); got != 30 {
		t.Errorf("TotalCounts: got %v, want %v", got, 30)
	}
}

func TestContractCase02(t *testing.T) {
	if got := subject.TotalCounts(nil); got != 0 {
		t.Errorf("TotalCounts: got %v, want %v", got, 0)
	}
}

func TestContractCase03(t *testing.T) {
	if got := subject.TotalCounts([]int{}); got != 0 {
		t.Errorf("TotalCounts: got %v, want %v", got, 0)
	}
}

func TestContractCase04(t *testing.T) {
	if got := subject.TotalCounts([]int{2}); got != 2 {
		t.Errorf("TotalCounts: got %v, want %v", got, 2)
	}
}

func TestContractCase05(t *testing.T) {
	if got := subject.TotalCounts([]int{4, 0, 7}); got != 11 {
		t.Errorf("TotalCounts: got %v, want %v", got, 11)
	}
}

func TestContractCase06(t *testing.T) {
	input := []int{3, 1, 4}
	subject.TotalCounts(input)
	if !slices.Equal(input, []int{3, 1, 4}) {
		t.Errorf("input changed: got %v want [3 1 4]", input)
	}
}
