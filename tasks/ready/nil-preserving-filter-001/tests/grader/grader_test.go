package subject_test

import (
	subject "task.local/nil-preserving-filter-001/subject"
	"testing"
)

func TestFilterRepresentations(t *testing.T) {
	for pattern := 0; pattern < 81; pattern++ {
		in := make([]int, 4)
		want := []int{}
		n := pattern
		for i := range in {
			in[i] = n%3 - 1
			n /= 3
			if in[i] >= 0 {
				want = append(want, in[i])
			}
		}
		before := append([]int(nil), in...)
		got := subject.KeepNonnegative(in)
		if got == nil || len(got) != len(want) {
			t.Fatalf("input %v: got %v (nil=%t), want present %v", before, got, got == nil, want)
		}
		for i := range want {
			if got[i] != want[i] {
				t.Fatalf("input %v: got %v, want %v", before, got, want)
			}
		}
		for i := range in {
			if in[i] != before[i] {
				t.Fatalf("input modified: got %v want %v", in, before)
			}
		}
	}
	backing := [3]int{5, -7, 9}
	empty := backing[1:1]
	got := subject.KeepNonnegative(empty)
	if got == nil || len(got) != 0 {
		t.Fatalf("empty subslice: got %v nil=%t; want non-nil empty result", got, got == nil)
	}
}
