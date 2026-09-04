package subject

func RuneOffsets(text string) []int {
	count := len([]rune(text))
	offsets := make([]int, count+1)
	for index := range offsets {
		offsets[index] = index
	}
	return offsets
}
