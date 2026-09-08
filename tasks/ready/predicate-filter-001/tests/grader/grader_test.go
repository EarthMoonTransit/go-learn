package subject_test

import (
	"task.local/predicate-filter-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	got := subject.Filter([]subject.Record{}, func(r subject.Record) bool { t.Error("unexpected empty predicate call"); return true })
	if len(got) != 0 {
		t.Errorf("empty got %+v", got)
	}
}
