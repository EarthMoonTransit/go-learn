package subject_test

import (
	"task.local/select-ready-case-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	stop := make(chan struct{}, 1)
	stop <- struct{}{}
	if n, s := subject.TryReceive(nil, stop); n != 0 || s != "canceled" {
		t.Error("readable cancel ignored")
	}
	if len(stop) != 0 {
		t.Error("cancel value not received")
	}
}
