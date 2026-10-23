package subject

import "sync"

type Transfer struct {
	Source      <-chan int
	Destination chan<- int
	Offset      int
}

func StartTransfers(transfers []Transfer, start <-chan struct{}) <-chan struct{} {
	done := make(chan struct{}, 1)
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
	go func() { workers.Wait(); done <- struct{}{}; close(done) }()
	return done
}
