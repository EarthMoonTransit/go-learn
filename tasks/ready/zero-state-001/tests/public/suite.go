package publictests

import "testing"

type State struct {
	Count   int
	Enabled bool
	Label   string
	Data    []byte
}

type API struct {
	InitialState func() State
	IsZero       func(State) bool
}

func Run(t *testing.T, api API) {
	t.Helper()
	initial := api.InitialState()
	if !api.IsZero(initial) {
		t.Errorf("InitialState returned a non-zero state: %+v", initial)
	}
	if api.IsZero(State{Count: 1}) {
		t.Error("state with Count=1 was reported as zero")
	}
	if api.IsZero(State{Data: []byte{}}) {
		t.Error("allocated empty Data was reported as a nil zero value")
	}
}
