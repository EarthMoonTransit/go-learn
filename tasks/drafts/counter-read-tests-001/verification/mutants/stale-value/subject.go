package subject

var saved int
var seen bool

func CounterValue(counter *int) int {
	if counter == nil {
		return 0
	}
	if !seen {
		saved = *counter
		seen = true
	}
	return saved
}
