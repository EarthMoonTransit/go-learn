package subject

import "sync"

func StartVisits(values []int, start <-chan struct{}, visit func(int)) <-chan struct{} {
	done := make(chan struct{})
	var workers sync.WaitGroup
	workers.Add(len(values))
	for index := range values {
		go func(argument int) {
			defer workers.Done()
			<-start
			visit(argument)
		}(index)
	}
	go func() { workers.Wait(); close(done) }()
	return done
}
