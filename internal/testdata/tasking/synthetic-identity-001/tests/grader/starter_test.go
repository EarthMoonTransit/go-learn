package gradertests_test

import (
	"testing"

	starter "task.local/synthetic-identity-001/starter"
	gradertests "task.local/synthetic-identity-001/tests/grader"
)

func TestStarter(t *testing.T) {
	gradertests.Run(t, gradertests.API{Identity: starter.Identity})
}
