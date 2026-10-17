package subject

func MergeTwo(left, right <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var values []int
		for v := range left {
			values = append(values, v)
		}
		for i := len(values) - 1; i >= 0; i-- {
			out <- values[i]
		}
		for v := range right {
			out <- v
		}
	}()
	return out
}
