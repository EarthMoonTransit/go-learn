package gradertests

import "testing"

type API struct {
	Identity func(int) int
}

func Run(t *testing.T, api API) {
	t.Helper()
	for _, value := range []int{0, -9, 42} {
		if got := api.Identity(value); got != value {
			t.Errorf("Identity(%d) = %d, want %d", value, got, value)
		}
	}
}
