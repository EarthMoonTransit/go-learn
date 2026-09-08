package subject

func Start(work func()) <-chan struct{} {
	done := make(chan struct{})
	work()
	close(done)
	return done
}
