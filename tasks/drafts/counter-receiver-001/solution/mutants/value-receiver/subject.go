package subject

type Counter struct {
	Value int
}

type counterAdder interface {
	add(int)
}

func (c *Counter) Add(delta int) {
	var target counterAdder = c
	target.add(delta)
}
func (c Counter) add(delta int) {
	c.Value += delta
}
