package subject_test

import (
	"slices"
	subject "task.local/owned-batch-channel-001/subject"
	"testing"
	"testing/synctest"
)

func exercise(t *testing.T, values []int, signal bool) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		ready := make(chan struct{}, 1)
		before := append([]int(nil), values...)
		returned := make(chan (<-chan int), 1)
		go func() { returned <- subject.Publish(values, ready) }()
		synctest.Wait()
		var out <-chan int
		received := false
		select {
		case out = <-returned:
			received = true
		default:
			t.Error("Publish waited for ready; want output returned before release")
		}
		if received && out != nil {
			select {
			case v, ok := <-out:
				t.Errorf("output before release: value=%d open=%t; want blocked", v, ok)
			default:
			}
		}
		if signal {
			ready <- struct{}{}
		} else {
			close(ready)
		}
		synctest.Wait()
		if !received {
			select {
			case out = <-returned:
				received = true
			default:
				t.Error("Publish did not return after release")
			}
		}
		if out == nil {
			t.Error("output=nil; want a private non-nil channel")
			return
		}
		got := []int{}
		closed := false
		for i := 0; i < 2*len(values)+4; i++ {
			synctest.Wait()
			select {
			case v, ok := <-out:
				if !ok {
					closed = true
				} else {
					got = append(got, v)
				}
			default:
			}
			if closed {
				break
			}
		}
		if !closed {
			t.Error("producer output remained open; want closed after final value")
		}
		if !slices.Equal(got, before) {
			t.Errorf("output=%v; want all values in order %v", got, before)
		}
		if !slices.Equal(values, before) {
			t.Errorf("input=%v; want unchanged %v", values, before)
		}
		if signal {
			select {
			case _, ok := <-ready:
				if !ok {
					t.Error("caller gate closed; want caller ownership")
				}
			default:
			}
		}
	})
}
func TestOwnedOutput(t *testing.T) {
	for _, c := range []struct {
		name   string
		values []int
	}{
		{"nil", nil}, {"empty", []int{}}, {"one", []int{4}}, {"mixed", []int{3, 0, -2, 3, 8}},
	} {
		t.Run(c.name, func(t *testing.T) { exercise(t, c.values, true) })
	}
	t.Run("closed-gate", func(t *testing.T) { exercise(t, []int{-5, 0, 2}, false) })
}
func TestIndependentOutputs(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ready := make(chan struct{})
		close(ready)
		a := subject.Publish(nil, ready)
		b := subject.Publish(nil, ready)
		synctest.Wait()
		if a == b {
			t.Error("calls share output; want distinct owned channels")
		}
		for _, out := range []<-chan int{a, b} {
			select {
			case v, ok := <-out:
				if ok {
					t.Errorf("empty output value=%d; want closed", v)
				}
			default:
				t.Error("empty output is not closed")
			}
		}
	})
}

func TestActiveCallIndependence(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		gateA, gateB := make(chan struct{}), make(chan struct{})
		returnedA, returnedB := make(chan (<-chan int), 1), make(chan (<-chan int), 1)
		go func() { returnedA <- subject.Publish([]int{2, 6}, gateA) }()
		synctest.Wait()
		go func() { returnedB <- subject.Publish([]int{8, -1, 0}, gateB) }()
		synctest.Wait()
		var a, b <-chan int
		haveA, haveB := false, false
		select {
		case a = <-returnedA:
			haveA = true
		default:
			t.Error("active A did not return before release")
		}
		select {
		case b = <-returnedB:
			haveB = true
		default:
			t.Error("active B did not return before release")
		}
		if haveA && haveB && a == b {
			t.Error("active calls share output; want separate channels")
		}
		drain := func(name string, out <-chan int, want []int) {
			got := []int{}
			closed := false
			for i := 0; i < 32; i++ {
				synctest.Wait()
				select {
				case v, ok := <-out:
					if ok {
						got = append(got, v)
					} else {
						closed = true
					}
				default:
				}
				if closed {
					break
				}
			}
			if !closed {
				t.Errorf("active %s output did not close", name)
			}
			if !slices.Equal(got, want) {
				t.Errorf("active %s output=%v; want its own values %v", name, got, want)
			}
		}
		close(gateA)
		synctest.Wait()
		if !haveA {
			select {
			case a = <-returnedA:
			default:
				t.Error("active A still did not return")
			}
		}
		drain("A", a, []int{2, 6})
		select {
		case v, ok := <-b:
			t.Errorf("active B output before its release: value=%d open=%t; want blocked", v, ok)
		default:
		}
		close(gateB)
		synctest.Wait()
		if !haveB {
			select {
			case b = <-returnedB:
			default:
				t.Error("active B still did not return")
			}
		}
		drain("B", b, []int{8, -1, 0})
	})
}
