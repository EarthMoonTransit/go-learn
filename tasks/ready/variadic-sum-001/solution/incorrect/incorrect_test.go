package variadicsum_test

import (
	"testing"

	solution "go-practice-bank/tasks/variadic-sum-001/solution/incorrect"
	gradertests "go-practice-bank/tasks/variadic-sum-001/tests/grader"
)

func TestIncorrectAgainstGrader(t *testing.T) {
	gradertests.Run(t, gradertests.API{Sum: solution.Sum, SumSlice: solution.SumSlice})
}
