package subject_test

import (
	"slices"
	subject "task.local/snapshot-transfers-001/subject"
	"testing"
	"testing/synctest"
)

type transferSource struct {
	channel chan int
	closed  bool
}
type transferOutput struct {
	channel chan int
	values  []int
	closed  bool
}
type transferGate struct {
	channel chan struct{}
	closed  bool
}
type transferProducer struct {
	abort chan struct{}
	done  chan struct{}
}
type transferCall struct {
	input, before []subject.Transfer
	returned      chan (<-chan struct{})
	handle        <-chan struct{}
	received      bool
}
type transferHarness struct {
	t         *testing.T
	sources   []*transferSource
	outputs   []*transferOutput
	gates     []*transferGate
	producers []transferProducer
	calls     []*transferCall
}

func (h *transferHarness) source(capacity int, values []int, closed bool) *transferSource {
	s := &transferSource{channel: make(chan int, capacity), closed: closed}
	for _, v := range values {
		s.channel <- v
	}
	if closed {
		close(s.channel)
	}
	h.sources = append(h.sources, s)
	return s
}
func (h *transferHarness) output(capacity int, pending ...int) *transferOutput {
	o := &transferOutput{channel: make(chan int, capacity)}
	for _, v := range pending {
		o.channel <- v
	}
	h.outputs = append(h.outputs, o)
	return o
}
func (h *transferHarness) gate() *transferGate {
	g := &transferGate{channel: make(chan struct{})}
	h.gates = append(h.gates, g)
	return g
}
func (g *transferGate) open() {
	if !g.closed {
		close(g.channel)
		g.closed = true
	}
}
func (h *transferHarness) launch(input []subject.Transfer, g *transferGate) *transferCall {
	c := &transferCall{input: input, before: append([]subject.Transfer(nil), input...), returned: make(chan (<-chan struct{}), 1)}
	h.calls = append(h.calls, c)
	go func() { c.returned <- subject.StartTransfers(input, g.channel) }()
	return c
}
func (h *transferHarness) send(s *transferSource, value int) {
	p := transferProducer{abort: make(chan struct{}), done: make(chan struct{})}
	h.producers = append(h.producers, p)
	go func() {
		defer close(p.done)
		select {
		case s.channel <- value:
		case <-p.abort:
		}
	}()
}
func (c *transferCall) capture(t *testing.T, label string) {
	t.Helper()
	if !c.received {
		select {
		case c.handle = <-c.returned:
			c.received = true
		default:
		}
	}
	if !c.received {
		t.Errorf("%s launcher has not returned; want completion before gate release", label)
	} else if c.handle == nil {
		t.Errorf("%s completion=nil; want nonnil", label)
	}
}
func (c *transferCall) completion(t *testing.T, label string, closed bool) {
	t.Helper()
	c.capture(t, label)
	if !c.received || c.handle == nil {
		return
	}
	select {
	case _, ok := <-c.handle:
		if ok {
			t.Errorf("%s completion sent payload; want closure only", label)
		}
		if !closed {
			t.Errorf("%s completion readable before required sends; want pending", label)
		}
	default:
		if closed {
			t.Errorf("%s completion remains open; want closed after all sends", label)
		}
	}
}
func (c *transferCall) preserved(t *testing.T, label string) {
	t.Helper()
	if !slices.Equal(c.input, c.before) {
		t.Errorf("%s descriptors changed; want caller-owned descriptors preserved", label)
	}
}
func (o *transferOutput) drain() {
	if o.closed {
		return
	}
	for n := 0; n < 32; n++ {
		select {
		case value, ok := <-o.channel:
			if !ok {
				o.closed = true
				return
			}
			o.values = append(o.values, value)
		default:
			return
		}
	}
}
func (o *transferOutput) want(t *testing.T, label string, values []int) {
	t.Helper()
	o.drain()
	if !slices.Equal(o.values, values) {
		t.Errorf("%s output=%v; want %v", label, o.values, values)
	}
	if o.closed {
		t.Errorf("%s caller destination closed; want open", label)
	}
}
func (h *transferHarness) cleanup() {
	// All wrong workers perform at most two receives/sends. Rescue reaches every endpoint.
	for _, g := range h.gates {
		g.open()
	}
	for _, p := range h.producers {
		close(p.abort)
	}
	synctest.Wait()
	for _, p := range h.producers {
		select {
		case <-p.done:
		default:
			h.t.Error("fixture sender not joined after abort")
		}
	}
	// No declared wrong closes Sources. Abort/join precedes these caller closes.
	for _, s := range h.sources {
		if !s.closed {
			close(s.channel)
			s.closed = true
		}
	}
	for step := 0; step < 16; step++ {
		for _, o := range h.outputs {
			o.drain()
		}
		synctest.Wait()
	}
	for _, c := range h.calls {
		c.capture(h.t, "after finite rescue")
	}
}
func TestCapturedTransferDescriptors(t *testing.T) {
	for _, replace := range []bool{false, true} {
		label := "preserved"
		if replace {
			label = "caller-replaced"
		}
		t.Run(label, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				h := &transferHarness{t: t}
				defer h.cleanup()
				g := h.gate()
				input := make([]subject.Transfer, 3)
				originalSources := make([]*transferSource, 3)
				originalOutputs := make([]*transferOutput, 3)
				values, offsets := []int{0, -7, 4}, []int{5, 2, -3}
				for i := range input {
					originalSources[i] = h.source(2, []int{values[i], 40 + i}, false)
					originalOutputs[i] = h.output(3)
					input[i] = subject.Transfer{Source: originalSources[i].channel, Destination: originalOutputs[i].channel, Offset: offsets[i]}
				}
				call := h.launch(input, g)
				synctest.Wait()
				call.capture(t, "before start")
				call.completion(t, "nonempty before start", false)
				call.preserved(t, "before caller replacement")
				for i, source := range originalSources {
					if remaining := len(source.channel); remaining != 2 {
						t.Errorf("source%d consumed before start: remaining=%d; want2", i, remaining)
					}
				}
				for i, o := range originalOutputs {
					o.want(t, "before start endpoint"+string(rune('A'+i)), nil)
				}
				decoys := make([]*transferOutput, 0, 3)
				if replace {
					for i := range input {
						s := h.source(2, []int{100 + i, 200 + i}, false)
						o := h.output(4)
						decoys = append(decoys, o)
						input[i] = subject.Transfer{Source: s.channel, Destination: o.channel, Offset: 30 + i}
					}
					call.before = append([]subject.Transfer(nil), input...)
				}
				g.open()
				synctest.Wait()
				call.completion(t, "captured transfers finished", true)
				for i, o := range originalOutputs {
					o.want(t, "captured endpoint"+string(rune('A'+i)), []int{values[i] + offsets[i]})
				}
				for i, o := range decoys {
					o.want(t, "unused decoy"+string(rune('A'+i)), nil)
				}
				call.preserved(t, "after transfers")
				for i, s := range originalSources {
					select {
					case value, ok := <-s.channel:
						if !ok || value != 40+i {
							t.Errorf("source%d remaining=(%d,%v); want (%d,true)", i, value, ok, 40+i)
						}
					default:
						t.Errorf("source%d lost second queued value; want%d pending", i, 40+i)
					}
				}
			})
		})
	}
}
func TestIndependentTransferBlockingPhases(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		h := &transferHarness{t: t}
		defer h.cleanup()
		g := h.gate()
		fullSource, waitingSource, readySource := h.source(1, []int{5}, false), h.source(1, nil, false), h.source(1, []int{0}, false)
		fullOut, waitingOut, readyOut := h.output(1, 91), h.output(1), h.output(1)
		c := h.launch([]subject.Transfer{{fullSource.channel, fullOut.channel, 4}, {waitingSource.channel, waitingOut.channel, 3}, {readySource.channel, readyOut.channel, -2}}, g)
		synctest.Wait()
		c.capture(t, "blocking setup")
		g.open()
		synctest.Wait()
		c.completion(t, "full destination and empty source held", false)
		readyOut.want(t, "ready peer while other transfers blocked", []int{-2})
		waitingOut.want(t, "empty source before input", nil)
		waitingSource.channel <- (-4)
		synctest.Wait()
		waitingOut.want(t, "released input", []int{-1})
		c.completion(t, "full destination still held", false)
		fullOut.drain()
		synctest.Wait()
		fullOut.want(t, "full destination after caller drain", []int{91, 9})
		c.completion(t, "all blocking phases released", true)
		c.preserved(t, "blocking phases")
	})
}
func TestUnbufferedTransferEndpoints(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		h := &transferHarness{t: t}
		defer h.cleanup()
		g := h.gate()
		input, ready := h.source(0, nil, false), h.source(1, []int{9}, false)
		receivedOut, rendezvousOut := h.output(1), h.output(0)
		c := h.launch([]subject.Transfer{{input.channel, receivedOut.channel, 5}, {ready.channel, rendezvousOut.channel, -1}}, g)
		synctest.Wait()
		c.capture(t, "unbuffered setup")
		g.open()
		synctest.Wait()
		receivedOut.want(t, "unbuffered receive waiting", nil)
		c.completion(t, "unbuffered endpoints held", false)
		h.send(input, -5)
		synctest.Wait()
		receivedOut.want(t, "unbuffered receive paired", []int{0})
		c.completion(t, "unbuffered destination still held", false)
		rendezvousOut.drain()
		synctest.Wait()
		rendezvousOut.want(t, "unbuffered send paired", []int{8})
		c.completion(t, "unbuffered transfers finished", true)
	})
}
func TestClosedTransferSources(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		h := &transferHarness{t: t}
		defer h.cleanup()
		g := h.gate()
		empty, queued := h.source(0, nil, true), h.source(2, []int{0, 6}, true)
		a, b := h.output(2), h.output(3)
		c := h.launch([]subject.Transfer{{empty.channel, a.channel, 12}, {queued.channel, b.channel, -4}}, g)
		synctest.Wait()
		c.capture(t, "closed setup")
		g.open()
		synctest.Wait()
		a.want(t, "closed empty source", nil)
		b.want(t, "closed queued zero", []int{-4})
		c.completion(t, "closed sources finished", true)
		select {
		case value, ok := <-queued.channel:
			if !ok || value != 6 {
				t.Errorf("closed source next=(%d,%v); want(6,true)", value, ok)
			}
		default:
			t.Error("closed source lost queued second value")
		}
	})
}
func TestEmptyAndIndependentTransferCompletions(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		h := &transferHarness{t: t}
		defer h.cleanup()
		g := h.gate()
		a, b := h.launch(nil, g), h.launch([]subject.Transfer{}, g)
		synctest.Wait()
		a.completion(t, "nil batch before start", true)
		b.completion(t, "empty batch before start", true)
		if a.received && b.received && a.handle != nil && a.handle == b.handle {
			t.Error("empty batches share completion; want fresh handles")
		}
	})
}
func TestTransferCallIsolation(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		h := &transferHarness{t: t}
		defer h.cleanup()
		g := h.gate()
		pending, ready := h.source(1, nil, false), h.source(1, []int{8}, false)
		oa, ob := h.output(1), h.output(1)
		a, b := h.launch([]subject.Transfer{{pending.channel, oa.channel, 2}}, g), h.launch([]subject.Transfer{{ready.channel, ob.channel, -1}}, g)
		synctest.Wait()
		a.capture(t, "first call")
		b.capture(t, "second call")
		if a.received && b.received && a.handle != nil && a.handle == b.handle {
			t.Error("active calls share completion; want independent handles")
		}
		g.open()
		synctest.Wait()
		a.completion(t, "first source still held", false)
		b.completion(t, "second finished", true)
		ob.want(t, "independent second", []int{7})
		pending.channel <- 3
		synctest.Wait()
		oa.want(t, "independent first", []int{5})
		a.completion(t, "first finished", true)
	})
}
