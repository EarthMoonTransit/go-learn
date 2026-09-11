package subject_test

import (
	"slices"
	"task.local/bounded-slice-view-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	in := []int{1, 2, 3, 4}
	before := slices.Clone(in)
	v := subject.CappedView(in, 1, 3)
	if !slices.Equal(in, before) {
		t.Errorf("call changed input got%v want%v", in, before)
	}
	if !slices.Equal(v, []int{2, 3}) || len(v) != 2 || cap(v) != 2 {
		t.Fatalf("view got%v len%d cap%d", v, len(v), cap(v))
	}
	v[0] = 9
	if in[1] != 9 {
		t.Errorf("view detached: %v", in)
	}
	v = append(v, 8)
	if in[3] != 4 {
		t.Errorf("append overwrote outside: %v", in)
	}
}

func TestPublicCase02(t *testing.T) {
	for _, bounds := range [][2]int{{0, 0}, {2, 2}, {3, 3}} {
		in := []int{1, 2, 3}
		v := subject.CappedView(in, bounds[0], bounds[1])
		if len(v) != 0 || cap(v) != 0 || !slices.Equal(in, []int{1, 2, 3}) {
			t.Errorf("empty %v len%d cap%d input%v", bounds, len(v), cap(v), in)
		}
	}
}

func TestPublicNilView(t *testing.T) {
	got := subject.CappedView(nil, 0, 0)
	if len(got) != 0 || cap(got) != 0 {
		t.Errorf("CappedView(nil, 0, 0): len=%d cap=%d; want len=0 cap=0", len(got), cap(got))
	}
}
