package pipeline_test

import (
	"bytes"
	"strings"
	"testing"

	"task.local/stream-normalize-001/pipeline"
)

func TestRunNormalizesAndOmitsBlankLines(t *testing.T) {
	var output bytes.Buffer
	err := pipeline.Run(strings.NewReader(" a \n\n b \n"), &output)
	if err != nil {
		t.Fatalf("Run returned %v", err)
	}
	if got := output.String(); got != "A\nB\n" {
		t.Fatalf("output = %q", got)
	}
}
