package subject

import "sync"

var paceMu sync.Mutex
var activePace <-chan int

func Pace(input <-chan int, release <-chan struct{}) <-chan int {
	paceMu.Lock()
	defer paceMu.Unlock()
	if activePace != nil {
		return activePace
	}
	output := make(chan int)
	activePace = output
	go func() {
		defer func() { paceMu.Lock(); activePace = nil; paceMu.Unlock(); close(output) }()
		for value := range input {
			<-release
			output <- value
		}
	}()
	return output
}
