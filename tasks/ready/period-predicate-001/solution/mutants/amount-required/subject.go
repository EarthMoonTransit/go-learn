package subject

type Record struct {
	Day      int
	Amount   float64
	Category string
}
type DaysPeriod struct{ From, To int }

func WithinDays(p DaysPeriod) func(Record) bool {
	return func(r Record) bool { return r.Amount > 0 && r.Day >= p.From && r.Day <= p.To }
}
