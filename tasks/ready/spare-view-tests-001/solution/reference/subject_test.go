package subject_test

import (
	"reflect"
	subject "task.local/spare-view-tests-001/subject"
	"testing"
)

func TestSpareViewBoundaries(t *testing.T) {
	cases := []struct {
		name                   string
		size, low, high, spare int
	}{
		{"interior", 5, 1, 3, 2}, {"zero-spare", 5, 1, 3, 0}, {"hidden-capacity", 3, 1, 3, 3},
		{"empty-interior", 5, 2, 2, 2}, {"empty-start", 0, 0, 0, 3}, {"empty-end", 5, 5, 5, 2},
		{"all-remaining", 5, 1, 5, 3}, {"exact-end", 8, 2, 8, 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			storage := []int{3, -5, 7, 0, 11, 13, -17, 19}
			before := append([]int(nil), storage...)
			in := storage[:c.size]
			got := subject.SpareView(in, c.low, c.high, c.spare)
			wantLen := c.high - c.low
			if len(got) != wantLen || cap(got) != wantLen+c.spare {
				t.Fatalf("len/cap=%d/%d; want %d/%d", len(got), cap(got), wantLen, wantLen+c.spare)
			}
			if !reflect.DeepEqual(storage, before) {
				t.Fatalf("call changed backing=%v; want %v", storage, before)
			}
			for i, v := range before[c.low:c.high] {
				if got[i] != v {
					t.Fatalf("element %d=%d; want %d", i, got[i], v)
				}
			}
			if len(got) > 0 {
				got[0] = 101
				if storage[c.low] != 101 {
					t.Fatalf("view write not shared: backing[%d]=%d; want 101", c.low, storage[c.low])
				}
				storage[c.high-1] = -103
				if got[len(got)-1] != -103 {
					t.Fatalf("input write not shared: got %d; want -103", got[len(got)-1])
				}
			}
			for i := 0; i < c.spare; i++ {
				got = append(got, 211+i)
				if storage[c.high+i] != 211+i {
					t.Fatalf("append within spare %d not reused: got backing %d; want %d", i, storage[c.high+i], 211+i)
				}
			}
			atLimit := append([]int(nil), storage...)
			got = append(got, 307)
			if !reflect.DeepEqual(storage, atLimit) {
				t.Fatalf("append beyond limit changed backing=%v; want %v", storage, atLimit)
			}
			for i := range got {
				got[i] = -401 - i
			}
			if !reflect.DeepEqual(storage, atLimit) {
				t.Fatalf("grown result still shares input: %v; want %v", storage, atLimit)
			}
			if len(in) != c.size || cap(in) != 8 {
				t.Fatalf("caller header changed: len/cap=%d/%d; want %d/8", len(in), cap(in), c.size)
			}
		})
	}
}
func TestSpareViewNilAndEmpty(t *testing.T) {
	for _, in := range [][]int{nil, make([]int, 0)} {
		got := subject.SpareView(in, 0, 0, 0)
		if len(got) != 0 || cap(got) != 0 {
			t.Fatalf("empty len/cap=%d/%d; want 0/0", len(got), cap(got))
		}
	}
}
