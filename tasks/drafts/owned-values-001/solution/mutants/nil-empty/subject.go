package subject

func OwnedValues(values []int) <-chan int {
	if len(values) == 0 {
		return nil
	}
	out := make(chan int, len(values))
	for _, v := range values {
		out <- v
	}
	close(out)
	return out
}
