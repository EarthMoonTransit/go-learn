package subject_test

import (
	"fmt"
	"slices"
	"task.local/generic-reduce-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	in := []int{1, 2, 3}
	calls := []int{}
	got := subject.Reduce(in, func(s string, n int) string { calls = append(calls, n); return fmt.Sprintf("%s%d", s, n) }, "x")
	if got != "x123" || !slices.Equal(calls, []int{1, 2, 3}) || !slices.Equal(in, []int{1, 2, 3}) {
		t.Errorf("ordered reduce got%q calls%v input%v", got, calls, in)
	}
}

func TestPublicCase02(t *testing.T) {
	calls := 0
	got := subject.Reduce[int, string](nil, func(s string, n int) string { calls++; return "bad" }, "initial")
	if got != "initial" || calls != 0 {
		t.Errorf("empty got%q calls%d", got, calls)
	}
	v := subject.Reduce([]int{2, 3}, func(a, b int) int { return a - b }, 10)
	if v != 5 {
		t.Errorf("carry initial got%d want5", v)
	}
}

func TestPublicCase03(t *testing.T) {
	type count int
	got := subject.Reduce([]count{1, 2}, func(a []int, v count) []int { return append(a, int(v)) }, []int{9})
	if !slices.Equal(got, []int{9, 1, 2}) {
		t.Errorf("noncomparable accumulator got%v", got)
	}
}
