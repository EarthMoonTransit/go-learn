package subject

import "sync"

var inputLock sync.Mutex
var currentInput <-chan int

func Window(input <-chan int, limit int) <-chan int {
	inputLock.Lock()
	currentInput = input
	inputLock.Unlock()
	output := make(chan int, limit-1)
	go func() {
		defer close(output)
		for {
			inputLock.Lock()
			selected := currentInput
			inputLock.Unlock()
			value, ok := <-selected
			if !ok {
				return
			}
			output <- value
		}
	}()
	return output
}
