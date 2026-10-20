package subject_test

import (
	subject "task.local/array-window-update-001/subject"
	"testing"
)

func TestPatchWindow(t *testing.T) {
	base := [8]int{2, -3, 5, 7, 11, 0, -13, 17}
	cases := []struct {
		name          string
		lo, hi, value int
		want          [8]int
		ok            bool
	}{
		{"middle", 2, 5, 9, [8]int{2, -3, 9, 9, 9, 0, -13, 17}, true},
		{"first", 0, 1, -4, [8]int{-4, -3, 5, 7, 11, 0, -13, 17}, true},
		{"last", 7, 8, 0, [8]int{2, -3, 5, 7, 11, 0, -13, 0}, true},
		{"all", 0, 8, 6, [8]int{6, 6, 6, 6, 6, 6, 6, 6}, true},
		{"empty-middle", 3, 3, 99, base, true}, {"empty-end", 8, 8, 99, base, true},
		{"negative", -1, 3, 99, base, false}, {"reversed", 5, 2, 99, base, false},
		{"past-end", 2, 9, 99, base, false}, {"both-past", 9, 9, 99, base, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, ok := subject.PatchWindow(base, c.lo, c.hi, c.value)
			if got != c.want || ok != c.ok {
				t.Fatalf("PatchWindow(%v,%d,%d,%d) = (%v,%t); want (%v,%t)", base, c.lo, c.hi, c.value, got, ok, c.want, c.ok)
			}
		})
	}
}
