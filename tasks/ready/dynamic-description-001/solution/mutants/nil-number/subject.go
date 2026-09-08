package subject

import (
	"fmt"
	"strconv"
)

type NumberBox interface{ Number() int }
type FancyNumberBox interface{ Value() string }
type FancyNumber struct{ Text string }

func (n FancyNumber) Value() string { return n.Text }
func Describe(v any) string {
	switch n := v.(type) {
	case nil:
		return "number 0.0"
	case int:
		return fmt.Sprintf("number %.1f", float64(n))
	case float64:
		return fmt.Sprintf("number %.1f", n)
	case NumberBox:
		return fmt.Sprintf("box %.1f", float64(n.Number()))
	case FancyNumberBox:
		value := 0
		if f, ok := n.(FancyNumber); ok {
			if parsed, e := strconv.Atoi(f.Text); e == nil {
				value = parsed
			}
		}
		return fmt.Sprintf("fancy %.1f", float64(value))
	default:
		return "unsupported"
	}
}
