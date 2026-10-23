package subject_test

import (
	"sync/atomic"
	subject "task.local/periodic-budget-001/subject"
	"testing"
	"testing/synctest"
	"time"
)

type visitRecord struct{ elapsed, remaining time.Duration }

func advance(t time.Duration) { clock := time.NewTimer(t); <-clock.C; clock.Stop(); synctest.Wait() }
func startPeriodic(period, budget time.Duration, visit func(time.Duration, time.Duration) bool) chan bool {
	done := make(chan bool, 1)
	go func() { done <- subject.RunPeriodic(period, budget, visit) }()
	synctest.Wait()
	return done
}
func noPeriodicResult(t *testing.T, done chan bool, phase string) {
	t.Helper()
	select {
	case result := <-done:
		t.Errorf("%s: returned %t before required completion", phase, result)
		done <- result
	default:
	}
}
func finishPeriodic(t *testing.T, done chan bool, rescue *atomic.Bool, period, budget time.Duration, want bool, phase string) {
	t.Helper()
	select {
	case got := <-done:
		if got != want {
			t.Errorf("%s: budgetDone=%t; want %t", phase, got, want)
		}
	default:
		t.Errorf("%s: RunPeriodic has not returned at required boundary", phase)
		// All callback gates are released by caller before cleanup; known wrong loops
		// retain either the finite original budget or callback-stop escape.
		rescue.Store(true)
		advance(2*budget + 4*period + time.Millisecond)
		select {
		case got := <-done:
			if got != want {
				t.Errorf("%s after cleanup: budgetDone=%t; want %t", phase, got, want)
			}
		default:
			t.Error("RunPeriodic still running after finite budget/callback rescue")
		}
	}
	synctest.Wait()
}
func visitAt(t *testing.T, records chan visitRecord, wantElapsed, wantRemaining time.Duration, phase string) {
	t.Helper()
	select {
	case got := <-records:
		if got.elapsed != wantElapsed || got.remaining != wantRemaining {
			t.Errorf("%s: visit=(%v,%v); want (%v,%v)", phase, got.elapsed, got.remaining, wantElapsed, wantRemaining)
		}
	default:
		t.Errorf("%s: required callback has not started", phase)
	}
}
func TestPublicFirstPeriodAndRemaining(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var rescue atomic.Bool
		records := make(chan visitRecord, 128)
		done := startPeriodic(2*time.Millisecond, 11*time.Millisecond, func(e, r time.Duration) bool { records <- visitRecord{e, r}; return !rescue.Load() })
		advance(time.Millisecond)
		if n := len(records); n != 0 {
			t.Errorf("visits before first period=%d; want 0", n)
		}
		noPeriodicResult(t, done, "before first period")
		advance(1500 * time.Microsecond)
		visitAt(t, records, 2*time.Millisecond, 9*time.Millisecond, "first visit")
		rescue.Store(true)
		advance(4 * time.Millisecond)
		finishPeriodic(t, done, &rescue, 2*time.Millisecond, 11*time.Millisecond, false, "callback stop")
	})
}
func slowPeriodic(t *testing.T, period, budget, firstObservation, releaseAt time.Duration) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		var calls atomic.Int32
		var rescue atomic.Bool
		records := make(chan visitRecord, 128)
		done := startPeriodic(period, budget, func(e, r time.Duration) bool {
			n := calls.Add(1)
			records <- visitRecord{e, r}
			if n == 1 {
				<-release
				return !rescue.Load()
			}
			return false
		})
		advance(firstObservation)
		visitAt(t, records, period, budget-period, "held first visit")
		advance(releaseAt - firstObservation)
		noPeriodicResult(t, done, "callback still held")
		if n := calls.Load(); n != 1 {
			t.Errorf("callbacks started while first was held=%d; want 1", n)
		}
		close(release)
		synctest.Wait()
		visitAt(t, records, releaseAt, budget-releaseAt, "delayed next visit")
		finishPeriodic(t, done, &rescue, period, budget, false, "second callback stop")
	})
}
func TestPublicDelayedVisitUsesCurrentElapsed(t *testing.T) {
	slowPeriodic(t, 2*time.Millisecond, 11*time.Millisecond, 2500*time.Microsecond, 5500*time.Microsecond)
}
func budgetRun(t *testing.T, period, budget time.Duration) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		var rescue atomic.Bool
		var calls atomic.Int32
		done := startPeriodic(period, budget, func(e, r time.Duration) bool {
			calls.Add(1)
			if e >= budget || r != budget-e {
				t.Errorf("budget callback durations=(%v,%v); budget=%v", e, r, budget)
			}
			return !rescue.Load()
		})
		advance(budget - time.Microsecond)
		noPeriodicResult(t, done, "before total budget")
		advance(501 * time.Microsecond)
		finishPeriodic(t, done, &rescue, period, budget, true, "total budget")
		if period > budget && calls.Load() != 0 {
			t.Errorf("callbacks before a period longer than budget=%d; want 0", calls.Load())
		}
	})
}
func TestPublicIndependentTotalBudget(t *testing.T) {
	budgetRun(t, 5*time.Millisecond, 3*time.Millisecond)
	budgetRun(t, 2*time.Millisecond, 11*time.Millisecond)
}
func TestPublicHeldCallbackOutlivesBudget(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		release := make(chan struct{})
		var rescue atomic.Bool
		var calls atomic.Int32
		done := startPeriodic(2*time.Millisecond, 11*time.Millisecond, func(e, r time.Duration) bool {
			n := calls.Add(1)
			if n == 1 {
				<-release
			}
			return !rescue.Load()
		})
		advance(12500 * time.Microsecond)
		noPeriodicResult(t, done, "callback held beyond budget")
		close(release)
		synctest.Wait()
		finishPeriodic(t, done, &rescue, 2*time.Millisecond, 11*time.Millisecond, true, "budget after callback release")
		if n := calls.Load(); n != 1 {
			t.Errorf("callbacks after expired held visit=%d; want 1", n)
		}
	})
}
func TestPublicImmediateCallbackStop(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var rescue atomic.Bool
		done := startPeriodic(2*time.Millisecond, 11*time.Millisecond, func(e, r time.Duration) bool { return false })
		advance(2500 * time.Microsecond)
		finishPeriodic(t, done, &rescue, 2*time.Millisecond, 11*time.Millisecond, false, "first callback requested stop")
	})
}
