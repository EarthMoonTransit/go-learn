package subject_test

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"task.local/cleanup-order-001/subject"
	"testing"
)

type cleanupScope struct {
	helpers  int
	cleanups []func()
	errors   []string
}

func (s *cleanupScope) Helper()          { s.helpers++ }
func (s *cleanupScope) Cleanup(f func()) { s.cleanups = append(s.cleanups, f) }
func (s *cleanupScope) Errorf(format string, args ...any) {
	s.errors = append(s.errors, fmt.Sprintf(format, args...))
}
func (s *cleanupScope) finish() {
	for i := len(s.cleanups) - 1; i >= 0; i-- {
		s.cleanups[i]()
	}
	s.cleanups = nil
}

type closeProbe struct {
	name  string
	order *[]string
	calls int
	err   error
}

func (r *closeProbe) Close() error {
	r.calls++
	if r.order != nil {
		*r.order = append(*r.order, r.name)
	}
	return r.err
}

func TestPublicCase01(t *testing.T) {
	s := &cleanupScope{}
	r := &closeProbe{}
	calls := 0
	got, e := subject.AcquireForTest(s, func() (io.Closer, error) { calls++; return r, nil })
	if got != r || e != nil || calls != 1 || s.helpers != 1 || len(s.cleanups) != 1 || r.calls != 0 {
		t.Fatalf("setup got%v,%v calls%d helpers%d cleanup%d closes%d", got, e, calls, s.helpers, len(s.cleanups), r.calls)
	}
	s.finish()
	if r.calls != 1 || len(s.errors) != 0 {
		t.Errorf("cleanup closes%d errors%v", r.calls, s.errors)
	}
}

func TestPublicCase02(t *testing.T) {
	cause := errors.New("acquire failed")
	s := &cleanupScope{}
	unused := &closeProbe{}
	calls := 0
	got, e := subject.AcquireForTest(s, func() (io.Closer, error) { calls++; return unused, cause })
	if got != nil || e != cause || calls != 1 || s.helpers != 1 || len(s.cleanups) != 0 || unused.calls != 0 {
		t.Errorf("failed acquisition got%v,%v calls%d cleanups%d closes%d", got, e, calls, len(s.cleanups), unused.calls)
	}
}

func TestPublicCase03(t *testing.T) {
	var order []string
	cause := errors.New("close failed")
	s := &cleanupScope{}
	a := &closeProbe{name: "a", order: &order}
	b := &closeProbe{name: "b", order: &order, err: cause}
	for _, r := range []*closeProbe{a, b} {
		if _, e := subject.AcquireForTest(s, func() (io.Closer, error) { return r, nil }); e != nil {
			t.Fatal(e)
		}
	}
	s.finish()
	if !slices.Equal(order, []string{"b", "a"}) || a.calls != 1 || b.calls != 1 || !slices.Equal(s.errors, []string{"cleanup: close failed"}) {
		t.Errorf("order%v closes%d,%d errors%v", order, a.calls, b.calls, s.errors)
	}
}
