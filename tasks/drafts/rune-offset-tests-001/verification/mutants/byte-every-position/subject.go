package subject

func RuneOffsets(text string) []int {
	offsets := make([]int, len(text)+1)
	for index := range offsets {
		offsets[index] = index
	}
	return offsets
}
