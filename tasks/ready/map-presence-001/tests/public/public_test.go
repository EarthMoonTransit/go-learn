package subject_test

import (
	"task.local/map-presence-001/subject"
	"testing"
)

func TestPresenceExamples(t *testing.T) {
	input := map[string]int{"zero": 0, "other": 7}
	if got := subject.LookupOr(input, "zero", 9); got != 0 {
		t.Fatalf("stored zero replaced with %d", got)
	}
	if got := subject.LookupOr(input, "absent", 9); got != 9 {
		t.Fatalf("missing key=%d, want fallback 9", got)
	}
}
