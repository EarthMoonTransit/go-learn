package gradertests_test

import (
	"testing"

	starter "go-practice-bank/tasks/variadic-sum-001/starter"
	gradertests "go-practice-bank/tasks/variadic-sum-001/tests/grader"
)

func TestStarter(t *testing.T) {
	gradertests.Run(t, gradertests.API{Sum: starter.Sum, SumSlice: starter.SumSlice})
}
