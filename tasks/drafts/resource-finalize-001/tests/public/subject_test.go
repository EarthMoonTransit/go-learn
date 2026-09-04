package subject_test

import (
	"errors"
	"testing"

	"task.local/resource-finalize-001/subject"
)

var errWrite = errors.New("write failed")

type failingResource struct {
	closeCalls int
}

func (resource *failingResource) Write([]byte) (int, error) {
	return 0, errWrite
}

func (resource *failingResource) Close() error {
	resource.closeCalls++
	return nil
}

func TestStoreClosesAfterWriteFailure(t *testing.T) {
	resource := &failingResource{}
	err := subject.Store(resource, []byte("data"))
	if !errors.Is(err, errWrite) {
		t.Fatalf("Store returned %v; want write failure", err)
	}
	if resource.closeCalls != 1 {
		t.Fatalf("Close calls = %d; want 1", resource.closeCalls)
	}
}
