package subject_test

import (
	"task.local/receiver-reset-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	e := subject.Entry{Name: "x"}
	e.Clear()
	if e.Name != "" {
		t.Errorf("name got %q want empty", e.Name)
	}
}
