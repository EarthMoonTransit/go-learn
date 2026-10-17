package subject

import (
	"sync"
)

type Result struct {
	Index int
	Value int
	Err   error
}

func CollectResults(jobs []func() (int, error)) []Result {
	out := make([]Result, len(jobs))
	var workers sync.WaitGroup
	for i, job := range jobs {
		if i == len(jobs)-1 {
			continue
		}
		workers.Add(1)
		go func(index int, operation func() (int, error)) {
			defer workers.Done()
			value, err := operation()
			out[index] = Result{Index: index, Value: value, Err: err}
		}(i, job)
	}
	workers.Wait()
	return out
}
