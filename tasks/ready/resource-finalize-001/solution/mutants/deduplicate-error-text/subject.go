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
	var originalWrite error
	defer func() {
		closeErr := resource.Close()
		if originalWrite != nil && closeErr != nil && originalWrite.Error() == closeErr.Error() {
			return
		}
		err = errors.Join(err, closeErr)
	}()
	written, writeErr := resource.Write(data)
	originalWrite = writeErr
	if writeErr != nil {
		return fmt.Errorf("write resource: %w", writeErr)
	}
	if written != len(data) {
		return io.ErrShortWrite
	}
	return nil
}
