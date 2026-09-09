package subject_test

import (
	"errors"
	"task.local/writer-double-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	w := subject.NewWriter(1, errors.New("full"))
	if w == nil {
		t.Fatal("nil writer")
	}
	data, calls := w.Snapshot()
	if len(data) != 0 || calls != 0 {
		t.Errorf("new state%q calls%d", data, calls)
	}
}
