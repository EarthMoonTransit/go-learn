package subject_test

import "bufio"

import (
	"bytes"
	"task.local/buffered-flush-001/subject"
	"testing"
)

func TestGraderCase01(t *testing.T) {
	var out bytes.Buffer
	if e := subject.WriteBuffered(bufio.NewWriter(&out), "a\nb"); e != nil || out.String() != "a\nb" {
		t.Errorf("newline got%q,%v", out.String(), e)
	}
}
