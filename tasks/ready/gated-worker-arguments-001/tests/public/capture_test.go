package subject_test

import (
	"slices"
	"sort"
	"sync"
	subject "task.local/gated-worker-arguments-001/subject"
	"testing"
	"testing/synctest"
)

type p19Batch struct {
	input, want                          []int
	start, release                       chan struct{}
	returned                             chan (<-chan struct{})
	done                                 <-chan struct{}
	received, startClosed, releaseClosed bool
	mutex                                sync.Mutex
	calls                                []int
}

func p19NewBatch(values []int) *p19Batch {
	b := &p19Batch{input: values, want: append([]int(nil), values...), start: make(chan struct{}), release: make(chan struct{}), returned: make(chan (<-chan struct{}), 1)}
	go func() {
		b.returned <- subject.StartVisits(values, b.start, func(value int) {
			b.mutex.Lock()
			b.calls = append(b.calls, value)
			b.mutex.Unlock()
			<-b.release
		})
	}()
	return b
}
func (b *p19Batch) openStart() {
	if !b.startClosed {
		close(b.start)
		b.startClosed = true
	}
}
func (b *p19Batch) releaseCallbacks() {
	if !b.releaseClosed {
		close(b.release)
		b.releaseClosed = true
	}
}
func (b *p19Batch) cleanup() { b.openStart(); b.releaseCallbacks(); synctest.Wait() }
func (b *p19Batch) captureReturn(t *testing.T, label string, required bool) {
	t.Helper()
	if !b.received {
		select {
		case b.done = <-b.returned:
			b.received = true
		default:
		}
	}
	if required && !b.received {
		t.Errorf("%s launcher has not returned; want a completion channel before start release", label)
	}
}
func (b *p19Batch) assertDone(t *testing.T, label string, wantClosed bool) {
	t.Helper()
	if !b.received {
		if wantClosed {
			t.Errorf("%s launcher has not returned; want completed handle", label)
		}
		return
	}
	if b.done == nil {
		t.Errorf("%s completion channel is nil, want non-nil", label)
		return
	}
	select {
	case _, ok := <-b.done:
		if ok {
			t.Errorf("%s completion sent a value, want closure without payload", label)
		}
		if !wantClosed {
			t.Errorf("%s completion closed while callbacks remain pending; want every callback returned first", label)
		}
	default:
		if wantClosed {
			t.Errorf("%s completion is still open, want closed after every callback returned", label)
		}
	}
}
func (b *p19Batch) snapshot() []int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return append([]int(nil), b.calls...)
}
func (b *p19Batch) assertStarted(t *testing.T, label string, want int) {
	t.Helper()
	if got := len(b.snapshot()); got != want {
		t.Errorf("%s callback starts = %d, want %d", label, got, want)
	}
}
func (b *p19Batch) assertValuesAndInput(t *testing.T, label string) {
	t.Helper()
	got, want := b.snapshot(), append([]int(nil), b.want...)
	sort.Ints(got)
	sort.Ints(want)
	if !slices.Equal(got, want) {
		t.Errorf("%s callback argument multiset = %v, want %v including each repeated element", label, got, want)
	}
	if !slices.Equal(b.input, b.want) {
		t.Errorf("%s input after completion = %v, want unchanged %v", label, b.input, b.want)
	}
}
func p19Exercise(t *testing.T, values []int) {
	t.Helper()
	b := p19NewBatch(values)
	defer b.cleanup()
	synctest.Wait()
	b.captureReturn(t, "before start", true)
	b.assertStarted(t, "before start closes", 0)
	b.assertDone(t, "before start", len(values) == 0)
	b.openStart()
	synctest.Wait()
	b.captureReturn(t, "after start", true)
	b.assertStarted(t, "all callbacks held", len(b.want))
	if len(b.want) > 0 {
		b.assertDone(t, "callbacks held", false)
	}
	b.releaseCallbacks()
	synctest.Wait()
	b.captureReturn(t, "after callbacks", true)
	b.assertDone(t, "after callbacks", true)
	b.assertValuesAndInput(t, "completed batch")
}
func TestCapturedArguments(t *testing.T) {
	for _, c := range []struct {
		name   string
		values []int
	}{{"distinct", []int{-4, 0, 9, 6}}, {"repeated", []int{5, 5, -2, 0, -2}}, {"singleton", []int{8}}} {
		t.Run(c.name, func(t *testing.T) { synctest.Test(t, func(t *testing.T) { p19Exercise(t, c.values) }) })
	}
}
func TestEmptyBatches(t *testing.T) {
	for _, c := range []struct {
		name   string
		values []int
	}{{"nil", nil}, {"empty", []int{}}} {
		t.Run(c.name, func(t *testing.T) { synctest.Test(t, func(t *testing.T) { p19Exercise(t, c.values) }) })
	}
}
func TestIndependentConcurrentBatches(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		first, second := p19NewBatch([]int{0, -7, 0}), p19NewBatch([]int{4, 11})
		defer first.cleanup()
		defer second.cleanup()
		synctest.Wait()
		first.captureReturn(t, "first before start", true)
		second.captureReturn(t, "second before start", true)
		first.assertStarted(t, "first before start", 0)
		second.assertStarted(t, "second before start", 0)
		first.openStart()
		second.openStart()
		synctest.Wait()
		first.captureReturn(t, "first started", true)
		second.captureReturn(t, "second started", true)
		first.assertStarted(t, "first callbacks held", 3)
		second.assertStarted(t, "second callbacks held", 2)
		first.assertDone(t, "first callbacks held", false)
		second.assertDone(t, "second callbacks held", false)
		first.releaseCallbacks()
		synctest.Wait()
		first.captureReturn(t, "first finished", true)
		first.assertDone(t, "first finished", true)
		second.assertDone(t, "second while first completed", false)
		second.releaseCallbacks()
		synctest.Wait()
		second.captureReturn(t, "second finished", true)
		second.assertDone(t, "second finished", true)
		first.assertValuesAndInput(t, "independent first")
		second.assertValuesAndInput(t, "independent second")
	})
}
