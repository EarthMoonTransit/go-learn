package subject_test

import (
	"task.local/encoded-width-001/subject"
	"testing"
)

func TestDispatchExamples(t *testing.T) {
	for _, c := range []struct {
		tag string
		w   int
		ok  bool
	}{{"byte", 1, true}, {"word", 2, true}, {"", 0, false}, {"BYTE", 0, false}} {
		got, ok := subject.Width(c.tag)
		if got != c.w || ok != c.ok {
			t.Errorf("Width(%q)=(%d,%v), want (%d,%v)", c.tag, got, ok, c.w, c.ok)
		}
	}
}
