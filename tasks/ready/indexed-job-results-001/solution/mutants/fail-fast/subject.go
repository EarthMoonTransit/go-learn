package subject

type Result struct {
	Index int
	Value int
	Err   error
}

func CollectResults(jobs []func() (int, error)) []Result {
	completed := make(chan Result, len(jobs))
	for i, job := range jobs {
		go func(index int, operation func() (int, error)) {
			value, err := operation()
			completed <- Result{Index: index, Value: value, Err: err}
		}(i, job)
	}
	out := make([]Result, len(jobs))
	for range jobs {
		result := <-completed
		out[result.Index] = result
		if result.Err != nil {
			return out
		}
	}
	return out
}
