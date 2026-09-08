package subject_test

import (
	"task.local/sentinel-lookup-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	sentinel := subject.ErrMissing
	if sentinel == nil {
		t.Fatal("ErrMissing: got nil want original sentinel")
	}
	defer func() {
		if subject.ErrMissing != sentinel {
			t.Error("ErrMissing changed during call; want original supplied sentinel")
		}
	}()

	v, e := subject.Dictionary{}.Search("x")
	if v != "" || e != sentinel {
		t.Errorf("empty dictionary got (%q,%v)", v, e)
	}
}
