package subject_test

import (
	"task.local/range-offset-001/subject"
	"testing"
)

func TestRangeAndWindowBoundaries(t *testing.T) {
	subject.OffsetByIndex(nil)
	for n := 0; n <= 64; n++ {
		base := make([]int, n+2)
		base[0] = 71
		base[n+1] = 72
		for i := 0; i < n; i++ {
			base[i+1] = (i*97)%2001 - 1000
		}
		original := append([]int(nil), base...)
		window := base[1 : n+1]
		alias := window
		subject.OffsetByIndex(window)
		if len(window) != n || base[0] != 71 || base[n+1] != 72 {
			t.Fatal("slice window boundaries changed")
		}
		for i := range window {
			if alias[i] != original[i+1]+i+1 {
				t.Fatalf("length %d: element %d has incorrect update", n, i)
			}
		}
	}
}
