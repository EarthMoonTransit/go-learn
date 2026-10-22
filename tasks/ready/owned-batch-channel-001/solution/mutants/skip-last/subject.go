package subject

func Publish(values []int, ready <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		<-ready
		for i, v := range values {
			if i+1 < len(values) {
				out <- v
			}
		}
		close(out)
	}()
	return out
}
