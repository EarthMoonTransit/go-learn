package subject

type Window struct{ Start, End int }

func SelectedCopies(input []int, windows []Window) [][]int {
	count := 0
	for _, window := range windows {
		if window.Start < window.End {
			count++
		}
	}
	out := make([][]int, 0, count)
	for _, window := range windows {
		if window.Start == window.End {
			continue
		}
		n := window.End - window.Start
		owned := make([]int, n)
		for i := range owned {
			owned[i] = input[window.End-1-i]
		}
		out = append(out, owned)
	}
	return out
}
