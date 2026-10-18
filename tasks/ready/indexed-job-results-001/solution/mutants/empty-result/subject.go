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
	if len(jobs) == 0 {
		return []Result{{Index: 0}}
	}
	out := make([]Result, len(jobs))
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
