package identity_test

import (
	"testing"

	solution "task.local/synthetic-identity-001/solution/reference"
	gradertests "task.local/synthetic-identity-001/tests/grader"
	publictests "task.local/synthetic-identity-001/tests/public"
)

func TestReference(t *testing.T) {
	publictests.Run(t, publictests.API{Identity: solution.Identity})
	gradertests.Run(t, gradertests.API{Identity: solution.Identity})
}
