// Package protocol defines the bounded trusted frontend helper protocol.
package protocol

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"regexp"

	"go-practice-bank/internal/fuzznative/frame"
)

const Version = 5
const MaxRequest = 2 << 20
const MaxResponse = 4 << 20

type Request struct {
	Identity     frame.Identity    `json:"identity"`
	Runtime      Runtime           `json:"runtime"`
	KernelDigest string            `json:"kernel_digest"`
	Version      int               `json:"version"`
	Sources      map[string]string `json:"sources"`
	Exports      map[string]Export `json:"exports"`
}
type Response struct {
	Identity      frame.Identity    `json:"identity"`
	Version       int               `json:"version"`
	Runtime       Runtime           `json:"runtime"`
	KernelDigest  string            `json:"kernel_digest"`
	RequestSHA256 string            `json:"request_sha256"`
	SourceSHA256  map[string]string `json:"source_sha256"`
	Accepted      bool              `json:"accepted"`
	Diagnostic    string            `json:"diagnostic"`
	Metadata      *frame.Inspection `json:"metadata"`
	Rewritten     map[string]string `json:"rewritten"`
	Methods       map[string]int    `json:"methods"`
	Alias         string            `json:"alias"`
}

var runtimePattern = regexp.MustCompile(`^go1\.27(?:\.[0-9]+)?$`)
var digestPattern = regexp.MustCompile(`^[0-9a-f]{64}$`)

func (r Request) ValidateIdentity() error {
	if r.Version != Version || r.Runtime.Validate() != nil || !digestPattern.MatchString(r.KernelDigest) {
		return fmt.Errorf("helper request version/runtime/source binding")
	}
	return r.Identity.Validate()
}

func Hash(b []byte) string { s := sha256.Sum256(b); return hex.EncodeToString(s[:]) }
func SourceHashes(sources map[string]string) map[string]string {
	out := map[string]string{}
	for name, s := range sources {
		out[name] = Hash([]byte(s))
	}
	return out
}

// Decode rejects duplicate keys at every level, unknown/case-aliased struct
// fields, trailing JSON and excessive nesting before decoding the typed value.
func Decode(b []byte, target any) error {
	if !json.Valid(b) {
		return fmt.Errorf("invalid JSON")
	}
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	var value func(int) error
	value = func(depth int) error {
		if depth > 32 {
			return fmt.Errorf("JSON nesting limit")
		}
		token, err := dec.Token()
		if err != nil {
			return err
		}
		switch token {
		case json.Delim('{'):
			seen := map[string]bool{}
			for dec.More() {
				key, err := dec.Token()
				if err != nil {
					return err
				}
				s, ok := key.(string)
				if !ok || seen[s] {
					return fmt.Errorf("duplicate JSON key")
				}
				seen[s] = true
				if err := value(depth + 1); err != nil {
					return err
				}
			}
			end, err := dec.Token()
			if err != nil || end != json.Delim('}') {
				return fmt.Errorf("JSON object end")
			}
		case json.Delim('['):
			for dec.More() {
				if err := value(depth + 1); err != nil {
					return err
				}
			}
			end, err := dec.Token()
			if err != nil || end != json.Delim(']') {
				return fmt.Errorf("JSON array end")
			}
		case json.Delim('}'), json.Delim(']'):
			return fmt.Errorf("unexpected JSON delimiter")
		}
		return nil
	}
	if err := value(0); err != nil {
		return err
	}
	if _, err := dec.Token(); err != io.EOF {
		return fmt.Errorf("trailing JSON")
	}
	// DisallowUnknownFields is case-insensitive. Enforce canonical exact field
	// spelling by round-tripping typed objects through their marshaled key tree.
	dec = json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields()
	if err := dec.Decode(target); err != nil {
		return err
	}
	canonical, err := json.Marshal(target)
	if err != nil {
		return err
	}
	var got, want any
	mustDecode := func(data []byte, v *any) error {
		d := json.NewDecoder(bytes.NewReader(data))
		d.UseNumber()
		return d.Decode(v)
	}
	if err := mustDecode(b, &got); err != nil {
		return err
	}
	if err := mustDecode(canonical, &want); err != nil {
		return err
	}
	return exactKeys(got, want)
}
func exactKeys(got, want any) error {
	switch actual := got.(type) {
	case map[string]any:
		expected, ok := want.(map[string]any)
		if !ok || len(actual) != len(expected) {
			return fmt.Errorf("missing/unknown canonical JSON keys")
		}
		for key, v := range actual {
			w, ok := expected[key]
			if !ok {
				return fmt.Errorf("noncanonical JSON key %s", key)
			}
			if err := exactKeys(v, w); err != nil {
				return err
			}
		}
	case []any:
		expected, ok := want.([]any)
		if !ok || len(actual) != len(expected) {
			return fmt.Errorf("JSON array type")
		}
		for i, v := range actual {
			if err := exactKeys(v, expected[i]); err != nil {
				return err
			}
		}
	case nil:
		if want != nil {
			return fmt.Errorf("null scalar or container")
		}
	}
	return nil
}
func ValidateResponse(raw []byte, requestBytes []byte, request Request) (Response, error) {
	var r Response
	if err := request.ValidateIdentity(); err != nil {
		return r, err
	}
	if len(raw) == 0 || len(raw) > MaxResponse || raw[len(raw)-1] != '\n' {
		return r, fmt.Errorf("incomplete/over-cap helper result")
	}
	if err := Decode(raw, &r); err != nil {
		return Response{}, err
	}
	if r.Version != Version || r.Runtime != request.Runtime || r.KernelDigest != request.KernelDigest || r.Identity != request.Identity || r.RequestSHA256 != Hash(requestBytes) || !maps.Equal(r.SourceSHA256, SourceHashes(request.Sources)) {
		return Response{}, fmt.Errorf("helper result version/runtime/original binding")
	}
	if !r.Accepted {
		if r.Diagnostic == "" || r.Metadata != nil || r.Rewritten != nil || r.Methods != nil || r.Alias != "" {
			return Response{}, fmt.Errorf("invalid rejected helper result")
		}
		return r, nil
	}
	if r.Diagnostic != "" || r.Metadata == nil || r.Metadata.SeedCount < 1 || r.Metadata.SeedCount > 64 || r.Metadata.SeedCount != len(r.Metadata.SeedValues) || r.Metadata.HelperCount < 0 || r.Metadata.HelperCount > 16 || r.Alias == "" || len(r.Rewritten) != len(request.Sources) || r.Methods["Add"] < 1 || r.Methods["Register"] != 1 || r.Methods["Encode"]+r.Methods["WrapEncode"] < 1 || r.Methods["Decode"]+r.Methods["WrapDecode"] < 1 {
		return Response{}, fmt.Errorf("invalid admitted helper result")
	}
	if r.Metadata.AnalysisSteps <= 0 || r.Metadata.AnalysisSteps > frame.MaxSteps {
		return Response{}, fmt.Errorf("admitted analysis step bound")
	}
	for _, seed := range r.Metadata.SeedValues {
		if len(seed) > frame.MaxLiteralLength {
			return Response{}, fmt.Errorf("admitted seed byte bound")
		}
	}
	for name := range request.Sources {
		if r.Rewritten[name] == "" {
			return Response{}, fmt.Errorf("missing rewritten original file")
		}
	}
	for method, count := range r.Methods {
		switch method {
		case "Add", "Register", "Error", "Errorf", "Fatal", "Fatalf", "Encode", "Decode", "WrapEncode", "WrapDecode":
		default:
			return Response{}, fmt.Errorf("unknown rewrite method")
		}
		if count < 0 {
			return Response{}, fmt.Errorf("negative rewrite method count")
		}
	}
	return r, nil
}
