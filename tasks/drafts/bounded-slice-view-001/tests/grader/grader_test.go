package subject_test

import (
	"task.local/bounded-slice-view-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	if v := subject.CappedView(nil, 0, 0); len(v) != 0 || cap(v) != 0 {
		t.Errorf("nil got%v cap%d", v, cap(v))
	}
}
