package learner

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"go-practice-bank/internal/safeio"
	"go-practice-bank/internal/toolchainpolicy"
	"gopkg.in/yaml.v3"
)

var topicIDPattern = regexp.MustCompile(`^[a-z][a-z0-9]*(?:-[a-z0-9]+)*$`)
var theoryAnchorPattern = regexp.MustCompile(`^#[a-zA-Z][a-zA-Z0-9_-]*$`)
var theoryLinkPattern = regexp.MustCompile(`\[[^\]\n]+\]\(([^)\n]*)\)`)
var theoryHTMLPattern = regexp.MustCompile(`(?i)<\s*(?:/?[a-z][a-z0-9]*(?:\s|/?>)|!--|!doctype)`)

type TheoryLink struct {
	URL         string `yaml:"url" json:"url"`
	Description string `yaml:"description" json:"description"`
}
type Topic struct {
	ID            string       `yaml:"id" json:"id"`
	SectionID     string       `yaml:"section_id" json:"section_id"`
	Title         string       `yaml:"title" json:"title"`
	Path          string       `yaml:"path" json:"-"`
	Prerequisites []string     `yaml:"prerequisites" json:"prerequisites"`
	TaskIDs       []string     `yaml:"task_ids" json:"task_ids"`
	GoVersion     string       `yaml:"go_version" json:"go_version"`
	ReviewedAt    string       `yaml:"reviewed_at" json:"reviewed_at"`
	OfficialLinks []TheoryLink `yaml:"official_links" json:"official_links"`
	Status        string       `yaml:"-" json:"status"`
}
type TheoryView struct {
	Topic    Topic        `json:"topic"`
	Markdown string       `json:"markdown"`
	Lessons  []LessonView `json:"lessons"`
}
type topicDocument struct {
	SchemaVersion int     `yaml:"schema_version"`
	Topics        []Topic `yaml:"topics"`
}

func safeTheoryURL(raw string, official bool) bool {
	if !official && theoryAnchorPattern.MatchString(raw) {
		return true
	}
	if strings.ContainsAny(raw, "\\\r\n\t ") {
		return false
	}
	u, err := url.Parse(raw)
	if err != nil || u.Scheme != "https" || u.Hostname() == "" || u.User != nil || u.Opaque != "" || u.Port() != "" {
		return false
	}
	return !official || u.Host == "go.dev" || u.Host == "pkg.go.dev"
}

func validateTheoryMarkdown(content []byte) error {
	if !utf8.Valid(content) || bytes.ContainsRune(content, 0) {
		return errors.New("theory must be UTF-8 text")
	}
	fenced := false
	for _, line := range strings.Split(string(content), "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			fenced = !fenced
			continue
		}
		if fenced {
			continue
		}
		if theoryHTMLPattern.MatchString(line) {
			return errors.New("raw HTML is not supported in theory")
		}
		for _, match := range theoryLinkPattern.FindAllStringSubmatch(line, -1) {
			if !safeTheoryURL(match[1], false) {
				return errors.New("unsafe theory link")
			}
		}
	}
	if fenced {
		return errors.New("unclosed theory code fence")
	}
	return nil
}

func readTheoryPage(root, path string) (string, error) {
	if !fs.ValidPath(path) || !strings.HasPrefix(path, "theory/") || !strings.HasSuffix(path, ".md") || strings.Contains(path, "\\") {
		return "", errors.New("theory path must be a Markdown file below theory/")
	}
	data, err := safeio.ReadFile(root, path, 64<<10)
	if err != nil {
		return "", err
	}
	if err := validateTheoryMarkdown(data); err != nil {
		return "", err
	}
	return string(data), nil
}

// ValidateTheory validates optional course metadata without opening learner state.
func ValidateTheory(root string) error {
	_, err := (&Service{root: root}).Topics()
	return err
}

func (s *Service) Topics() ([]Topic, error) {
	data, err := safeio.ReadFile(s.root, "curriculum/topics.yaml", 128<<10)
	if errors.Is(err, os.ErrNotExist) {
		return []Topic{}, nil
	}
	if err != nil {
		return nil, err
	}
	if !utf8.Valid(data) {
		return nil, errors.New("topic metadata must be UTF-8")
	}
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)
	var doc topicDocument
	if err := decoder.Decode(&doc); err != nil {
		return nil, err
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return nil, errors.New("topics require exactly one YAML document")
	}
	if doc.SchemaVersion != 1 || len(doc.Topics) > 200 {
		return nil, errors.New("unsupported topic version or count")
	}
	if len(doc.Topics) == 0 {
		return []Topic{}, nil
	}
	route, err := loadRoute(s.root)
	if err != nil {
		return nil, err
	}
	sections := map[string]bool{}
	if route != nil {
		for _, b := range route.Blocks {
			sections[b.ID] = true
		}
	}
	policy, err := toolchainpolicy.Load(s.root)
	if err != nil {
		return nil, err
	}
	known := map[string]int{}
	paths := map[string]bool{}
	for i := range doc.Topics {
		t := &doc.Topics[i]
		if !topicIDPattern.MatchString(t.ID) || strings.TrimSpace(t.Title) == "" || len(t.Title) > 160 || !sections[t.SectionID] {
			return nil, fmt.Errorf("invalid topic identity or section: %s", t.ID)
		}
		if _, exists := known[t.ID]; exists || paths[t.Path] {
			return nil, errors.New("duplicate topic ID or path")
		}
		known[t.ID] = i
		paths[t.Path] = true
		if t.GoVersion != policy.GoLanguageVersion {
			return nil, fmt.Errorf("topic %s needs review for Go %s", t.ID, policy.GoLanguageVersion)
		}
		if _, err := time.Parse("2006-01-02", t.ReviewedAt); err != nil {
			return nil, errors.New("invalid topic review date")
		}
		if len(t.TaskIDs) < 1 || len(t.TaskIDs) > 3 || len(t.OfficialLinks) < 1 || len(t.OfficialLinks) > 3 {
			return nil, errors.New("topic needs 1–3 tasks and official links")
		}
		seen := map[string]bool{}
		for _, id := range t.TaskIDs {
			if seen[id] {
				return nil, errors.New("duplicate topic task")
			}
			seen[id] = true
			if _, err := s.Show(id); err != nil {
				return nil, fmt.Errorf("topic %s task: %w", t.ID, err)
			}
		}
		seen = map[string]bool{}
		for _, link := range t.OfficialLinks {
			if !safeTheoryURL(link.URL, true) || strings.TrimSpace(link.Description) == "" || seen[link.URL] {
				return nil, errors.New("invalid or duplicate official link")
			}
			seen[link.URL] = true
		}
		_, err := readTheoryPage(s.root, t.Path)
		t.Status = "ready"
		if errors.Is(err, os.ErrNotExist) {
			t.Status = "planned"
		} else if err != nil {
			return nil, fmt.Errorf("topic %s: %w", t.ID, err)
		}
	}
	visiting, visited := map[string]bool{}, map[string]bool{}
	var visit func(string) error
	visit = func(id string) error {
		index, ok := known[id]
		if !ok {
			return fmt.Errorf("unknown prerequisite topic %s", id)
		}
		if visiting[id] {
			return errors.New("topic prerequisite cycle")
		}
		if visited[id] {
			return nil
		}
		visiting[id] = true
		seen := map[string]bool{}
		for _, dep := range doc.Topics[index].Prerequisites {
			if seen[dep] {
				return errors.New("duplicate topic prerequisite")
			}
			seen[dep] = true
			if err := visit(dep); err != nil {
				return err
			}
		}
		visiting[id] = false
		visited[id] = true
		return nil
	}
	for id := range known {
		if err := visit(id); err != nil {
			return nil, err
		}
	}
	return doc.Topics, nil
}

func (s *Service) TheoryView(id string) (TheoryView, error) {
	topics, err := s.Topics()
	if err != nil {
		return TheoryView{}, err
	}
	for _, topic := range topics {
		if topic.ID != id {
			continue
		}
		view := TheoryView{Topic: topic, Lessons: []LessonView{}}
		if topic.Status == "ready" {
			view.Markdown, err = readTheoryPage(s.root, topic.Path)
			if err != nil {
				return TheoryView{}, err
			}
		}
		for _, taskID := range topic.TaskIDs {
			task, err := s.Show(taskID)
			if err != nil {
				return TheoryView{}, err
			}
			number, err := s.LessonNumber(taskID)
			if err != nil {
				return TheoryView{}, err
			}
			view.Lessons = append(view.Lessons, LessonView{ID: taskID, Title: task.Title, Number: number, Available: true})
		}
		return view, nil
	}
	return TheoryView{}, fmt.Errorf("unknown topic %q", id)
}
