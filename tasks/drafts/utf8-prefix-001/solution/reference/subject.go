package subject

import "unicode/utf8"

func PrefixRunes(text string, limit int) string {
	if limit <= 0 {
		return ""
	}
	offset := 0
	for count := 0; count < limit && offset < len(text); count++ {
		_, size := utf8.DecodeRuneInString(text[offset:])
		offset += size
	}
	return text[:offset]
}
