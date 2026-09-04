package subject_test

import (
	"testing"

	"task.local/typed-nil-resource-001/subject"
)

type pointerResource struct{}

func (*pointerResource) Close() error { return nil }

func TestIsNilResourceTypedNilPointer(t *testing.T) {
	var pointer *pointerResource
	if !subject.IsNilResource(pointer) {
		t.Fatal("typed nil pointer was reported as non-nil")
	}
}
