package subject_test

import (
	"slices"
	subject "task.local/idle-stream-001/subject"
	"testing"
	"testing/synctest"
	"time"
)

type idleResult struct {
	values    []int
	exhausted bool
}

func startIdle(input <-chan int, quiet time.Duration) chan idleResult {
	done := make(chan idleResult, 1)
	go func() { values, exhausted := subject.UntilIdle(input, quiet); done <- idleResult{values, exhausted} }()
	synctest.Wait()
	return done
}
func advanceIdleClock(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
	timer.Stop()
	synctest.Wait()
}
func checkStillCollecting(t *testing.T, done chan idleResult, phase string) {
	t.Helper()
	select {
	case got := <-done:
		t.Errorf("%s: returned early with (%v,%t); want collection still active", phase, got.values, got.exhausted)
		done <- got
	default:
	}
}
func readIdleResult(t *testing.T, done chan idleResult, input chan int, want []int, exhausted bool, phase string) {
	t.Helper()
	var got idleResult
	select {
	case got = <-done:
	default:
		t.Errorf("%s: UntilIdle did not return at the required boundary", phase)
		close(input)
		synctest.Wait()
		select {
		case got = <-done:
		default:
			advanceIdleClock(100 * time.Second)
			select {
			case got = <-done:
			default:
				t.Fatal("UntilIdle still active after cleanup")
			}
		}
	}
	if !slices.Equal(got.values, want) || got.exhausted != exhausted {
		t.Errorf("%s: result=(%v,%t); want (%v,%t)", phase, got.values, got.exhausted, want, exhausted)
	}
}
func TestQuietIntervalStartsAtCall(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		input := make(chan int, 16)
		done := startIdle(input, 10*time.Second)
		advanceIdleClock(9 * time.Second)
		checkStillCollecting(t, done, "before initial quiet interval")
		advanceIdleClock(2 * time.Second)
		readIdleResult(t, done, input, nil, false, "empty open input became idle")
	})
}
func TestEveryArrivalRestartsQuietInterval(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		input := make(chan int, 16)
		done := startIdle(input, 10*time.Second)
		advanceIdleClock(6 * time.Second)
		input <- 3
		synctest.Wait()
		advanceIdleClock(6 * time.Second)
		checkStillCollecting(t, done, "after original deadline and before reset deadline")
		input <- 0
		synctest.Wait()
		advanceIdleClock(6 * time.Second)
		checkStillCollecting(t, done, "after zero restarted timer")
		input <- 3
		synctest.Wait()
		advanceIdleClock(9 * time.Second)
		checkStillCollecting(t, done, "nine seconds after duplicate arrival")
		advanceIdleClock(2 * time.Second)
		readIdleResult(t, done, input, []int{3, 0, 3}, false, "quiet after final arrival")
	})
}
func checkClosedInput(t *testing.T, values []int) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		input := make(chan int, 16)
		for _, value := range values {
			input <- value
		}
		close(input)
		done := startIdle(input, 10*time.Second)
		var got idleResult
		select {
		case got = <-done:
		default:
			t.Error("closed input: UntilIdle must return without waiting for quiet interval")
			advanceIdleClock(20 * time.Second)
			select {
			case got = <-done:
			default:
				t.Fatal("closed input did not finish after cleanup clock advance")
			}
		}
		if !slices.Equal(got.values, values) || !got.exhausted {
			t.Errorf("closed input result=(%v,%t); want (%v,true)", got.values, got.exhausted, values)
		}
	})
}
func TestClosedInputPreservesAllValues(t *testing.T) {
	checkClosedInput(t, nil)
	checkClosedInput(t, []int{7, 0, 7, -2})
}
func TestCloseAfterAcceptedValueReturnsPromptly(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		input := make(chan int, 16)
		done := startIdle(input, 10*time.Second)
		advanceIdleClock(4 * time.Second)
		input <- 11
		synctest.Wait()
		close(input)
		synctest.Wait()
		select {
		case got := <-done:
			if !slices.Equal(got.values, []int{11}) || !got.exhausted {
				t.Errorf("close after value result=(%v,%t); want ([11],true)", got.values, got.exhausted)
			}
		default:
			t.Error("close after value: UntilIdle waited instead of returning promptly")
			advanceIdleClock(20 * time.Second)
			select {
			case <-done:
			default:
				t.Fatal("UntilIdle still active after closed-input cleanup")
			}
		}
	})
}
