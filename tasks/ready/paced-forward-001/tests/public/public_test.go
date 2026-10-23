package subject_test

import (
	"slices"
	"task.local/paced-forward-001/subject"
	"testing"
	"testing/synctest"
)

type pacingProbe struct {
	input                      chan int
	release                    chan struct{}
	handle                     chan (<-chan int)
	output                     <-chan int
	returned                   bool
	inputClosed, releaseClosed bool
	seen                       []int
}

func beginPace(t *testing.T, values []int) *pacingProbe {
	t.Helper()
	p := &pacingProbe{input: make(chan int, 8), release: make(chan struct{}, 12), handle: make(chan (<-chan int), 1)}
	for _, value := range values {
		p.input <- value
	}
	go func() { p.handle <- subject.Pace(p.input, p.release) }()
	synctest.Wait()
	select {
	case p.output = <-p.handle:
		p.returned = true
	default:
		t.Error("Pace did not return before input EOF")
	}
	if p.returned {
		if p.output == nil {
			t.Error("Pace output is nil; want a private channel")
		}
		if cap(p.output) != 0 {
			t.Errorf("output capacity=%d; want 0", cap(p.output))
		}
	}
	return p
}

func (p *pacingProbe) endInput() {
	if !p.inputClosed {
		close(p.input)
		p.inputClosed = true
	}
}
func (p *pacingProbe) endRelease() {
	if !p.releaseClosed {
		close(p.release)
		p.releaseClosed = true
	}
}

func (p *pacingProbe) noValue(t *testing.T, label string) {
	t.Helper()
	synctest.Wait()
	select {
	case value, ok := <-p.output:
		if ok {
			p.seen = append(p.seen, value)
			t.Errorf("%s: output=%d before permission; want no value", label, value)
		} else {
			t.Errorf("%s: output closed before input EOF", label)
		}
	default:
	}
}

func (p *pacingProbe) take(t *testing.T, want int, label string) {
	t.Helper()
	synctest.Wait()
	select {
	case value, ok := <-p.output:
		if !ok {
			t.Errorf("%s: output closed; want value %d", label, want)
			return
		}
		p.seen = append(p.seen, value)
		if value != want {
			t.Errorf("%s: value=%d; want %d", label, value, want)
		}
	default:
		t.Errorf("%s: no permitted output; want %d before EOF", label, want)
	}
}

func (p *pacingProbe) finish(t *testing.T, want []int) {
	t.Helper()
	// Rescue belongs to the fixture, including after a failed prompt/permit assertion.
	p.endInput()
	p.endRelease()
	synctest.Wait()
	if !p.returned {
		select {
		case p.output = <-p.handle:
			p.returned = true
		default:
		}
	}
	closed := false
	for attempts := 0; attempts < 40; attempts++ {
		synctest.Wait()
		select {
		case value, ok := <-p.output:
			if !ok {
				closed = true
			} else {
				p.seen = append(p.seen, value)
			}
		default:
		}
		if closed {
			break
		}
	}
	if !closed {
		t.Error("output remains open after caller EOF, release closure and drain")
	}
	if !slices.Equal(p.seen, want) {
		t.Errorf("complete stream=%v; want %v", p.seen, want)
	}
}

func TestPaceRepeatedPermissions(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		values := []int{0, 4, 4, -3}
		p := beginPace(t, values)
		defer p.finish(t, values)
		if len(p.input) != 3 {
			t.Errorf("paused input remaining=%d; want 3 (one held value)", len(p.input))
		}
		for i, value := range values {
			p.noValue(t, "paused value")
			p.release <- struct{}{}
			p.take(t, value, "one-token delivery")
			synctest.Wait()
			if len(p.input) != max(0, len(values)-i-2) {
				t.Errorf("after delivery %d input remaining=%d; want %d", i, len(p.input), max(0, len(values)-i-2))
			}
		}
	})
}

func TestPacePreservesIdleToken(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		p := beginPace(t, nil)
		defer p.finish(t, []int{7, 8})
		p.release <- struct{}{}
		synctest.Wait()
		if len(p.release) != 1 {
			t.Errorf("idle token count=%d; want 1 until input exists", len(p.release))
		}
		p.input <- 7
		p.take(t, 7, "retained early token")
		p.input <- 8
		p.noValue(t, "second value needs a new token")
		p.release <- struct{}{}
		p.take(t, 8, "second explicit token")
	})
}

func TestPacePreservesSecondTokenWhileSendHeld(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		p := beginPace(t, []int{11, 22})
		defer p.finish(t, []int{11, 22})
		p.release <- struct{}{}
		synctest.Wait() // No output receiver: the first permitted send is still held.
		p.release <- struct{}{}
		synctest.Wait()
		if len(p.release) != 1 {
			t.Errorf("held-send token count=%d; want 1 preserved for next value", len(p.release))
		}
		if len(p.input) != 1 {
			t.Errorf("held-send input remaining=%d; want 1", len(p.input))
		}
		p.take(t, 11, "first held send")
		p.take(t, 22, "preserved second token")
	})
}

func TestPaceReleaseClosure(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		values := []int{-2, 0, -2, 5}
		p := beginPace(t, values)
		defer p.finish(t, values)
		p.noValue(t, "pending before release closure")
		p.endRelease()
		for _, value := range values {
			p.take(t, value, "release EOF permits remaining value")
		}
	})
}

func TestPaceInputEOFWithPending(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		p := beginPace(t, []int{6})
		defer p.finish(t, []int{6})
		p.endInput()
		p.noValue(t, "input EOF preserves pending unpermitted value")
		p.release <- struct{}{}
		p.take(t, 6, "pending value after input EOF")
	})
}

func TestPaceEmptyInputEOF(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		p := beginPace(t, nil)
		defer p.finish(t, nil)
		p.endInput()
		synctest.Wait()
		select {
		case _, ok := <-p.output:
			if ok {
				t.Error("empty input produced a value")
			}
		default:
			t.Error("empty input EOF did not close output while release remained open")
		}
	})
}

func TestPaceIndependentOverlappingCalls(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		a := beginPace(t, []int{31, 32})
		b := beginPace(t, []int{71, 72})
		if a.returned && b.returned && a.output == b.output {
			t.Error("overlapping calls share output; want private channels")
		}
		a.finish(t, []int{31, 32}) // B's input and permission remain open and held.
		if len(b.input) != 1 {
			t.Errorf("independent B input remaining=%d; want 1 held value", len(b.input))
		}
		b.noValue(t, "independent B remains paused")
		b.finish(t, []int{71, 72})
	})
}
