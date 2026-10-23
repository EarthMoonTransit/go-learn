package subject

import "sync"

type Transfer struct {
	Source      <-chan int
	Destination chan<- int
	Offset      int
}

func StartTransfers(transfers []Transfer, start <-chan struct{}) <-chan struct{} {
	owner := completionFor(start)
	done := owner.handle
	var workers sync.WaitGroup
	for _, transfer := range transfers {
		workers.Add(1)
		go func(item Transfer) {
			defer workers.Done()
			<-start
			value, ok := <-item.Source
			if ok {
				item.Destination <- value + item.Offset
			}
		}(transfer)
	}
	go func() { workers.Wait(); owner.once.Do(func() { close(done) }) }()
	return done
}

type completionState struct {
	handle chan struct{}
	once   sync.Once
}

var completionMu sync.Mutex
var completions = make(map[<-chan struct{}]*completionState)

func completionFor(start <-chan struct{}) *completionState {
	completionMu.Lock()
	defer completionMu.Unlock()
	current := completions[start]
	if current == nil {
		current = &completionState{handle: make(chan struct{})}
		completions[start] = current
	}
	return current
}
