package subject_test

import (
	"errors"
	"io"
	"slices"
	"testing"

	"task.local/resource-finalize-001/subject"
)

var errClose = errors.New("close failed")

type recordingResource struct {
	written   int
	writeErr  error
	closeErr  error
	events    []string
	writeData []byte
}

func (resource *recordingResource) Write(data []byte) (int, error) {
	resource.events = append(resource.events, "write")
	resource.writeData = append([]byte(nil), data...)
	return resource.written, resource.writeErr
}

func (resource *recordingResource) Close() error {
	resource.events = append(resource.events, "close")
	return resource.closeErr
}

func TestStoreSuccessAndShortWrite(t *testing.T) {
	success := &recordingResource{written: 4}
	if err := subject.Store(success, []byte("data")); err != nil {
		t.Fatalf("success returned %v", err)
	}
	if !slices.Equal(success.events, []string{"write", "close"}) || !slices.Equal(success.writeData, []byte("data")) {
		t.Fatalf("success events/data = %v, %q", success.events, success.writeData)
	}

	short := &recordingResource{written: 2}
	if err := subject.Store(short, []byte("data")); !errors.Is(err, io.ErrShortWrite) {
		t.Fatalf("short write returned %v", err)
	}
	if !slices.Equal(short.events, []string{"write", "close"}) {
		t.Fatalf("short write events = %v", short.events)
	}
}

func TestStorePreservesCloseAndWriteErrors(t *testing.T) {
	closeOnly := &recordingResource{written: 4, closeErr: errClose}
	if err := subject.Store(closeOnly, []byte("data")); !errors.Is(err, errClose) {
		t.Fatalf("close failure returned %v", err)
	}

	both := &recordingResource{writeErr: errWrite, closeErr: errClose}
	err := subject.Store(both, []byte("data"))
	if !errors.Is(err, errWrite) || !errors.Is(err, errClose) {
		t.Fatalf("combined failure returned %v", err)
	}
	if !slices.Equal(both.events, []string{"write", "close"}) {
		t.Fatalf("combined failure events = %v", both.events)
	}
}
