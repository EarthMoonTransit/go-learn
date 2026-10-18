package tasking

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"

	"go-practice-bank/internal/taxonomy"
	"gopkg.in/yaml.v3"
)

var taskDigestContracts = []string{
	"LICENSE",
	"THIRD_PARTY_NOTICES",
	"schemas/task.schema.json",
	"schemas/task-v2.schema.json",
	"schemas/task-lifecycle.schema.json",
	"schemas/task-review.schema.json",
	"schemas/task-provenance.schema.json",
	"schemas/task-provenance-v2.schema.json",
	"schemas/project-identity.schema.json",
	"schemas/outbound-license-policy.schema.json",
	"schemas/task-quality-report.schema.json",
	"schemas/task-check-collection.schema.json",
	"schemas/taxonomy-v3.schema.json",
	"schemas/toolchain-policy.schema.json",
	"config/toolchain.yaml",
	"config/project-identity.yaml",
	"config/outbound-license.yaml",
	"curriculum/skills.yaml",
	"sources/registry.yaml",
}

func ToolchainPolicyDigest(root string) (string, error) {
	content, err := os.ReadFile(filepath.Join(root, "config", "toolchain.yaml"))
	if err != nil {
		return "", fmt.Errorf("read toolchain policy: %w", err)
	}
	digest := sha256.Sum256(content)
	return hex.EncodeToString(digest[:]), nil
}

func TaskInputDigest(root string, located LocatedTask) (string, error) {
	hash := sha256.New()
	writeDigestField(hash, "contract", []byte(CheckContractVersion))
	writeDigestField(hash, "state", []byte(located.State))
	contracts := taskDigestContracts
	taxonomyBytes, err := readDigestFile(filepath.Join(root, "curriculum", "skills.yaml"))
	if err != nil {
		return "", fmt.Errorf("read taxonomy version for collection: %w", err)
	}
	var header struct {
		Version int `yaml:"version"`
	}
	if err := yaml.Unmarshal(taxonomyBytes, &header); err != nil {
		return "", fmt.Errorf("decode taxonomy version for collection: %w", err)
	}
	if header.Version == 4 {
		contracts = append(append([]string(nil), contracts...), taxonomy.ContractPaths()...)
		contracts = append(contracts, "schemas/coverage-report.schema.json")
	}
	for _, relative := range contracts {
		content, err := readDigestFile(filepath.Join(root, filepath.FromSlash(relative)))
		if err != nil {
			return "", fmt.Errorf("digest contract %s: %w", relative, err)
		}
		writeDigestField(hash, "contract-file:"+relative, content)
	}
	var paths []string
	err = filepath.WalkDir(located.Path, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if path == located.Path {
			return nil
		}
		relative, err := filepath.Rel(located.Path, path)
		if err != nil || !fs.ValidPath(filepath.ToSlash(relative)) {
			return fmt.Errorf("invalid task input path %q", relative)
		}
		info, err := os.Lstat(path)
		if err != nil {
			return err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("symlink is not allowed in task input: %s", relative)
		}
		if entry.IsDir() {
			return nil
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("non-regular task input: %s", relative)
		}
		paths = append(paths, filepath.ToSlash(relative))
		return nil
	})
	if err != nil {
		return "", err
	}
	sort.Strings(paths)
	for _, relative := range paths {
		path := filepath.Join(located.Path, filepath.FromSlash(relative))
		info, err := os.Lstat(path)
		if err != nil {
			return "", err
		}
		content, err := readDigestFile(path)
		if err != nil {
			return "", err
		}
		writeDigestField(hash, "task-mode:"+relative, []byte(fmt.Sprintf("%04o", info.Mode().Perm())))
		writeDigestField(hash, "task-file:"+relative, content)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func readDigestFile(path string) ([]byte, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return nil, fmt.Errorf("digest input must be a regular file: %s", path)
	}
	return os.ReadFile(path)
}

func writeDigestField(writer io.Writer, name string, content []byte) {
	_ = binary.Write(writer, binary.BigEndian, uint64(len(name)))
	_, _ = io.WriteString(writer, name)
	_ = binary.Write(writer, binary.BigEndian, uint64(len(content)))
	_, _ = writer.Write(content)
}

func MarshalCheckCollection(collection CheckCollection) ([]byte, error) {
	return yaml.Marshal(collection)
}

func LoadCheckCollection(root, path string) (CheckCollection, error) {
	if path == "" {
		return CheckCollection{}, errors.New("check collection path is required")
	}
	resolved := path
	if !filepath.IsAbs(path) {
		if err := cleanRelativePath(filepath.ToSlash(path)); err != nil {
			return CheckCollection{}, err
		}
		resolved = filepath.Join(root, filepath.FromSlash(path))
	}
	info, err := os.Lstat(resolved)
	if err != nil {
		return CheckCollection{}, err
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return CheckCollection{}, fmt.Errorf("check collection must be a regular file: %s", path)
	}
	var collection CheckCollection
	if err := decodeStrictFile(resolved, &collection); err != nil {
		return CheckCollection{}, fmt.Errorf("decode check collection: %w", err)
	}
	if err := ValidateCheckCollection(root, collection, nil); err != nil {
		return CheckCollection{}, err
	}
	return collection, nil
}

func ValidateCheckCollection(root string, collection CheckCollection, states []State) error {
	if err := validateSchema(root, "task-check-collection.schema.json", collection); err != nil {
		return fmt.Errorf("check collection schema: %w", err)
	}
	if collection.ContractVersion != CheckContractVersion {
		return fmt.Errorf("stale check contract %q; require %q", collection.ContractVersion, CheckContractVersion)
	}
	toolchainDigest, err := ToolchainPolicyDigest(root)
	if err != nil {
		return err
	}
	if collection.ToolchainPolicyDigest != toolchainDigest {
		return errors.New("stale check collection: toolchain policy digest changed")
	}
	expected, err := DiscoverTasks(root, states)
	if err != nil {
		return err
	}
	expectedByID := make(map[string]LocatedTask, len(expected))
	for _, located := range expected {
		expectedByID[located.ID] = located
	}
	seen := make(map[string]bool, len(collection.Tasks))
	for _, checked := range collection.Tasks {
		if seen[checked.TaskID] {
			return fmt.Errorf("duplicate task result %q in check collection", checked.TaskID)
		}
		seen[checked.TaskID] = true
		if _, ok := expectedByID[checked.TaskID]; !ok {
			return fmt.Errorf("unknown task result %q in check collection", checked.TaskID)
		}
	}
	for _, located := range expected {
		if !seen[located.ID] {
			return fmt.Errorf("missing task result %q in check collection", located.ID)
		}
	}
	if len(collection.Tasks) != len(expected) {
		return fmt.Errorf("stale check collection: contains %d tasks, repository requires %d", len(collection.Tasks), len(expected))
	}
	for index, located := range expected {
		checked := collection.Tasks[index]
		if checked.TaskID != located.ID || checked.State != located.State {
			return fmt.Errorf("check collection task %d is %s/%s; require %s/%s in deterministic order", index, checked.TaskID, checked.State, located.ID, located.State)
		}
		if checked.Report.TaskID != checked.TaskID || checked.Report.State != checked.State {
			return fmt.Errorf("forged check collection: report identity differs for %s", checked.TaskID)
		}
		digest, err := TaskInputDigest(root, located)
		if err != nil {
			return err
		}
		if checked.InputDigest != digest {
			return fmt.Errorf("stale check collection: task input digest changed for %s", checked.TaskID)
		}
	}
	if collection.Execution.TaskChecks != len(collection.Tasks) {
		return errors.New("forged check collection: execution task count differs from results")
	}
	if len(collection.Tasks) != 0 && collection.Execution.Sessions != 1 {
		return errors.New("forged check collection: a non-empty full pass must use one execution session")
	}
	if collection.Execution.ChildGoRuns != collection.Execution.VariantRuns+collection.Execution.LegacyCommands || collection.Execution.Workspaces != collection.Execution.ChildGoRuns {
		return errors.New("forged check collection: execution topology is inconsistent")
	}
	return nil
}

func DecodeCheckCollection(content []byte) (CheckCollection, error) {
	var collection CheckCollection
	decoder := yaml.NewDecoder(bytes.NewReader(content))
	decoder.KnownFields(true)
	if err := decoder.Decode(&collection); err != nil {
		return CheckCollection{}, err
	}
	var trailing any
	if err := decoder.Decode(&trailing); err == nil {
		return CheckCollection{}, errors.New("multiple YAML documents are not allowed")
	} else if !errors.Is(err, io.EOF) {
		return CheckCollection{}, err
	}
	return collection, nil
}
