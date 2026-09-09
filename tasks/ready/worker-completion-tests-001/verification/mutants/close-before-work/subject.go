package subject

func Start(work func()) <-chan struct{} {
	done := make(chan struct{})
	go func() { close(done); work() }()
	return done
}
