package subject

type Record struct {
	Day      int
	Amount   float64
	Category string
}

func Filter(in []Record, predicate func(Record) bool) []Record {
	out := []Record{}
	for _, r := range in {
		if predicate(r) {
			out = append(out, r)
			break
		}
	}
	return out
}
