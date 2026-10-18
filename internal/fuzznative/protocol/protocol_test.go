package protocol

import (
	"encoding/json"
	"strings"
	"testing"

	"go-practice-bank/internal/fuzznative/frame"
)

func rejection(t *testing.T) (Request, []byte, Response) {
	t.Helper()
	request := Request{Version: Version, Identity: frame.Identity{ModulePath: "task.local/protocol-hex-001", Grammar: "hex-codec-v1", Target: "FuzzHex"}, Runtime: Runtime{Version: "go1.27.1", GOOS: "linux", GOARCH: "amd64"}, KernelDigest: strings.Repeat("a", 64), Sources: map[string]string{"fuzz_test.go": "package subject_test"}, Exports: map[string]Export{}}
	raw, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
	}
	raw = append(raw, '\n')
	response := Response{Version: Version, Identity: request.Identity, Runtime: request.Runtime, KernelDigest: request.KernelDigest, RequestSHA256: Hash(raw), SourceSHA256: SourceHashes(request.Sources), Diagnostic: "frame.import: missing testing import"}
	return request, raw, response
}

func TestResponseBindsExactRequestAndExecutionIdentity(t *testing.T) {
	for _, test := range []struct {
		name   string
		mutate func(*Response)
	}{
		{"matching-rejection", func(*Response) {}},
		{"helper-version", func(r *Response) { r.Version-- }},
		{"runtime-version", func(r *Response) { r.Runtime.Version = "go1.27.2" }},
		{"runtime-platform", func(r *Response) { r.Runtime.GOOS = "darwin" }},
		{"kernel", func(r *Response) { r.KernelDigest = strings.Repeat("b", 64) }},
		{"module", func(r *Response) { r.Identity.ModulePath = "task.local/another-001" }},
		{"target", func(r *Response) { r.Identity.Target = "FuzzOther" }},
		{"grammar", func(r *Response) { r.Identity.Grammar = "other" }},
		{"request-bytes", func(r *Response) { r.RequestSHA256 = strings.Repeat("0", 64) }},
		{"source-bytes", func(r *Response) { r.SourceSHA256["fuzz_test.go"] = strings.Repeat("0", 64) }},
		{"extra-source", func(r *Response) { r.SourceSHA256["extra_test.go"] = Hash(nil) }},
		{"malformed-acceptance", func(r *Response) { r.Accepted = true; r.Diagnostic = "" }},
		{"silent-rejection", func(r *Response) { r.Diagnostic = "" }},
		{"rejected-rewrite", func(r *Response) { r.Rewritten = map[string]string{"fuzz_test.go": "package subject_test"} }},
	} {
		t.Run(test.name, func(t *testing.T) {
			request, raw, response := rejection(t)
			test.mutate(&response)
			output, err := json.Marshal(response)
			if err != nil {
				t.Fatal(err)
			}
			output = append(output, '\n')
			_, err = ValidateResponse(output, raw, request)
			if test.name == "matching-rejection" {
				if err != nil {
					t.Fatal(err)
				}
			} else if err == nil {
				t.Fatal("accepted mismatched helper response")
			}
		})
	}
}

func TestClosedJSONAndCompleteResponse(t *testing.T) {
	request, raw, response := rejection(t)
	encoded, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	for _, data := range [][]byte{
		encoded,
		[]byte(strings.Replace(string(encoded), `"version":5`, `"Version":5`, 1) + "\n"),
		[]byte(strings.Replace(string(encoded), `"version":5`, `"version":5,"version":5`, 1) + "\n"),
		[]byte(strings.Replace(string(encoded), `"identity":{`, `"identity":{"unknown":true,`, 1) + "\n"),
		[]byte(strings.Replace(string(encoded), `"goos":"linux"`, `"goos":"linux","goos":"linux"`, 1) + "\n"),
		append(append([]byte(nil), encoded...), []byte("\n{}\n")...),
	} {
		if _, err := ValidateResponse(data, raw, request); err == nil {
			t.Fatalf("accepted incomplete/noncanonical response: %.160s", data)
		}
	}
}

func TestRuntimeRequiresPinnedFamilyAndConcreteTarget(t *testing.T) {
	valid := Runtime{Version: "go1.27.1", GOOS: "linux", GOARCH: "amd64"}
	if err := valid.Validate(); err != nil {
		t.Fatal(err)
	}
	for _, r := range []Runtime{{Version: "go1.28.1", GOOS: "linux", GOARCH: "amd64"}, {Version: "devel go1.27", GOOS: "linux", GOARCH: "amd64"}, {Version: "go1.27.1", GOOS: "", GOARCH: "amd64"}, {Version: "go1.27.1", GOOS: "linux", GOARCH: "../amd64"}} {
		if r.Validate() == nil {
			t.Fatalf("accepted unsupported runtime %+v", r)
		}
	}
}
