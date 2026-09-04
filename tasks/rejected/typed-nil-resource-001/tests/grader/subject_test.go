package subject_test

import (
	"testing"

	"task.local/typed-nil-resource-001/subject"
)

type valueResource struct{}

func (valueResource) Close() error { return nil }

type sliceResource []byte

func (sliceResource) Close() error { return nil }

func TestIsNilResourceKinds(t *testing.T) {
	var nilInterface subject.Resource
	if !subject.IsNilResource(nilInterface) {
		t.Error("nil interface reported as non-nil")
	}
	var nilSlice sliceResource
	if !subject.IsNilResource(nilSlice) {
		t.Error("typed nil slice reported as non-nil")
	}
	if subject.IsNilResource(valueResource{}) {
		t.Error("value implementation reported as nil")
	}
	pointer := &pointerResource{}
	if subject.IsNilResource(pointer) {
		t.Error("non-nil pointer reported as nil")
	}
}
