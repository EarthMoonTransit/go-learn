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

func TestPublicZeroExtraCase01(t *testing.T) {
	in := []int{1}
	out := subject.ReservedCopy(in, 0)
	if len(out) != 1 || cap(out) != 1 || out[0] != 1 {
		t.Fatalf("zero extra got%v cap%d", out, cap(out))
	}
	out[0] = 7
	if in[0] != 1 {
		t.Error("zero extra shares input")
	}
}
