package subject

func Start(work func()) <-chan struct{} {
	done := make(chan struct{}, 1)
	go func() { work(); done <- struct{}{}; close(done) }()
	return done
}
