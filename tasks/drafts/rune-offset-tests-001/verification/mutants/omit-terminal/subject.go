package subject

func RuneOffsets(text string) []int {
	offsets := make([]int, 0, len(text))
	for offset := range text {
		offsets = append(offsets, offset)
	}
	return offsets
}
