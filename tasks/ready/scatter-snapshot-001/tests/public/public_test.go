package subject_test

import (
	subject "task.local/scatter-snapshot-001/subject"
	"testing"
)

func TestGatherSnapshot(t *testing.T) {
	source := [8]int{3, -5, 7, 0, 11, 13, -17, 19}
	cases := []struct {
		name          string
		indexes, want [8]int
		ok            bool
	}{
		{"identity", [8]int{0, 1, 2, 3, 4, 5, 6, 7}, source, true},
		{"cycle", [8]int{1, 2, 0, 4, 5, 6, 7, 3}, [8]int{-5, 7, 3, 11, 13, -17, 19, 0}, true},
		{"repeats", [8]int{7, 7, 0, 0, 3, 3, 7, 0}, [8]int{19, 19, 3, 3, 0, 0, 19, 3}, true},
		{"negative-first", [8]int{-1, 1, 2, 3, 4, 5, 6, 7}, source, false},
		{"invalid-last", [8]int{7, 6, 5, 4, 3, 2, 1, 8}, source, false},
		{"negative-last", [8]int{2, 2, 2, 2, 2, 2, 2, -1}, source, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, ok := subject.Gather(source, c.indexes)
			if got != c.want || ok != c.ok {
				t.Fatalf("Gather(%v,%v)=(%v,%t); want (%v,%t)", source, c.indexes, got, ok, c.want, c.ok)
			}
		})
	}
}
