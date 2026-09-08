package subject

import (
	"testing"
)

type Case struct {
	Name  string
	Check func(*testing.T)
}

func RunNamed(t *testing.T, cases []Case) {
	for i, c := range cases {
		if i+1 < len(cases) {
			t.Run(c.Name, c.Check)
		}
	}
}
