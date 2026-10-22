package subject

import "sync"

type Result struct {
	Index int
	Value int
	Err   error
}

func RunPool(jobs []func() (int, error), workers int) []Result {
	queue := make(chan int, len(jobs))
	for i := range jobs {
		queue <- i
	}
	close(queue)
	merged := make(chan Result)
	var running sync.WaitGroup
	count := min(workers, len(jobs))
	running.Add(count)
	for range count {
		go func() {
			defer running.Done()
			for i := range queue {
				value, _ := jobs[i]()
				merged <- Result{Index: i, Value: value, Err: nil}
			}
		}()
	}
	go func() { running.Wait(); close(merged) }()
	var results []Result
	for item := range merged {
		results = append(results, item)
	}
	return results
}
