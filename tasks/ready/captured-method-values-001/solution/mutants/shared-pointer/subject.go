package subject

type Counter struct{ Value int }

func (c Counter) ReadValue() int    { return c.Value }
func (c *Counter) ReadPointer() int { return c.Value }

var latest *Counter

func Capture(c *Counter) (func() int, func() int) {
	saved := c.Value
	latest = c
	return func() int { return saved }, func() int { return latest.Value }
}
