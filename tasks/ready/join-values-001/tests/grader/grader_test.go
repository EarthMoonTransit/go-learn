package subject_test

import (
	"task.local/join-values-001/subject"
	"testing"
)

func TestListConstructionBoundaries(t *testing.T) {
	cases := []struct {
		prefix string
		values []int
		want   string
	}{
		{"", nil, "[]"}, {"p", []int{}, "p[]"}, {"", []int{0}, "[0]"},
		{"%s", []int{-19}, "%s[-19]"}, {"[", []int{9, 2, 9}, "[[9,2,9]"},
		{"\x00\xff", []int{-120, 5, 0}, "\x00\xff[-120,5,0]"},
		{" x ", []int{10, 20, 30, 40}, " x [10,20,30,40]"},
	}
	for _, c := range cases {
		before := append([]int(nil), c.values...)
		got := subject.FormatValues(c.prefix, c.values)
		if got != c.want {
			t.Errorf("list=%q, want %q", got, c.want)
		}
		for i, want := range before {
			if c.values[i] != want {
				t.Errorf("input index %d changed", i)
			}
		}
	}
}
