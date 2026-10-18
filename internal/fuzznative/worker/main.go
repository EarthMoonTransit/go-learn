// Trusted admission/rewrite helper. Original learner packages are never executed.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/importer"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"

	"go-practice-bank/internal/fuzznative/frame"
	"go-practice-bank/internal/fuzznative/protocol"
)

// Filled by the checker when it builds these embedded sources in isolation.
var kernelDigest string

func fail(err error) { fmt.Fprintln(os.Stderr, err); os.Exit(97) }
func inside(root, p string) (string, error) {
	root, err := filepath.EvalSymlinks(root)
	if err != nil {
		return "", err
	}
	p, err = filepath.EvalSymlinks(p)
	if err != nil {
		return "", err
	}
	relative, err := filepath.Rel(root, p)
	if err != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("export escaped disposable root")
	}
	return p, nil
}
func main() {
	if len(os.Args) != 1 {
		fail(fmt.Errorf("helper does not accept arguments"))
	}
	raw, err := io.ReadAll(io.LimitReader(os.Stdin, protocol.MaxRequest+1))
	if err != nil {
		fail(err)
	}
	if len(raw) > protocol.MaxRequest {
		fail(fmt.Errorf("helper request cap"))
	}
	var req protocol.Request
	if err := protocol.Decode(raw, &req); err != nil {
		fail(err)
	}
	if req.Version != protocol.Version || req.Runtime != protocol.CurrentRuntime() || req.KernelDigest != kernelDigest {
		fail(fmt.Errorf("helper request version/runtime/source binding"))
	}
	if err := req.ValidateIdentity(); err != nil {
		fail(err)
	}
	root := filepath.Dir(os.Getenv("HOME"))
	if root == "." || os.Getenv("GOCACHE") != filepath.Join(root, "cache") {
		fail(fmt.Errorf("helper environment"))
	}
	fset := token.NewFileSet()
	var exportError error
	imp := importer.ForCompiler(fset, "gc", func(path string) (io.ReadCloser, error) {
		ref, ok := req.Exports[path]
		if !ok {
			exportError = fmt.Errorf("missing pinned export %s", path)
			return nil, exportError
		}
		p, err := inside(root, ref.Path)
		if err != nil {
			exportError = err
			return nil, exportError
		}
		f, err := os.Open(p)
		if err != nil {
			exportError = err
			return nil, err
		}
		defer f.Close()
		info, err := f.Stat()
		if err != nil || !info.Mode().IsRegular() || info.Size() > 64<<20 {
			exportError = fmt.Errorf("invalid pinned export %s", path)
			return nil, exportError
		}
		data, err := io.ReadAll(io.LimitReader(f, (64<<20)+1))
		if err != nil || len(data) > 64<<20 || protocol.Hash(data) != ref.SHA256 {
			exportError = fmt.Errorf("pinned export digest mismatch for %s", path)
			return nil, exportError
		}
		return io.NopCloser(bytes.NewReader(data)), nil
	})
	r := protocol.Response{Identity: req.Identity, Version: protocol.Version, Runtime: protocol.CurrentRuntime(), KernelDigest: kernelDigest, RequestSHA256: protocol.Hash(raw), SourceSHA256: protocol.SourceHashes(req.Sources)}
	inspection, err := frame.Analyze(fset, imp, req.Sources, req.Identity)
	if exportError != nil {
		fail(exportError)
	}
	if err != nil {
		r.Diagnostic = err.Error()
	} else {
		rewrite, err := inspection.Rewrite(req.Identity.ObserverPath())
		if err != nil {
			fail(err)
		}
		r.Accepted = true
		r.Metadata = &inspection
		r.Rewritten = rewrite.Sources
		r.Methods = rewrite.Methods
		r.Alias = rewrite.ImportAlias
	}
	b, err := json.Marshal(r)
	if err != nil {
		fail(err)
	}
	b = append(b, '\n')
	n, err := os.Stdout.Write(b)
	if err != nil || n != len(b) {
		fail(fmt.Errorf("helper response write"))
	}
}
