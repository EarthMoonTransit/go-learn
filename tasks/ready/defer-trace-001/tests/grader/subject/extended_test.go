package subject_test

import (
	"slices"
	"task.local/defer-trace-001/subject"
	"testing"
)

func TestTraceRangeAndIndependence(t *testing.T) {
	for seed := -128; seed <= 128; seed++ {
		want := []int{seed + 9, subject.Checkpoint, seed + 4, seed}
		if got := subject.Trace(seed); !slices.Equal(got, want) {
			t.Fatalf("Trace(%d) = %v, want %v", seed, got, want)
		}
	}
	first := subject.Trace(7)
	subject.Trace(-8)
	want := []int{16, subject.Checkpoint, 11, 7}
	if !slices.Equal(first, want) {
		t.Errorf("earlier trace changed = %v, want %v", first, want)
	}
}
