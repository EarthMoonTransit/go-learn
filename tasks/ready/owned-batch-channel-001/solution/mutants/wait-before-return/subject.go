package subject

func Publish(values []int, ready <-chan struct{}) <-chan int {
	out := make(chan int, len(values))
	func() {
		<-ready
		for _, v := range values {
			out <- v
		}
		close(out)
	}()
	return out
}
