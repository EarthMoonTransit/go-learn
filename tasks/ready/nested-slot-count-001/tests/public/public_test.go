package subject_test

import (
	"slices"
	"task.local/nested-slot-count-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	if got := subject.CountSlots(map[string][]bool{"A": {true, false, true}, "B": {false, false}}); got != 5 {
		t.Errorf("CountSlots: got %v, want %v", got, 5)
	}
}

func TestPublicCase02(t *testing.T) {
	if got := subject.CountSlots(map[string][]bool{"A": {false, false, false}}); got != 3 {
		t.Errorf("CountSlots: got %v, want %v", got, 3)
	}
}

func TestPublicCase03(t *testing.T) {
	if got := subject.CountSlots(nil); got != 0 {
		t.Errorf("CountSlots: got %v, want %v", got, 0)
	}
}

func TestPublicCase04(t *testing.T) {
	if got := subject.CountSlots(map[string][]bool{"A": nil, "B": {}}); got != 0 {
		t.Errorf("CountSlots: got %v, want %v", got, 0)
	}
}

func TestPublicCase05(t *testing.T) {
	g := map[string][]bool{"A": {true, false}}
	subject.CountSlots(g)
	if len(g) != 1 || !slices.Equal(g["A"], []bool{true, false}) {
		t.Errorf("grid changed: %v", g)
	}
}

func TestPublicCase06(t *testing.T) {
	if got := subject.CountSlots(map[string][]bool{"A": {true}, "B": {true, true}, "C": {true, true, true}}); got != 6 {
		t.Errorf("CountSlots: got %v, want %v", got, 6)
	}
}
