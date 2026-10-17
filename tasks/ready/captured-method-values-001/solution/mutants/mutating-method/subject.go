package subject

type Counter struct{ Value int }

func (c Counter) ReadValue() int    { return c.Value }
func (c *Counter) ReadPointer() int { value := c.Value; c.Value++; return value }

func Capture(c *Counter) (func() int, func() int) {
	return c.ReadValue, c.ReadPointer
}
