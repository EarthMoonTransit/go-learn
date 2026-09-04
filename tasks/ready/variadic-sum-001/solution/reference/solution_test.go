package variadicsum_test

import (
	"testing"

	solution "go-practice-bank/tasks/variadic-sum-001/solution/reference"
	gradertests "go-practice-bank/tasks/variadic-sum-001/tests/grader"
	publictests "go-practice-bank/tasks/variadic-sum-001/tests/public"
)

func TestReference(t *testing.T) {
	publictests.Run(t, publictests.API{Sum: solution.Sum, SumSlice: solution.SumSlice})
	gradertests.Run(t, gradertests.API{Sum: solution.Sum, SumSlice: solution.SumSlice})
}
