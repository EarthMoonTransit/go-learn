package subject_test

import (
	"bytes"
	"errors"
	"io"
	"task.local/writer-greeting-001/subject"
	"testing"
)

type writerProbe struct {
	closes int
	n      int
	err    error
	text   string
	calls  int
}

func (w *writerProbe) Write(p []byte) (int, error) {
	w.calls++
	w.text += string(p)
	if w.n < 0 {
		return len(p), w.err
	}
	return w.n, w.err
}

func TestPublicCase01(t *testing.T) {
	var b bytes.Buffer
	e := subject.WriteGreeting(&b, "Chris")
	if e != nil || b.String() != "Hello, Chris" {
		t.Errorf("greeting got%q,%v", b.String(), e)
	}
}

func TestPublicCase02(t *testing.T) {
	for _, name := range []string{"", " λ "} {
		w := &writerProbe{n: -1}
		e := subject.WriteGreeting(w, name)
		if w.closes != 0 {
			t.Errorf("closes got %d want 0", w.closes)
		}
		if e != nil || w.text != "Hello, "+name || (w.calls != 1) {
			t.Errorf("name%q got%q,%v calls%d", name, w.text, e, w.calls)
		}
	}
}

func TestPublicCase03(t *testing.T) {
	sentinel := errors.New("write failed")
	w := &writerProbe{n: 2, err: sentinel}
	if e := subject.WriteGreeting(w, "x"); e != sentinel || (w.calls != 1) {
		t.Errorf("failure got%v calls%d", e, w.calls)
	}
}

func TestPublicCase04(t *testing.T) {
	w := &writerProbe{n: 2}
	if e := subject.WriteGreeting(w, "x"); e != io.ErrShortWrite || (w.calls != 1) {
		t.Errorf("short got%v calls%d", e, w.calls)
	}
}

func (w *writerProbe) Close() error { w.closes++; return nil }

func TestWriterRemainsOpen(t *testing.T) {
	for _, c := range []struct {
		n int
		e error
	}{{-1, nil}, {2, errors.New("failed")}, {2, nil}} {
		w := &writerProbe{n: c.n, err: c.e}
		_ = subject.WriteGreeting(w, "x")
		if w.closes != 0 {
			t.Errorf("closes got %d want 0 (count%d error%v)", w.closes, c.n, c.e)
		}
	}
}
