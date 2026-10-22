package subject

func Collect(in <-chan int) []int {
	out := make([]int, 0)
	for v := range in {
		if v == 0 {
			return out
		}
		out = append(out, v)
	}
	return out
}
