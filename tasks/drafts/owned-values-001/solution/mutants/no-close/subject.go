package subject

func OwnedValues(values []int) <-chan int {
	out := make(chan int, len(values))
	for _, v := range values {
		out <- v
	}
	return out
}
