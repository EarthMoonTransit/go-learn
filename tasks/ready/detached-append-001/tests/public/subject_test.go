package subject_test

import (
	"slices"
	"testing"

	"task.local/detached-append-001/subject"
)

func TestAppendDetachedDoesNotAliasBase(t *testing.T) {
	base := make([]int, 2, 4)
	copy(base, []int{1, 2})
	got := subject.AppendDetached(base, 3)
	if !slices.Equal(got, []int{1, 2, 3}) {
		t.Fatalf("got %v", got)
	}
	got[0] = 99
	if base[0] != 1 {
		t.Fatalf("changing result changed base to %v", base)
	}
}
