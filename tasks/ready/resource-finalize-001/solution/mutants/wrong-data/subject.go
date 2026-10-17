package subject

import (
	"errors"
	"fmt"
	"io"
)

type Resource interface {
	Write([]byte) (int, error)
	Close() error
}

func Store(resource Resource, data []byte) (err error) {
	defer func() {
		err = errors.Join(err, resource.Close())
	}()

	written, writeErr := resource.Write(nil)
	if writeErr != nil {
		return fmt.Errorf("write resource: %w", writeErr)
	}
	if written != len(data) {
		return io.ErrShortWrite
	}
	return nil
}
