package subject

func WithinLimit(log string, limit int) bool { return len(log) <= limit }
