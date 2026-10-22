package subject

type Result struct {
	Value    int
	Err      error
	Panicked bool
}

func Launch(work func() int) <-chan Result { out := make(chan Result); close(out); return out }
