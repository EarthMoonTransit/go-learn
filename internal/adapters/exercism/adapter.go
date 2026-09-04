package exercism

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"go-practice-bank/internal/catalog"
	"go-practice-bank/internal/sources"
)

const (
	maxConfigBytes = 8 << 20
	maxFileBytes   = 64 << 20
	maxTotalBytes  = 512 << 20
	maxFiles       = 20_000
)

var (
	slugPattern     = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
	topicPattern    = regexp.MustCompile(`^[a-z0-9]+(?:[-_][a-z0-9]+)*$`)
	generatedSource = regexp.MustCompile(`(?m)^// Source: exercism/problem-specifications\r?$`)
	generatedCommit = regexp.MustCompile(`(?m)^// Commit: ([0-9a-f]{7,40})(?:[[:space:]]|$)`)
)

type Config struct {
	Slug      string          `json:"slug"`
	Files     FilesConfig     `json:"files"`
	Concepts  []ConceptConfig `json:"concepts"`
	Exercises ExerciseGroups  `json:"exercises"`
}

type FilesConfig struct {
	Solution    []string `json:"solution"`
	Test        []string `json:"test"`
	Example     []string `json:"example"`
	Exemplar    []string `json:"exemplar"`
	Editor      []string `json:"editor"`
	Invalidator []string `json:"invalidator"`
}

type ConceptConfig struct {
	Slug string `json:"slug"`
}

type ExerciseGroups struct {
	Concept  []ExerciseConfig `json:"concept"`
	Practice []ExerciseConfig `json:"practice"`
}

type ExerciseConfig struct {
	Slug          string   `json:"slug"`
	Status        string   `json:"status"`
	Difficulty    *int     `json:"difficulty"`
	Concepts      []string `json:"concepts"`
	Practices     []string `json:"practices"`
	Prerequisites []string `json:"prerequisites"`
	Topics        []string `json:"topics"`
}

type ExerciseMetadata struct {
	Files FilesConfig `json:"files"`
}

type Adapter struct{}

func (Adapter) Harvest(checkoutRoot string, source sources.Source, companion *sources.Source) ([]catalog.Candidate, error) {
	if source.Adapter != "exercism-go-v1" {
		return nil, fmt.Errorf("source %q selects unsupported adapter %q", source.ID, source.Adapter)
	}
	if source.ID != "exercism-go" || source.UsageMode != "adapt" {
		return nil, fmt.Errorf("Exercism adapter requires the registered adapt source exercism-go")
	}
	root, err := canonicalDirectory(checkoutRoot)
	if err != nil {
		return nil, err
	}
	config, err := loadConfig(root)
	if err != nil {
		return nil, err
	}
	knownConcepts := map[string]bool{}
	for _, concept := range config.Concepts {
		if !slugPattern.MatchString(concept.Slug) {
			return nil, fmt.Errorf("config.json contains invalid concept slug %q", concept.Slug)
		}
		knownConcepts[concept.Slug] = true
		conceptDir, err := safePath(root, "concepts", concept.Slug)
		if err != nil {
			return nil, err
		}
		info, err := os.Lstat(conceptDir)
		if err != nil || !info.IsDir() {
			return nil, fmt.Errorf("concept directory is missing or invalid: concepts/%s", concept.Slug)
		}
	}

	type scheduled struct {
		kind     string
		exercise ExerciseConfig
	}
	all := make([]scheduled, 0, len(config.Exercises.Concept)+len(config.Exercises.Practice))
	if len(config.Exercises.Concept)+len(config.Exercises.Practice) > 10_000 {
		return nil, errors.New("config.json contains too many exercises")
	}
	for _, exercise := range config.Exercises.Concept {
		all = append(all, scheduled{kind: "concept", exercise: exercise})
	}
	for _, exercise := range config.Exercises.Practice {
		all = append(all, scheduled{kind: "practice", exercise: exercise})
	}
	sort.Slice(all, func(i, j int) bool {
		if all[i].kind != all[j].kind {
			return all[i].kind < all[j].kind
		}
		return all[i].exercise.Slug < all[j].exercise.Slug
	})
	seen := map[string]bool{}
	candidates := make([]catalog.Candidate, 0, len(all))
	for _, item := range all {
		key := item.kind + "/" + item.exercise.Slug
		if seen[key] {
			return nil, fmt.Errorf("config.json contains duplicate exercise %q", key)
		}
		seen[key] = true
		candidate, err := inspectExercise(root, source, companion, config.Files, knownConcepts, item.kind, item.exercise)
		if err != nil {
			return nil, err
		}
		candidates = append(candidates, candidate)
	}
	return candidates, nil
}

func loadConfig(root string) (Config, error) {
	configPath, err := safePath(root, "config.json")
	if err != nil {
		return Config{}, err
	}
	b, err := readRegularFile(configPath, maxConfigBytes)
	if err != nil {
		return Config{}, fmt.Errorf("read config.json: %w", err)
	}
	dec := json.NewDecoder(bufio.NewReader(strings.NewReader(string(b))))
	var config Config
	if err := dec.Decode(&config); err != nil {
		return Config{}, fmt.Errorf("decode config.json: %w", err)
	}
	var trailing any
	if err := dec.Decode(&trailing); !errors.Is(err, io.EOF) {
		if err == nil {
			return Config{}, errors.New("config.json contains more than one JSON value")
		}
		return Config{}, fmt.Errorf("decode config.json trailer: %w", err)
	}
	if config.Slug != "go" {
		return Config{}, fmt.Errorf("config.json slug is %q, want go", config.Slug)
	}
	return config, nil
}

func inspectExercise(root string, source sources.Source, companion *sources.Source, files FilesConfig, knownConcepts map[string]bool, kind string, exercise ExerciseConfig) (catalog.Candidate, error) {
	if !slugPattern.MatchString(exercise.Slug) {
		return catalog.Candidate{}, fmt.Errorf("invalid exercise slug %q", exercise.Slug)
	}
	if kind != "concept" && kind != "practice" {
		return catalog.Candidate{}, fmt.Errorf("invalid exercise kind %q", kind)
	}
	for field, values := range map[string][]string{"concepts": exercise.Concepts, "practices": exercise.Practices, "prerequisites": exercise.Prerequisites, "topics": exercise.Topics} {
		if len(values) > 1_000 {
			return catalog.Candidate{}, fmt.Errorf("exercise %q has too many %s", exercise.Slug, field)
		}
		for _, value := range values {
			pattern := slugPattern
			if field == "topics" {
				pattern = topicPattern
			}
			if !pattern.MatchString(value) || len(value) > 200 {
				return catalog.Candidate{}, fmt.Errorf("exercise %q has invalid %s value %q", exercise.Slug, field, value)
			}
		}
	}
	sourcePath := path.Join("exercises", kind, exercise.Slug)
	exerciseDir, err := safePath(root, filepath.FromSlash(sourcePath))
	if err != nil {
		return catalog.Candidate{}, err
	}
	candidate := catalog.Candidate{
		SchemaVersion:          1,
		SourceID:               source.ID,
		SourceCommit:           source.Commit,
		SourcePath:             sourcePath,
		Slug:                   exercise.Slug,
		ExerciseKind:           kind,
		UpstreamStatus:         exercise.Status,
		UpstreamDifficulty:     exercise.Difficulty,
		UpstreamConcepts:       sortedUnique(append(append([]string{}, exercise.Concepts...), exercise.Practices...)),
		UpstreamPrerequisites:  sortedUnique(exercise.Prerequisites),
		UpstreamTopics:         sortedUnique(exercise.Topics),
		StatementPaths:         []string{},
		DocumentationPaths:     []string{},
		StarterFilePaths:       []string{},
		PublicTestPaths:        []string{},
		GeneratedTestDataPaths: []string{},
		ReferenceSolutionPaths: []string{},
		MetadataPaths:          []string{},
		AdditionalFilePaths:    []string{},
		FileHashes:             []catalog.FileHash{},
		License:                source.License.Expression,
		UsageMode:              source.UsageMode,
		ImportWarnings:         []string{},
		ProvenanceRefs:         []catalog.ProvenanceRef{},
	}
	info, err := os.Lstat(exerciseDir)
	if errors.Is(err, os.ErrNotExist) {
		candidate.ImportWarnings = append(candidate.ImportWarnings, "exercise directory is missing")
		addMetadataWarnings(&candidate, knownConcepts)
		sortCandidate(&candidate)
		return candidate, nil
	}
	if err != nil {
		return catalog.Candidate{}, fmt.Errorf("inspect %s: %w", sourcePath, err)
	}
	if !info.IsDir() {
		return catalog.Candidate{}, fmt.Errorf("exercise path %s is not a directory", sourcePath)
	}
	metadata, err := loadExerciseMetadata(exerciseDir)
	if err != nil {
		return catalog.Candidate{}, fmt.Errorf("inspect %s metadata: %w", sourcePath, err)
	}
	if exercise.Status != "" {
		candidate.ImportWarnings = append(candidate.ImportWarnings, "upstream exercise status: "+exercise.Status)
	}
	declared, err := validateDeclaredFiles(exerciseDir, metadata.Files)
	if err != nil {
		return catalog.Candidate{}, fmt.Errorf("inspect %s declared files: %w", sourcePath, err)
	}
	effectiveFiles := mergeFiles(files, metadata.Files)
	var totalBytes int64
	fileCount := 0
	err = filepath.WalkDir(exerciseDir, func(current string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relCheckout, err := filepath.Rel(root, current)
		if err != nil || escapes(relCheckout) {
			return fmt.Errorf("path escapes checkout: %s", current)
		}
		relCheckout = filepath.ToSlash(relCheckout)
		if current == exerciseDir {
			return nil
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("source path contains forbidden symlink: %s", relCheckout)
		}
		if entry.IsDir() {
			return nil
		}
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("source path is not a regular file: %s", relCheckout)
		}
		fileCount++
		if fileCount > maxFiles {
			return fmt.Errorf("exercise %s exceeds %d files", exercise.Slug, maxFiles)
		}
		if info.Size() > maxFileBytes {
			return fmt.Errorf("source file exceeds %d bytes: %s", maxFileBytes, relCheckout)
		}
		totalBytes += info.Size()
		if totalBytes > maxTotalBytes {
			return fmt.Errorf("exercise %s exceeds %d total bytes", exercise.Slug, maxTotalBytes)
		}
		digest, err := hashRegularFile(current, info)
		if err != nil {
			return fmt.Errorf("hash %s: %w", relCheckout, err)
		}
		relExercise, err := filepath.Rel(exerciseDir, current)
		if err != nil || escapes(relExercise) {
			return fmt.Errorf("file escapes exercise: %s", current)
		}
		relExercise = filepath.ToSlash(relExercise)
		roles := classify(relExercise, effectiveFiles)
		candidate.FileHashes = append(candidate.FileHashes, catalog.FileHash{Path: relCheckout, SHA256: digest, Roles: roles})
		classifyPath(&candidate, relCheckout, roles)
		if contains(roles, "generated-test-data") && path.Base(relExercise) == "cases_test.go" {
			reference, warning, err := generatedProvenance(current, info, relCheckout, companion)
			if err != nil {
				return err
			}
			if reference != nil {
				candidate.ProvenanceRefs = append(candidate.ProvenanceRefs, *reference)
			}
			if warning != "" {
				candidate.ImportWarnings = append(candidate.ImportWarnings, warning)
			}
		}
		if contains(roles, "public-test") && !declared.tests[relExercise] {
			candidate.ImportWarnings = append(candidate.ImportWarnings, "public test not declared in .meta/config.json: "+relExercise)
		}
		return nil
	})
	if err != nil {
		return catalog.Candidate{}, fmt.Errorf("inspect %s: %w", sourcePath, err)
	}
	addMetadataWarnings(&candidate, knownConcepts)
	if len(candidate.StatementPaths) == 0 {
		candidate.ImportWarnings = append(candidate.ImportWarnings, "no statement file found")
	}
	if len(candidate.StarterFilePaths) == 0 {
		candidate.ImportWarnings = append(candidate.ImportWarnings, "no starter Go file found")
	}
	if len(candidate.PublicTestPaths) == 0 {
		candidate.ImportWarnings = append(candidate.ImportWarnings, "no public test file found")
	}
	if len(candidate.ReferenceSolutionPaths) == 0 {
		candidate.ImportWarnings = append(candidate.ImportWarnings, "no reference solution found")
	}
	if len(candidate.GeneratedTestDataPaths) > 0 {
		if len(candidate.ProvenanceRefs) > 0 {
			candidate.ImportWarnings = append(candidate.ImportWarnings, "generated case provenance resolved to a registered exact problem-specifications commit")
		} else {
			candidate.ImportWarnings = append(candidate.ImportWarnings, "generated test data has no exact companion commit in embedded metadata")
		}
	}
	sortCandidate(&candidate)
	return candidate, nil
}

func classify(rel string, files FilesConfig) []string {
	roles := map[string]bool{}
	base := path.Base(rel)
	if strings.HasPrefix(rel, ".docs/") {
		roles["documentation"] = true
		if base == "instructions.md" || base == "introduction.md" {
			roles["statement"] = true
		}
	}
	if strings.HasPrefix(rel, ".meta/") {
		roles["metadata"] = true
	}
	if matchesAny(rel, append(append([]string{}, files.Example...), files.Exemplar...)) || strings.HasPrefix(rel, ".meta/") && (base == "example.go" || base == "exemplar.go" || base == "solution.go") {
		roles["reference-solution"] = true
	}
	if matchesAny(rel, files.Test) || strings.HasSuffix(base, "_test.go") && !strings.HasPrefix(rel, ".meta/") {
		roles["public-test"] = true
	}
	if isGeneratedTestData(rel) {
		roles["generated-test-data"] = true
	}
	if (matchesAny(rel, append(append([]string{}, files.Solution...), files.Editor...)) || strings.HasSuffix(base, ".go")) && !roles["public-test"] && !strings.HasPrefix(rel, ".meta/") {
		roles["starter"] = true
	}
	if len(roles) == 0 {
		roles["additional"] = true
	}
	result := make([]string, 0, len(roles))
	for role := range roles {
		result = append(result, role)
	}
	sort.Strings(result)
	return result
}

func classifyPath(candidate *catalog.Candidate, path string, roles []string) {
	for _, role := range roles {
		switch role {
		case "statement":
			candidate.StatementPaths = append(candidate.StatementPaths, path)
		case "documentation":
			candidate.DocumentationPaths = append(candidate.DocumentationPaths, path)
		case "starter":
			candidate.StarterFilePaths = append(candidate.StarterFilePaths, path)
		case "public-test":
			candidate.PublicTestPaths = append(candidate.PublicTestPaths, path)
		case "generated-test-data":
			candidate.GeneratedTestDataPaths = append(candidate.GeneratedTestDataPaths, path)
		case "reference-solution":
			candidate.ReferenceSolutionPaths = append(candidate.ReferenceSolutionPaths, path)
		case "metadata":
			candidate.MetadataPaths = append(candidate.MetadataPaths, path)
		case "additional":
			candidate.AdditionalFilePaths = append(candidate.AdditionalFilePaths, path)
		}
	}
}

func addMetadataWarnings(candidate *catalog.Candidate, knownConcepts map[string]bool) {
	for _, concept := range append(append([]string{}, candidate.UpstreamConcepts...), candidate.UpstreamPrerequisites...) {
		if !knownConcepts[concept] {
			candidate.ImportWarnings = append(candidate.ImportWarnings, "unknown upstream concept: "+concept)
		}
	}
}

func sortCandidate(candidate *catalog.Candidate) {
	for _, values := range []*[]string{
		&candidate.UpstreamConcepts,
		&candidate.UpstreamPrerequisites,
		&candidate.UpstreamTopics,
		&candidate.StatementPaths,
		&candidate.DocumentationPaths,
		&candidate.StarterFilePaths,
		&candidate.PublicTestPaths,
		&candidate.GeneratedTestDataPaths,
		&candidate.ReferenceSolutionPaths,
		&candidate.MetadataPaths,
		&candidate.AdditionalFilePaths,
		&candidate.ImportWarnings,
	} {
		*values = sortedUnique(*values)
	}
	sort.Slice(candidate.FileHashes, func(i, j int) bool { return candidate.FileHashes[i].Path < candidate.FileHashes[j].Path })
	sort.Slice(candidate.ProvenanceRefs, func(i, j int) bool {
		return candidate.ProvenanceRefs[i].FilePath < candidate.ProvenanceRefs[j].FilePath
	})
}

func generatedProvenance(file string, expected os.FileInfo, relative string, companion *sources.Source) (*catalog.ProvenanceRef, string, error) {
	prefix, err := readRegularPrefix(file, expected, 8<<10)
	if err != nil {
		return nil, "", err
	}
	if !generatedSource.Match(prefix) {
		return nil, "generated case file has no exercism/problem-specifications source header: " + relative, nil
	}
	match := generatedCommit.FindSubmatch(prefix)
	if len(match) != 2 {
		return nil, "generated case file has no parseable problem-specifications commit: " + relative, nil
	}
	abbreviated := string(match[1])
	if companion == nil {
		return nil, "generated case commit is unresolved because the companion source is not registered: " + abbreviated, nil
	}
	commit, evidence, err := resolveProvenanceCommit(*companion, abbreviated)
	if err != nil {
		return nil, "", err
	}
	return &catalog.ProvenanceRef{
		FilePath:          relative,
		SourceID:          companion.ID,
		Repository:        companion.Repository,
		Commit:            commit,
		AbbreviatedCommit: abbreviated,
		License:           companion.License.Expression,
		Evidence:          evidence,
	}, "", nil
}

func resolveProvenanceCommit(source sources.Source, abbreviated string) (string, string, error) {
	type revision struct{ commit, evidence string }
	byCommit := map[string]revision{
		source.Commit: {commit: source.Commit, evidence: source.Repository + "/commit/" + source.Commit},
	}
	for _, item := range source.ProvenanceCommits {
		byCommit[item.Commit] = revision{commit: item.Commit, evidence: item.Evidence}
	}
	var matches []revision
	for _, item := range byCommit {
		if strings.HasPrefix(item.commit, abbreviated) {
			matches = append(matches, item)
		}
	}
	if len(matches) != 1 {
		return "", "", fmt.Errorf("problem-specifications commit %q resolves to %d registered revisions", abbreviated, len(matches))
	}
	return matches[0].commit, matches[0].evidence, nil
}

func isGeneratedTestData(rel string) bool {
	base := path.Base(rel)
	return strings.HasPrefix(rel, ".meta/testdata/") ||
		base == "tests.toml" ||
		base == "cases_test.go" ||
		strings.HasPrefix(base, "cases.") ||
		strings.HasPrefix(base, "test_cases.") ||
		strings.HasPrefix(base, "canonical-data.")
}

type declaredFiles struct {
	tests map[string]bool
}

func loadExerciseMetadata(exerciseDir string) (ExerciseMetadata, error) {
	metadataPath, err := safePath(exerciseDir, ".meta", "config.json")
	if err != nil {
		return ExerciseMetadata{}, err
	}
	b, err := readRegularFile(metadataPath, maxConfigBytes)
	if err != nil {
		return ExerciseMetadata{}, err
	}
	var metadata ExerciseMetadata
	decoder := json.NewDecoder(bytes.NewReader(b))
	if err := decoder.Decode(&metadata); err != nil {
		return ExerciseMetadata{}, err
	}
	var trailing any
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) {
		return ExerciseMetadata{}, errors.New(".meta/config.json contains trailing JSON")
	}
	return metadata, nil
}

func validateDeclaredFiles(exerciseDir string, files FilesConfig) (declaredFiles, error) {
	result := declaredFiles{tests: map[string]bool{}}
	groups := map[string][]string{
		"solution":    files.Solution,
		"test":        files.Test,
		"example":     files.Example,
		"exemplar":    files.Exemplar,
		"editor":      files.Editor,
		"invalidator": files.Invalidator,
	}
	for role, paths := range groups {
		for _, relative := range paths {
			if relative == "" || path.IsAbs(relative) || path.Clean(relative) != relative || relative == ".." || strings.HasPrefix(relative, "../") {
				return declaredFiles{}, fmt.Errorf("invalid %s path %q", role, relative)
			}
			resolved, err := safePath(exerciseDir, filepath.FromSlash(relative))
			if err != nil {
				return declaredFiles{}, err
			}
			info, err := os.Lstat(resolved)
			if err != nil {
				return declaredFiles{}, fmt.Errorf("declared %s file %q is missing", role, relative)
			}
			if !info.Mode().IsRegular() {
				return declaredFiles{}, fmt.Errorf("declared %s path %q is not regular", role, relative)
			}
			if role == "test" {
				result.tests[relative] = true
			}
		}
	}
	return result, nil
}

func mergeFiles(left, right FilesConfig) FilesConfig {
	return FilesConfig{
		Solution:    append(append([]string{}, left.Solution...), right.Solution...),
		Test:        append(append([]string{}, left.Test...), right.Test...),
		Example:     append(append([]string{}, left.Example...), right.Example...),
		Exemplar:    append(append([]string{}, left.Exemplar...), right.Exemplar...),
		Editor:      append(append([]string{}, left.Editor...), right.Editor...),
		Invalidator: append(append([]string{}, left.Invalidator...), right.Invalidator...),
	}
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func LoadConfig(checkoutRoot string) (Config, error) {
	root, err := canonicalDirectory(checkoutRoot)
	if err != nil {
		return Config{}, err
	}
	return loadConfig(root)
}

func matchesAny(name string, patterns []string) bool {
	for _, pattern := range patterns {
		matched, err := path.Match(pattern, name)
		if err == nil && matched {
			return true
		}
	}
	return false
}

func sortedUnique(values []string) []string {
	seen := map[string]bool{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		if value != "" && !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	sort.Strings(result)
	return result
}

func canonicalDirectory(root string) (string, error) {
	absolute, err := filepath.Abs(root)
	if err != nil {
		return "", fmt.Errorf("resolve checkout: %w", err)
	}
	resolved, err := filepath.EvalSymlinks(absolute)
	if err != nil {
		return "", fmt.Errorf("resolve checkout symlinks: %w", err)
	}
	info, err := os.Stat(resolved)
	if err != nil || !info.IsDir() {
		return "", fmt.Errorf("checkout is not a directory: %s", resolved)
	}
	return resolved, nil
}

func safePath(root string, elements ...string) (string, error) {
	candidate := filepath.Join(append([]string{root}, elements...)...)
	rel, err := filepath.Rel(root, candidate)
	if err != nil || escapes(rel) {
		return "", fmt.Errorf("path escapes checkout: %q", candidate)
	}
	current := root
	parts := strings.Split(filepath.Clean(rel), string(filepath.Separator))
	for _, part := range parts {
		if part == "." || part == "" {
			continue
		}
		current = filepath.Join(current, part)
		info, err := os.Lstat(current)
		if errors.Is(err, os.ErrNotExist) {
			return candidate, nil
		}
		if err != nil {
			return "", fmt.Errorf("inspect path component %s: %w", current, err)
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return "", fmt.Errorf("source path contains forbidden symlink component: %s", current)
		}
	}
	return candidate, nil
}

func escapes(rel string) bool {
	return rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) || filepath.IsAbs(rel)
}

func readRegularFile(file string, limit int64) ([]byte, error) {
	info, err := os.Lstat(file)
	if err != nil {
		return nil, err
	}
	if !info.Mode().IsRegular() || info.Mode()&os.ModeSymlink != 0 {
		return nil, errors.New("file is not a regular non-symlink file")
	}
	if info.Size() > limit {
		return nil, fmt.Errorf("file exceeds %d bytes", limit)
	}
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	opened, err := f.Stat()
	if err != nil || !os.SameFile(info, opened) {
		return nil, errors.New("file changed while opening")
	}
	b, err := io.ReadAll(io.LimitReader(f, limit+1))
	if err != nil {
		return nil, err
	}
	if int64(len(b)) > limit {
		return nil, fmt.Errorf("file exceeds %d bytes", limit)
	}
	return b, nil
}

func hashRegularFile(file string, expected os.FileInfo) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	opened, err := f.Stat()
	if err != nil || !opened.Mode().IsRegular() || !os.SameFile(expected, opened) {
		return "", errors.New("file changed while opening")
	}
	hash := sha256.New()
	n, err := io.Copy(hash, io.LimitReader(f, maxFileBytes+1))
	if err != nil {
		return "", err
	}
	if n > maxFileBytes {
		return "", fmt.Errorf("file exceeds %d bytes", maxFileBytes)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func readRegularPrefix(file string, expected os.FileInfo, limit int64) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	opened, err := f.Stat()
	if err != nil || !opened.Mode().IsRegular() || !os.SameFile(expected, opened) {
		return nil, errors.New("file changed while opening")
	}
	return io.ReadAll(io.LimitReader(f, limit))
}
