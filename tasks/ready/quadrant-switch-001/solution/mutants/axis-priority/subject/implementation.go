package subject

func Locate(x, y int) string {
	switch {
	case y == 0 && x > 0:
		return "positive-x"
	case y == 0 && x < 0:
		return "negative-x"
	case x == 0 && y > 0:
		return "positive-y"
	case x == 0 && y < 0:
		return "negative-y"
	case x > 0 && y > 0:
		return "quadrant-i"
	case x < 0 && y > 0:
		return "quadrant-ii"
	case x < 0 && y < 0:
		return "quadrant-iii"
	default:
		return "quadrant-iv"
	}
}
