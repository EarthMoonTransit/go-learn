package subject_test

import (
	"slices"
	"task.local/rank-occupancy-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	if got := subject.CountRank(map[string][]bool{"A": {true, false, true, false, false, false, false, true}, "B": {false, false, false, false, true, false, false, false}}, 1); got != 1 {
		t.Errorf("CountRank: got %v, want %v", got, 1)
	}
}

func TestContractCase02(t *testing.T) {
	if got := subject.CountRank(map[string][]bool{"A": {true, true, false, false, false, false, false, false}, "B": {false, true, false, false, false, false, false, true}}, 2); got != 2 {
		t.Errorf("CountRank: got %v, want %v", got, 2)
	}
}

func TestContractCase03(t *testing.T) {
	if got := subject.CountRank(map[string][]bool{"A": {false, false, false, false, false, false, false, true}}, 8); got != 1 {
		t.Errorf("CountRank: got %v, want %v", got, 1)
	}
}

func TestContractCase04(t *testing.T) {
	if got := subject.CountRank(map[string][]bool{"A": {true, false, false, false, false, false, false, true}}, 0); got != 0 {
		t.Errorf("CountRank: got %v, want %v", got, 0)
	}
}

func TestContractCase05(t *testing.T) {
	if got := subject.CountRank(map[string][]bool{"A": {true, false, false, false, false, false, false, true}}, 9); got != 0 {
		t.Errorf("CountRank: got %v, want %v", got, 0)
	}
}

func TestContractCase06(t *testing.T) {
	if got := subject.CountRank(nil, 1); got != 0 {
		t.Errorf("CountRank: got %v, want %v", got, 0)
	}
}

func TestContractCase07(t *testing.T) {
	if got := subject.CountRank(map[string][]bool{}, 8); got != 0 {
		t.Errorf("CountRank: got %v, want %v", got, 0)
	}
}

func TestContractCase08(t *testing.T) {
	g := map[string][]bool{"A": {true, false, false, false, false, false, false, false}}
	subject.CountRank(g, 1)
	if len(g) != 1 || !slices.Equal(g["A"], []bool{true, false, false, false, false, false, false, false}) {
		t.Errorf("grid changed: %v", g)
	}
}

func TestContractCase09(t *testing.T) {
	if got := subject.CountRank(map[string][]bool{"A": {true, false, true, false, false, false, false, true}}, 4); got != 0 {
		t.Errorf("CountRank: got %v, want %v", got, 0)
	}
}
