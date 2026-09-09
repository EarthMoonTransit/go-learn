package subject

func OwnedValues(values []int) <-chan int {
	out := make(chan int, len(values))
	for _, v := range values {
		out <- v
	}
	if len(values) > 0 {
		values[0] = 0
	}
	close(out)
	return out
}
