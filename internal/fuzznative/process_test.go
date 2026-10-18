package fuzznative

import (
	"os"
	"path/filepath"
	"testing"

	"go-practice-bank/internal/isolation"
)

func TestCaptureDrainsWithoutKeepingUnboundedOutput(t *testing.T) {
	output := &capture{limit: 4}
	for _, value := range []string{"abc", "defgh", "ij"} {
		n, err := output.Write([]byte(value))
		if err != nil || n != len(value) {
			t.Fatalf("capture did not drain process: n=%d err=%v", n, err)
		}
	}
	if string(output.b) != "abcd" || !output.truncated {
		t.Fatalf("capture exceeded bound or hid truncation: %q %v", output.b, output.truncated)
	}
}

func TestTraceReadsFailClosedAtFilesystemBoundary(t *testing.T) {
	environment, err := isolation.NewEnvironment("fuzz-trace-unit")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := environment.Close(); err != nil {
			t.Error(err)
		}
	})
	path := filepath.Join(environment.Root, "trace.jsonl")
	if _, _, err := readTrace(environment, path, 8); err == nil {
		t.Fatal("accepted missing trace")
	}
	if err := os.WriteFile(path, []byte("12345678"), 0600); err != nil {
		t.Fatal(err)
	}
	data, truncated, err := readTrace(environment, path, 8)
	if err != nil || truncated || string(data) != "12345678" {
		t.Fatalf("exact cap lost: %q %v %v", data, truncated, err)
	}
	if err := os.WriteFile(path, []byte("123456789"), 0600); err != nil {
		t.Fatal(err)
	}
	data, truncated, err = readTrace(environment, path, 8)
	if err != nil || !truncated || string(data) != "12345678" {
		t.Fatalf("over cap not marked: %q %v %v", data, truncated, err)
	}
	outside := filepath.Join(t.TempDir(), "outside")
	if err := os.WriteFile(outside, []byte("x"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, _, err := readTrace(environment, outside, 8); err == nil {
		t.Fatal("accepted outside trace")
	}
	link := filepath.Join(environment.Root, "link")
	if err := os.Symlink(outside, link); err == nil {
		if _, _, err := readTrace(environment, link, 8); err == nil {
			t.Fatal("accepted symlink trace")
		}
	}
	if _, _, err := readTrace(environment, environment.Root, 8); err == nil {
		t.Fatal("accepted directory trace")
	}
}
