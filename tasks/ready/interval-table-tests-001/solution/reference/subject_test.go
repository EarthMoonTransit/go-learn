package subject_test

import (
	subject "task.local/interval-table-tests-001/subject"
	"testing"
)

func TestNamedPartitions(t *testing.T) {
	cases := []struct {
		name      string
		lo, hi, x int
		want      bool
	}{
		{"below", 2, 6, 1, false}, {"lower", 2, 6, 2, true}, {"interior", 2, 6, 4, true},
		{"upper", 2, 6, 6, false}, {"above", 2, 6, 7, false}, {"empty", 3, 3, 3, false},
		{"reversed", 6, 2, 4, false}, {"negative-interior", -6, -2, -4, true},
		{"negative-upper", -6, -2, -2, false},
	}
	for _, c := range cases {
		got := subject.Contains(c.lo, c.hi, c.x)
		if got != c.want {
			t.Errorf("case %s: Contains(%d,%d,%d)=%t; want %t", c.name, c.lo, c.hi, c.x, got, c.want)
		}
	}
}
