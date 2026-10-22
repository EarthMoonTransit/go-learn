package subject

import "sync"

var pendingMu sync.Mutex
var pending []int

func Publish(values []int, ready <-chan struct{}) <-chan int {
	pendingMu.Lock()
	pending = values
	pendingMu.Unlock()
	out := make(chan int)
	go func() {
		<-ready
		pendingMu.Lock()
		current := pending
		pendingMu.Unlock()
		for _, v := range current {
			out <- v
		}
		close(out)
	}()
	return out
}
