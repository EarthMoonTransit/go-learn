package subject

import "sync"

func StartVisits(values []int, start <-chan struct{}, visit func(int)) <-chan struct{} {
	done := make(chan struct{})
	var workers sync.WaitGroup
	workers.Add(1)
	go func() {
		defer workers.Done()
		if len(values) > 0 {
			<-start
		}
		for _, value := range values {
			visit(value)
		}
	}()
	go func() { workers.Wait(); close(done) }()
	return done
}
