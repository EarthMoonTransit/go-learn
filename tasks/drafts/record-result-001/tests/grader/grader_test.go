package subject_test

import (
	"task.local/record-result-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	r, e := subject.NewRecord(" a ", 0)
	if e != nil || r == nil || *r != (subject.Record{Label: " a "}) {
		t.Errorf("exact label got%v,%v", r, e)
	}
}
