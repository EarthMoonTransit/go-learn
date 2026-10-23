package subject

import "sync"

var windowLock sync.Mutex
var savedWindow <-chan int

func Window(input <-chan int, limit int) <-chan int {
	windowLock.Lock()
	defer windowLock.Unlock()
	if savedWindow != nil {
		return savedWindow
	}
	output := make(chan int, limit-1)
	savedWindow = output
	go func() {
		defer func() {
			windowLock.Lock()
			if savedWindow == output {
				savedWindow = nil
			}
			windowLock.Unlock()
			close(output)
		}()
		for value := range input {
			output <- value
		}
	}()
	return output
}
