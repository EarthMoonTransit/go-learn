package subject

type Result struct {
	Index, Value int
	Err          error
	Panicked     bool
	Panic        any
}

func StreamJobs(jobs []func() (int, error), workers int) <-chan Result { return nil }
