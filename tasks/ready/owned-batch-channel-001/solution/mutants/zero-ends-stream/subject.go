package subject

func Publish(values []int, ready <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		<-ready
		for _, v := range values {
			if v == 0 {
				break
			}
			out <- v
		}
		close(out)
	}()
	return out
}
