package subject

import "sync"

type Result struct {
	Index, Value int
	Err          error
	Panicked     bool
	Panic        any
}

func runJob(index int, work func() (int, error)) (result Result) {
	result.Index = index
	defer func() {
		if caught := recover(); caught != nil {
			result = Result{Index: index, Panicked: true, Panic: caught}
		}
	}()
	result.Value, result.Err = work()
	return
}
func StreamJobs(jobs []func() (int, error), workers int) <-chan Result {
	output := make(chan Result)
	if len(jobs) == 0 {
		close(output)
		return output
	}
	merged := make(chan Result, len(jobs))
	indices := make(chan int, len(jobs))
	for index := range jobs {
		indices <- index
	}
	close(indices)
	count := min(workers, len(jobs))
	var joined sync.WaitGroup
	joined.Add(count)
	for range count {
		go func() {
			defer joined.Done()
			for index := range indices {
				result := runJob(index, jobs[index])
				merged <- result
			}
		}()
	}
	go func() {
		joined.Wait()
		close(merged)
		for result := range merged {
			output <- result
		}
		close(output)
	}()
	return output
}
