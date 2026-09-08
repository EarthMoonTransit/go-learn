package subject_test

import (
	"errors"
	"task.local/writer-double-001/subject"
	"testing"
)

func TestPublicCase01(t *testing.T) {
	cause := errors.New("limit")
	w := subject.NewWriter(5, cause)
	if w == nil {
		t.Fatal("nil writer")
	}
	initial, initialCalls := w.Snapshot()
	if len(initial) != 0 || initialCalls != 0 {
		t.Errorf("initial snapshot %q calls%d want empty/0", initial, initialCalls)
	}
	n, e := w.Write([]byte("ab"))
	if n != 2 || e != nil {
		t.Errorf("initial got%d,%v", n, e)
	}
	n, e = w.Write([]byte("cdef"))
	data, calls := w.Snapshot()
	if n != 3 || e != cause || string(data) != "abcde" || calls != 2 {
		t.Errorf("partial got%d,%v data%q calls%d", n, e, data, calls)
	}
	n, e = w.Write([]byte("x"))
	if n != 0 || e != cause {
		t.Errorf("full got%d,%v", n, e)
	}
	n, e = w.Write(nil)
	_, calls = w.Snapshot()
	if n != 0 || e != nil || calls != 4 {
		t.Errorf("empty got%d,%v calls%d", n, e, calls)
	}
}

func TestPublicCase02(t *testing.T) {
	cause := errors.New("limit")
	a := subject.NewWriter(2, cause)
	b := subject.NewWriter(0, cause)
	if a == nil || b == nil {
		t.Fatal("nil writer")
	}
	input := []byte("ab")
	n, e := a.Write(input)
	if n != 2 || e != nil {
		t.Errorf("exact limit got%d,%v", n, e)
	}
	input[0] = 'x'
	data, calls := a.Snapshot()
	if string(data) != "ab" || calls != 1 {
		t.Errorf("input alias data%q calls%d", data, calls)
	}
	data[0] = 'z'
	next, _ := a.Snapshot()
	if string(next) != "ab" {
		t.Errorf("snapshot alias %q", next)
	}
	if n, e := b.Write([]byte("x")); n != 0 || e != cause {
		t.Errorf("zero limit got%d,%v", n, e)
	}
	next, calls = b.Snapshot()
	if len(next) != 0 || calls != 1 {
		t.Errorf("independent writer %q calls%d", next, calls)
	}
}
