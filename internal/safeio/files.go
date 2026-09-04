package safeio

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(root, relative string, limit int64) ([]byte, error) {
	path, err := resolve(root, relative, false)
	if err != nil {
		return nil, err
	}
	info, err := os.Lstat(path)
	if err != nil {
		return nil, err
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return nil, fmt.Errorf("path is not a regular file: %s", relative)
	}
	if info.Size() > limit {
		return nil, fmt.Errorf("file %s exceeds %d bytes", relative, limit)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if int64(len(b)) > limit {
		return nil, fmt.Errorf("file %s exceeds %d bytes", relative, limit)
	}
	return b, nil
}

func AtomicWrite(root, relative string, content []byte, limit int64) (bool, error) {
	if int64(len(content)) > limit {
		return false, fmt.Errorf("output %s exceeds %d bytes", relative, limit)
	}
	path, err := resolve(root, relative, true)
	if err != nil {
		return false, err
	}
	if info, err := os.Lstat(path); err == nil {
		if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
			return false, fmt.Errorf("output path is not a regular file: %s", relative)
		}
		if info.Size() > limit {
			return false, fmt.Errorf("existing output %s exceeds %d bytes", relative, limit)
		}
		existing, err := os.ReadFile(path)
		if err != nil {
			return false, err
		}
		if bytes.Equal(existing, content) {
			return false, nil
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return false, err
	}
	directory := filepath.Dir(path)
	temporary, err := os.CreateTemp(directory, ".atomic-*")
	if err != nil {
		return false, fmt.Errorf("create temporary output: %w", err)
	}
	temporaryPath := temporary.Name()
	remove := true
	defer func() {
		if remove {
			_ = os.Remove(temporaryPath)
		}
	}()
	if err := temporary.Chmod(0o644); err != nil {
		_ = temporary.Close()
		return false, err
	}
	if _, err := temporary.Write(content); err != nil {
		_ = temporary.Close()
		return false, err
	}
	if err := temporary.Sync(); err != nil {
		_ = temporary.Close()
		return false, err
	}
	if err := temporary.Close(); err != nil {
		return false, err
	}
	if err := os.Rename(temporaryPath, path); err != nil {
		return false, fmt.Errorf("publish output: %w", err)
	}
	remove = false
	if dir, err := os.Open(directory); err == nil {
		_ = dir.Sync()
		_ = dir.Close()
	}
	return true, nil
}

func resolve(root, relative string, createParents bool) (string, error) {
	if !fs.ValidPath(relative) || relative == "." || strings.ContainsRune(relative, '\\') {
		return "", fmt.Errorf("invalid repository-relative path %q", relative)
	}
	absoluteRoot, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}
	canonicalRoot, err := filepath.EvalSymlinks(absoluteRoot)
	if err != nil {
		return "", fmt.Errorf("resolve repository root: %w", err)
	}
	rootInfo, err := os.Lstat(canonicalRoot)
	if err != nil || rootInfo.Mode()&os.ModeSymlink != 0 || !rootInfo.IsDir() {
		return "", errors.New("repository root is not a real directory")
	}
	parts := strings.Split(relative, "/")
	current := canonicalRoot
	for index, part := range parts[:len(parts)-1] {
		current = filepath.Join(current, part)
		info, err := os.Lstat(current)
		if errors.Is(err, os.ErrNotExist) && createParents {
			if err := os.Mkdir(current, 0o755); err != nil && !errors.Is(err, os.ErrExist) {
				return "", fmt.Errorf("create output directory %s: %w", strings.Join(parts[:index+1], "/"), err)
			}
			info, err = os.Lstat(current)
		}
		if err != nil {
			return "", fmt.Errorf("inspect output directory %s: %w", current, err)
		}
		if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return "", fmt.Errorf("output ancestor is not a real directory: %s", current)
		}
	}
	return filepath.Join(current, parts[len(parts)-1]), nil
}
