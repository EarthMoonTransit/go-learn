package subject

type Record struct {
	Day      int
	Amount   float64
	Category string
}

func Filter(in []Record, predicate func(Record) bool) []Record {
	out := in[:0]
	for _, r := range in {
		if predicate(r) {
			out = append(out, r)
		}
	}
	return out
}
