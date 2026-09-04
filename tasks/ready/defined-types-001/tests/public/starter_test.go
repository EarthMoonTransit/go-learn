package publictests_test

import (
	"testing"

	starter "go-practice-bank/tasks/defined-types-001/starter"
	publictests "go-practice-bank/tasks/defined-types-001/tests/public"
)

func TestStarter(t *testing.T) {
	publictests.Run(t, publictests.API{
		SamplesZero: starter.Samples(0),
		BytesZero:   starter.Bytes(0),
		EncodedSize: func(count, bytesPerSample int) int {
			return int(starter.EncodedSize(starter.Samples(count), bytesPerSample))
		},
		Combine: func(left, right int) int {
			return int(starter.Combine(starter.Samples(left), starter.Samples(right)))
		},
	})
}
