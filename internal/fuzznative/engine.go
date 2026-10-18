package fuzznative

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"go-practice-bank/internal/fuzznative/frame"
	"go-practice-bank/internal/fuzznative/observe"
	"go-practice-bank/internal/fuzznative/protocol"
	"go-practice-bank/internal/fuzznative/trace"
	"go-practice-bank/internal/isolation"
)

const compilerOutputCap = 4 << 20

// Engine reuses only its explicitly supplied disposable environment. Every
// native invocation gets a new module; no success cache survives this engine.
type Engine struct {
	mu          sync.Mutex
	environment *isolation.Environment
	metrics     Metrics
	runtime     Runtime
	goTool      os.FileInfo
	helper      string
	exports     map[string]map[string]protocol.Export
}

func NewEngine(environment *isolation.Environment) *Engine {
	return &Engine{environment: environment, exports: map[string]map[string]protocol.Export{}}
}

func (e *Engine) Metrics() Metrics {
	e.mu.Lock()
	defer e.mu.Unlock()
	return e.metrics
}

func (e *Engine) workspace(name string) (string, error) {
	path, err := e.environment.NewWorkspace(name)
	if err == nil {
		e.metrics.Workspaces++
	}
	return path, err
}

func (e *Engine) checkGoTool() error {
	if e.goTool != nil {
		current, err := os.Stat(goExecutable())
		if err != nil || !os.SameFile(e.goTool, current) || e.goTool.Size() != current.Size() || !e.goTool.ModTime().Equal(current.ModTime()) {
			return fmt.Errorf("Go toolchain changed during native fuzz operation")
		}
	}
	return nil
}

func (e *Engine) goCommand(work string, args []string, mode isolation.ExecutionMode, timeout time.Duration) (Process, []byte, error) {
	if err := e.checkGoTool(); err != nil {
		return Process{}, nil, err
	}
	e.metrics.GoRuns++
	return runProcess(e.environment, work, goExecutable(), args, nil, mode, timeout, compilerOutputCap, false, nil)
}

func processFailure(p Process, output []byte, err error) string {
	if err != nil {
		return err.Error()
	}
	if p.TimedOut {
		return "native fuzz command exceeded its deadline"
	}
	if p.Truncated {
		return "native fuzz command exceeded its capture limit"
	}
	if p.Exit != 0 {
		return fmt.Sprintf("native fuzz command exited %d: %s", p.Exit, strings.TrimSpace(string(output)))
	}
	return ""
}

func (e *Engine) prepare(version string, timeout time.Duration) error {
	if e.helper != "" {
		return nil
	}
	work, err := e.workspace("fuzz-helper")
	if err != nil {
		return err
	}
	p, output, err := e.goCommand(work, []string{"env", "-json", "GOVERSION", "GOOS", "GOARCH"}, isolation.ExecutionModeNormal, timeout)
	if failure := processFailure(p, output, err); failure != "" {
		return fmt.Errorf("query toolchain: %s", failure)
	}
	var values struct{ GOVERSION, GOOS, GOARCH string }
	if err := json.Unmarshal(output, &values); err != nil {
		return fmt.Errorf("decode Go runtime: %w", err)
	}
	e.runtime = Runtime{Version: values.GOVERSION, GOOS: values.GOOS, GOARCH: values.GOARCH}
	if err := e.runtime.Validate(); err != nil {
		return err
	}
	e.goTool, err = os.Stat(goExecutable())
	if err != nil {
		return err
	}
	if err := materializeHelper(work, version); err != nil {
		return err
	}
	helper := filepath.Join(work, "fuzz-admit"+executableSuffix())
	p, output, err = e.goCommand(work, []string{"build", "-ldflags=-X=main.kernelDigest=" + SourceDigest(), "-o", helper, "./internal/fuzznative/worker"}, isolation.ExecutionModeNormal, timeout)
	if failure := processFailure(p, output, err); failure != "" {
		return fmt.Errorf("build trusted helper: %s", failure)
	}
	e.helper = helper
	return nil
}

func executableSuffix() string {
	if protocol.CurrentRuntime().GOOS == "windows" {
		return ".exe"
	}
	return ""
}

func writeSources(work string, identity Identity, version string, implementation, sources map[string]string) error {
	if err := os.WriteFile(filepath.Join(work, "go.mod"), []byte("module "+identity.ModulePath+"\n\ngo "+version+"\n"), 0600); err != nil {
		return err
	}
	dir := filepath.Join(work, "subject")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	for _, files := range []map[string]string{implementation, sources} {
		for name, source := range files {
			if err := os.WriteFile(filepath.Join(dir, name), []byte(source), 0600); err != nil {
				return err
			}
		}
	}
	return nil
}

func implementationKey(request Request) string {
	var names []string
	for name := range request.Implementation {
		names = append(names, name)
	}
	sort.Strings(names)
	var fields []string
	fields = append(fields, request.Identity.ModulePath, request.GoVersion)
	for _, name := range names {
		fields = append(fields, name, protocol.Hash([]byte(request.Implementation[name])))
	}
	return protocol.Hash([]byte(strings.Join(fields, "\x00")))
}

func (e *Engine) pinnedExports(request Request) (map[string]protocol.Export, error) {
	key := implementationKey(request)
	if exports := e.exports[key]; exports != nil {
		return exports, nil
	}
	work, err := e.workspace("fuzz-exports")
	if err != nil {
		return nil, err
	}
	if err := writeSources(work, request.Identity, request.GoVersion, request.Implementation, nil); err != nil {
		return nil, err
	}
	p, output, err := e.goCommand(work, []string{"list", "-export", "-deps", "-json", "testing", "bytes", "encoding/hex", request.Identity.SubjectPath()}, isolation.ExecutionModeNormal, request.Timeout)
	if failure := processFailure(p, output, err); failure != "" {
		return nil, fmt.Errorf("compile pinned exports: %s", failure)
	}
	exports := map[string]protocol.Export{}
	decoder := json.NewDecoder(bytes.NewReader(output))
	for {
		var item struct{ ImportPath, Export string }
		err := decoder.Decode(&item)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("decode pinned exports: %w", err)
		}
		if item.Export == "" {
			continue
		}
		path, err := contained(e.environment.Root, item.Export)
		if err != nil {
			return nil, err
		}
		info, err := os.Lstat(path)
		if err != nil {
			return nil, err
		}
		if !info.Mode().IsRegular() || info.Size() > 64<<20 {
			return nil, fmt.Errorf("invalid pinned export file")
		}
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		data, readErr := io.ReadAll(io.LimitReader(file, (64<<20)+1))
		closeErr := file.Close()
		if readErr != nil {
			return nil, readErr
		}
		if closeErr != nil {
			return nil, closeErr
		}
		if len(data) > 64<<20 {
			return nil, fmt.Errorf("pinned export exceeds 64 MiB")
		}
		exports[item.ImportPath] = protocol.Export{Path: path, SHA256: protocol.Hash(data)}
	}
	if _, ok := exports[request.Identity.SubjectPath()]; !ok {
		return nil, fmt.Errorf("missing subject export")
	}
	e.exports[key] = exports
	return exports, nil
}

func validateRequest(request Request) error {
	if err := request.Identity.Validate(); err != nil {
		return err
	}
	if request.GoVersion != "1.27" || request.Timeout < time.Second || request.Timeout > 60*time.Second {
		return fmt.Errorf("unsupported native fuzz language or command limit")
	}
	if len(request.Sources) < 1 || len(request.Sources) > frame.MaxFiles {
		return fmt.Errorf("frame.resources: require 1..8 test files")
	}
	total := 0
	for name, source := range request.Sources {
		total += len(source)
		if filepath.Base(name) != name || !strings.HasSuffix(name, "_test.go") || !utf8.ValidString(source) || total > frame.MaxSourceBytes {
			return fmt.Errorf("frame.resources: require UTF-8 basename test files within 256 KiB")
		}
	}
	if len(request.Implementation) == 0 {
		return fmt.Errorf("missing trusted subject implementation")
	}
	for name, source := range request.Implementation {
		if filepath.Base(name) != name || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") || len(source) > 8<<20 {
			return fmt.Errorf("invalid trusted implementation file")
		}
		if _, ok := request.Sources[name]; ok {
			return fmt.Errorf("test and implementation paths overlap")
		}
	}
	return nil
}

func (e *Engine) Verify(request Request) Result {
	e.mu.Lock()
	defer e.mu.Unlock()
	result := Result{}
	fail := func(phase, diagnostic string) Result {
		result.Valid = false
		result.Detected = false
		result.Canonical = ""
		result.Phase = phase
		if e.environment != nil {
			diagnostic = strings.ReplaceAll(diagnostic, e.environment.Root, "<isolated>")
		}
		if len(diagnostic) > 8192 {
			diagnostic = diagnostic[:8192] + "…"
		}
		result.Diagnostic = diagnostic
		return result
	}
	if err := validateRequest(request); err != nil {
		return fail("admission", err.Error())
	}
	if e.environment == nil {
		return fail("infrastructure", "isolated environment is not initialized")
	}
	if err := e.prepare(request.GoVersion, request.Timeout); err != nil {
		return fail("infrastructure", err.Error())
	}
	if request.Race && !isolation.RaceDetectorSupported(e.runtime.GOOS, e.runtime.GOARCH) {
		result.RaceUnsupported = true
		return fail("infrastructure", "race detector is unsupported on this native target")
	}
	exports, err := e.pinnedExports(request)
	if err != nil {
		return fail("exports", err.Error())
	}
	input := protocol.Request{Version: protocol.Version, Identity: request.Identity, Runtime: e.runtime, KernelDigest: SourceDigest(), Sources: request.Sources, Exports: exports}
	raw, err := json.Marshal(input)
	if err != nil {
		return fail("infrastructure", err.Error())
	}
	raw = append(raw, '\n')
	if len(raw) > protocol.MaxRequest {
		return fail("admission", "native fuzz helper request exceeds its bound")
	}
	e.metrics.HelperRuns++
	p, output, err := runProcess(e.environment, filepath.Dir(e.helper), e.helper, nil, raw, isolation.ExecutionModeNormal, request.Timeout, protocol.MaxResponse, true, nil)
	if failure := processFailure(p, output, err); failure != "" {
		return fail("admission-helper", failure)
	}
	response, err := protocol.ValidateResponse(output, raw, input)
	if err != nil {
		return fail("admission-helper", err.Error())
	}
	if !response.Accepted {
		return fail("admission", response.Diagnostic)
	}
	work, err := e.workspace("fuzz-native")
	if err != nil {
		return fail("infrastructure", err.Error())
	}
	if err := writeSources(work, request.Identity, request.GoVersion, request.Implementation, response.Rewritten); err != nil {
		return fail("infrastructure", err.Error())
	}
	observer, err := sourceFiles.ReadFile("observe/observe.go")
	if err != nil {
		return fail("infrastructure", err.Error())
	}
	dir := filepath.Join(work, "internal", "observe")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fail("infrastructure", err.Error())
	}
	if err := os.WriteFile(filepath.Join(dir, "observe.go"), observer, 0600); err != nil {
		return fail("infrastructure", err.Error())
	}
	if request.Vet {
		p, output, err := e.goCommand(work, []string{"vet", "./subject"}, isolation.ExecutionModeNormal, request.Timeout)
		if failure := processFailure(p, output, err); failure != "" {
			return fail("vet", failure)
		}
		result.Vetted = true
	}
	mode := isolation.ExecutionModeNormal
	args := []string{"test", "-c", "-vet=off"}
	if request.Race {
		mode = isolation.ExecutionModeRace
		args = append(args, "-race")
	}
	binary := filepath.Join(work, "fuzz.test"+executableSuffix())
	args = append(args, "-o", binary, "./subject")
	p, output, err = e.goCommand(work, args, mode, request.Timeout)
	if failure := processFailure(p, output, err); failure != "" {
		return fail("compile", failure)
	}
	tracePath := filepath.Join(work, "observer.jsonl")
	e.metrics.NativeRuns++
	native, output, err := runProcess(e.environment, work, binary, []string{"-test.run=^FuzzHex$", "-test.count=1", "-test.v=test2json", "-test.timeout=" + request.Timeout.String()}, nil, mode, request.Timeout, trace.MaxNativeBytes, true, map[string]string{"GO_LEARN_CODEC_TRACE": tracePath})
	if err != nil {
		return fail("infrastructure", err.Error())
	}
	if err := e.checkGoTool(); err != nil {
		return fail("infrastructure", err.Error())
	}
	e.metrics.GoRuns++
	converted, jsonOutput, err := runProcess(e.environment, work, goExecutable(), []string{"tool", "test2json", "-p", request.Identity.SubjectPath()}, output, isolation.ExecutionModeNormal, request.Timeout, trace.MaxConvertedBytes, false, nil)
	if err != nil {
		return fail("infrastructure", err.Error())
	}
	data, truncated, readErr := readTrace(e.environment, tracePath, observe.MaxTraceBytes)
	if readErr != nil {
		return fail("native", "read complete observer trace: "+readErr.Error())
	}
	outcome := trace.Validate(request.Identity, response.Metadata.SeedValues, data, jsonOutput, native.Exit, converted.Exit, native.TimedOut || converted.TimedOut, native.Truncated || converted.Truncated || truncated)
	if !outcome.Valid {
		return fail("native", outcome.Diagnostic)
	}
	result.Valid, result.Detected, result.Canonical = outcome.Valid, outcome.Detected, outcome.Canonical
	result.Callbacks, result.Calls, result.Assertions = outcome.Callbacks, outcome.Entries, outcome.Assertions
	if outcome.Detected {
		result.Diagnostic = strings.TrimSpace(string(output))
		if len(result.Diagnostic) > 8192 {
			result.Diagnostic = result.Diagnostic[:8192] + "…"
		}
		result.Diagnostic = strings.ReplaceAll(result.Diagnostic, e.environment.Root, "<isolated>")
	}
	return result
}
