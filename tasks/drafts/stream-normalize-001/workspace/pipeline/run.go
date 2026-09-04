package pipeline

import (
	"io"

	"task.local/stream-normalize-001/transform"
)

func Run(reader io.Reader, writer io.Writer) error {
	return transform.NormalizeLines(reader, writer)
}
