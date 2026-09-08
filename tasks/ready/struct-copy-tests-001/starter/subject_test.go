package subject_test

import (
	"slices"
	"task.local/struct-copy-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	input := subject.Snapshot{Label: " x ", Count: 7, Items: []int{1, 2}}
	got := subject.CopySnapshot(input)
	if got.Label != input.Label || got.Count != 7 || !slices.Equal(got.Items, input.Items) {
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

// Add cases for the remaining requirements in TASK.md.
