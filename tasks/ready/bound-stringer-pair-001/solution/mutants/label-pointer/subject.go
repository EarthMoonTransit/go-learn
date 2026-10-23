package subject

import "fmt"

type Label struct{ Text string }
type Counter struct{ Value int }

func (l Label) String() string    { return "label=" + l.Text }
func (c *Counter) String() string { return fmt.Sprintf("count=%d", c.Value) }
func Bind(label Label, counter *Counter) (fmt.Stringer, fmt.Stringer) {
	if counter == nil {
		return &label, nil
	}
	return &label, counter
}
