package subject

type Entry struct {
	Key   string
	Value int
}

func Entries(values map[string]int) []Entry {
	out := make([]Entry, 0, len(values))
	for k, v := range values {
		if v > 0 {
			out = append(out, Entry{Key: k, Value: v})
		}
	}
	return out
}
