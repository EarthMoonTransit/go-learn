package subject

func Collect(in <-chan int) []int {
	out := make([]int, 0)
	for v := range in {
		seen := false
		for _, x := range out {
			if x == v {
				seen = true
			}
		}
		if !seen {
			out = append(out, v)
		}
	}
	return out
}
