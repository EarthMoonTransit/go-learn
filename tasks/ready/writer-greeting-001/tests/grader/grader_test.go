package subject_test

import (
	"io"
	"task.local/writer-greeting-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	w := &writerProbe{n: 0}
	if e := subject.WriteGreeting(w, ""); e != io.ErrShortWrite {
		t.Errorf("zero write got%v wantShortWrite", e)
	}
}
