package subject_test

import (
	subject "task.local/embedded-field-update-001/subject"
	"testing"
)

func TestEachUpdateDestination(t *testing.T) {
	for a := -2; a <= 2; a++ {
		for r := -2; r <= 2; r++ {
			for u := -2; u <= 2; u++ {
				in := subject.Layer{Counters: subject.Counters{Accepted: 19, Rejected: -13, Kept: 31}, Rejected: 7, Tag: "tag"}
				want := subject.Layer{Counters: subject.Counters{Accepted: 19 + a, Rejected: -13 + r, Kept: 31}, Rejected: 7 + u, Tag: "tag"}
				if got := subject.ApplyCounts(in, a, r, u); got != want {
					t.Fatalf("increments (%d,%d,%d): got %+v; want %+v", a, r, u, got, want)
				}
			}
		}
	}
}
