package subject

import (
	"sync"
	"sync/atomic"
)

// Job and its handle are supplied scaffolding; edit only StartJob.
type Job struct {
	Run    func()
	Handle Handle
}
type Handle struct{ state *jobState }
type jobState struct {
	release     chan struct{}
	done        chan struct{}
	releaseOnce sync.Once
	finishOnce  sync.Once
	calls       atomic.Int64
}

func NewJob() Job {
	state := &jobState{release: make(chan struct{}), done: make(chan struct{})}
	return Job{Handle: Handle{state: state}, Run: func() {
		state.calls.Add(1)
		<-state.release
		state.finishOnce.Do(func() { close(state.done) })
	}}
}
func (h Handle) Release()              { h.state.releaseOnce.Do(func() { close(h.state.release) }) }
func (h Handle) Done() <-chan struct{} { return h.state.done }
func (h Handle) Calls() int            { return int(h.state.calls.Load()) }
