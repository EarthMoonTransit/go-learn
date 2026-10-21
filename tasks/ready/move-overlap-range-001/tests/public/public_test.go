package subject_test

import (
	"reflect"
	subject "task.local/move-overlap-range-001/subject"
	"testing"
)

func TestMoveRangeSnapshots(t *testing.T) {
	cases := []struct {
		name        string
		dst, src, n int
		want        []int
	}{
		{"right-overlap", 2, 0, 4, []int{3, -5, 3, -5, 7, 11}},
		{"left-overlap", 0, 2, 4, []int{7, 11, 13, 17, 13, 17}},
		{"separate", 4, 0, 2, []int{3, -5, 7, 11, 3, -5}},
		{"same", 1, 1, 4, []int{3, -5, 7, 11, 13, 17}},
		{"one", 5, 0, 1, []int{3, -5, 7, 11, 13, 3}},
		{"empty-at-end", 6, 6, 0, []int{3, -5, 7, 11, 13, 17}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			backing := []int{3, -5, 7, 11, 13, 17, 41, 43}
			buf := backing[:6]
			alias := buf
			ok := subject.MoveRange(buf, c.dst, c.src, c.n)
			if !ok || !reflect.DeepEqual(alias, c.want) || !reflect.DeepEqual(backing[6:], []int{41, 43}) {
				t.Fatalf("MoveRange(%d,%d,%d): ok=%t values=%v spare=%v; want true %v [41 43]", c.dst, c.src, c.n, ok, alias, backing[6:], c.want)
			}
		})
	}
}
func TestMoveRangeInvalidAtomicity(t *testing.T) {
	max := int(^uint(0) >> 1)
	for _, c := range [][3]int{{-1, 0, 1}, {0, -1, 1}, {0, 0, -1}, {1, 0, 6}, {0, 2, 5}, {6, 0, 1}, {0, 6, 1}, {7, 0, 0}, {0, 7, 0}, {1, 2, max}, {max, 0, 1}, {0, max, 1}} {
		storage := []int{2, 3, 5, 7, 11, 13, 17, 19}
		before := append([]int(nil), storage...)
		ok := subject.MoveRange(storage[:6], c[0], c[1], c[2])
		if ok || !reflect.DeepEqual(storage, before) {
			t.Fatalf("invalid range %v: ok=%t backing=%v; want false and %v", c, ok, storage, before)
		}
	}
	if !subject.MoveRange(nil, 0, 0, 0) {
		t.Fatal("nil storage empty range: got false; want true")
	}
}
