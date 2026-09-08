package subject_test

import (
	"task.local/map-entry-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	got := subject.Entries(map[string]int{"a": 1})
	if len(got) != 1 || got[0] != (subject.Entry{Key: "a", Value: 1}) {
		t.Errorf("single got%v", got)
	}
}

// Add cases for the remaining requirements in TASK.md.
