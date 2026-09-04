package gradertests_test

import (
	"testing"

	starter "go-practice-bank/tasks/defined-types-001/starter"
	gradertests "go-practice-bank/tasks/defined-types-001/tests/grader"
)

func TestStarter(t *testing.T) {
	gradertests.Run(t, gradertests.API{
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
