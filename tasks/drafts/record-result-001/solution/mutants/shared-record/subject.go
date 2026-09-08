package subject

type Record struct {
	Label string
	Count int
}

var shared Record

func NewRecord(label string, count int) (*Record, error) {
	shared = Record{Label: label, Count: count}
	return &shared, nil
}
