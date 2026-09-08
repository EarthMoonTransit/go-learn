package subject_test

import (
	"task.local/dictionary-update-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	d := subject.Dictionary{"k": "v"}
	if e := d.Update("k", "v"); e != nil || len(d) != 1 || d["k"] != "v" {
		t.Errorf("same value got %v,%v", d, e)
	}
}
