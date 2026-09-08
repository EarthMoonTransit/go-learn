package subject_test

import (
	"task.local/dictionary-add-tests-001/subject"
	"testing"
)

func TestStartingCase01(t *testing.T) {
	d := subject.Dictionary{}
	e := d.Add("test", "this is just a test")
	if e != nil || d["test"] != "this is just a test" {
		t.Errorf("insert got %v,%v", d, e)
	}
}

// Add cases for the remaining requirements in TASK.md.
