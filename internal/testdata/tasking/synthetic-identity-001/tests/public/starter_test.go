package publictests_test

import (
	"testing"

	starter "task.local/synthetic-identity-001/starter"
	publictests "task.local/synthetic-identity-001/tests/public"
)

func TestStarter(t *testing.T) {
	publictests.Run(t, publictests.API{Identity: starter.Identity})
}
