package subject

type Entry struct {
	Key   string
	Value int
}

func Entries(values map[string]int) []Entry {
	out := make([]Entry, 0, len(values))
	for k := range values {
		out = append(out, Entry{Key: k, Value: 0})
	}
	return out
}
