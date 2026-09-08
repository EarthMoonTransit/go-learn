package subject

import (
	"strings"
)

func Application(log string) string {
	for _, m := range []rune{'❗', '🔍', '☀'} {
		if strings.ContainsRune(log, m) {
			switch m {
			case '❗':
				return "recommendation"
			case '🔍':
				return "search"
			default:
				return "weather"
			}
		}
	}
	return "default"
}
