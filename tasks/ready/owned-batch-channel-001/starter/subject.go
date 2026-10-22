package subject

func Publish(values []int, ready <-chan struct{}) <-chan int {
	out := make(chan int)
	close(out)
	return out
}
