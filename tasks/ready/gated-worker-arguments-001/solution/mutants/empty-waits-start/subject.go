package subject

import "sync"

func StartVisits(values []int, start <-chan struct{}, visit func(int)) <-chan struct{} {
	done := make(chan struct{})
	var workers sync.WaitGroup
	workers.Add(len(values))
	for _, value := range values {
		go func(argument int) {
			defer workers.Done()
			<-start
			visit(argument)
		}(value)
	}
	go func() {
		workers.Wait()
		if len(values) == 0 {
			<-start
		}
		close(done)
	}()
	return done
}
