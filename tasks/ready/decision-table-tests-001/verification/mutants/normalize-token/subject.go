package subject

import (
	"strings"
)

func tokenValue(token string) int {
	switch strings.ToLower(strings.TrimSpace(token)) {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	}
	return 0
}
func SelectAction(first, second, visible string) string {
	sum := tokenValue(first) + tokenValue(second)
	v := tokenValue(visible)
	if sum == 22 {
		return "P"
	}
	if sum == 21 {
		if v >= 10 {
			return "S"
		}
		return "W"
	}
	if sum >= 17 {
		return "S"
	}
	if sum >= 12 && v < 7 {
		return "S"
	}
	return "H"
}
