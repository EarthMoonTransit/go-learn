package sources

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	sourceIDPattern = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	commitPattern   = regexp.MustCompile(`^[0-9a-f]{40}$`)
)

type Registry struct {
	Version   int      `yaml:"version"`
	AuditedAt string   `yaml:"audited_at"`
	Policy    string   `yaml:"policy"`
	Sources   []Source `yaml:"sources"`
}

type Source struct {
	ID                string             `yaml:"id"`
	Repository        string             `yaml:"repository"`
	DefaultBranch     string             `yaml:"default_branch"`
	Commit            string             `yaml:"commit"`
	License           SourceLicense      `yaml:"license"`
	UsageMode         string             `yaml:"usage_mode"`
	Adapter           string             `yaml:"adapter"`
	Scope             string             `yaml:"scope"`
	Attribution       []string           `yaml:"attribution"`
	Evidence          []string           `yaml:"evidence"`
	Risks             []string           `yaml:"risks"`
	ProvenanceCommits []ProvenanceCommit `yaml:"provenance_commits,omitempty"`
}

type ProvenanceCommit struct {
	Commit   string `yaml:"commit"`
	Evidence string `yaml:"evidence"`
}

type SourceLicense struct {
	Expression   string `yaml:"expression"`
	EvidencePath string `yaml:"evidence_path"`
	Notes        string `yaml:"notes"`
}

func LoadRegistry(root string) (Registry, error) {
	path := filepath.Join(root, "sources", "registry.yaml")
	b, err := readBounded(path, 4<<20)
	if err != nil {
		return Registry{}, fmt.Errorf("read source registry: %w", err)
	}
	dec := yaml.NewDecoder(bytes.NewReader(b))
	dec.KnownFields(true)
	var registry Registry
	if err := dec.Decode(&registry); err != nil {
		return Registry{}, fmt.Errorf("decode source registry: %w", err)
	}
	var extra any
	if err := dec.Decode(&extra); err == nil {
		return Registry{}, errors.New("source registry contains multiple YAML documents")
	} else if !errors.Is(err, io.EOF) {
		return Registry{}, fmt.Errorf("decode source registry trailer: %w", err)
	}
	if registry.Version != 1 {
		return Registry{}, fmt.Errorf("source registry version must be 1, got %d", registry.Version)
	}
	if registry.AuditedAt == "" {
		return Registry{}, errors.New("source registry audited_at is required")
	}
	if registry.Policy != "docs/LICENSE_POLICY.md" {
		return Registry{}, fmt.Errorf("source registry policy must be %q", "docs/LICENSE_POLICY.md")
	}
	seen := map[string]bool{}
	for _, source := range registry.Sources {
		if err := Validate(source); err != nil {
			return Registry{}, err
		}
		if seen[source.ID] {
			return Registry{}, fmt.Errorf("duplicate source id %q", source.ID)
		}
		seen[source.ID] = true
	}
	return registry, nil
}

func (r Registry) Find(id string) (Source, error) {
	for _, source := range r.Sources {
		if source.ID == id {
			return source, nil
		}
	}
	return Source{}, fmt.Errorf("source %q is not registered", id)
}

func (r Registry) ExercismCompanion(primary Source) (Source, bool) {
	if primary.ID != "exercism-go" {
		return Source{}, false
	}
	for _, source := range r.Sources {
		u, err := url.Parse(source.Repository)
		if err != nil {
			continue
		}
		if u.Host == "github.com" && u.Path == "/exercism/problem-specifications" {
			return source, true
		}
	}
	return Source{}, false
}

func Validate(source Source) error {
	if !sourceIDPattern.MatchString(source.ID) {
		return fmt.Errorf("source %q has invalid id", source.ID)
	}
	if !commitPattern.MatchString(source.Commit) {
		return fmt.Errorf("source %q commit must be a lowercase 40-character SHA", source.ID)
	}
	if strings.ContainsAny(source.Repository, "\r\n\x00") {
		return fmt.Errorf("source %q repository contains control characters", source.ID)
	}
	u, err := url.Parse(source.Repository)
	if err != nil || u.Scheme != "https" || u.Host == "" || u.User != nil || u.RawQuery != "" || u.Fragment != "" {
		return fmt.Errorf("source %q repository must be an absolute HTTPS URL without credentials, query, or fragment", source.ID)
	}
	switch source.UsageMode {
	case "adapt", "reference-only", "concepts-only", "blocked":
	default:
		return fmt.Errorf("source %q has invalid usage mode %q", source.ID, source.UsageMode)
	}
	if source.Adapter == "" || source.License.Expression == "" || source.License.EvidencePath == "" || source.DefaultBranch == "" || source.Scope == "" || len(source.Evidence) == 0 || len(source.Risks) == 0 {
		return fmt.Errorf("source %q must declare branch, adapter, license evidence, scope, audit evidence, and risks", source.ID)
	}
	seenProvenance := map[string]bool{}
	for _, revision := range source.ProvenanceCommits {
		if !commitPattern.MatchString(revision.Commit) || revision.Evidence == "" {
			return fmt.Errorf("source %q has invalid provenance commit", source.ID)
		}
		if seenProvenance[revision.Commit] {
			return fmt.Errorf("source %q has duplicate provenance commit %q", source.ID, revision.Commit)
		}
		seenProvenance[revision.Commit] = true
		evidenceURL, err := url.Parse(revision.Evidence)
		if err != nil || evidenceURL.Scheme != "https" || evidenceURL.Host == "" {
			return fmt.Errorf("source %q provenance commit %q requires an absolute HTTPS evidence URL", source.ID, revision.Commit)
		}
	}
	return nil
}

func readBounded(path string, limit int64) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	b, err := io.ReadAll(io.LimitReader(f, limit+1))
	if err != nil {
		return nil, err
	}
	if int64(len(b)) > limit {
		return nil, fmt.Errorf("file exceeds %d bytes", limit)
	}
	return b, nil
}

func ValidSourceID(id string) bool {
	return sourceIDPattern.MatchString(id)
}
