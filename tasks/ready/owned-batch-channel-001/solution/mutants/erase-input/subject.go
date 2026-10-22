package subject

func Publish(values []int, ready <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		<-ready
		for i, v := range values {
			out <- v
			values[i] = 0
		}
		close(out)
	}()
	return out
}
