package subject_test

import (
	"context"
	"errors"
	"sync"
	subject "task.local/deadline-capped-batch-001/subject"
	"testing"
	"testing/synctest"
	"time"
)

type budgetKey struct{}

func origin() time.Time { return time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC) }

type observation struct {
	index       int
	ctx         context.Context
	deadline    time.Time
	hasDeadline bool
	done        <-chan struct{}
	marker      any
	entryErr    error
}

func observe(index int, ctx context.Context) observation {
	o := observation{index: index, ctx: ctx}
	if ctx != nil {
		o.deadline, o.hasDeadline = ctx.Deadline()
		o.done = ctx.Done()
		o.marker = ctx.Value(budgetKey{})
		o.entryErr = ctx.Err()
	}
	return o
}

type observations struct {
	mu     sync.Mutex
	values []observation
}

func (o *observations) add(v observation) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.values = append(o.values, v)
}
func (o *observations) copy() []observation {
	o.mu.Lock()
	defer o.mu.Unlock()
	return append([]observation(nil), o.values...)
}

type budgetCase struct {
	name              string
	count             int
	budget            time.Duration
	ceiling           time.Duration
	hasCeiling        bool
	parentDeadline    time.Duration
	hasParentDeadline bool
	canceled          bool
	errorAt           int
}

func checkBudgetCase(t *testing.T, c budgetCase) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.WithValue(context.Background(), budgetKey{}, "marker"))
		defer cancel()
		if c.hasParentDeadline {
			var stop context.CancelFunc
			parent, stop = context.WithDeadline(parent, origin().Add(c.parentDeadline))
			defer stop()
		}
		if c.canceled {
			cancel()
		}
		var ceiling time.Time
		if c.hasCeiling {
			ceiling = origin().Add(c.ceiling)
		}
		wantDeadline := origin().Add(c.budget)
		if c.hasCeiling && ceiling.Before(wantDeadline) {
			wantDeadline = ceiling
		}
		if c.hasParentDeadline && origin().Add(c.parentDeadline).Before(wantDeadline) {
			wantDeadline = origin().Add(c.parentDeadline)
		}
		var wantEntry error
		if c.canceled {
			wantEntry = context.Canceled
		} else if !wantDeadline.After(origin()) {
			wantEntry = context.DeadlineExceeded
		}
		sentinels := []error{errors.New("same text"), errors.New("same text"), errors.New("last")}
		values := []int{-7, 0, 19}
		want := make([]subject.Result, c.count)
		var log observations
		operations := make([]func(context.Context) (int, error), c.count)
		for i := range operations {
			index := i
			var err error
			if i == c.errorAt {
				err = sentinels[i]
			}
			want[i] = subject.Result{Value: values[i], Err: err}
			operations[i] = func(ctx context.Context) (int, error) { log.add(observe(index, ctx)); return values[index], err }
		}
		got := subject.RunBudgeted(parent, c.budget, ceiling, operations)
		synctest.Wait()
		if len(got) != len(want) {
			t.Errorf("result count=%d; want %d", len(got), len(want))
		}
		for i := range got {
			if i >= len(want) {
				break
			}
			if got[i].Value != want[i].Value {
				t.Errorf("result[%d].Value=%d; want %d including values with errors", i, got[i].Value, want[i].Value)
			}
			if got[i].Err != want[i].Err {
				t.Errorf("result[%d].Err identity=%v; want exact callback error %v", i, got[i].Err, want[i].Err)
			}
		}
		observed := log.copy()
		if len(observed) != c.count {
			t.Errorf("callback count=%d; want %d even after errors or expiry", len(observed), c.count)
		}
		var common <-chan struct{}
		for i, o := range observed {
			if o.index != i {
				t.Errorf("callback order[%d]=%d; want %d", i, o.index, i)
			}
			if o.ctx == nil {
				t.Errorf("callback %d received nil context", o.index)
				continue
			}
			if !o.hasDeadline || !o.deadline.Equal(wantDeadline) {
				t.Errorf("callback %d effective deadline=(%v,%t); want earliest %v", o.index, o.deadline, o.hasDeadline, wantDeadline)
			}
			if o.marker != "marker" {
				t.Errorf("callback %d lost parent value: got %v", o.index, o.marker)
			}
			if o.entryErr != wantEntry {
				t.Errorf("callback %d entry error=%v; want %v", o.index, o.entryErr, wantEntry)
			}
			if wantEntry == nil {
				if o.done == nil || o.done == parent.Done() {
					t.Errorf("callback %d operation scope is parent or has nil Done", o.index)
				}
				if i == 0 {
					common = o.done
				} else if o.done != common {
					t.Errorf("callback %d changed operation scope; want shared live Done", o.index)
				}
			}
			select {
			case <-o.ctx.Done():
			default:
				t.Errorf("operation context remains live after return at callback %d", o.index)
			}
		}
		if wantEntry == nil && parent.Err() != nil {
			t.Errorf("parent unexpectedly canceled after return: %v", parent.Err())
		}
	})
}
func TestBudgetPublicCompositionAndResults(t *testing.T) {
	cases := []budgetCase{
		{name: "nil", count: 0, budget: 5 * time.Second, errorAt: -1},
		{name: "one-success", count: 1, budget: 5 * time.Second, errorAt: -1},
		{name: "all-success", count: 3, budget: 5 * time.Second, errorAt: -1},
		{name: "error-prefix", count: 3, budget: 5 * time.Second, errorAt: 0},
		{name: "error-middle", count: 3, budget: 5 * time.Second, errorAt: 1},
		{name: "error-last", count: 3, budget: 5 * time.Second, errorAt: 2},
		{name: "earlier-ceiling", count: 3, budget: 8 * time.Second, hasCeiling: true, ceiling: 3 * time.Second, errorAt: 0},
		{name: "later-ceiling", count: 3, budget: 4 * time.Second, hasCeiling: true, ceiling: 9 * time.Second, errorAt: -1},
		{name: "earliest-parent", count: 3, budget: 8 * time.Second, hasCeiling: true, ceiling: 6 * time.Second, hasParentDeadline: true, parentDeadline: 2 * time.Second, errorAt: 2},
		{name: "earliest-budget", count: 3, budget: 2 * time.Second, hasCeiling: true, ceiling: 6 * time.Second, hasParentDeadline: true, parentDeadline: 8 * time.Second, errorAt: -1},
		{name: "zero-budget", count: 3, budget: 0, errorAt: 0},
		{name: "negative-budget", count: 3, budget: -time.Second, errorAt: 2},
		{name: "past-ceiling", count: 3, budget: 5 * time.Second, hasCeiling: true, ceiling: -time.Second, errorAt: -1},
		{name: "canceled-parent", count: 3, budget: 5 * time.Second, canceled: true, errorAt: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) { checkBudgetCase(t, c) })
	}
}

// All callbacks record data only; finite asynchronous defects cannot report after their test ends.
func timedBudget(t *testing.T, parentCancel bool) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.WithValue(context.Background(), budgetKey{}, "marker"))
		defer cancel()
		rescue, cleanup := make(chan struct{}), make(chan struct{})
		released := false
		release := func() {
			if !released {
				released = true
				close(rescue)
				close(cleanup)
			}
		}
		defer release()
		atFour, atSix := time.NewTimer(4*time.Second), time.NewTimer(6*time.Second)
		defer atFour.Stop()
		defer atSix.Stop()
		entered, observed := make(chan observation, 16), make(chan error, 16)
		finished := make(chan struct{}, 16)
		var log observations
		operations := []func(context.Context) (int, error){
			func(ctx context.Context) (int, error) {
				log.add(observe(0, ctx))
				timer := time.NewTimer(2 * time.Second)
				defer timer.Stop()
				select {
				case <-timer.C:
				case <-rescue:
				}
				return -4, nil
			},
			func(ctx context.Context) (int, error) {
				o := observe(1, ctx)
				log.add(o)
				entered <- o
				var done <-chan struct{}
				if ctx != nil {
					done = ctx.Done()
				}
				select {
				case <-done:
					observed <- ctx.Err()
				case <-rescue:
					observed <- nil
				}
				<-cleanup
				finished <- struct{}{}
				return 11, nil
			},
		}
		returned := make(chan []subject.Result, 1)
		wrapperDone := make(chan struct{})
		go func() {
			defer close(wrapperDone)
			returned <- subject.RunBudgeted(parent, 5*time.Second, origin().Add(20*time.Second), operations)
		}()
		var o observation
		haveEntry, haveResult := false, false
		var got []subject.Result
		select {
		case o = <-entered:
			haveEntry = true
		case got = <-returned:
			haveResult = true
			t.Error("returned before callback cleanup or skipped second callback")
		}
		if haveEntry {
			if o.ctx == nil {
				t.Error("second callback received nil context")
			} else if !o.hasDeadline || !o.deadline.Equal(origin().Add(5*time.Second)) {
				t.Errorf("second callback deadline=%v; want shared entry deadline %v; budget restarted after first callback", o.deadline, origin().Add(5*time.Second))
			}
			if parentCancel {
				cancel()
				synctest.Wait()
			} else {
				<-atFour.C
				synctest.Wait()
				if o.ctx != nil && o.ctx.Err() != nil {
					t.Errorf("shared scope canceled before five-second budget: %v", o.ctx.Err())
				}
				<-atSix.C
				synctest.Wait()
			}
			expected := error(context.DeadlineExceeded)
			if parentCancel {
				expected = context.Canceled
			}
			select {
			case err := <-observed:
				if err != expected {
					t.Errorf("observed cancellation=%v; want %v", err, expected)
				}
			default:
				if parentCancel {
					t.Error("parent cancellation did not reach active callback")
				} else {
					t.Error("shared five-second budget did not expire after earlier work")
				}
			}
			select {
			case got = <-returned:
				haveResult = true
				t.Error("returned before callback cleanup")
			default:
			}
		}
		release()
		cancel()
		synctest.Wait()
		if !haveResult {
			select {
			case got = <-returned:
				haveResult = true
			default:
				t.Error("wrapper did not return after independent cleanup release")
			}
		}
		<-wrapperDone
		synctest.Wait()
		if len(got) != 2 {
			t.Errorf("timed result count=%d; want 2", len(got))
		} else if got[0].Value != -4 || got[1].Value != 11 || got[0].Err != nil || got[1].Err != nil {
			t.Errorf("timed results=%v; want exact [(-4,nil),(11,nil)]", got)
		}
		seen := log.copy()
		if len(seen) != 2 {
			t.Errorf("timed callback count=%d; want 2", len(seen))
		}
		// All finite wrong variants have their rescue/cleanup gates released before bubble exit.
	})
}
func TestSharedBudgetIsNotRenewedAfterWork(t *testing.T) { timedBudget(t, false) }
func TestParentCancellationWaitsForCleanup(t *testing.T) { timedBudget(t, true) }
