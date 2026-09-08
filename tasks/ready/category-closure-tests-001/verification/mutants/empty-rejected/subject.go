package subject

type Record struct {
	Day      int
	Amount   float64
	Category string
}

func InCategory(category string) func(Record) bool {
	return func(r Record) bool { return category != "" && r.Category == category }
}
