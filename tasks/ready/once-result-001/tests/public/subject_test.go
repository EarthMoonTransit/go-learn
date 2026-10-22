package subject_test

import (
	"errors"
	"sync/atomic"
	subject "task.local/once-result-001/subject"
	"testing"
	"testing/synctest"
)

func TestLazyAndCompletedPublication(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		var count atomic.Int32
		marker := errors.New("first error")
		made := make(chan func() (int, error), 1)
		load := func() (int, error) { count.Add(1); <-release; return 7, marker }
		go func() { made <- subject.Memoize(load) }()
		synctest.Wait()
		var memo func() (int, error)
		released := false
		select {
		case memo = <-made:
		default:
			t.Error("Memoize must return lazily without calling load")
			close(release)
			released = true
			synctest.Wait()
			memo = <-made
		}
		if n := count.Load(); n != 0 {
			t.Errorf("construction load calls=%d; want0", n)
		}
		type result struct {
			v int
			e error
		}
		done := make(chan result, 4)
		// Observe one held call before introducing Once mutex contenders.
		go func() { v, e := memo(); done <- result{v, e} }()
		synctest.Wait()
		var results []result
		if !released {
			select {
			case x := <-done:
				results = append(results, x)
				t.Errorf("call returned (%d,%v) before initializer completed", x.v, x.e)
			default:
			}
		}
		if n := count.Load(); n != 1 {
			t.Errorf("initializer calls while first call held=%d; want1", n)
		}
		started := make(chan struct{}, 3)
		for range 3 {
			go func() { started <- struct{}{}; v, e := memo(); done <- result{v, e} }()
		}
		for range 3 {
			<-started
		}
		// Mutex waits are not durable synctest blocking. Release before Wait.
		// Exact values/count below and -race assess concurrent publication;
		// method review separately establishes all callers wait on Once.
		if !released {
			close(release)
		}
		synctest.Wait()
		for len(results) < 4 {
			select {
			case x := <-done:
				results = append(results, x)
			default:
				t.Error("call did not finish after initializer release")
				return
			}
		}
		for i, x := range results {
			if x.v != 7 || x.e != marker {
				t.Errorf("call%d=(%d,%v); want(7, exact first error)", i, x.v, x.e)
			}
		}
		v, e := memo()
		if v != 7 || e != marker {
			t.Errorf("later call=(%d,%v); want cached(7,exact first error)", v, e)
		}
		if n := count.Load(); n != 1 {
			t.Errorf("total load calls=%d; want1 including errors", n)
		}
	})
}
func TestIndependentInstancesAndZero(t *testing.T) {
	calls := 0
	a := subject.Memoize(func() (int, error) { calls++; return 0, nil })
	b := subject.Memoize(func() (int, error) { calls++; return -4, nil })
	for range 2 {
		av, ae := a()
		bv, be := b()
		if av != 0 || ae != nil || bv != -4 || be != nil {
			t.Errorf("independent instances got(%d,%v),(%d,%v); want(0,nil),(-4,nil)", av, ae, bv, be)
		}
	}
	if calls != 2 {
		t.Errorf("two instances load calls=%d; want2", calls)
	}
}
