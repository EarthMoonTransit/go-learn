package subject

import "fmt"

type Label struct{ Text string }
type Counter struct{ Value int }

func (l Label) String() string { return "label=" + l.Text }
func (c *Counter) String() string {
	value := c.Value
	if value < 0 {
		value = -value
	}
	return fmt.Sprintf("count=%d", value)
}
func Bind(label Label, counter *Counter) (fmt.Stringer, fmt.Stringer) {
	if counter == nil {
		return label, nil
	}
	return label, counter
}
