package subject_test

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"task.local/bounded-copy-001/subject"
	"testing"
)

type trackedReader struct {
	data          []byte
	err           error
	reads, closes int
}

func (r *trackedReader) Read(p []byte) (int, error) {
	r.reads++
	n := copy(p, r.data)
	r.data = r.data[n:]
	if len(r.data) == 0 {
		if r.err != nil {
			return n, r.err
		}
		return n, io.EOF
	}
	return n, nil
}
func (r *trackedReader) Close() error { r.closes++; return nil }

type trackedWriter struct {
	data           []byte
	short          bool
	err            error
	writes, closes int
}

func (w *trackedWriter) Write(p []byte) (int, error) {
	w.writes++
	n := len(p)
	if w.short {
		n /= 2
	}
	w.data = append(w.data, p[:n]...)
	return n, w.err
}
func (w *trackedWriter) Close() error { w.closes++; return nil }

func TestPublicCase01(t *testing.T) {
	for _, c := range []struct {
		input string
		limit int64
		want  string
	}{{"abcdef", 3, "abc"}, {"ab", 9, "ab"}, {"", 2, ""}} {
		var out bytes.Buffer
		n, e := subject.CopyAtMost(&out, strings.NewReader(c.input), c.limit)
		if e != nil || n != int64(len(c.want)) || out.String() != c.want {
			t.Errorf("copy got%d,%v,%q want%q", n, e, out.String(), c.want)
		}
	}
}

func TestPublicCase02(t *testing.T) {
	for _, limit := range []int64{-1, 0} {
		r := &trackedReader{data: []byte("abc")}
		w := &trackedWriter{}
		n, e := subject.CopyAtMost(w, r, limit)
		want := error(nil)
		if limit < 0 {
			want = subject.ErrLimit
		}
		if n != 0 || e != want || r.reads != 0 || w.writes != 0 || r.closes != 0 || w.closes != 0 {
			t.Errorf("limit%d got%d,%v reads%d writes%d closes%d,%d", limit, n, e, r.reads, w.writes, r.closes, w.closes)
		}
	}
}

func TestPublicCase03(t *testing.T) {
	cause := errors.New("read failed")
	r := &trackedReader{data: []byte("abc"), err: cause}
	w := &trackedWriter{}
	n, e := subject.CopyAtMost(w, r, 10)
	if n != 3 || e != cause || string(w.data) != "abc" || r.closes != 0 || w.closes != 0 {
		t.Errorf("read error got%d,%v data%q closes%d,%d", n, e, w.data, r.closes, w.closes)
	}
}

func TestPublicCase04(t *testing.T) {
	for _, cause := range []error{nil, errors.New("write failed")} {
		r := &trackedReader{data: []byte("abcd")}
		w := &trackedWriter{short: true, err: cause}
		n, e := subject.CopyAtMost(w, r, 10)
		want := cause
		if want == nil {
			want = io.ErrShortWrite
		}
		if n != 2 || e != want || string(w.data) != "ab" || r.closes != 0 || w.closes != 0 {
			t.Errorf("partial write got%d,%v data%q closes%d,%d", n, e, w.data, r.closes, w.closes)
		}
	}
}

var originalLimit = subject.ErrLimit

func TestPreservedLimitSentinel(t *testing.T) {
	if originalLimit == nil {
		t.Fatal("ErrLimit must be non-nil")
	}
	for _, limit := range []int64{-1, 0, 1, 9} {
		var out bytes.Buffer
		_, e := subject.CopyAtMost(&out, strings.NewReader("abc"), limit)
		if subject.ErrLimit != originalLimit || (limit < 0 && e != originalLimit) {
			t.Errorf("ErrLimit sentinel identity changed for limit %d", limit)
		}
	}
}
func TestSuccessfulCopyKeepsStreamsOpen(t *testing.T) {
	for _, limit := range []int64{2, 10} {
		r := &trackedReader{data: []byte("abc")}
		w := &trackedWriter{}
		n, e := subject.CopyAtMost(w, r, limit)
		want := "abc"
		if limit == 2 {
			want = "ab"
		}
		if n != int64(len(want)) || e != nil || string(w.data) != want || r.closes != 0 || w.closes != 0 {
			t.Errorf("successful copy got%d,%v,%q; closes got%d,%d want0,0", n, e, w.data, r.closes, w.closes)
		}
	}
}
