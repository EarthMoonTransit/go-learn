package subject_test

import (
	"slices"
	"task.local/detached-prefix-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	in := []int{2, 3, 4, 5}
	out := subject.DetachedPrefix(in, 2)
	if !slices.Equal(out, []int{2, 3}) || cap(out) != 2 || !slices.Equal(in, []int{2, 3, 4, 5}) {
		t.Fatalf("prefix got%v cap%d input%v", out, cap(out), in)
	}
	out[0] = 9
	if in[0] != 2 {
		t.Error("output shares input")
	}
	in[1] = 8
	if out[1] != 3 {
		t.Error("input shares output")
	}
}

func TestPublicCase02(t *testing.T) {
	in := make([]int, 2, 100)
	out := subject.DetachedPrefix(in, 0)
	if len(out) != 0 || cap(out) != 0 {
		t.Errorf("empty retained cap%d", cap(out))
	}
}
