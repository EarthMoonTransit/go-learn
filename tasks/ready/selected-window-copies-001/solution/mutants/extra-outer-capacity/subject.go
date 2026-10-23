package subject

type Window struct{ Start, End int }

func SelectedCopies(input []int, windows []Window) [][]int {
	count := 0
	for _, window := range windows {
		if window.Start < window.End {
			count++
		}
	}
	out := make([][]int, 0, count+1)
	for _, window := range windows {
		if window.Start == window.End {
			continue
		}
		n := window.End - window.Start
		owned := make([]int, n)
		copy(owned, input[window.Start:window.End])
		out = append(out, owned)
	}
	return out
}
