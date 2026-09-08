package subject

import (
	"testing"
)

type Case struct {
	Name  string
	Check func(*testing.T)
}

func RunNamed(t *testing.T, cases []Case) {
	for _, c := range cases {
		t.Run(c.Name, c.Check)
	}
}
