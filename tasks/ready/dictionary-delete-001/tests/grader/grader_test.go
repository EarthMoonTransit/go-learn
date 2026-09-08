package subject_test

import (
	"task.local/dictionary-delete-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	d := subject.Dictionary{"k": "v"}
	_ = d.Delete("k")
	if e := d.Delete("k"); e != subject.ErrAbsent {
		t.Errorf("repeated got%v wantErrAbsent", e)
	}
}
