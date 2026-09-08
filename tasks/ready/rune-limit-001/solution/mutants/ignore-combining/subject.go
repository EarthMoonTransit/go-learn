package subject

func WithinLimit(log string, limit int) bool {
	n := 0
	for _, r := range log {
		if r < 0x300 || r > 0x36f {
			n++
		}
	}
	return n <= limit
}
