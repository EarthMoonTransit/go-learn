package fuzznative

import (
	"crypto/sha256"
	"embed"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

// The worker and its callers share these canonical sources. Tests never enter
// the helper build or the execution-source digest.
//
//go:embed *.go frame/*.go protocol/*.go worker/*.go observe/*.go trace/*.go
var sourceFiles embed.FS

func sourcePaths() []string {
	var paths []string
	err := fs.WalkDir(sourceFiles, ".", func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() && strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	sort.Strings(paths)
	return paths
}

var sourceDigestOnce sync.Once
var sourceDigest string

func SourceDigest() string {
	sourceDigestOnce.Do(func() { sourceDigest = computeSourceDigest() })
	return sourceDigest
}

func computeSourceDigest() string {
	h := sha256.New()
	for _, path := range sourcePaths() {
		data, err := sourceFiles.ReadFile(path)
		if err != nil {
			panic(err)
		}
		_ = binary.Write(h, binary.BigEndian, uint64(len(path)))
		_, _ = h.Write([]byte(path))
		_ = binary.Write(h, binary.BigEndian, uint64(len(data)))
		_, _ = h.Write(data)
	}
	return hex.EncodeToString(h.Sum(nil))
}

func materializeHelper(workspace, version string) error {
	for _, path := range sourcePaths() {
		prefix := strings.SplitN(path, "/", 2)[0]
		if prefix != "frame" && prefix != "protocol" && prefix != "worker" {
			continue
		}
		data, err := sourceFiles.ReadFile(path)
		if err != nil {
			return err
		}
		target := filepath.Join(workspace, "internal", "fuzznative", filepath.FromSlash(path))
		if err := os.MkdirAll(filepath.Dir(target), 0700); err != nil {
			return err
		}
		if err := os.WriteFile(target, data, 0600); err != nil {
			return err
		}
	}
	return os.WriteFile(filepath.Join(workspace, "go.mod"), []byte(fmt.Sprintf("module go-practice-bank\n\ngo %s\n", version)), 0600)
}
