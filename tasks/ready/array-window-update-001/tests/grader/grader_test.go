package subject_test

import (
	subject "task.local/array-window-update-001/subject"
	"testing"
)

func TestWindowBoundaries(t *testing.T) {
	for lo := -2; lo <= 10; lo++ {
		for hi := -2; hi <= 10; hi++ {
			for _, v := range []int{-19, 0, 23} {
				a := [8]int{0, 1, -2, 3, -4, 5, -6, 7}
				want := a
				valid := lo >= 0 && hi >= lo && hi <= 8
				for i := range want {
					if valid && i >= lo && i < hi {
						want[i] = v
					}
				}
				got, ok := subject.PatchWindow(a, lo, hi, v)
				if got != want || ok != valid {
					t.Fatalf("bounds lo=%d hi=%d value=%d: got (%v,%t), want (%v,%t)", lo, hi, v, got, ok, want, valid)
				}
			}
		}
	}
}
