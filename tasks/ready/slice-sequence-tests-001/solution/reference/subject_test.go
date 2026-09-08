package subject_test

import (
	"slices"
	"task.local/slice-sequence-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	n, _, _ := subject.Sequences()
	if !slices.Equal(n, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		t.Errorf("numbers got%v", n)
	}
}

func TestContractCase02(t *testing.T) {
	n, w, v := subject.Sequences()
	m, x, y := subject.Sequences()
	if !slices.Equal(w, []string{"a", "b", "c", "d", "e"}) || !slices.Equal(v, []string{"b", "c"}) || !slices.Equal(x, w) || !slices.Equal(y, v) || len(n) != 10 || len(m) != 10 {
		t.Fatalf("sequences n%v w%v v%v", n, w, v)
	}
	v[0] = "changed"
	if w[1] != "changed" || w[0] != "a" || x[1] != "b" || y[0] != "b" {
		t.Error("subslice alias or independence incorrect")
	}
	w[2] = "back"
	if v[1] != "back" {
		t.Error("reverse alias missing")
	}
	n[0] = 99
	if m[0] != 0 {
		t.Error("number storage shared")
	}
}
