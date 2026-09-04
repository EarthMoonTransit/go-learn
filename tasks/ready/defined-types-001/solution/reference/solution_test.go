package definedtypes_test

import (
	"testing"

	solution "go-practice-bank/tasks/defined-types-001/solution/reference"
	gradertests "go-practice-bank/tasks/defined-types-001/tests/grader"
	publictests "go-practice-bank/tasks/defined-types-001/tests/public"
)

func TestReference(t *testing.T) {
	publictests.Run(t, publictests.API{
		SamplesZero: solution.Samples(0),
		BytesZero:   solution.Bytes(0),
		EncodedSize: func(count, bytesPerSample int) int {
			return int(solution.EncodedSize(solution.Samples(count), bytesPerSample))
		},
		Combine: func(left, right int) int {
			return int(solution.Combine(solution.Samples(left), solution.Samples(right)))
		},
	})
	gradertests.Run(t, gradertests.API{
		SamplesZero: solution.Samples(0),
		BytesZero:   solution.Bytes(0),
		EncodedSize: func(count, bytesPerSample int) int {
			return int(solution.EncodedSize(solution.Samples(count), bytesPerSample))
		},
		Combine: func(left, right int) int {
			return int(solution.Combine(solution.Samples(left), solution.Samples(right)))
		},
	})
}
