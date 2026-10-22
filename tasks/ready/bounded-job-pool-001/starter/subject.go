package subject

type Result struct {
	Index int
	Value int
	Err   error
}

func RunPool(jobs []func() (int, error), workers int) []Result { return nil }
