package subject_test

import (
	"task.local/category-closure-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	f := subject.InCategory("a")
	if f == nil {
		t.Fatal("nil predicate")
	}
	if !f(subject.Record{Category: "a"}) || f(subject.Record{Category: "b"}) {
		t.Error("predicate did not match exact category")
	}
}

// Add cases for the remaining requirements in TASK.md.
