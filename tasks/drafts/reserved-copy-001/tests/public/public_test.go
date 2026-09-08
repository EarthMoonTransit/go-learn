package subject_test

import (
	"slices"
	"task.local/reserved-copy-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	in := []int{3, -2, 0}
	out := subject.ReservedCopy(in, 4)
	if !slices.Equal(out, in) || len(out) != 3 || cap(out) != 7 {
		t.Fatalf("copy got%v len%d cap%d", out, len(out), cap(out))
	}
	out[0] = 9
	out = append(out, 8)
	if !slices.Equal(in, []int{3, -2, 0}) {
		t.Errorf("input changed %v", in)
	}
}

func TestPublicCase02(t *testing.T) {
	for _, extra := range []int{0, 3} {
		out := subject.ReservedCopy(nil, extra)
		if len(out) != 0 || cap(out) != extra {
			t.Errorf("empty extra%d len%d cap%d", extra, len(out), cap(out))
		}
	}
}
