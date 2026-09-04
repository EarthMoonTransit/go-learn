package zerostate_test

import (
	"testing"

	solution "go-practice-bank/tasks/zero-state-001/solution/reference"
	gradertests "go-practice-bank/tasks/zero-state-001/tests/grader"
	publictests "go-practice-bank/tasks/zero-state-001/tests/public"
)

func TestReference(t *testing.T) {
	publictests.Run(t, publictests.API{
		InitialState: func() publictests.State {
			state := solution.InitialState()
			return publictests.State{Count: state.Count, Enabled: state.Enabled, Label: state.Label, Data: state.Data}
		},
		IsZero: func(state publictests.State) bool {
			return solution.IsZero(solution.BufferState{
				Count: state.Count, Enabled: state.Enabled, Label: state.Label, Data: state.Data,
			})
		},
	})
	gradertests.Run(t, gradertests.API{IsZero: func(state gradertests.State) bool {
		return solution.IsZero(solution.BufferState{
			Count: state.Count, Enabled: state.Enabled, Label: state.Label, Data: state.Data,
		})
	}})
}
