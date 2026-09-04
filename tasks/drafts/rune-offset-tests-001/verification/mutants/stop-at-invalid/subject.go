package subject

import "unicode/utf8"

func RuneOffsets(text string) []int {
	offsets := make([]int, 0, len(text)+1)
	for offset := 0; offset < len(text); {
		value, size := utf8.DecodeRuneInString(text[offset:])
		if value == utf8.RuneError {
			break
		}
		offsets = append(offsets, offset)
		offset += size
	}
	return append(offsets, len(text))
}
