package publictests_test

import (
	"testing"

	starter "go-practice-bank/tasks/variadic-sum-001/starter"
	publictests "go-practice-bank/tasks/variadic-sum-001/tests/public"
)

func TestStarter(t *testing.T) {
	publictests.Run(t, publictests.API{Sum: starter.Sum, SumSlice: starter.SumSlice})
}
