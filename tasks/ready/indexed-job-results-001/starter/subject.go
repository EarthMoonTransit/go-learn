package subject

type Result struct {
	Index int
	Value int
	Err   error
}

func CollectResults(jobs []func() (int, error)) []Result { return nil }
