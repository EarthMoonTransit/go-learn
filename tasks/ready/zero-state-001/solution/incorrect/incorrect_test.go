package zerostate_test

import (
	"testing"

	solution "go-practice-bank/tasks/zero-state-001/solution/incorrect"
	gradertests "go-practice-bank/tasks/zero-state-001/tests/grader"
)

func TestIncorrectAgainstGrader(t *testing.T) {
	gradertests.Run(t, gradertests.API{IsZero: func(state gradertests.State) bool {
		return solution.IsZero(solution.BufferState{
			Count: state.Count, Enabled: state.Enabled, Label: state.Label, Data: state.Data,
		})
	}})
}
