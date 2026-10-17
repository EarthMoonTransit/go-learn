package subject

import "sync"

func StartVisits(values []int, start <-chan struct{}, visit func(int)) <-chan struct{} {
	done := make(chan struct{}, 1)
	var workers sync.WaitGroup
	workers.Add(len(values))
	for _, value := range values {
		go func(argument int) {
			defer workers.Done()
			<-start
			visit(argument)
		}(value)
	}
	go func() { workers.Wait(); done <- struct{}{}; close(done) }()
	return done
}
