package subject

func MergeTwo(left, right <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range left {
			out <- v
		}
		for v := range right {
			out <- v
		}
	}()
	return out
}
