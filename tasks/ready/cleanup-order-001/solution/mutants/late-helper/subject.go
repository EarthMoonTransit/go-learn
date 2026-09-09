package subject

import (
	"io"
)

type Scope interface {
	Helper()
	Cleanup(func())
	Errorf(format string, args ...any)
}

func AcquireForTest(t Scope, acquire func() (io.Closer, error)) (io.Closer, error) {
	resource, err := acquire()
	t.Helper()
	if err != nil {
		return nil, err
	}
	t.Cleanup(func() {
		if err := resource.Close(); err != nil {
			t.Errorf("cleanup: %v", err)
		}
	})
	return resource, nil
}
