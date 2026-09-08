package subject

type Entry struct {
	Key   string
	Value int
}

func Entries(values map[string]int) []Entry {
	out := make([]Entry, 0, len(values))
	for k, v := range values {
		out = append(out, Entry{Key: k, Value: v})
	}
	if len(out) > 1 {
		out[1] = out[0]
	}
	return out
}
