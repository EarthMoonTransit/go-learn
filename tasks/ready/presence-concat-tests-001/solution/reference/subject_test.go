package subject_test

import (
	subject "task.local/presence-concat-tests-001/subject"
	"testing"
)

func TestConcatenatedPresence(t *testing.T) {
	cases := []struct {
		name  string
		parts [][]int
		want  []int
	}{
		{"no-parts", nil, nil}, {"all-absent", [][]int{nil, nil}, nil},
		{"one-present-empty", [][]int{[]int{}}, []int{}},
		{"empty-after-nil", [][]int{nil, []int{}, nil}, []int{}},
		{"empty-before-nil", [][]int{[]int{}, nil}, []int{}},
		{"later-data", [][]int{nil, []int{4, 0, -7}, nil}, []int{4, 0, -7}},
		{"ordered-parts", [][]int{[]int{5, -2}, nil, []int{}, []int{0, 5}, []int{9}}, []int{5, -2, 0, 5, 9}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			before := make([][]int, len(c.parts))
			for i, p := range c.parts {
				before[i] = append([]int(nil), p...)
			}
			got := subject.ConcatPresent(c.parts...)
			if (got == nil) != (c.want == nil) || len(got) != len(c.want) {
				t.Fatalf("ConcatPresent: got %v nil=%t; want %v nil=%t", got, got == nil, c.want, c.want == nil)
			}
			for i, v := range c.want {
				if got[i] != v {
					t.Errorf("result[%d]=%d; want %d", i, got[i], v)
				}
			}
			for p, part := range c.parts {
				for i, v := range part {
					if v != before[p][i] {
						t.Errorf("input part %d element %d changed to %d; want %d", p, i, v, before[p][i])
					}
				}
			}
		})
	}
}
