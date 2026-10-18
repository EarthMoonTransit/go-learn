// Package observe is a native hex codec boundary recorder. Its source
// is copied into a disposable module and is added only after original-source
// admission; learner source cannot import this package in the input grammar.
package observe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"testing"
)

const Version = 3
const MaxTraceBytes = 16 << 20
const MaxRecords = 8192
const MaxAPIValue = 8192

type Record struct {
	Version    int    `json:"version"`
	Seq        int    `json:"seq"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Input      []byte `json:"input"`
	Output     []byte `json:"output"`
	Count      int    `json:"count"`
	Invocation int    `json:"invocation"`
	Method     string `json:"method"`
	OK         bool   `json:"ok"`
	Assertions int    `json:"assertions"`
	Trusted    bool   `json:"trusted"`
	End        string `json:"end"`
}

type callbackState struct {
	t                 *testing.T
	name              string
	calls, assertions int
	fatal             bool
}

var state struct {
	sync.Mutex
	seq, written, adds, callbacks int
	parent                        *testing.F
	registered                    bool
	active                        *callbackState
}

func ioFailure(err error) {
	fmt.Fprintln(os.Stderr, "native codec observer failure:", err)
	os.Exit(97)
}

func emitLocked(r Record) {
	state.seq++
	r.Version, r.Seq = Version, state.seq
	b, err := json.Marshal(r)
	if err != nil {
		ioFailure(err)
	}
	b = append(b, '\n')
	if state.seq > MaxRecords || state.written+len(b) > MaxTraceBytes {
		ioFailure(fmt.Errorf("native codec trace capture cap"))
	}
	f, err := os.OpenFile(os.Getenv("GO_LEARN_CODEC_TRACE"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		ioFailure(err)
	}
	n, writeErr := f.Write(b)
	closeErr := f.Close()
	if writeErr != nil || closeErr != nil || n != len(b) {
		ioFailure(fmt.Errorf("trace write %d/%d: %v; close: %v", n, len(b), writeErr, closeErr))
	}
	state.written += n
}

func Add(f *testing.F, input []byte) {
	state.Lock()
	if state.parent == nil {
		state.parent = f
	}
	trusted := state.parent == f && !state.registered && state.active == nil
	emitLocked(Record{Kind: "add", Name: f.Name(), Input: bytes.Clone(input), Count: state.adds, Trusted: trusted})
	state.adds++
	state.Unlock()
	f.Add(input)
}

func Register(f *testing.F, callback func(*testing.T, []byte)) {
	state.Lock()
	trusted := state.parent == f && !state.registered && state.active == nil
	emitLocked(Record{Kind: "register", Name: f.Name(), Count: state.adds, Trusted: trusted})
	state.registered = true
	state.Unlock()
	f.Fuzz(func(t *testing.T, input []byte) {
		state.Lock()
		trusted := state.active == nil
		c := &callbackState{t: t, name: t.Name()}
		state.active = c
		emitLocked(Record{Kind: "begin", Name: c.name, Input: bytes.Clone(input), Count: state.callbacks, Trusted: trusted})
		state.Unlock()
		returned := false
		defer func() {
			problem := recover()
			state.Lock()
			end := "unexpected"
			if problem != nil {
				end = "panic"
			} else if returned {
				end = "normal"
			} else if c.fatal {
				end = "fatal"
			}
			emitLocked(Record{Kind: "end", Name: c.name, Count: c.calls, Assertions: c.assertions, Trusted: state.active == c, End: end})
			state.active = nil
			state.callbacks++
			state.Unlock()
			if problem != nil {
				panic(problem)
			}
		}()
		callback(t, input)
		returned = true
	})
	state.Lock()
	emitLocked(Record{Kind: "parent-return", Name: f.Name(), Count: state.callbacks, Trusted: state.parent == f && state.registered && state.active == nil})
	state.Unlock()
}

func apiEntry(method string, input []byte) (int, *callbackState) {
	state.Lock()
	defer state.Unlock()
	c := state.active
	name := ""
	trusted := false
	invocation := 0
	if c != nil {
		name = c.name
		c.calls++
		invocation = c.calls
		trusted = true
	}
	emitLocked(Record{Kind: "api-entry", Name: name, Input: bytes.Clone(input), Invocation: invocation, Method: method, Trusted: trusted})
	return invocation, c
}

func apiReturn(method string, output []byte, ok bool, invocation int, c *callbackState) {
	if len(output) > MaxAPIValue {
		ioFailure(fmt.Errorf("codec API output exceeds %d bytes", MaxAPIValue))
	}
	state.Lock()
	defer state.Unlock()
	trusted := c != nil && state.active == c
	name := ""
	if c != nil {
		name = c.name
	}
	emitLocked(Record{Kind: "api-return", Name: name, Output: bytes.Clone(output), Invocation: invocation, Method: method, OK: ok, Trusted: trusted})
}

// Encode and Decode wrap the original, statically resolved instructor API
// function value. The trace records actual argument/result values at the call
// boundary, including the Decode success bit.
func Encode(fn func([]byte) []byte, input []byte) []byte {
	id, c := apiEntry("Encode", input)
	output := fn(input)
	apiReturn("Encode", output, false, id, c)
	return output
}

func Decode(fn func([]byte) ([]byte, bool), input []byte) ([]byte, bool) {
	id, c := apiEntry("Decode", input)
	output, ok := fn(input)
	apiReturn("Decode", output, ok, id, c)
	return output, ok
}

func WrapEncode(fn func([]byte) []byte) func([]byte) []byte {
	return func(input []byte) []byte { return Encode(fn, input) }
}

func WrapDecode(fn func([]byte) ([]byte, bool)) func([]byte) ([]byte, bool) {
	return func(input []byte) ([]byte, bool) { return Decode(fn, input) }
}

func assertion(t *testing.T, method string) {
	state.Lock()
	defer state.Unlock()
	name := ""
	if t != nil {
		name = t.Name()
	}
	c := state.active
	if c == nil {
		emitLocked(Record{Kind: "assert", Name: name, Method: method})
		return
	}
	c.assertions++
	if method == "Fatal" || method == "Fatalf" {
		c.fatal = true
	}
	emitLocked(Record{Kind: "assert", Name: name, Count: c.assertions, Method: method, Trusted: t == c.t})
}

func Error(t *testing.T, args ...any) { assertion(t, "Error"); t.Helper(); t.Error(args...) }
func Errorf(t *testing.T, format string, args ...any) {
	assertion(t, "Errorf")
	t.Helper()
	t.Errorf(format, args...)
}
func Fatal(t *testing.T, args ...any) { assertion(t, "Fatal"); t.Helper(); t.Fatal(args...) }
func Fatalf(t *testing.T, format string, args ...any) {
	assertion(t, "Fatalf")
	t.Helper()
	t.Fatalf(format, args...)
}
