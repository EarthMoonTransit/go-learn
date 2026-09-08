package subject_test

import (
	"slices"
	"task.local/generic-find-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	v, ok := subject.Find([]int{1, 2, 4}, func(x int) bool { return x%2 == 0 })
	if v != 2 || !ok {
		t.Errorf("first even got%d,%v", v, ok)
	}
}

func TestContractCase02(t *testing.T) {
	visited := []int{}
	in := []int{3, 0, 5}
	v, ok := subject.Find(in, func(x int) bool { visited = append(visited, x); return x == 0 })
	if v != 0 || !ok || !slices.Equal(visited, []int{3, 0}) || !slices.Equal(in, []int{3, 0, 5}) {
		t.Errorf("zero/visits/preservation got%d,%v %v %v", v, ok, visited, in)
	}
}

func TestContractCase03(t *testing.T) {
	calls := 0
	v, ok := subject.Find[string](nil, func(x string) bool { calls++; return true })
	if v != "" || ok || calls != 0 {
		t.Errorf("empty got%q,%v calls%d", v, ok, calls)
	}
	v, ok = subject.Find([]string{"a", "b"}, func(string) bool { return false })
	if v != "" || ok {
		t.Errorf("no match got%q,%v", v, ok)
	}
}

func TestContractCase04(t *testing.T) {
	type row struct{ Data []int }
	v, ok := subject.Find([]row{{[]int{1}}, {[]int{2, 3}}}, func(x row) bool { return len(x.Data) == 2 })
	if !ok || !slices.Equal(v.Data, []int{2, 3}) {
		t.Errorf("noncomparable got%v,%v", v, ok)
	}
}

func TestMatchingString(t *testing.T) {
	v, ok := subject.Find([]string{"a", "abc", "later"}, func(s string) bool { return len(s) > 1 })
	if v != "abc" || !ok {
		t.Errorf("string match got%q,%v wantabc,true", v, ok)
	}
}
