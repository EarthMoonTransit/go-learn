package subject

type Window struct{ Start, End int }

func SelectedCopies(input []int, windows []Window) [][]int {
	count := 0
	for range windows {
		count++
	}
	out := make([][]int, 0, count)
	for _, window := range windows {

		n := window.End - window.Start
		owned := make([]int, n)
		copy(owned, input[window.Start:window.End])
		out = append(out, owned)
	}
	return out
}
