package gradertests_test

import (
	"testing"

	starter "go-practice-bank/tasks/zero-state-001/starter"
	gradertests "go-practice-bank/tasks/zero-state-001/tests/grader"
)

func TestStarter(t *testing.T) {
	gradertests.Run(t, gradertests.API{IsZero: func(state gradertests.State) bool {
		return starter.IsZero(starter.BufferState{
			Count: state.Count, Enabled: state.Enabled, Label: state.Label, Data: state.Data,
		})
	}})
}
