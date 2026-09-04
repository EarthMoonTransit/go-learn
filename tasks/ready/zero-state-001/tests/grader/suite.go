package gradertests

import "testing"

type State struct {
	Count   int
	Enabled bool
	Label   string
	Data    []byte
}

type API struct {
	IsZero func(State) bool
}

func Run(t *testing.T, api API) {
	t.Helper()
	tests := []struct {
		name  string
		state State
		want  bool
	}{
		{name: "literal zero", state: State{}, want: true},
		{name: "enabled", state: State{Enabled: true}, want: false},
		{name: "label", state: State{Label: "set"}, want: false},
		{name: "data", state: State{Data: []byte{0}}, want: false},
		{name: "allocated empty data", state: State{Data: make([]byte, 0, 8)}, want: false},
	}
	for _, test := range tests {
		if got := api.IsZero(test.state); got != test.want {
			t.Errorf("%s: got %t, want %t", test.name, got, test.want)
		}
	}
}
