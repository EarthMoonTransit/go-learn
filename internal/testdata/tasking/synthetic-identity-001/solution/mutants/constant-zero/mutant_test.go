package identity_test

import (
	"testing"

	mutant "task.local/synthetic-identity-001/solution/mutants/constant-zero"
	gradertests "task.local/synthetic-identity-001/tests/grader"
	publictests "task.local/synthetic-identity-001/tests/public"
)

func TestMutant(t *testing.T) {
	publictests.Run(t, publictests.API{Identity: mutant.Identity})
	gradertests.Run(t, gradertests.API{Identity: mutant.Identity})
}
