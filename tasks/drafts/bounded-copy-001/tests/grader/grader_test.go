package subject_test

import (
	"bytes"
	"strings"
	"task.local/bounded-copy-001/subject"
	"testing"
)

func TestGraderBoundary(t *testing.T) {
	var out bytes.Buffer
	n, e := subject.CopyAtMost(&out, strings.NewReader("abc"), 1)
	if n != 1 || e != nil || out.String() != "a" {
		t.Errorf("one byte got%d,%v,%q", n, e, out.String())
	}
}
