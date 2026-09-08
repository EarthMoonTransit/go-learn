package subject_test

import (
	"task.local/guarded-item-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if got := subject.GetItem([]int{0}, 0); got != 0 {
		t.Errorf("GetItem: got %v, want %v", got, 0)
	}
}

func TestGraderCase02(t *testing.T) {
	if got := subject.GetItem([]int{}, -1); got != -1 {
		t.Errorf("GetItem: got %v, want %v", got, -1)
	}
}

func TestGraderCase03(t *testing.T) {
	if got := subject.GetItem([]int{2}, 99); got != -1 {
		t.Errorf("GetItem: got %v, want %v", got, -1)
	}
}
