package subject_test

import (
	"task.local/struct-transition-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	m := subject.Motor{Speed: 5, Drain: 2, Battery: 100}
	want := subject.Motor{Speed: 5, Drain: 2, Battery: 98, Distance: 5}
	if got := subject.Advance(m); got != want {
		t.Errorf("advance got %+v want %+v", got, want)
	}
}

// Add cases for the remaining requirements in TASK.md.
