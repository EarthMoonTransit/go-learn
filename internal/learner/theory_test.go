package learner

import (
	"bytes"
	"encoding/json"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

const syntheticTopics = `schema_version: 1
topics:
  - id: first-check
    section_id: first-steps
    title: First check
    path: theory/en/first-check.md
    prerequisites: []
    task_ids: [literal-bindings-001]
    go_version: "1.27"
    reviewed_at: "2026-09-10"
    official_links:
      - url: https://pkg.go.dev/testing#T
        description: Read testing.T
`

func theoryFixture(t *testing.T) *Service {
	t.Helper()
	s := uiTestService(t, "")
	root := t.TempDir()
	// Copy metadata and task inputs only; tests never execute authored code.
	for _, relative := range []string{"curriculum", "config", "sources", "schemas", "docs", "tasks/ready/literal-bindings-001"} {
		source := filepath.Join(s.root, relative)
		if err := filepath.WalkDir(source, func(path string, entry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(s.root, path)
			if err != nil {
				return err
			}
			target := filepath.Join(root, rel)
			if entry.IsDir() {
				return os.MkdirAll(target, 0755)
			}
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			return os.WriteFile(target, data, 0644)
		}); err != nil {
			t.Fatal(err)
		}
	}
	s.root = root
	putTheory(t, s, syntheticTopics, "# First check\n\nRead [testing](https://pkg.go.dev/testing#T).\n")
	return s
}
func putTheory(t *testing.T, s *Service, metadata, page string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Join(s.root, "theory/en"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(s.root, "curriculum/topics.yaml"), []byte(metadata), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(s.root, "theory/en/first-check.md"), []byte(page), 0644); err != nil {
		t.Fatal(err)
	}
}
func homeSnapshot(t *testing.T, home string) map[string]string {
	t.Helper()
	snapshot := map[string]string{}
	if err := filepath.WalkDir(home, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(home, path)
		snapshot[rel] = string(data)
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	return snapshot
}
func TestTheoryReadOnlyCLIAndHTTP(t *testing.T) {
	s := theoryFixture(t)
	if _, err := s.Start("literal-bindings-001"); err != nil {
		t.Fatal(err)
	}
	before := homeSnapshot(t, s.home)
	view, err := s.TheoryView("first-check")
	if err != nil {
		t.Fatal(err)
	}
	if view.Topic.Status != "ready" || view.Lessons[0].Number != "01.01" {
		t.Fatalf("bad view: %+v", view)
	}
	var output, stderr bytes.Buffer
	if code := RunCLI(Config{Root: s.root, Home: s.home}, []string{"theory", "first-check"}, &output, &stderr); code != 0 {
		t.Fatalf("CLI %d: %s", code, &stderr)
	}
	if !strings.Contains(output.String(), view.Markdown) || !strings.Contains(output.String(), "learn show literal-bindings-001") {
		t.Fatal("CLI/view disagree")
	}
	u := testLocalUI(t, s)
	w := uiHTTP(u, "GET", "/api/theory?id=first-check", "", nil)
	var got TheoryView
	view.Topic.Path = "" // Filesystem paths are deliberately not exposed by the API.
	if w.Code != 200 || json.Unmarshal(w.Body.Bytes(), &got) != nil || !reflect.DeepEqual(view, got) {
		t.Fatalf("API/view disagree: %s", w.Body)
	}
	if w := uiHTTP(u, "GET", "/api/topics", "", nil); w.Code != 200 {
		t.Fatal(w.Body)
	}
	for _, test := range []struct {
		path   string
		status int
		change func(*http.Request)
	}{
		{"/api/topics?path=secret", 400, nil}, {"/api/theory?id=../secret", 400, nil}, {"/api/theory?id=first-check&id=second", 400, nil},
		{"/api/theory?id=unknown", 404, nil}, {"/api/topics", 401, func(r *http.Request) { r.Header.Del("Authorization") }},
		{"/api/theory?id=first-check", 403, func(r *http.Request) { r.Header.Set("Origin", "https://evil.example") }},
	} {
		if w := uiHTTP(u, "GET", test.path, "", test.change); w.Code != test.status {
			t.Fatalf("%s: %d", test.path, w.Code)
		}
	}
	if !reflect.DeepEqual(before, homeSnapshot(t, s.home)) {
		t.Fatal("Theory mutated learner files, markers, progress or clocks")
	}
}
func TestTheoryRejectsUnsafeInputs(t *testing.T) {
	s := theoryFixture(t)
	for _, test := range []struct{ name, metadata, page string }{
		{"unknown field", syntheticTopics + "unexpected: true\n", "safe"},
		{"version", strings.Replace(syntheticTopics, "version: 1", "version: 2", 1), "safe"},
		{"task", strings.Replace(syntheticTopics, "literal-bindings-001", "unknown-001", 1), "safe"},
		{"section", strings.Replace(syntheticTopics, "first-steps", "absent", 1), "safe"},
		{"traversal", strings.Replace(syntheticTopics, "theory/en/first-check.md", "theory/../secret.md", 1), "safe"},
		{"absolute", strings.Replace(syntheticTopics, "theory/en/first-check.md", "/tmp/secret.md", 1), "safe"},
		{"cycle", strings.Replace(syntheticTopics, "prerequisites: []", "prerequisites: [first-check]", 1), "safe"},
		{"missing prerequisite", strings.Replace(syntheticTopics, "prerequisites: []", "prerequisites: [missing]", 1), "safe"},
		{"official host", strings.Replace(syntheticTopics, "pkg.go.dev/testing", "evil.example/testing", 1), "safe"},
		{"credentials", strings.Replace(syntheticTopics, "https://pkg.go.dev", "https://user@pkg.go.dev", 1), "safe"},
		{"raw HTML", syntheticTopics, "<script>alert(1)</script>"},
		{"javascript", syntheticTopics, "[run](javascript:alert)"},
		{"data", syntheticTopics, "[run](data:text/plain,x)"},
		{"relative URL", syntheticTopics, "[run](../../secret)"},
		{"oversize page", syntheticTopics, strings.Repeat("x", (64<<10)+1)},
		{"invalid UTF8", syntheticTopics, string([]byte{255})},
		{"oversize metadata", syntheticTopics + strings.Repeat(" ", 128<<10), "safe"},
		{"extra document", syntheticTopics + "---\n{}\n", "safe"},
	} {
		t.Run(test.name, func(t *testing.T) {
			putTheory(t, s, test.metadata, test.page)
			if _, err := s.Topics(); err == nil {
				t.Fatal("invalid theory accepted")
			}
		})
	}
}
func TestTheoryMissingAndSymlinkPages(t *testing.T) {
	s := theoryFixture(t)
	path := filepath.Join(s.root, "theory/en/first-check.md")
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
	view, err := s.TheoryView("first-check")
	if err != nil || view.Topic.Status != "planned" {
		t.Fatalf("missing page: %+v %v", view, err)
	}
	outside := filepath.Join(t.TempDir(), "secret.md")
	if err := os.WriteFile(outside, []byte("secret"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, path); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Topics(); err == nil {
		t.Fatal("symlink file accepted")
	}
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(filepath.Dir(path)); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Dir(outside), filepath.Dir(path)); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Topics(); err == nil {
		t.Fatal("symlink ancestor accepted")
	}
	if err := os.Remove(filepath.Join(s.root, "curriculum/topics.yaml")); err != nil {
		t.Fatal(err)
	}
	topics, err := s.Topics()
	if err != nil || len(topics) != 0 {
		t.Fatalf("missing metadata: %v", err)
	}
	if _, err := s.Show("literal-bindings-001"); err != nil {
		t.Fatal("theory absence broke task", err)
	}
}
