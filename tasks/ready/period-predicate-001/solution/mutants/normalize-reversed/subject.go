package subject

type Record struct {
	Day      int
	Amount   float64
	Category string
}
type DaysPeriod struct{ From, To int }

func WithinDays(p DaysPeriod) func(Record) bool {
	if p.From > p.To {
		p.From, p.To = p.To, p.From
	}
	return func(r Record) bool { return r.Day >= p.From && r.Day <= p.To }
}
