package subject

type Counter struct{ Value int }

func (c Counter) ReadValue() int    { return c.Value }
func (c *Counter) ReadPointer() int { return c.Value }

func Capture(c *Counter) (func() int, func() int) {
	return nil, nil
}
