package subject

func Start(work func()) <-chan struct{} {
	done := make(chan struct{})
	go func() { defer close(done) }()
	return done
}
