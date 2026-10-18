package fuzznative

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestEmbeddedHelperIsStandaloneAndDigestBound(t *testing.T) {
	if digest := SourceDigest(); len(digest) != 64 || digest != computeSourceDigest() {
		t.Fatalf("unstable source digest %q", digest)
	}
	work := t.TempDir()
	if err := materializeHelper(work, "1.27"); err != nil {
		t.Fatal(err)
	}
	files := 0
	err := filepath.WalkDir(work, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		files++
		relative, err := filepath.Rel(work, path)
		if err != nil {
			return err
		}
		if relative == "go.mod" {
			return nil
		}
		prefix := "internal/fuzznative/"
		slash := filepath.ToSlash(relative)
		if !strings.HasPrefix(slash, prefix) || strings.HasSuffix(slash, "_test.go") {
			t.Fatalf("unexpected helper material %s", relative)
		}
		name := strings.TrimPrefix(slash, prefix)
		if !strings.HasPrefix(name, "frame/") && !strings.HasPrefix(name, "protocol/") && !strings.HasPrefix(name, "worker/") {
			t.Fatalf("host/observer leaked into helper: %s", name)
		}
		original, err := sourceFiles.ReadFile(name)
		if err != nil {
			return err
		}
		copied, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if string(original) != string(copied) {
			t.Fatalf("helper copy differs from canonical source: %s", name)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if files < 10 {
		t.Fatalf("incomplete trusted helper: %d files", files)
	}
}

func TestInvalidRequestCannotStartAnyProcess(t *testing.T) {
	engine := NewEngine(nil)
	request := Request{Identity: Identity{ModulePath: "task.local/request-hex-001", Grammar: "hex-codec-v1", Target: "FuzzHex"}, GoVersion: "1.27", Timeout: time.Second, Implementation: map[string]string{"hex.go": "package subject"}, Sources: map[string]string{"fuzz_test.go": "package subject_test"}}
	for _, mutate := range []func(*Request){
		func(r *Request) { r.Identity.ModulePath = "task.local/../escape" },
		func(r *Request) { r.Identity.Target = "FuzzOther" },
		func(r *Request) { r.GoVersion = "1.28" },
		func(r *Request) { r.Timeout = 0 },
		func(r *Request) { r.Sources = map[string]string{"../fuzz_test.go": "package subject_test"} },
		func(r *Request) { r.Sources = map[string]string{"fuzz_test.go": strings.Repeat("a", 256<<10+1)} },
	} {
		value := request
		mutate(&value)
		result := engine.Verify(value)
		if result.Valid || result.Detected || result.Phase != "admission" || result.Diagnostic == "" {
			t.Fatalf("invalid request passed: %+v", result)
		}
	}
	if metrics := engine.Metrics(); metrics != (Metrics{}) {
		t.Fatalf("invalid admission started processes: %+v", metrics)
	}
}
