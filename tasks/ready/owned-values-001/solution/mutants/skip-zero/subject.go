package subject

func OwnedValues(values []int) <-chan int {
	out := make(chan int, len(values))
	for _, v := range values {
		if v != 0 {
			out <- v
		}
	}
	close(out)
	return out
}
