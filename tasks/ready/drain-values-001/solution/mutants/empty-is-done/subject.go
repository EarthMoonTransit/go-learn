package subject

func Collect(in <-chan int) []int {
	out := []int{}
	for {
		select {
		case v, ok := <-in:
			if !ok {
				return out
			}
			out = append(out, v)
		default:
			return out
		}
	}
}
