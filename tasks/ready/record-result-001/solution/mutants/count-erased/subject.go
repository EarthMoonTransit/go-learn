package subject

type Record struct {
	Label string
	Count int
}

func NewRecord(label string, count int) (*Record, error) { return &Record{Label: label, Count: 0}, nil }
