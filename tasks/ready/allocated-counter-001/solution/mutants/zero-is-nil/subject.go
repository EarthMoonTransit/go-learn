package subject

func NewCounter(initial int) *int {
	if initial == 0 {
		return nil
	}
	return &initial
}
