package subject_test

import (
	"fmt"
	"slices"
	subject "task.local/drain-values-001/subject"
	"testing"
	"testing/synctest"
)

func checkOpen(t *testing.T, values []int) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		in := make(chan int, 32)
		done := make(chan []int, 1)
		go func() { done <- subject.Collect(in) }()
		returned := false
		var got []int
		checkHeld := func(phase string) {
			synctest.Wait()
			if !returned {
				select {
				case got = <-done:
					returned = true
					t.Errorf("Collect returned during %s while input open; want wait for producer close", phase)
				default:
				}
			}
		}
		checkHeld("initial empty interval")
		for i, v := range values {
			in <- v
			checkHeld(fmt.Sprintf("after value %d (%d)", i, v))
		}
		close(in)
		synctest.Wait()
		if !returned {
			select {
			case got = <-done:
				returned = true
			default:
				t.Error("Collect did not return after producer close")
			}
		}
		if returned && got == nil {
			t.Error("Collect returned nil; want non-nil result even for empty input")
		}
		if !slices.Equal(got, values) {
			t.Errorf("Collect result=%v; want every value in order %v", got, values)
		}
	})
}
func TestWaitForProducerClose(t *testing.T) {
	t.Run("empty", func(t *testing.T) { checkOpen(t, nil) })
	t.Run("gated-values", func(t *testing.T) { checkOpen(t, []int{4, 0, -3, 4, 2}) })
}
func TestAlreadyClosedBufferedInput(t *testing.T) {
	for _, values := range [][]int{nil, {}, {0}, {5, -1, 5, 0, 8}} {
		in := make(chan int, 32)
		for _, v := range values {
			in <- v
		}
		close(in)
		got := subject.Collect(in)
		if got == nil {
			t.Errorf("Collect(%v)=nil; want non-nil slice", values)
		}
		if !slices.Equal(got, values) {
			t.Errorf("Collect(%v)=%v; want unchanged order and duplicates", values, got)
		}
	}
}
func TestCallsKeepOwnResults(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		a, b := make(chan int, 32), make(chan int, 32)
		ad, bd := make(chan []int, 1), make(chan []int, 1)
		go func() { ad <- subject.Collect(a) }()
		synctest.Wait()
		a <- 3
		synctest.Wait()
		go func() { bd <- subject.Collect(b) }()
		synctest.Wait()
		b <- 9
		synctest.Wait()
		a <- 3
		synctest.Wait()
		close(a)
		synctest.Wait()
		select {
		case got := <-ad:
			if !slices.Equal(got, []int{3, 3}) {
				t.Errorf("call A while B held=%v; want own values [3 3]", got)
			}
		default:
			t.Error("call A did not finish after its close while B held")
		}
		select {
		case got := <-bd:
			t.Errorf("call B returned %v before its producer close", got)
		default:
		}
		b <- 0
		close(b)
		synctest.Wait()
		select {
		case got := <-bd:
			if !slices.Equal(got, []int{9, 0}) {
				t.Errorf("call B=%v; want own values [9 0]", got)
			}
		default:
			t.Error("call B did not finish after its close")
		}
	})
}
