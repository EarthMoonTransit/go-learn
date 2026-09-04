package pipeline_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"task.local/stream-normalize-001/pipeline"
)

var errInput = errors.New("input failed")
var errOutput = errors.New("output failed")

type dataErrorReader struct {
	done bool
}

func (reader *dataErrorReader) Read(buffer []byte) (int, error) {
	if reader.done {
		return 0, io.EOF
	}
	reader.done = true
	return copy(buffer, " first \n second \n"), errInput
}

type errorWriter struct{}

func (errorWriter) Write([]byte) (int, error) {
	return 0, errOutput
}

func TestRunPreservesOutputBeforeReadError(t *testing.T) {
	var output bytes.Buffer
	err := pipeline.Run(&dataErrorReader{}, &output)
	if !errors.Is(err, errInput) {
		t.Fatalf("Run returned %v; want input error", err)
	}
	if got := output.String(); got != "FIRST\nSECOND\n" {
		t.Fatalf("partial output = %q", got)
	}
}

func TestRunReturnsWriteError(t *testing.T) {
	err := pipeline.Run(bytes.NewBufferString("line\n"), errorWriter{})
	if !errors.Is(err, errOutput) {
		t.Fatalf("Run returned %v; want output error", err)
	}
}
