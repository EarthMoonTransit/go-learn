package subject

func PrefixRunes(text string, limit int) string {
	if limit <= 0 {
		return ""
	}
	if limit >= len(text) {
		return text
	}
	return text[:limit]
}
