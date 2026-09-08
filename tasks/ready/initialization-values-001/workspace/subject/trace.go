package subject

var events []string

func record(step string, value int) int { events = append(events, step); return value }
func Trace() []string                   { return append([]string(nil), events...) }
