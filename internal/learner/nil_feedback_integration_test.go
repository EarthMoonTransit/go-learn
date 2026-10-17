//go:build taskintegration

package learner

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-practice-bank/internal/isolation"
	"go-practice-bank/internal/tasking"
)

const nilPublicAssertion = `
func TestPublicNilView(t *testing.T) {
	got := subject.CappedView(nil, 0, 0)
	if len(got) != 0 || cap(got) != 0 {
		t.Errorf("CappedView(nil, 0, 0): len=%d cap=%d; want len=0 cap=0", len(got), cap(got))
	}
}
`
const nilCapacityDefect = `package subject

func CappedView(in []int, start, end int) []int {
	if in == nil {
		return make([]int, 0, 1)
	}
	return in[start:end:end]
}
`

func TestNilViewPublicFeedbackAndCompatibility(t *testing.T) {
	const id = "bounded-slice-view-001"
	canonicalRoot, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	canonical := filepath.Join(canonicalRoot, "tasks/ready", id)
	before, err := tasking.TaskContentDigest(canonical)
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	for _, rel := range []string{"config", "schemas", "sources", "curriculum", "docs", "tasks/ready/" + id} {
		if err := isolation.CopyTree(filepath.Join(canonicalRoot, rel), filepath.Join(root, rel)); err != nil {
			t.Fatal(err)
		}
	}
	taskDir := filepath.Join(root, "tasks/ready", id)
	notes, err := os.ReadFile(filepath.Join(taskDir, "review/notes.md"))
	if err != nil {
		t.Fatal(err)
	}
	if before == boundedBeforeNilCheck && !strings.Contains(string(notes), "P4C: nil public contract") {
		manifest := filepath.Join(taskDir, "task.yaml")
		manifestBytes, err := os.ReadFile(manifest)
		if err != nil {
			t.Fatal(err)
		}
		manifestBytes = []byte(strings.Replace(string(manifestBytes), "        - prefix-cleared", "        - prefix-cleared\n        - nil-spare-cap", 1))
		if err := os.WriteFile(manifest, manifestBytes, 0600); err != nil {
			t.Fatal(err)
		}
		// P4T precedes its separate content PR. Only this disposable pre-release
		// fixture receives the projected assertion. Once P4C's review marker exists,
		// always exercise canonical public input, including omission experiments.
		public := filepath.Join(taskDir, "tests/public/public_test.go")
		data, err := os.ReadFile(public)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(public, append(data, []byte(nilPublicAssertion)...), 0600); err != nil {
			t.Fatal(err)
		}
		mutant := filepath.Join(taskDir, "solution/mutants/nil-spare-cap/subject.go")
		if err := os.MkdirAll(filepath.Dir(mutant), 0700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(mutant, []byte(nilCapacityDefect), 0600); err != nil {
			t.Fatal(err)
		}
	}
	digest, err := tasking.TaskContentDigest(taskDir)
	if err != nil {
		t.Fatal(err)
	}
	provenancePath := filepath.Join(taskDir, "provenance.yaml")
	data, err := os.ReadFile(provenancePath)
	if err != nil {
		t.Fatal(err)
	}
	var provenance map[string]any
	if err := json.Unmarshal(data, &provenance); err != nil {
		t.Fatal(err)
	}
	provenance["project"].(map[string]any)["content_digest"] = digest
	data, err = json.Marshal(provenance)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(provenancePath, data, 0600); err != nil {
		t.Fatal(err)
	}
	s, err := New(Config{Root: root, Home: t.TempDir()})
	if err != nil {
		t.Fatal(err)
	}
	// Reuse this test's disposable compiler cache across its public/full checks.
	// Each verification still assembles a fresh execution module.
	s.verify = integrationVerifier(t, root)
	workspace, err := s.Start(id)
	if err != nil {
		t.Fatal(err)
	}
	markerPath := filepath.Join(workspace, learnerWorkspaceMarker)
	markerBytes, err := os.ReadFile(markerPath)
	if err != nil {
		t.Fatal(err)
	}
	var marker WorkspaceManifest
	if err := json.Unmarshal(markerBytes, &marker); err != nil {
		t.Fatal(err)
	}
	marker.ContentDigest = boundedBeforeNilCheck
	markerBytes, err = json.Marshal(marker)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(markerPath, 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(markerPath, markerBytes, 0600); err != nil {
		t.Fatal(err)
	}
	subject := filepath.Join(workspace, "subject/subject.go")
	if err := os.WriteFile(subject, []byte(nilCapacityDefect), 0600); err != nil {
		t.Fatal(err)
	}
	// With a removed assertion, keep the current fixture marker so failure is
	// specifically unexpected public success, not an unrelated digest rejection.
	if digest != boundedWithNilCheck {
		marker.ContentDigest = digest
		markerBytes, _ = json.Marshal(marker)
		if err := os.WriteFile(markerPath, markerBytes, 0600); err != nil {
			t.Fatal(err)
		}
	}
	progress := persistedBytes(t, s)
	view, err := s.TaskView(id)
	if err != nil || view.WorkspaceUnavailable {
		t.Fatalf("old workspace unavailable: %+v %v", view, err)
	}
	for path, want := range map[string][]byte{markerPath: markerBytes, subject: []byte(nilCapacityDefect), filepath.Join(s.home, "progress.json"): progress} {
		got, err := os.ReadFile(path)
		if err != nil || !bytes.Equal(got, want) {
			t.Fatalf("read rewrote %s", path)
		}
	}
	result, err := s.Check()
	if err != nil || result.Passed || !strings.Contains(result.Diagnostic, "len=0 cap=1; want len=0 cap=0") {
		t.Fatalf("nil public defect: %+v %v", result, err)
	}
	reference, err := os.ReadFile(filepath.Join(taskDir, "solution/reference/subject.go"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(subject, reference, 0600); err != nil {
		t.Fatal(err)
	}
	result, err = s.Check()
	if err != nil || !result.Passed {
		t.Fatalf("reference public: %+v %v", result, err)
	}
	result, err = s.Complete()
	if err != nil || !result.Passed {
		t.Fatalf("reference full: %+v %v", result, err)
	}
	// A non-nil empty result with zero capacity is contractually valid.
	alternative := strings.Replace(string(nilCapacityDefect), "make([]int, 0, 1)", "make([]int, 0)", 1)
	if err := os.WriteFile(subject, []byte(alternative), 0600); err != nil {
		t.Fatal(err)
	}
	result, err = s.Check()
	if err != nil || !result.Passed {
		t.Fatalf("valid empty representation: %+v %v", result, err)
	}
	after, err := tasking.TaskContentDigest(canonical)
	if err != nil || after != before {
		t.Fatal("canonical task changed")
	}
}
