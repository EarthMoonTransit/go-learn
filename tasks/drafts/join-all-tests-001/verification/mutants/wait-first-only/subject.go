package subject

func JoinAll(work []func()) <-chan struct{} {
	done := make(chan struct{})
	if len(work) == 0 {
		close(done)
		return done
	}
	for i, run := range work {
		go func(index int, f func()) {
			f()
			if index == 0 {
				close(done)
			}
		}(i, run)
	}
	return done
}
