package subject

func RuneOffsets(text string) []int {
	offsets := make([]int, 0, len(text)+1)
	for offset := range text {
		offsets = append(offsets, offset)
	}
	return append(offsets, len(text))
}
