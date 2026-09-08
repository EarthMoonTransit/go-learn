package subject

type Record struct {
	Day      int
	Amount   float64
	Category string
}

var selected string

func InCategory(category string) func(Record) bool {
	selected = category
	return func(r Record) bool { return r.Category == selected }
}
