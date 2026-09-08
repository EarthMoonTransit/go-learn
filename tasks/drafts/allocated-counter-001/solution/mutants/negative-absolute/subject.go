package subject

func NewCounter(initial int) *int {
	if initial < 0 {
		initial = -initial
	}
	return &initial
}
