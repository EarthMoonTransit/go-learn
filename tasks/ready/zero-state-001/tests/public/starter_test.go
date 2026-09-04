package publictests_test

import (
	"testing"

	starter "go-practice-bank/tasks/zero-state-001/starter"
	publictests "go-practice-bank/tasks/zero-state-001/tests/public"
)

func TestStarter(t *testing.T) {
	publictests.Run(t, publictests.API{
		InitialState: func() publictests.State {
			return fromStarter(starter.InitialState())
		},
		IsZero: func(state publictests.State) bool {
			return starter.IsZero(toStarter(state))
		},
	})
}

func fromStarter(state starter.BufferState) publictests.State {
	return publictests.State{Count: state.Count, Enabled: state.Enabled, Label: state.Label, Data: state.Data}
}

func toStarter(state publictests.State) starter.BufferState {
	return starter.BufferState{Count: state.Count, Enabled: state.Enabled, Label: state.Label, Data: state.Data}
}
