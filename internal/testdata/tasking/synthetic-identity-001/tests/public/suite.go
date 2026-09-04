package publictests

import "testing"

type API struct {
	Identity func(int) int
}

func Run(t *testing.T, api API) {
	t.Helper()
	if got := api.Identity(7); got != 7 {
		t.Errorf("Identity(7) = %d, want 7", got)
	}
}
