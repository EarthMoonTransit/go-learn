package subject_test

import (
	subject "task.local/recursive-total-001/subject"
	"testing"
)

func TestTotalsAndPreservation(t *testing.T) {
	cases := []struct {
		name   string
		values []int
		want   int
	}{
		{"nil", nil, 0}, {"empty", []int{}, 0}, {"single", []int{-7}, -7},
		{"minimal-split", []int{2, 9}, 11}, {"unequal", []int{4, -8, 0, 3, 12}, 11},
		{"zero", []int{0}, 0}, {"cancellation", []int{-6, 6, 0, -2, 2}, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			before := append([]int(nil), c.values...)
			got := subject.Total(c.values)
			if got != c.want {
				t.Errorf("Total(%v)=%d; want %d", before, got, c.want)
			}
			for i, v := range before {
				if c.values[i] != v {
					t.Errorf("input[%d]=%d; want unchanged %d", i, c.values[i], v)
				}
			}
		})
	}
}
