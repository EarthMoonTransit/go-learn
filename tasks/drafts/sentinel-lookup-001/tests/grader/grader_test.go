package subject_test

import (
	"task.local/sentinel-lookup-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	v, e := subject.Dictionary{}.Search("x")
	if v != "" || e != subject.ErrMissing {
		t.Errorf("empty dictionary got (%q,%v)", v, e)
	}
}
