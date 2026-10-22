package subject

func Collect(in <-chan int) []int {
	out := []int{}
	if v, ok := <-in; ok {
		out = append(out, v)
	}
	return out
}
