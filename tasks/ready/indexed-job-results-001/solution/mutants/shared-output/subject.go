package subject

import (
	"sync"
)

type Result struct {
	Index int
	Value int
	Err   error
}

var p17ResultCache []Result

func CollectResults(jobs []func() (int, error)) []Result {
	if cap(p17ResultCache) < len(jobs) {
		p17ResultCache = make([]Result, len(jobs))
	}
	out := p17ResultCache[:len(jobs)]
	var workers sync.WaitGroup
	workers.Add(len(jobs))
	for i, job := range jobs {
		go func(index int, operation func() (int, error)) {
			defer workers.Done()
			value, err := operation()
			out[index] = Result{Index: index, Value: value, Err: err}
		}(i, job)
	}
	workers.Wait()
	return out
}
