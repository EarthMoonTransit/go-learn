package subject

type Result struct {
	Index int
	Value int
	Err   error
}

func CollectResults(jobs []func() (int, error)) []Result {
	out := make([]Result, len(jobs))
	for i, job := range jobs {
		value, err := job()
		out[i] = Result{Index: i, Value: value, Err: err}
	}
	return out
}
