package subject

func Application(log string) string {
	for _, b := range []byte(log) {
		switch rune(b) {
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}
