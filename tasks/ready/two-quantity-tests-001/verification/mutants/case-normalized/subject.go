package subject

import (
	"strings"
)

func Quantities(kinds []string) (int, float64) {
	a, b := 0, 0
	for _, k := range kinds {
		k = strings.ToLower(k)
		if k == "solid" {
			a++
		}
		if k == "liquid" {
			b++
		}
	}
	return a * 50, float64(b) * 0.2
}
