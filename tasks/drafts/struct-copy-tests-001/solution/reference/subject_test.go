package subject_test

import (
	"slices"
	"task.local/struct-copy-tests-001/subject"
	"testing"
)

func TestContractCase01(t *testing.T) {
	input := subject.Snapshot{Label: " x ", Count: 7, Items: []int{1, 2}}
	got := subject.CopySnapshot(input)
	if got.Label != input.Label || got.Count != 7 || !slices.Equal(got.Items, []int{1, 2}) || !slices.Equal(input.Items, []int{1, 2}) {
		t.Fatalf("copy got%+v", got)
	}
	got.Label = "other"
	got.Count = 9
	if input.Label != " x " || input.Count != 7 {
		t.Error("scalar values changed")
	}
	got.Items[0] = 8
	if input.Items[0] != 8 {
		t.Error("slice backing storage not shared")
	}
	input.Items[1] = 6
	if got.Items[1] != 6 {
		t.Error("reverse sharing missing")
	}
}

func TestContractCase02(t *testing.T) {
	for _, items := range [][]int{nil, make([]int, 0), make([]int, 0, 4), make([]int, 2, 5)} {
		got := subject.CopySnapshot(subject.Snapshot{Items: items})
		if (got.Items == nil) != (items == nil) || len(got.Items) != len(items) || cap(got.Items) != cap(items) {
			t.Errorf("slice header got nil%v len%d cap%d want nil%v len%d cap%d", got.Items == nil, len(got.Items), cap(got.Items), items == nil, len(items), cap(items))
		}
	}
}
