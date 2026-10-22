package subject

func Collect(in <-chan int) []int {
	out := make([]int, 0)
	for v := range in {
		out = append(out, v)
	}
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}
