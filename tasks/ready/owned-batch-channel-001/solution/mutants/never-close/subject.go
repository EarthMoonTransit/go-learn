package subject

func Publish(values []int, ready <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		<-ready
		for _, v := range values {
			out <- v
		}

	}()
	return out
}
