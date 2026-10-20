package subject_test

import (
	subject "task.local/scatter-snapshot-001/subject"
	"testing"
)

func TestGatherAllPositions(t *testing.T) {
	source := [8]int{-29, 0, 5, 18, -3, 22, 41, 8}
	for pos := 0; pos < 8; pos++ {
		for index := -2; index < 10; index++ {
			indexes := [8]int{7, 6, 5, 4, 3, 2, 1, 0}
			indexes[pos] = index
			want := source
			ok := index >= 0 && index < 8
			if ok {
				for i, j := range indexes {
					want[i] = source[j]
				}
			}
			got, valid := subject.Gather(source, indexes)
			if got != want || valid != ok {
				t.Fatalf("position=%d index=%d: got (%v,%t); want (%v,%t)", pos, index, got, valid, want, ok)
			}
		}
	}
}
