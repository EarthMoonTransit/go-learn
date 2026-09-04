package tasking

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	taskIDPattern = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*-[0-9]{3}$`)
	slugPattern   = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
)

var orderedStates = []State{StateDraft, StateInReview, StateReady, StateRejected, StateRetired}

var stateDirectories = map[State]string{
	StateDraft:    "drafts",
	StateInReview: "review",
	StateReady:    "ready",
	StateRejected: "rejected",
	StateRetired:  "retired",
}

var qualityByState = map[State]string{
	StateDraft:    "draft",
	StateInReview: "review",
	StateReady:    "ready",
	StateRejected: "quarantined",
	StateRetired:  "retired",
}

func ParseState(value string) (State, error) {
	state := State(value)
	if _, ok := stateDirectories[state]; !ok {
		return "", fmt.Errorf("invalid lifecycle state %q", value)
	}
	return state, nil
}

func packagePath(root string, state State, id string) (string, error) {
	directory, ok := stateDirectories[state]
	if !ok {
		return "", fmt.Errorf("invalid lifecycle state %q", state)
	}
	if !taskIDPattern.MatchString(id) {
		return "", fmt.Errorf("invalid task id %q", id)
	}
	return filepath.Join(root, "tasks", directory, id), nil
}

type LocatedTask struct {
	ID    string
	State State
	Path  string
}

func FindTask(root, id string) (LocatedTask, error) {
	if !taskIDPattern.MatchString(id) {
		return LocatedTask{}, fmt.Errorf("invalid task id %q", id)
	}
	var matches []LocatedTask
	for _, state := range orderedStates {
		path, _ := packagePath(root, state, id)
		info, err := os.Lstat(path)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return LocatedTask{}, fmt.Errorf("inspect %s: %w", path, err)
		}
		if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return LocatedTask{}, fmt.Errorf("task path must be a real directory: %s", path)
		}
		matches = append(matches, LocatedTask{ID: id, State: state, Path: path})
	}
	if len(matches) == 0 {
		return LocatedTask{}, fmt.Errorf("task %q was not found", id)
	}
	if len(matches) != 1 {
		return LocatedTask{}, fmt.Errorf("duplicate task id %q exists in %d lifecycle directories", id, len(matches))
	}
	return matches[0], nil
}

func inspectTree(root string) error {
	rootInfo, err := os.Lstat(root)
	if err != nil {
		return err
	}
	if rootInfo.Mode()&os.ModeSymlink != 0 || !rootInfo.IsDir() {
		return errors.New("task package must be a real directory")
	}
	return filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if path == root {
			return nil
		}
		relative, err := filepath.Rel(root, path)
		if err != nil || !fs.ValidPath(filepath.ToSlash(relative)) {
			return fmt.Errorf("invalid task path %q", relative)
		}
		info, err := os.Lstat(path)
		if err != nil {
			return err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("symlink is not allowed in task package: %s", relative)
		}
		if !info.IsDir() && entry.Name() == "go.work" {
			return fmt.Errorf("go.work is not allowed in task package: %s", relative)
		}
		if !info.IsDir() && !info.Mode().IsRegular() {
			return fmt.Errorf("non-regular file is not allowed in task package: %s", relative)
		}
		return nil
	})
}

func listTaskIDs(root string) ([]string, error) {
	seen := map[string]State{}
	var result []string
	for _, state := range orderedStates {
		directory := filepath.Join(root, "tasks", stateDirectories[state])
		entries, err := os.ReadDir(directory)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			if !taskIDPattern.MatchString(entry.Name()) {
				continue
			}
			path := filepath.Join(directory, entry.Name())
			info, err := os.Lstat(path)
			if err != nil || info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
				return nil, fmt.Errorf("task path must be a real directory: %s", path)
			}
			if first, ok := seen[entry.Name()]; ok {
				return nil, fmt.Errorf("duplicate task id %q in %s and %s", entry.Name(), first, state)
			}
			seen[entry.Name()] = state
			result = append(result, entry.Name())
		}
	}
	sort.Strings(result)
	return result, nil
}

func cleanRelativePath(value string) error {
	if value == "" || filepath.IsAbs(value) || strings.ContainsRune(value, '\\') || strings.ContainsAny(value, "\r\n\x00") || !fs.ValidPath(value) {
		return fmt.Errorf("invalid repository-relative path %q", value)
	}
	cleaned := filepath.ToSlash(filepath.Clean(value))
	if cleaned == "." || cleaned == ".." || strings.HasPrefix(cleaned, "../") || cleaned != value {
		return fmt.Errorf("invalid repository-relative path %q", value)
	}
	return nil
}
