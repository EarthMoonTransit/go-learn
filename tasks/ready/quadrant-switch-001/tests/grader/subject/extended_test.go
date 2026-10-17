package subject_test

import (
	"task.local/quadrant-switch-001/subject"
	"testing"
)

func TestCoordinateBoundaries(t *testing.T) {
	max := int(^uint(0) >> 1)
	values := []int{-max - 1, -1000000, -1, 0, 1, 1000000, max}
	names := [3][3]string{
		{"quadrant-iii", "negative-x", "quadrant-ii"},
		{"negative-y", "origin", "positive-y"},
		{"quadrant-iv", "positive-x", "quadrant-i"},
	}
	for _, x := range values {
		for _, y := range values {
			ix, iy := 1, 1
			if x < 0 {
				ix = 0
			} else if x > 0 {
				ix = 2
			}
			if y < 0 {
				iy = 0
			} else if y > 0 {
				iy = 2
			}
			want := names[ix][iy]
			if got := subject.Locate(x, y); got != want {
				t.Errorf("Locate(%d, %d) = %q, want %q", x, y, got, want)
			}
		}
	}
}
