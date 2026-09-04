package toolchainpolicy

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

var versionPattern = regexp.MustCompile(`^1\.[0-9]+$`)

type Policy struct {
	SchemaVersion            int          `yaml:"schema_version"`
	GoLanguageVersion        string       `yaml:"go_language_version"`
	CIToolchain              string       `yaml:"ci_toolchain"`
	ModulePrefix             string       `yaml:"module_prefix"`
	ForbidToolchainDirective bool         `yaml:"forbid_toolchain_directive"`
	LegacyTasks              []LegacyTask `yaml:"legacy_tasks"`
}

type LegacyTask struct {
	ID                string `yaml:"id"`
	GoLanguageVersion string `yaml:"go_language_version"`
}

func Load(root string) (Policy, error) {
	path := filepath.Join(root, "config", "toolchain.yaml")
	content, err := os.ReadFile(path)
	if err != nil {
		return Policy{}, fmt.Errorf("read toolchain policy: %w", err)
	}
	decoder := yaml.NewDecoder(bytes.NewReader(content))
	decoder.KnownFields(true)
	var policy Policy
	if err := decoder.Decode(&policy); err != nil {
		return Policy{}, fmt.Errorf("decode toolchain policy: %w", err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err == nil {
		return Policy{}, errors.New("toolchain policy contains multiple YAML documents")
	} else if !errors.Is(err, io.EOF) {
		return Policy{}, fmt.Errorf("decode toolchain policy trailer: %w", err)
	}
	if err := policy.Validate(); err != nil {
		return Policy{}, err
	}
	return policy, nil
}

func (p Policy) Validate() error {
	if p.SchemaVersion != 1 || !versionPattern.MatchString(p.GoLanguageVersion) {
		return errors.New("toolchain policy requires schema_version 1 and a Go 1.x language version")
	}
	if p.CIToolchain != p.GoLanguageVersion+".x" {
		return errors.New("ci_toolchain must be the patch line for go_language_version")
	}
	if p.ModulePrefix != "task.local/" || !p.ForbidToolchainDirective {
		return errors.New("toolchain policy must retain the task module prefix and forbid toolchain directives")
	}
	if len(p.LegacyTasks) == 0 {
		return errors.New("toolchain policy must explicitly identify legacy task compatibility")
	}
	seen := map[string]bool{}
	last := ""
	for _, legacy := range p.LegacyTasks {
		if legacy.ID == "" || !versionPattern.MatchString(legacy.GoLanguageVersion) || seen[legacy.ID] {
			return fmt.Errorf("invalid or duplicate legacy toolchain entry %q", legacy.ID)
		}
		if last != "" && legacy.ID <= last {
			return errors.New("legacy toolchain entries must be sorted by task id")
		}
		seen[legacy.ID] = true
		last = legacy.ID
	}
	return nil
}

func (p Policy) VersionForTask(id string, legacy bool) (string, error) {
	if !legacy {
		return p.GoLanguageVersion, nil
	}
	for _, item := range p.LegacyTasks {
		if item.ID == id {
			return item.GoLanguageVersion, nil
		}
	}
	return "", fmt.Errorf("task %q is not in the legacy toolchain allowlist", id)
}

func ValidateGoMod(content []byte, modulePath, version string, forbidToolchain bool) []string {
	var problems []string
	moduleSeen := false
	versionSeen := false
	for _, rawLine := range strings.Split(string(content), "\n") {
		line := strings.TrimSpace(rawLine)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		switch fields[0] {
		case "module":
			moduleSeen = len(fields) == 2 && fields[1] == modulePath
			if !moduleSeen {
				problems = append(problems, "go.mod module must be "+modulePath)
			}
		case "go":
			versionSeen = len(fields) == 2 && fields[1] == version
			if !versionSeen {
				problems = append(problems, "go.mod must declare go "+version)
			}
		case "toolchain":
			if forbidToolchain {
				problems = append(problems, "go.mod toolchain directive is forbidden")
			}
		case "require", "replace", "exclude", "retract", "tool", "godebug":
			problems = append(problems, fmt.Sprintf("go.mod directive %s is not allowed", fields[0]))
		}
	}
	if !moduleSeen {
		problems = append(problems, "go.mod is missing the required module directive")
	}
	if !versionSeen {
		problems = append(problems, "go.mod is missing the required go directive")
	}
	sort.Strings(problems)
	return compact(problems)
}

func ValidateRepository(root string, policy Policy) error {
	content, err := os.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil {
		return err
	}
	var problems []string
	moduleSeen := false
	versionSeen := false
	for _, rawLine := range strings.Split(string(content), "\n") {
		fields := strings.Fields(strings.TrimSpace(rawLine))
		if len(fields) == 0 || strings.HasPrefix(fields[0], "//") {
			continue
		}
		switch fields[0] {
		case "module":
			moduleSeen = len(fields) == 2 && fields[1] == "go-practice-bank"
			if !moduleSeen {
				problems = append(problems, "root module must be go-practice-bank")
			}
		case "go":
			versionSeen = len(fields) == 2 && fields[1] == policy.GoLanguageVersion
			if !versionSeen {
				problems = append(problems, "root go.mod must declare go "+policy.GoLanguageVersion)
			}
		case "toolchain":
			if policy.ForbidToolchainDirective {
				problems = append(problems, "root go.mod toolchain directive is forbidden")
			}
		}
	}
	if !moduleSeen {
		problems = append(problems, "root go.mod is missing the required module directive")
	}
	if !versionSeen {
		problems = append(problems, "root go.mod is missing the required go directive")
	}
	sort.Strings(problems)
	problems = compact(problems)
	if len(problems) != 0 {
		return fmt.Errorf("root go.mod violates toolchain policy: %s", strings.Join(problems, "; "))
	}
	return nil
}

func compact(values []string) []string {
	if len(values) < 2 {
		return values
	}
	result := values[:1]
	for _, value := range values[1:] {
		if value != result[len(result)-1] {
			result = append(result, value)
		}
	}
	return result
}
