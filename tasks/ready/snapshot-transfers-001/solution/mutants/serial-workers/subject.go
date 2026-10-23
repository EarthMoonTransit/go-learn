package subject

type Transfer struct {
	Source      <-chan int
	Destination chan<- int
	Offset      int
}

func StartTransfers(transfers []Transfer, start <-chan struct{}) <-chan struct{} {
	done := make(chan struct{})
	captured := append([]Transfer(nil), transfers...)
	go func() {
		defer close(done)
		for _, item := range captured {
			<-start
			value, ok := <-item.Source
			if ok {
				item.Destination <- value + item.Offset
			}
		}
	}()
	return done
}
