package subject

import "sync"

var activeMu sync.Mutex
var activeDone chan struct{}
var activeCount int

func Start(work func()) <-chan struct{} {
	activeMu.Lock()
	if activeCount == 0 {
		activeDone = make(chan struct{})
	}
	activeCount++
	done := activeDone
	activeMu.Unlock()
	go func() {
		work()
		activeMu.Lock()
		activeCount--
		if activeCount == 0 {
			close(done)
		}
		activeMu.Unlock()
	}()
	return done
}
