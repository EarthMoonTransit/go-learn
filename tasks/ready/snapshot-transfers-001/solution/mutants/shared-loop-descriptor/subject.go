package subject

import "sync"

type Transfer struct {
	Source      <-chan int
	Destination chan<- int
	Offset      int
}

func StartTransfers(transfers []Transfer, start <-chan struct{}) <-chan struct{} {
	done := make(chan struct{})
	var workers sync.WaitGroup
	var item Transfer
	for _, transfer := range transfers {
		item = transfer
		workers.Add(1)
		go func() {
			defer workers.Done()
			<-start
			value, ok := <-item.Source
			if ok {
				item.Destination <- value + item.Offset
			}
		}()
	}
	go func() { workers.Wait(); close(done) }()
	return done
}
