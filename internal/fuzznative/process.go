package fuzznative

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"go-practice-bank/internal/isolation"
)

type Process struct {
	Executable     string   `json:"executable"`
	Cwd            string   `json:"cwd"`
	Args           []string `json:"args"`
	Env            []string `json:"env"`
	DeadlineMillis int64    `json:"deadline_millis"`
	Exit           int      `json:"exit"`
	TimedOut       bool     `json:"timed_out"`
	Error          string   `json:"error,omitempty"`
	Seconds        float64  `json:"seconds"`
	Cap            int      `json:"capture_cap"`
	Truncated      bool     `json:"truncated"`
	OutputSHA256   string   `json:"output_sha256"`
}
type capture struct {
	sync.Mutex
	b         []byte
	limit     int
	truncated bool
}

func (c *capture) Write(p []byte) (int, error) {
	c.Lock()
	defer c.Unlock()
	n := len(p)
	room := c.limit - len(c.b)
	if len(p) > room {
		p = p[:room]
		c.truncated = true
	}
	c.b = append(c.b, p...)
	return n, nil
}
func contained(root, p string) (string, error) {
	r, err := filepath.EvalSymlinks(root)
	if err != nil {
		return "", err
	}
	p, err = filepath.EvalSymlinks(p)
	if err != nil {
		return "", err
	}
	rel, err := filepath.Rel(r, p)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", fmt.Errorf("path escaped disposable root")
	}
	return p, nil
}

func goExecutable() string {
	name := "go"
	if runtime.GOOS == "windows" {
		name += ".exe"
	}
	return filepath.Join(runtime.GOROOT(), "bin", name)
}

func runProcess(env *isolation.Environment, cwd, executable string, args []string, stdin []byte, mode isolation.ExecutionMode, deadline time.Duration, cap int, native bool, extra map[string]string) (Process, []byte, error) {
	r := Process{Args: args, DeadlineMillis: deadline.Milliseconds(), Exit: -100, Cap: cap}
	var err error
	r.Cwd, err = contained(env.Root, cwd)
	if err != nil {
		return r, nil, err
	}
	if native {
		r.Executable, err = contained(env.Root, executable)
	} else {
		if executable != goExecutable() {
			return r, nil, fmt.Errorf("converter must use pinned Go")
		}
		r.Executable, err = filepath.EvalSymlinks(executable)
	}
	if err != nil {
		return r, nil, err
	}
	if cap < 1 || deadline <= 0 {
		return r, nil, fmt.Errorf("invalid capture cap/deadline")
	}
	r.Env, err = env.GoEnv(isolation.RunGoOptions{Mode: mode})
	if err != nil {
		return r, nil, err
	}
	if mode == isolation.ExecutionModeRace {
		r.Env = append(r.Env, "GORACE=halt_on_error=1 exitcode=98")
	}
	for k, v := range extra {
		if k != "GO_LEARN_CODEC_TRACE" {
			return r, nil, fmt.Errorf("unsupported observer environment key")
		}
		r.Env = append(r.Env, k+"="+v)
	}
	sort.Strings(r.Env)
	ctx, cancel := context.WithTimeout(context.Background(), deadline)
	defer cancel()
	cmd := exec.CommandContext(ctx, r.Executable, args...)
	cmd.Dir = r.Cwd
	cmd.Env = r.Env
	cmd.Stdin = bytes.NewReader(stdin)
	output := &capture{limit: cap}
	cmd.Stdout = output
	cmd.Stderr = output
	start := time.Now()
	runErr := cmd.Run()
	r.Seconds = time.Since(start).Seconds()
	r.TimedOut = ctx.Err() != nil
	if cmd.ProcessState != nil {
		r.Exit = cmd.ProcessState.ExitCode()
	}
	if runErr != nil {
		r.Error = runErr.Error()
	}
	r.Truncated = output.truncated
	s := sha256.Sum256(output.b)
	r.OutputSHA256 = hex.EncodeToString(s[:])
	return r, output.b, nil
}

// ReadTrace refuses symlinks and reads at most limit+1 bytes after process exit.
func readTrace(env *isolation.Environment, path string, limit int) ([]byte, bool, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return nil, false, err
	}
	if !info.Mode().IsRegular() {
		return nil, false, fmt.Errorf("trace is not a regular file")
	}
	path, err = contained(env.Root, path)
	if err != nil {
		return nil, false, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, false, err
	}
	defer f.Close()
	b := make([]byte, limit+1)
	n := 0
	for n < len(b) {
		got, e := f.Read(b[n:])
		n += got
		if e != nil {
			if e == io.EOF {
				break
			}
			return nil, false, e
		}
		if got == 0 {
			return nil, false, fmt.Errorf("trace read made no progress")
		}
	}
	if n > limit {
		return b[:limit], true, nil
	}
	return b[:n], false, nil
}
