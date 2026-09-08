package subject_test

import (
	"slices"
	"task.local/remove-item-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	if got := subject.RemoveItem([]int{3, 2, 6, 4, 8}, 2); !slices.Equal(got, []int{3, 2, 4, 8}) {
		t.Errorf("RemoveItem: got %v, want %v", got, []int{3, 2, 4, 8})
	}
}

func TestContractCase02(t *testing.T) {
	if got := subject.RemoveItem([]int{3, 2, 6}, 0); !slices.Equal(got, []int{2, 6}) {
		t.Errorf("RemoveItem: got %v, want %v", got, []int{2, 6})
	}
}

func TestContractCase03(t *testing.T) {
	if got := subject.RemoveItem([]int{3, 2, 6}, 2); !slices.Equal(got, []int{3, 2}) {
		t.Errorf("RemoveItem: got %v, want %v", got, []int{3, 2})
	}
}

func TestContractCase04(t *testing.T) {
	if got := subject.RemoveItem([]int{3, 2, 6}, -1); !slices.Equal(got, []int{3, 2, 6}) {
		t.Errorf("RemoveItem: got %v, want %v", got, []int{3, 2, 6})
	}
}

func TestContractCase05(t *testing.T) {
	if got := subject.RemoveItem([]int{3, 2, 6}, 3); !slices.Equal(got, []int{3, 2, 6}) {
		t.Errorf("RemoveItem: got %v, want %v", got, []int{3, 2, 6})
	}
}

func TestContractCase06(t *testing.T) {
	if got := subject.RemoveItem([]int{3, 2, 6}, 10); !slices.Equal(got, []int{3, 2, 6}) {
		t.Errorf("RemoveItem: got %v, want %v", got, []int{3, 2, 6})
	}
}

func TestContractCase07(t *testing.T) {
	if got := subject.RemoveItem(nil, 0); !slices.Equal(got, []int{}) {
		t.Errorf("RemoveItem: got %v, want %v", got, []int{})
	}
}

func TestContractCase08(t *testing.T) {
	if got := subject.RemoveItem([]int{}, -1); !slices.Equal(got, []int{}) {
		t.Errorf("RemoveItem: got %v, want %v", got, []int{})
	}
}

func TestContractCase09(t *testing.T) {
	if got := subject.RemoveItem([]int{7}, 0); !slices.Equal(got, []int{}) {
		t.Errorf("RemoveItem: got %v, want %v", got, []int{})
	}
}
