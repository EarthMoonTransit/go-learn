package subject

import "sync"

var p19SharedCompletion = make(chan struct{})
var p19CloseOnce sync.Once

func StartVisits(values []int, start <-chan struct{}, visit func(int)) <-chan struct{} {
	done := p19SharedCompletion
	var workers sync.WaitGroup
	workers.Add(len(values))
	for _, value := range values {
		go func(argument int) {
			defer workers.Done()
			<-start
			visit(argument)
		}(value)
	}
	go func() { workers.Wait(); p19CloseOnce.Do(func() { close(done) }) }()
	return done
}
