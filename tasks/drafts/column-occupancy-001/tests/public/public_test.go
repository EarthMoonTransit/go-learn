package subject_test

import (
	"slices"
	"task.local/column-occupancy-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.CountColumn(map[string][]bool{"A": {true, false, true, false, false, false, false, true}}, "A"); got != 3 {
		t.Errorf("CountColumn: got %v, want %v", got, 3)
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.CountColumn(map[string][]bool{"A": {true}, "B": {true, true}}, "A"); got != 1 {
		t.Errorf("CountColumn: got %v, want %v", got, 1)
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.CountColumn(map[string][]bool{"A": {true}}, "Z"); got != 0 {
		t.Errorf("CountColumn: got %v, want %v", got, 0)
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.CountColumn(nil, "A"); got != 0 {
		t.Errorf("CountColumn: got %v, want %v", got, 0)
	}
}

func TestPublicCase05(t *testing.T) {
	if got := subject.CountColumn(map[string][]bool{"A": {}}, "A"); got != 0 {
		t.Errorf("CountColumn: got %v, want %v", got, 0)
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.CountColumn(map[string][]bool{"A": {false, false}}, "A"); got != 0 {
		t.Errorf("CountColumn: got %v, want %v", got, 0)
	}
}

func TestPublicCase07(t *testing.T) {
	g := map[string][]bool{"A": {true, false, true}}
	subject.CountColumn(g, "A")
	if len(g) != 1 || !slices.Equal(g["A"], []bool{true, false, true}) {
		t.Errorf("grid changed: got %v want A:[true false true]", g)
	}
}
