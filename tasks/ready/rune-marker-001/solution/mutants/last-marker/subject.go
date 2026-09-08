package subject

func Application(log string) string {
	result := "default"
	for _, r := range log {
		switch r {
		case '❗':
			result = "recommendation"
		case '🔍':
			result = "search"
		case '☀':
			result = "weather"
		}
	}
	return result
}
