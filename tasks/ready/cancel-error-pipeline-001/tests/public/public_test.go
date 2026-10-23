package subject_test

import (
	"context"
	"errors"
	"slices"
	subject "task.local/cancel-error-pipeline-001/subject"
	"testing"
	"testing/synctest"
)

type pipeHandles struct {
	values   <-chan int
	terminal <-chan error
}
type pipeProbe struct {
	input                                               chan int
	cancel                                              context.CancelFunc
	handles                                             chan pipeHandles
	output                                              <-chan int
	terminal                                            <-chan error
	returned, inputClosed, outputClosed, terminalClosed bool
	values                                              []int
	reasons                                             []error
}

func startPipe(t *testing.T, parent context.Context, cancel context.CancelFunc, values []int, step func(int) (int, error)) *pipeProbe {
	t.Helper()
	p := &pipeProbe{input: make(chan int, len(values)), cancel: cancel, handles: make(chan pipeHandles, 1)}
	for _, v := range values {
		p.input <- v
	}
	go func() {
		out, term := subject.TransformUntil(parent, p.input, step)
		p.handles <- pipeHandles{out, term}
	}()
	synctest.Wait()
	p.capture()
	if !p.returned {
		t.Error("pipeline must return before EOF")
	} else {
		if p.output == nil || p.terminal == nil {
			t.Error("nil pipeline handle")
		}
		if cap(p.output) != 0 {
			t.Errorf("output capacity=%d; want 0", cap(p.output))
		}
		if cap(p.terminal) != 1 {
			t.Errorf("terminal capacity=%d; want 1", cap(p.terminal))
		}
	}
	return p
}
func (p *pipeProbe) capture() {
	if p.returned {
		return
	}
	select {
	case h := <-p.handles:
		p.output = h.values
		p.terminal = h.terminal
		p.returned = true
	default:
	}
}
func (p *pipeProbe) closeInput() {
	if !p.inputClosed {
		close(p.input)
		p.inputClosed = true
	}
}
func (p *pipeProbe) observe() {
	for range 80 {
		synctest.Wait()
		p.capture()
		if !p.outputClosed {
			select {
			case v, ok := <-p.output:
				if ok {
					p.values = append(p.values, v)
				} else {
					p.outputClosed = true
				}
			default:
			}
		}
		if !p.terminalClosed {
			select {
			case err, ok := <-p.terminal:
				if ok {
					p.reasons = append(p.reasons, err)
				} else {
					p.terminalClosed = true
				}
			default:
			}
		}
		if p.outputClosed && p.terminalClosed {
			return
		}
	}
}
func (p *pipeProbe) cleanup() { p.cancel(); p.closeInput(); p.observe() }
func (p *pipeProbe) check(t *testing.T, want []int, reason error) {
	t.Helper()
	if !p.outputClosed {
		t.Error("output remains open after termination")
	}
	if !p.terminalClosed {
		t.Error("terminal remains open after termination")
	}
	if !slices.Equal(p.values, want) {
		t.Errorf("mapped prefix=%v; want %v", p.values, want)
	}
	if len(p.reasons) != 1 {
		t.Errorf("terminal count=%d; want 1", len(p.reasons))
	} else if p.reasons[0] != reason {
		t.Errorf("terminal error=%v (%T); want exact %v (%T)", p.reasons[0], p.reasons[0], reason, reason)
	}
}
func pipelineCase(t *testing.T, values []int, failAt int) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.Background())
		marker := errors.New("mapped input rejected")
		calls := make(chan int, 16)
		index := 0
		p := startPipe(t, parent, cancel, values, func(v int) (int, error) {
			calls <- v
			at := index
			index++
			if at == failAt {
				return 900 + v, marker
			}
			return 3 * v, nil
		})
		defer p.cleanup()
		p.observe()
		count := len(values)
		var reason error
		if failAt >= 0 && failAt < len(values) {
			count = failAt
			reason = marker
			if !p.terminalClosed {
				t.Error("terminal still pending after step failure with input open")
			}
		} else if len(values) > 0 && len(p.values) == 0 {
			t.Error("no mapped value before input EOF")
		}
		p.closeInput()
		p.observe()
		expected := make([]int, count)
		for i := range expected {
			expected[i] = 3 * values[i]
		}
		p.check(t, expected, reason)
		wantCalls := len(values)
		if reason != nil {
			wantCalls = failAt + 1
		}
		if len(calls) != wantCalls {
			t.Errorf("step calls=%d; want %d (no calls after first error)", len(calls), wantCalls)
		}
		if parent.Err() != nil {
			t.Error("pipeline canceled caller parent")
		}
	})
}
func TestPipelinePublic(t *testing.T) {
	pipelineCase(t, nil, -1)
	pipelineCase(t, []int{0}, -1)
	pipelineCase(t, []int{-2, 0, 4, -2, 7}, -1)
	for _, at := range []int{0, 2, 4} {
		pipelineCase(t, []int{-2, 0, 4, -2, 7}, at)
	}
}
func TestCancelInputWait(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.Background())
		p := startPipe(t, parent, cancel, nil, func(v int) (int, error) { return v, nil })
		defer p.cleanup()
		cancel()
		p.observe()
		if !p.terminalClosed {
			t.Error("terminal still pending after cancellation while input open")
		}
		p.check(t, nil, context.Canceled)
	})
}
func TestCancelPausedSend(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.Background())
		called := make(chan struct{}, 1)
		p := startPipe(t, parent, cancel, []int{6}, func(v int) (int, error) { called <- struct{}{}; return v, nil })
		defer p.cleanup()
		select {
		case <-called:
		default:
			t.Error("step did not reach paused output send")
		}
		// No receiver exists when cancellation happens. Settle before observing output.
		cancel()
		synctest.Wait()
		select {
		case v, ok := <-p.output:
			if ok {
				p.values = append(p.values, v)
				t.Error("output remains open after paused-send cancellation")
			} else {
				p.outputClosed = true
			}
		default:
			t.Error("output remains open after paused-send cancellation")
		}
		p.observe()
		p.check(t, nil, context.Canceled)
	})
}
func TestJoinHeldStep(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.Background())
		entered := make(chan struct{}, 1)
		release := make(chan struct{})
		p := startPipe(t, parent, cancel, []int{8}, func(v int) (int, error) { entered <- struct{}{}; <-release; return v, nil })
		defer p.cleanup()
		select {
		case <-entered:
		default:
			t.Error("held step was not entered")
		}
		cancel()
		synctest.Wait()
		select {
		case err, ok := <-p.terminal:
			if ok {
				p.reasons = append(p.reasons, err)
			} else {
				p.terminalClosed = true
			}
			t.Error("terminal arrived before active step completed")
		default:
		}
		close(release)
		p.observe()
		p.check(t, nil, context.Canceled)
	})
}
func TestAlreadyCanceled(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		parent, cancel := context.WithCancel(context.Background())
		cancel()
		calls := make(chan int, 4)
		p := startPipe(t, parent, cancel, []int{5, 7}, func(v int) (int, error) { calls <- v; return v, nil })
		defer p.cleanup()
		p.observe()
		if len(p.input) != 2 || len(calls) != 0 {
			t.Errorf("pre-canceled call consumed input/called step: remaining=%d calls=%d", len(p.input), len(calls))
		}
		p.check(t, nil, context.Canceled)
	})
}
