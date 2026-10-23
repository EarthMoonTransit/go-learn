package subject_test

import (
	"slices"
	subject "task.local/buffer-window-stream-001/subject"
	"testing"
	"testing/synctest"
)

func drainWindow(output <-chan int) ([]int, bool) {
	var values []int
	for range 40 {
		synctest.Wait()
		select {
		case value, ok := <-output:
			if !ok {
				return values, true
			}
			values = append(values, value)
		default:
		}
	}
	return values, false
}

func windowCase(t *testing.T, values []int, limit int) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		input := make(chan int, len(values))
		for _, v := range values {
			input <- v
		}
		handles := make(chan (<-chan int), 1)
		go func() { handles <- subject.Window(input, limit) }()
		synctest.Wait()
		var output <-chan int
		returned := false
		select {
		case output = <-handles:
			returned = true
		default:
			t.Error("Window must return before EOF")
		}
		var got []int
		if returned {
			if output == nil {
				t.Error("output is nil")
			} else {
				if cap(output) != limit-1 {
					t.Errorf("capacity=%d; want limit-1=%d", cap(output), limit-1)
				}
				if used := len(values) - len(input); used > limit {
					t.Errorf("in-flight=%d; maximum=%d", used, limit)
				}
				if len(values) >= limit && len(output) != cap(output) {
					t.Errorf("filled buffer=%d; want capacity=%d", len(output), cap(output))
				}
				if len(values) > 0 {
					select {
					case v, ok := <-output:
						if !ok {
							t.Error("output closed before forwarding available input")
						} else {
							got = append(got, v)
						}
					default:
						t.Error("no forwarding progress before EOF")
					}
				}
			}
		}
		// Every required specimen can finish after this caller-owned release.
		close(input)
		synctest.Wait()
		if !returned {
			select {
			case output = <-handles:
				returned = true
			default:
				t.Error("Window did not return after EOF release")
			}
		}
		rest, closed := drainWindow(output)
		got = append(got, rest...)
		if !closed {
			t.Error("output remains open after input exhaustion")
		}
		if !slices.Equal(got, values) {
			t.Errorf("stream=%v; want %v", got, values)
		}
	})
}

func TestWindowPublic(t *testing.T) {
	for _, limit := range []int{1, 2, 3, 4} {
		t.Run(string(rune('0'+limit)), func(t *testing.T) {
			windowCase(t, []int{4, 0, -3, 4, 8, -1, 2, 7}, limit)
		})
	}
	windowCase(t, nil, 1)
	windowCase(t, []int{0}, 4)
	windowCase(t, []int{-5}, 2)
}

func TestPrivateOutputs(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		a, b := make(chan int), make(chan int)
		handles := make(chan (<-chan int), 2)
		go func() { handles <- subject.Window(a, 2) }()
		go func() { handles <- subject.Window(b, 2) }()
		synctest.Wait()
		var outputs []<-chan int
		for range 2 {
			select {
			case out := <-handles:
				outputs = append(outputs, out)
			default:
			}
		}
		if len(outputs) != 2 {
			t.Error("independent Window calls must return before EOF")
		}
		if len(outputs) == 2 && outputs[0] == outputs[1] {
			t.Error("calls share output")
		}
		close(a)
		close(b)
		synctest.Wait()
		for len(outputs) < 2 {
			select {
			case out := <-handles:
				outputs = append(outputs, out)
			default:
				goto released
			}
		}
	released:
		for _, out := range outputs {
			drainWindow(out)
		}
	})
}

func TestOverlappingWindowCalls(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		a, b := make(chan int, 3), make(chan int, 3)
		for _, v := range []int{11, 12, 13} {
			a <- v
		}
		for _, v := range []int{21, 22, 23} {
			b <- v
		}
		ha, hb := make(chan (<-chan int), 1), make(chan (<-chan int), 1)
		go func() { ha <- subject.Window(a, 1) }()
		synctest.Wait()
		var oa, ob <-chan int
		returnedA, returnedB := false, false
		select {
		case oa = <-ha:
			returnedA = true
		default:
			t.Error("overlapping A must return before EOF")
		}
		// A has settled before B exists. No receiver releases either first send yet.
		go func() { hb <- subject.Window(b, 1) }()
		synctest.Wait()
		select {
		case ob = <-hb:
			returnedB = true
		default:
			t.Error("overlapping B must return before EOF")
		}
		if returnedA && returnedB && oa == ob {
			t.Error("overlapping calls share output")
		}
		close(a)
		synctest.Wait()
		if !returnedA {
			select {
			case oa = <-ha:
				returnedA = true
			default:
			}
		}
		gotA, closedA := drainWindow(oa)
		if !closedA {
			t.Error("overlapping A did not close before B release")
		}
		if !slices.Equal(gotA, []int{11, 12, 13}) {
			t.Errorf("overlapping A stream=%v; want [11 12 13]", gotA)
		}
		// Even after assertions, release B's EOF and drain both outputs.
		close(b)
		synctest.Wait()
		if !returnedB {
			select {
			case ob = <-hb:
				returnedB = true
			default:
			}
		}
		gotB, closedB := drainWindow(ob)
		if !closedB {
			t.Error("overlapping B did not close after release")
		}
		if !slices.Equal(gotB, []int{21, 22, 23}) {
			t.Errorf("overlapping B stream=%v; want [21 22 23]", gotB)
		}
		drainWindow(oa)
		drainWindow(ob)
	})
}
