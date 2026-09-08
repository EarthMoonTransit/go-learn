package subject_test

import (
	"task.local/grouped-unit-cost-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.BatchCost(0); got != uint(0) {
		t.Errorf("BatchCost: got %v, want %v", got, uint(0))
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.BatchCost(1); got != uint(10000) {
		t.Errorf("BatchCost: got %v, want %v", got, uint(10000))
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.BatchCost(9); got != uint(90000) {
		t.Errorf("BatchCost: got %v, want %v", got, uint(90000))
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.BatchCost(10); got != uint(95000) {
		t.Errorf("BatchCost: got %v, want %v", got, uint(95000))
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.BatchCost(21); got != uint(200000) {
		t.Errorf("BatchCost: got %v, want %v", got, uint(200000))
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.BatchCost(37); got != uint(355000) {
		t.Errorf("BatchCost: got %v, want %v", got, uint(355000))
	}
}

func TestPublicCase07(t *testing.T) {
	if got := subject.BatchCost(100); got != uint(950000) {
		t.Errorf("BatchCost: got %v, want %v", got, uint(950000))
	}
}
