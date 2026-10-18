// Package trace validates the complete bounded native seed protocol.
// Source admission is a separate prerequisite and cannot be inferred here.
package trace

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	"go-practice-bank/internal/fuzznative/frame"
	"go-practice-bank/internal/fuzznative/observe"
)

const Parent = "FuzzHex"
const MaxNativeBytes = 8 << 20
const MaxConvertedBytes = 32 << 20

type Outcome struct {
	Valid      bool   `json:"valid"`
	Detected   bool   `json:"detected"`
	Diagnostic string `json:"diagnostic,omitempty"`
	Callbacks  int    `json:"callbacks"`
	Entries    int    `json:"entries"`
	Assertions int    `json:"assertions"`
	Canonical  string `json:"canonical,omitempty"`
}

// object checks duplicate keys, exact key spelling and one complete flat JSON
// object. This protocol deliberately has no nested objects or arrays.
func object(line []byte, allowed, required map[string]bool, target any) error {
	dec := json.NewDecoder(bytes.NewReader(line))
	t, err := dec.Token()
	if err != nil || t != json.Delim('{') {
		return fmt.Errorf("expected JSON object")
	}
	seen := map[string]bool{}
	for dec.More() {
		key, err := dec.Token()
		if err != nil {
			return err
		}
		k, ok := key.(string)
		if !ok || !allowed[k] || seen[k] {
			return fmt.Errorf("unknown/duplicate JSON key %q", key)
		}
		seen[k] = true
		value, err := dec.Token()
		if err != nil {
			return err
		}
		if value == nil && k != "input" && k != "output" {
			return fmt.Errorf("null protocol scalar %s", k)
		}
		if _, nested := value.(json.Delim); nested {
			return fmt.Errorf("nested protocol value")
		}
	}
	if _, err = dec.Token(); err != nil {
		return err
	}
	if _, err = dec.Token(); err != io.EOF {
		return fmt.Errorf("trailing JSON tokens")
	}
	for k := range required {
		if !seen[k] {
			return fmt.Errorf("missing JSON key %s", k)
		}
	}
	dec = json.NewDecoder(bytes.NewReader(line))
	dec.DisallowUnknownFields()
	if err := dec.Decode(target); err != nil {
		return err
	}
	return nil
}
func keys(names ...string) map[string]bool {
	out := map[string]bool{}
	for _, n := range names {
		out[n] = true
	}
	return out
}
func lines(data []byte, cap int) ([][]byte, error) {
	if len(data) == 0 || len(data) > cap || data[len(data)-1] != '\n' {
		return nil, fmt.Errorf("empty, over-cap or incomplete JSONL")
	}
	out := bytes.Split(data[:len(data)-1], []byte{'\n'})
	for _, line := range out {
		if len(line) == 0 {
			return nil, fmt.Errorf("empty protocol line")
		}
	}
	return out, nil
}

// recordEqual compares the complete record, including fields irrelevant to its
// kind. Only actual byte-value payload slots use nil/empty value equivalence.
func recordEqual(got, want observe.Record) bool {
	if got.Kind != want.Kind || got.Name != want.Name || got.Count != want.Count || got.Invocation != want.Invocation || got.Method != want.Method || got.OK != want.OK || got.Assertions != want.Assertions || got.Trusted != want.Trusted || got.End != want.End {
		return false
	}
	if !bytes.Equal(got.Input, want.Input) || !bytes.Equal(got.Output, want.Output) {
		return false
	}
	if want.Kind != "add" && want.Kind != "begin" && want.Kind != "api-entry" && got.Input != nil {
		return false
	}
	if want.Kind != "api-return" && got.Output != nil {
		return false
	}
	return true
}

type child struct {
	name                                string
	calls, assertions, encodes, decodes int
	fatal                               bool
}

func parseTrace(seeds [][]byte, data []byte) ([]child, error) {
	raw, err := lines(data, observe.MaxTraceBytes)
	if err != nil {
		return nil, err
	}
	if len(raw) > observe.MaxRecords {
		return nil, fmt.Errorf("trace record cap")
	}
	allowed := keys("version", "seq", "kind", "name", "input", "output", "count", "invocation", "method", "ok", "assertions", "trusted", "end")
	var records []observe.Record
	for i, line := range raw {
		var r observe.Record
		if err := object(line, allowed, allowed, &r); err != nil {
			return nil, fmt.Errorf("trace line %d: %w", i+1, err)
		}
		if r.Version != observe.Version || r.Seq != i+1 {
			return nil, fmt.Errorf("trace version/sequence at %d", i+1)
		}
		if len(r.Input) > observe.MaxAPIValue || len(r.Output) > observe.MaxAPIValue {
			return nil, fmt.Errorf("trace payload cap at %d", i+1)
		}
		records = append(records, r)
	}
	cursor := 0
	take := func(want observe.Record) error {
		if cursor >= len(records) || !recordEqual(records[cursor], want) {
			return fmt.Errorf("unexpected/missing %s at record %d", want.Kind, cursor+1)
		}
		cursor++
		return nil
	}
	for i, seed := range seeds {
		if err := take(observe.Record{Kind: "add", Name: Parent, Input: seed, Count: i, Trusted: true}); err != nil {
			return nil, err
		}
	}
	if err := take(observe.Record{Kind: "register", Name: Parent, Count: len(seeds), Trusted: true}); err != nil {
		return nil, err
	}
	var children []child
	for i, seed := range seeds {
		c := child{name: fmt.Sprintf("%s/seed#%d", Parent, i)}
		if err := take(observe.Record{Kind: "begin", Name: c.name, Input: seed, Count: i, Trusted: true}); err != nil {
			return nil, err
		}
		var pending *observe.Record
		var encoded [][]byte
		for cursor < len(records) && records[cursor].Kind != "end" {
			r := records[cursor]
			if c.fatal {
				return nil, fmt.Errorf("event after Fatal at %s", c.name)
			}
			switch r.Kind {
			case "api-entry":
				if pending != nil {
					return nil, fmt.Errorf("nested/pending invocation at %s", c.name)
				}
				c.calls++
				if c.calls > 8 {
					return nil, fmt.Errorf("API quota at %s", c.name)
				}
				switch r.Method {
				case "Encode":
					if !bytes.Equal(r.Input, seed) {
						return nil, fmt.Errorf("Encode argument differs from callback seed")
					}
				case "Decode":
					found := false
					for _, wire := range encoded {
						if bytes.Equal(wire, r.Input) {
							found = true
							break
						}
					}
					if !found {
						return nil, fmt.Errorf("Decode argument lacks completed same-callback Encode return")
					}
				default:
					return nil, fmt.Errorf("unsupported API method")
				}
				if err := take(observe.Record{Kind: "api-entry", Name: c.name, Input: r.Input, Invocation: c.calls, Method: r.Method, Trusted: true}); err != nil {
					return nil, err
				}
				entry := r
				pending = &entry
			case "api-return":
				if pending == nil {
					return nil, fmt.Errorf("orphan API return")
				}
				ok := r.OK
				if pending.Method == "Encode" {
					ok = false
				}
				if err := take(observe.Record{Kind: "api-return", Name: c.name, Output: r.Output, Invocation: pending.Invocation, Method: pending.Method, OK: ok, Trusted: true}); err != nil {
					return nil, err
				}
				if pending.Method == "Encode" {
					c.encodes++
					encoded = append(encoded, bytes.Clone(r.Output))
				} else {
					c.decodes++
				}
				pending = nil
			case "assert":
				if pending != nil {
					return nil, fmt.Errorf("assertion during API invocation")
				}
				c.assertions++
				if c.assertions > 64 {
					return nil, fmt.Errorf("assertion quota at %s", c.name)
				}
				switch r.Method {
				case "Error", "Errorf":
				case "Fatal", "Fatalf":
					c.fatal = true
				default:
					return nil, fmt.Errorf("unsupported assertion")
				}
				if err := take(observe.Record{Kind: "assert", Name: c.name, Count: c.assertions, Method: r.Method, Trusted: true}); err != nil {
					return nil, err
				}
			default:
				return nil, fmt.Errorf("unexpected callback record %s", r.Kind)
			}
		}
		if pending != nil {
			return nil, fmt.Errorf("incomplete API invocation at %s", c.name)
		}
		if c.encodes == 0 || (c.assertions == 0 && c.decodes == 0) {
			return nil, fmt.Errorf("missing completed codec participation at %s", c.name)
		}
		end := "normal"
		if c.fatal {
			end = "fatal"
		}
		if err := take(observe.Record{Kind: "end", Name: c.name, Count: c.calls, Assertions: c.assertions, Trusted: true, End: end}); err != nil {
			return nil, err
		}
		children = append(children, c)
	}
	if err := take(observe.Record{Kind: "parent-return", Name: Parent, Count: len(seeds), Trusted: true}); err != nil {
		return nil, err
	}
	if cursor != len(records) {
		return nil, fmt.Errorf("trace after parent finality")
	}
	return children, nil
}

type nativeEvent struct {
	Action     string
	Package    string
	Test       string
	Output     string
	OutputType string
}

func parseNative(data []byte, children []child, packagePath string) (bool, error) {
	raw, err := lines(data, MaxConvertedBytes)
	if err != nil {
		return false, err
	}
	// This pinned converter runs without -t. Other pinned event fields exist but
	// are ineligible in this exact mode (bench/build/attribute events excluded).
	allowed := keys("Action", "Package", "Test", "Output", "OutputType")
	required := keys("Action", "Package")
	started, parentOpen, parentDone, finished := false, false, false, false
	index, active := 0, false
	errorsSeen := make([]bool, len(children))
	failed := false
	errorContinuation := false
	for i, line := range raw {
		var e nativeEvent
		if err := object(line, allowed, required, &e); err != nil {
			return false, fmt.Errorf("native line %d: %w", i+1, err)
		}
		if e.Package != packagePath || finished {
			return false, fmt.Errorf("native package/finality")
		}
		if e.Action != "output" && (e.Output != "" || e.OutputType != "") {
			return false, fmt.Errorf("native non-output payload")
		}
		switch e.Action {
		case "start":
			if started || e.Test != "" || i != 0 {
				return false, fmt.Errorf("native start")
			}
			started = true
		case "run":
			errorContinuation = false
			if !started || parentDone {
				return false, fmt.Errorf("native run phase")
			}
			if !parentOpen {
				if e.Test != Parent || index != 0 {
					return false, fmt.Errorf("native parent run")
				}
				parentOpen = true
			} else {
				if active || index >= len(children) || e.Test != children[index].name {
					return false, fmt.Errorf("native child run order")
				}
				active = true
			}
		case "pass", "fail":
			errorContinuation = false
			if !started || !parentOpen {
				return false, fmt.Errorf("native end phase")
			}
			if active {
				c := children[index]
				want := "pass"
				if c.assertions > 0 {
					want = "fail"
					failed = true
				}
				if e.Test != c.name || e.Action != want || errorsSeen[index] != (c.assertions > 0) {
					return false, fmt.Errorf("native child/assertion mismatch at %s", c.name)
				}
				active = false
				index++
			} else if !parentDone {
				want := "pass"
				if failed {
					want = "fail"
				}
				if index != len(children) || e.Test != Parent || e.Action != want {
					return false, fmt.Errorf("native parent end")
				}
				parentDone = true
			} else {
				want := "pass"
				if failed {
					want = "fail"
				}
				if e.Test != "" || e.Action != want {
					return false, fmt.Errorf("native package end")
				}
				finished = true
			}
		case "output":
			if !started || e.Output == "" {
				return false, fmt.Errorf("native output phase")
			}
			switch e.OutputType {
			case "", "frame", "error", "error-continue":
			default:
				return false, fmt.Errorf("native OutputType")
			}
			if active {
				if e.Test != children[index].name {
					return false, fmt.Errorf("native output receiver")
				}
				if e.OutputType == "error-continue" && !errorContinuation {
					return false, fmt.Errorf("native orphan error continuation")
				}
				errorContinuation = e.OutputType == "error" || e.OutputType == "error-continue"
				if errorContinuation {
					errorsSeen[index] = true
				}
			} else if parentDone {
				if e.OutputType == "error" || e.OutputType == "error-continue" {
					return false, fmt.Errorf("native package error output")
				}
				if e.Test != "" {
					return false, fmt.Errorf("native package output receiver")
				}
			} else if !parentOpen || e.Test != Parent || e.OutputType == "error" || e.OutputType == "error-continue" {
				return false, fmt.Errorf("native parent output")
			}
		default:
			return false, fmt.Errorf("native unsupported action %q", e.Action)
		}
	}
	if !finished || active || index != len(children) {
		return false, fmt.Errorf("native incomplete finality")
	}
	return failed, nil
}

// Validate returns Detection only after both complete protocols and independent
// raw process status agree. It never treats infrastructure failures as kills.
func Validate(identity frame.Identity, seeds [][]byte, data, native []byte, nativeExit, converterExit int, timedOut, truncated bool) Outcome {
	o := Outcome{}
	fail := func(err error) Outcome { o.Diagnostic = err.Error(); return o }
	if err := identity.Validate(); err != nil {
		return fail(err)
	}
	if len(seeds) < 1 || len(seeds) > 64 {
		return fail(fmt.Errorf("static seed count"))
	}
	if timedOut || truncated || converterExit != 0 || (nativeExit != 0 && nativeExit != 1) {
		return fail(fmt.Errorf("invalid raw process result native=%d converter=%d timeout=%v truncated=%v", nativeExit, converterExit, timedOut, truncated))
	}
	for _, seed := range seeds {
		if len(seed) > 4096 {
			return fail(fmt.Errorf("static seed length"))
		}
	}
	children, err := parseTrace(seeds, data)
	if err != nil {
		return fail(err)
	}
	failed, err := parseNative(native, children, identity.SubjectPath())
	if err != nil {
		return fail(err)
	}
	if (nativeExit == 1) != failed {
		return fail(fmt.Errorf("raw status/native assertion mismatch"))
	}
	canonical, err := canonicalObservation(data, native)
	if err != nil {
		return fail(err)
	}
	o.Valid = true
	o.Detected = failed
	o.Callbacks = len(children)
	o.Canonical = canonical
	for _, c := range children {
		o.Entries += c.calls
		o.Assertions += c.assertions
	}
	return o
}

func canonicalObservation(data, native []byte) (string, error) {
	h := sha256.New()
	for _, line := range bytes.Split(bytes.TrimSuffix(data, []byte{'\n'}), []byte{'\n'}) {
		var r observe.Record
		if err := json.Unmarshal(line, &r); err != nil {
			return "", err
		}
		if len(r.Input) == 0 {
			r.Input = nil
		}
		if len(r.Output) == 0 {
			r.Output = nil
		}
		b, err := json.Marshal(r)
		if err != nil {
			return "", err
		}
		_, _ = h.Write(b)
		_, _ = h.Write([]byte{'\n'})
	}
	_, _ = h.Write([]byte("native\n"))
	for _, line := range bytes.Split(bytes.TrimSuffix(native, []byte{'\n'}), []byte{'\n'}) {
		var event nativeEvent
		if err := json.Unmarshal(line, &event); err != nil {
			return "", err
		}
		b, err := json.Marshal([3]string{event.Action, event.Package, event.Test})
		if err != nil {
			return "", err
		}
		_, _ = h.Write(b)
		_, _ = h.Write([]byte{'\n'})
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
