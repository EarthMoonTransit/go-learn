package subject_test

import (
	subject "task.local/nil-preserving-filter-001/subject"
	"testing"
)

func TestSlicePresence(t *testing.T) {
	cases := []struct {
		name     string
		in, want []int
	}{
		{"absent", nil, nil}, {"present-empty", []int{}, []int{}},
		{"no-matches", []int{-3, -1}, []int{}}, {"zero", []int{0}, []int{0}},
		{"mixed", []int{-2, 7, 0, -5, 7, 3}, []int{7, 0, 7, 3}},
		{"already-valid", []int{4, 0, 9}, []int{4, 0, 9}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			before := append([]int(nil), c.in...)
			got := subject.KeepNonnegative(c.in)
			if (got == nil) != (c.want == nil) {
				t.Fatalf("input %v (nil=%t): result nil=%t; want nil=%t", c.in, c.in == nil, got == nil, c.want == nil)
			}
			if len(got) != len(c.want) {
				t.Fatalf("input %v: got %v; want %v", before, got, c.want)
			}
			for i := range got {
				if got[i] != c.want[i] {
					t.Fatalf("result[%d]=%d; want %d (result %v)", i, got[i], c.want[i], got)
				}
			}
			for i := range before {
				if c.in[i] != before[i] {
					t.Fatalf("input[%d] changed to %d; want %d", i, c.in[i], before[i])
				}
			}
		})
	}
}
