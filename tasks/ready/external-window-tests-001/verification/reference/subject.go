package subject

type Window struct{ min, max int }

func NewWindow(min, max int) Window      { return Window{min: min, max: max} }
func (w Window) Contains(value int) bool { return value >= w.min && value < w.max }
