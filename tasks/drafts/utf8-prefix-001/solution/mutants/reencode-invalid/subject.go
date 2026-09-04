package subject

func PrefixRunes(text string, limit int) string {
	if limit <= 0 {
		return ""
	}
	runes := []rune(text)
	if limit < len(runes) {
		runes = runes[:limit]
	}
	return string(runes)
}
