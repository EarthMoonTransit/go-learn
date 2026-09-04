package learner

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const maxProgressBytes = 4 << 20

func DefaultHome() (string, error) {
	if configured := os.Getenv("GO_LEARN_HOME"); configured != "" {
		return filepath.Abs(configured)
	}
	base, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("resolve user config directory: %w", err)
	}
	return filepath.Join(base, "go-practice-bank", "learner"), nil
}

func prepareDirectory(path string) (string, error) {
	absolute, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(absolute, 0o700); err != nil {
		return "", err
	}
	originalInfo, err := os.Lstat(absolute)
	if err != nil || originalInfo.Mode()&os.ModeSymlink != 0 || !originalInfo.IsDir() {
		return "", errors.New("learner home must be a real directory")
	}
	resolved, err := filepath.EvalSymlinks(absolute)
	if err != nil {
		return "", err
	}
	info, err := os.Lstat(resolved)
	if err != nil || info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return "", errors.New("learner home must be a real directory")
	}
	if err := os.Chmod(resolved, 0o700); err != nil {
		return "", err
	}
	return resolved, nil
}

func preparePrivateSubdirectory(home, name string) (string, error) {
	if filepath.Base(name) != name || name == "." || name == ".." || name == "" {
		return "", errors.New("private subdirectory name must be a single path component")
	}
	path := filepath.Join(home, name)
	if err := os.Mkdir(path, 0o700); err != nil && !errors.Is(err, os.ErrExist) {
		return "", err
	}
	info, err := os.Lstat(path)
	if err != nil || info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return "", fmt.Errorf("learner %s directory must be a real directory", name)
	}
	resolved, err := filepath.EvalSymlinks(path)
	if err != nil {
		return "", err
	}
	relative, err := filepath.Rel(home, resolved)
	if err != nil || relative != name {
		return "", fmt.Errorf("learner %s directory escapes learner home", name)
	}
	if err := os.Chmod(resolved, 0o700); err != nil {
		return "", err
	}
	return resolved, nil
}

func loadProgress(home string) (Progress, error) {
	path := filepath.Join(home, "progress.json")
	info, err := os.Lstat(path)
	if errors.Is(err, os.ErrNotExist) {
		return Progress{SchemaVersion: progressSchemaVersion, Tasks: map[string]*TaskProgress{}}, nil
	}
	if err != nil {
		return Progress{}, err
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() || info.Size() > maxProgressBytes {
		return Progress{}, errors.New("progress.json must be a bounded regular file")
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return Progress{}, err
	}
	decoder := json.NewDecoder(bytes.NewReader(content))
	decoder.DisallowUnknownFields()
	var progress Progress
	if err := decoder.Decode(&progress); err != nil {
		return Progress{}, fmt.Errorf("decode progress.json: %w", err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err == nil {
		return Progress{}, errors.New("progress.json contains multiple JSON values")
	} else if !errors.Is(err, io.EOF) {
		return Progress{}, err
	}
	if progress.SchemaVersion != progressSchemaVersion || progress.Tasks == nil {
		return Progress{}, errors.New("progress.json has an unsupported schema")
	}
	return progress, nil
}

func saveProgress(home string, progress Progress) error {
	content, err := json.MarshalIndent(progress, "", "  ")
	if err != nil {
		return err
	}
	content = append(content, '\n')
	if len(content) > maxProgressBytes {
		return errors.New("progress state exceeds the local size limit")
	}
	path := filepath.Join(home, "progress.json")
	temporary, err := os.CreateTemp(home, ".progress-*")
	if err != nil {
		return err
	}
	temporaryPath := temporary.Name()
	remove := true
	defer func() {
		_ = temporary.Close()
		if remove {
			_ = os.Remove(temporaryPath)
		}
	}()
	if err := temporary.Chmod(0o600); err != nil {
		return err
	}
	if _, err := temporary.Write(content); err != nil {
		return err
	}
	if err := temporary.Sync(); err != nil {
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	if info, err := os.Lstat(path); err == nil && (info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular()) {
		return errors.New("refusing to replace non-regular progress.json")
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err := os.Rename(temporaryPath, path); err != nil {
		return err
	}
	remove = false
	if directory, err := os.Open(home); err == nil {
		_ = directory.Sync()
		_ = directory.Close()
	}
	return nil
}

func withProgressLock(home string, action func(*Progress) error) error {
	lockPath := filepath.Join(home, ".progress.lock")
	lock, err := os.OpenFile(lockPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return errors.New("another learn command is updating local progress")
		}
		return err
	}
	if err := lock.Close(); err != nil {
		_ = os.Remove(lockPath)
		return err
	}
	defer os.Remove(lockPath)
	progress, err := loadProgress(home)
	if err != nil {
		return err
	}
	if err := action(&progress); err != nil {
		return err
	}
	return saveProgress(home, progress)
}

func writePrivateFile(path string, content []byte, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, mode)
	if err != nil {
		return err
	}
	remove := true
	defer func() {
		_ = file.Close()
		if remove {
			_ = os.Remove(path)
		}
	}()
	if _, err := file.Write(content); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	remove = false
	return nil
}

func readPrivateRegular(path string, limit int64) ([]byte, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return nil, fmt.Errorf("path is not a regular file: %s", path)
	}
	if info.Size() > limit {
		return nil, fmt.Errorf("file exceeds %d bytes: %s", limit, path)
	}
	return os.ReadFile(path)
}

func atomicPrivateWrite(path string, content []byte) error {
	directory, err := prepareDirectory(filepath.Dir(path))
	if err != nil {
		return err
	}
	path = filepath.Join(directory, filepath.Base(path))
	if info, err := os.Lstat(path); err == nil && (info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular()) {
		return errors.New("refusing to replace a non-regular learner file")
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	temporary, err := os.CreateTemp(directory, ".write-*")
	if err != nil {
		return err
	}
	temporaryPath := temporary.Name()
	remove := true
	defer func() {
		_ = temporary.Close()
		if remove {
			_ = os.Remove(temporaryPath)
		}
	}()
	if err := temporary.Chmod(0o600); err != nil {
		return err
	}
	if _, err := temporary.Write(content); err != nil {
		return err
	}
	if err := temporary.Sync(); err != nil {
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	if err := os.Rename(temporaryPath, path); err != nil {
		return err
	}
	remove = false
	if handle, err := os.Open(directory); err == nil {
		_ = handle.Sync()
		_ = handle.Close()
	}
	return nil
}
