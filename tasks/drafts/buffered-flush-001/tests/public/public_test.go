package subject_test

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"task.local/buffered-flush-001/subject"
	"testing"
)

type bufferSink struct {
	short          bool
	err            error
	data           []byte
	writes, closes int
}

func (w *bufferSink) Write(p []byte) (int, error) {
	w.writes++
	n := len(p)
	if w.short {
		n /= 2
	}
	w.data = append(w.data, p[:n]...)
	return n, w.err
}
func (w *bufferSink) Close() error { w.closes++; return nil }

func TestPublicCase01(t *testing.T) {
	for _, text := range []string{"", "small", " λ ", strings.Repeat("x", 9000)} {
		var out bytes.Buffer
		e := subject.WriteBuffered(&out, text)
		if e != nil || out.String() != text {
			t.Errorf("buffer got%q,%v want%q", out.String(), e, text)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	cause := errors.New("flush failed")
	for _, c := range []struct {
		short bool
		err   error
		want  error
	}{{false, cause, cause}, {true, nil, io.ErrShortWrite}, {true, cause, cause}} {
		w := &bufferSink{short: c.short, err: c.err}
		e := subject.WriteBuffered(w, "abc")
		if e != c.want || w.closes != 0 {
			t.Errorf("flush error got%v closes%d want%v,0", e, w.closes, c.want)
		}
	}
}

func TestPublicCase03(t *testing.T) {
	w := &bufferSink{}
	if e := subject.WriteBuffered(w, ""); e != nil || w.writes != 0 || w.closes != 0 {
		t.Errorf("empty err%v writes%d closes%d", e, w.writes, w.closes)
	}
	e := subject.WriteBuffered(w, "text")
	if e != nil || string(w.data) != "text" || w.closes != 0 {
		t.Errorf("success err%v bytes%q closes%d", e, w.data, w.closes)
	}
}

func TestLargeWriteError(t *testing.T) {
	cause := errors.New("write failed")
	w := &bufferSink{err: cause}
	if e := subject.WriteBuffered(w, strings.Repeat("x", 9000)); e != cause || w.closes != 0 {
		t.Errorf("large write error got%v closes%d wantoriginal,0", e, w.closes)
	}
}
