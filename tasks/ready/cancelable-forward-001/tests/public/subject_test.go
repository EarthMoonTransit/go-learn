package subject_test

import (
	"context"
	"slices"
	"task.local/cancelable-forward-001/subject"
	"testing"
	"testing/synctest"
)

func closeOwnedInput(input chan int) {
	for {
		select {
		case _, open := <-input:
			if !open {
				return
			}
		default:
			close(input)
			return
		}
	}
}
func drainForward(out <-chan int, limit int) {
	for step := 0; step < limit; step++ {
		synctest.Wait()
		select {
		case _, open := <-out:
			if !open {
				return
			}
		default:
			return
		}
	}
}
func expectForwardClosed(t *testing.T, out <-chan int) {
	t.Helper()
	synctest.Wait()
	select {
	case got, open := <-out:
		if open {
			t.Errorf("after cancellation/EOF output = %d, want closed after queued values", got)
		}
	default:
		t.Error("after cancellation/EOF output = pending, want closed")
	}
}
func expectCallerInputOpen(t *testing.T, input chan int) bool {
	t.Helper()
	select {
	case _, open := <-input:
		if !open {
			t.Error("caller input = closed, want caller-owned open channel")
			return false
		}
	default:
	}
	return true
}

func TestForwardOrderAndEOF(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		input := make(chan int, 4)
		want := []int{0, 3, -2, 3}
		for _, v := range want {
			input <- v
		}
		close(input)
		out := subject.Forward(context.Background(), input)
		defer func() { drainForward(out, 8); synctest.Wait() }()
		if out == nil {
			t.Fatal("Forward output = nil, want non-nil channel")
		}
		var got []int
		for step := 0; step < len(want); step++ {
			synctest.Wait()
			select {
			case value, open := <-out:
				if open {
					got = append(got, value)
				}
			default:
			}
		}
		if !slices.Equal(got, want) {
			t.Errorf("forwarded order/occurrences = %v, want %v", got, want)
		}
		expectForwardClosed(t, out)
	})
}
func TestForwardCancellationPhases(t *testing.T) {
	for _, mode := range []string{"before-start", "waiting-input", "waiting-output"} {
		t.Run(mode, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				parent, cancel := context.WithCancel(context.Background())
				input := make(chan int)
				stop := make(chan struct{})
				if mode == "before-start" {
					cancel()
				}
				out := subject.Forward(parent, input)
				defer func() {
					cancel()
					close(stop)
					synctest.Wait()
					closeOwnedInput(input)
					drainForward(out, cap(out)+3)
					synctest.Wait()
				}()
				if out == nil {
					t.Fatal("Forward output = nil, want non-nil channel")
				}
				queued := 0
				if mode == "waiting-output" {
					// Fill any permitted output buffer and leave one send blocked. This tests
					// backpressure without adding an unbuffered-output requirement.
					supplied := make(chan struct{})
					go func() {
						defer close(supplied)
						for i := 0; i < cap(out)+1; i++ {
							select {
							case input <- 7:
							case <-stop:
								return
							}
						}
					}()
					synctest.Wait()
					select {
					case <-supplied:
					default:
						t.Error("output phase input not accepted, want one value beyond output capacity")
					}
					queued = len(out)
				} else {
					synctest.Wait()
				}
				cancel()
				synctest.Wait()
				for i := 0; i < queued; i++ {
					select {
					case _, open := <-out:
						if !open {
							t.Error("pre-cancel queued value absent, want retained output")
						}
					default:
						t.Error("pre-cancel queued value pending, want retained output")
					}
				}
				expectForwardClosed(t, out)
				if expectCallerInputOpen(t, input) {
					synctest.Wait()
					select {
					case input <- 99:
						t.Error("stage still receives caller input after cancellation, want no retained receiver")
					default:
					}
				}
				if parent.Err() != context.Canceled {
					t.Errorf("parent error = %v, want context.Canceled", parent.Err())
				}
			})
		})
	}
}

func TestForwardProgressBeforeInputClose(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.Background())
		input := make(chan int, 1)
		out := subject.Forward(parent, input)
		defer func() { cancel(); closeOwnedInput(input); drainForward(out, cap(out)+3); synctest.Wait() }()
		input <- 9
		synctest.Wait()
		select {
		case got, open := <-out:
			if !open || got != 9 {
				t.Errorf("live input output = (%d,%t), want (9,true) before EOF", got, open)
			}
		default:
			t.Error("live input output = pending, want value 9 before EOF")
		}
		cancel()
		expectForwardClosed(t, out)
	})
}
