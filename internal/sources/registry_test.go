package sources

import (
	"strings"
	"testing"
)

func TestValidateAcceptsExactProvenanceCommit(t *testing.T) {
	source := validTestSource()
	source.ProvenanceCommits = []ProvenanceCommit{{
		Commit:   strings.Repeat("b", 40),
		Evidence: "https://example.com/commit/" + strings.Repeat("b", 40),
	}}
	if err := Validate(source); err != nil {
		t.Fatal(err)
	}
}

func TestValidateRejectsAbbreviatedProvenanceCommit(t *testing.T) {
	source := validTestSource()
	source.ProvenanceCommits = []ProvenanceCommit{{
		Commit:   "abcdef0",
		Evidence: "https://example.com/commit/abcdef0",
	}}
	if err := Validate(source); err == nil || !strings.Contains(err.Error(), "invalid provenance commit") {
		t.Fatalf("expected invalid provenance commit error, got %v", err)
	}
}

func TestValidateRejectsDuplicateProvenanceCommit(t *testing.T) {
	source := validTestSource()
	revision := ProvenanceCommit{
		Commit:   strings.Repeat("b", 40),
		Evidence: "https://example.com/commit/" + strings.Repeat("b", 40),
	}
	source.ProvenanceCommits = []ProvenanceCommit{revision, revision}
	if err := Validate(source); err == nil || !strings.Contains(err.Error(), "duplicate provenance commit") {
		t.Fatalf("expected duplicate provenance commit error, got %v", err)
	}
}

func validTestSource() Source {
	return Source{
		ID:            "test-source",
		Repository:    "https://example.com/repository",
		DefaultBranch: "main",
		Commit:        strings.Repeat("a", 40),
		License: SourceLicense{
			Expression:   "MIT",
			EvidencePath: "LICENSE",
		},
		UsageMode: "reference-only",
		Adapter:   "test-adapter-v1",
		Scope:     "Test fixture.",
		Evidence:  []string{"https://example.com/evidence"},
		Risks:     []string{"Test risk."},
	}
}
