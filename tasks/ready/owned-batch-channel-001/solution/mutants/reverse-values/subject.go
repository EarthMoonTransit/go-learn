package subject

func Publish(values []int, ready <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		<-ready
		for i := len(values) - 1; i >= 0; i-- {
			out <- values[i]
		}
		close(out)
	}()
	return out
}
