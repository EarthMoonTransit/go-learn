package subject

func Start(work func()) <-chan struct{} {
	done := make(chan struct{})
	go func() { work() }()
	return done
}
