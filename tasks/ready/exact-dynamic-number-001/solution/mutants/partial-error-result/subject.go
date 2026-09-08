package subject

import (
	"strconv"
)

type FancyNumber struct{ Text string }

func (n FancyNumber) Value() string { return n.Text }

type FancyNumberBox interface{ Value() string }

func Extract(box FancyNumberBox) int {
	n, ok := box.(FancyNumber)
	if !ok {
		return 0
	}
	value, err := strconv.Atoi(n.Text)
	if err != nil {
		return value
	}
	return value
}
