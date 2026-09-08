package subject

type Record struct {
	Day      int
	Amount   float64
	Category string
}
type DaysPeriod struct{ From, To int }

var current DaysPeriod

func WithinDays(p DaysPeriod) func(Record) bool {
	current = p
	return func(r Record) bool { return r.Day >= current.From && r.Day <= current.To }
}
