package subject_test

import (
	"errors"
	"strings"
	subject "task.local/buffered-word-stream-001/subject"
	"testing"
)

func TestGraderChunkedUTF8AndBounds(t *testing.T) {
	normalWords(t, "中\u2002文\v終", "中\n文\n終\n", 1)
	normalWords(t, strings.Repeat("z ", 8192), strings.Repeat("z\n", 8192), 11)
	marker := errors.New("grader input terminal")
	dst := &wordDestination{quota: -1}
	src := &wordSource{data: []byte("one two "), terminal: marker, step: 2}
	if got := subject.WriteWords(dst, src); got != marker {
		t.Errorf("chunked source error=%v; want exact marker", got)
	}
	if string(dst.data) != "one\ntwo\n" {
		t.Errorf("chunked accepted prefix=%q; want ordered words", dst.data)
	}
	openWordStreams(t, src, dst)
}
