package subject_test

import (
	"reflect"
	subject "task.local/move-overlap-range-001/subject"
	"testing"
)

func TestMoveRangeAllSmallRegions(t *testing.T) {
	for dst := 0; dst <= 6; dst++ {
		for src := 0; src <= 6; src++ {
			for n := 0; n <= 6; n++ {
				buf := []int{9, 0, -2, 6, 6, 13}
				before := append([]int(nil), buf...)
				want := append([]int(nil), buf...)
				valid := n <= 6-dst && n <= 6-src
				if valid {
					for i := 0; i < n; i++ {
						want[dst+i] = before[src+i]
					}
				}
				ok := subject.MoveRange(buf, dst, src, n)
				if ok != valid || !reflect.DeepEqual(buf, want) {
					t.Fatalf("dst=%d src=%d n=%d: got (%v,%t); want (%v,%t)", dst, src, n, buf, ok, want, valid)
				}
			}
		}
	}
}
