package subject

import "sync"

var pendingMu sync.Mutex
var pending []int

func Collect(in <-chan int) []int {
	pendingMu.Lock()
	pending = nil
	pendingMu.Unlock()
	for v := range in {
		pendingMu.Lock()
		pending = append(pending, v)
		pendingMu.Unlock()
	}
	pendingMu.Lock()
	out := append([]int{}, pending...)
	pendingMu.Unlock()
	return out
}
