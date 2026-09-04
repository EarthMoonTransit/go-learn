package subject

type Counter struct {
	Value int
}

type counterAdder interface {
	add(int)
}

func (c *Counter) Add(delta int) {
	copied := *c
	var target counterAdder = &copied
	target.add(delta)
}
func (c *Counter) add(delta int) {
	c.Value += delta
}
