package subject

import (
	"strconv"
)

type FancyNumber struct{ Text string }

func (n FancyNumber) Value() string { return n.Text }

type FancyNumberBox interface{ Value() string }

func Extract(box FancyNumberBox) int {
	if p, ok := box.(*FancyNumber); ok && p != nil {
		v, _ := strconv.Atoi(p.Text)
		return v
	}
	n, ok := box.(FancyNumber)
	if !ok {
		return 0
	}
	value, err := strconv.Atoi(n.Text)
	if err != nil {
		return 0
	}
	return value
}
