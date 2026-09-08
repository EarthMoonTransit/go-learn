package subject_test

import (
	"slices"
	"task.local/prepend-values-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	if got := subject.PrependItems([]int{3, 2, 6}, 5); !slices.Equal(got, []int{5, 3, 2, 6}) {
		t.Errorf("PrependItems: got %v, want %v", got, []int{5, 3, 2, 6})
	}
}

func TestContractCase02(t *testing.T) {
	if got := subject.PrependItems([]int{3, 2, 6, 4, 8}, 5, 1); !slices.Equal(got, []int{5, 1, 3, 2, 6, 4, 8}) {
		t.Errorf("PrependItems: got %v, want %v", got, []int{5, 1, 3, 2, 6, 4, 8})
	}
}

func TestContractCase03(t *testing.T) {
	if got := subject.PrependItems([]int{3, 2, 6}); !slices.Equal(got, []int{3, 2, 6}) {
		t.Errorf("PrependItems: got %v, want %v", got, []int{3, 2, 6})
	}
}

func TestContractCase04(t *testing.T) {
	if got := subject.PrependItems(nil, 5, 1); !slices.Equal(got, []int{5, 1}) {
		t.Errorf("PrependItems: got %v, want %v", got, []int{5, 1})
	}
}

func TestContractCase05(t *testing.T) {
	if got := subject.PrependItems(nil); !slices.Equal(got, []int{}) {
		t.Errorf("PrependItems: got %v, want %v", got, []int{})
	}
}

func TestContractCase06(t *testing.T) {
	if got := subject.PrependItems([]int{}, 2); !slices.Equal(got, []int{2}) {
		t.Errorf("PrependItems: got %v, want %v", got, []int{2})
	}
}
