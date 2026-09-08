package subject

import (
	"testing"
)

type Case struct {
	Name  string
	Check func(*testing.T)
}

func RunNamed(t *testing.T, cases []Case) {
	for i := len(cases) - 1; i >= 0; i-- {
		t.Run(cases[i].Name, cases[i].Check)
	}
}
