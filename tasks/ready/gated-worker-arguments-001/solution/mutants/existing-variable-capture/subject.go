package subject

import "sync"

func StartVisits(values []int, start <-chan struct{}, visit func(int)) <-chan struct{} {
	done := make(chan struct{})
	var workers sync.WaitGroup
	workers.Add(len(values))
	var value int
	for _, value = range values {
		go func() {
			defer workers.Done()
			<-start
			visit(value)
		}()
	}
	go func() { workers.Wait(); close(done) }()
	return done
}
