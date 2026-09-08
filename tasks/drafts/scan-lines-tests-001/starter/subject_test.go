package subject_test

import (
	"slices"
	"strings"
	"task.local/scan-lines-tests-001/subject"
	"testing"
)

type lineReader struct {
	data   []byte
	err    error
	closes int
}

func (r *lineReader) Read(p []byte) (int, error) {
	n := copy(p, r.data)
	r.data = r.data[n:]
	if len(r.data) == 0 {
		return n, r.err
	}
	return n, nil
}
func (r *lineReader) Close() error { r.closes++; return nil }

func TestStartingCase01(t *testing.T) {
	v, e := subject.ScanLines(strings.NewReader("a\nb"))
	if e != nil || !slices.Equal(v, []string{"a", "b"}) {
		t.Errorf("final line got%q,%v", v, e)
	}
}

// Add cases for the remaining requirements in TASK.md.
