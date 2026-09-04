package subject

import (
	"fmt"
	"io"
)

type Resource interface {
	Write([]byte) (int, error)
	Close() error
}

func Store(resource Resource, data []byte) (err error) {
	defer func() {
		if closeErr := resource.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	written, writeErr := resource.Write(data)
	if writeErr != nil {
		return fmt.Errorf("write resource: %w", writeErr)
	}
	if written != len(data) {
		return io.ErrShortWrite
	}
	return nil
}
