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
	for index := range transfers {
		workers.Add(1)
		go func(index int) {
			defer workers.Done()
			<-start
			item := transfers[index]
			value, ok := <-item.Source
			if ok {
				item.Destination <- value + item.Offset
			}
		}(index)
	}
	go func() { workers.Wait(); close(done) }()
	return done
}
