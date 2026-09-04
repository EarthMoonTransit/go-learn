package subject_test

import (
	"slices"
	"testing"

	"task.local/detached-append-001/subject"
)

func TestAppendDetachedContentAndIsolation(t *testing.T) {
	base := make([]int, 2, 5)
	copy(base, []int{4, 5})
	got := subject.AppendDetached(base, 6, 7)
	if !slices.Equal(base, []int{4, 5}) {
		t.Fatalf("construction changed base: %v", base)
	}
	if !slices.Equal(got, []int{4, 5, 6, 7}) {
		t.Fatalf("got %v", got)
	}
	got[1] = 90
	if base[1] != 5 {
		t.Fatalf("result aliases base: %v", base)
	}
	base[0] = 80
	if got[0] != 4 {
		t.Fatalf("base aliases result: %v", got)
	}
}

func TestAppendDetachedNilAndNoValues(t *testing.T) {
	got := subject.AppendDetached(nil, 8)
	if !slices.Equal(got, []int{8}) {
		t.Fatalf("nil base: %v", got)
	}
	base := []int{3, 4}
	got = subject.AppendDetached(base)
	got[0] = 7
	if base[0] != 3 {
		t.Fatalf("no-value result aliases base")
	}
}
