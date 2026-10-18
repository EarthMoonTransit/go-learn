package subject

import "sync"

func StartVisits(values []int, start <-chan struct{}, visit func(int)) <-chan struct{} {
	done := make(chan struct{})
	var workers sync.WaitGroup
	workers.Add(len(values))
	for i, value := range values {
		values[i] = 0
		go func(argument int) {
			defer workers.Done()
			<-start
			visit(argument)
		}(value)
	}
	go func() { workers.Wait(); close(done) }()
	return done
}
