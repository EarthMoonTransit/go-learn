package trace

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"go-practice-bank/internal/fuzznative/frame"
	"go-practice-bank/internal/fuzznative/observe"
)

var testIdentity = frame.Identity{ModulePath: "task.local/trace-hex-001", Grammar: "hex-codec-v1", Target: "FuzzHex"}

// This is an offline protocol fixture, never evidence of an actual Go run.
func transcripts(t *testing.T, assertion bool) ([][]byte, []observe.Record, []nativeEvent) {
	t.Helper()
	name := Parent + "/seed#0"
	seeds := [][]byte{{0x80}}
	records := []observe.Record{
		{Kind: "add", Name: Parent, Input: seeds[0]},
		{Kind: "register", Name: Parent, Count: 1},
		{Kind: "begin", Name: name, Input: seeds[0]},
		{Kind: "api-entry", Name: name, Input: seeds[0], Method: "Encode", Invocation: 1},
		{Kind: "api-return", Name: name, Output: []byte("80"), Method: "Encode", Invocation: 1},
		{Kind: "api-entry", Name: name, Input: []byte("80"), Method: "Decode", Invocation: 2},
		{Kind: "api-return", Name: name, Output: seeds[0], Method: "Decode", Invocation: 2, OK: true},
	}
	events := []nativeEvent{{Action: "start", Package: testIdentity.SubjectPath()}, {Action: "run", Package: testIdentity.SubjectPath(), Test: Parent}, {Action: "run", Package: testIdentity.SubjectPath(), Test: name}}
	assertions := 0
	endAction := "pass"
	if assertion {
		assertions = 1
		endAction = "fail"
		records = append(records, observe.Record{Kind: "assert", Name: name, Count: 1, Method: "Error"})
		events = append(events, nativeEvent{Action: "output", Package: testIdentity.SubjectPath(), Test: name, Output: "    fuzz_test.go:9: decoded value mismatch\n", OutputType: "error"})
	}
	records = append(records, observe.Record{Kind: "end", Name: name, Count: 2, Assertions: assertions, End: "normal"}, observe.Record{Kind: "parent-return", Name: Parent, Count: 1})
	for i := range records {
		records[i].Version = observe.Version
		records[i].Seq = i + 1
		records[i].Trusted = true
	}
	events = append(events, nativeEvent{Action: endAction, Package: testIdentity.SubjectPath(), Test: name}, nativeEvent{Action: endAction, Package: testIdentity.SubjectPath(), Test: Parent}, nativeEvent{Action: endAction, Package: testIdentity.SubjectPath()})
	return seeds, records, events
}

func jsonLines[T any](t *testing.T, values []T) []byte {
	t.Helper()
	var out []byte
	for _, v := range values {
		line, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		out = append(out, line...)
		out = append(out, '\n')
	}
	return out
}

func TestCompleteNativeAgreementIsRequiredForDetection(t *testing.T) {
	for _, assertion := range []bool{false, true} {
		seeds, records, events := transcripts(t, assertion)
		exit := 0
		if assertion {
			exit = 1
		}
		observer, native := jsonLines(t, records), jsonLines(t, events)
		result := Validate(testIdentity, seeds, observer, native, exit, 0, false, false)
		if !result.Valid || result.Detected != assertion || len(result.Canonical) != 64 || result.Callbacks != 1 || result.Entries != 2 {
			t.Fatalf("complete protocol: %+v", result)
		}
		for _, fault := range []struct {
			name               string
			raw, converted     int
			timeout, truncated bool
			observer, native   []byte
		}{
			{"panic-after-assert", 2, 0, false, false, observer, native},
			{"race-after-assert", 98, 0, false, false, observer, native},
			{"observer-fault", 97, 0, false, false, observer, native},
			{"converter-fault", exit, 1, false, false, observer, native},
			{"deadline", exit, 0, true, false, observer, native},
			{"capture-limit", exit, 0, false, true, observer, native},
			{"missing-observer-finality", exit, 0, false, false, jsonLines(t, records[:len(records)-1]), native},
			{"missing-native-finality", exit, 0, false, false, observer, jsonLines(t, events[:len(events)-1])},
			{"partial-line", exit, 0, false, false, observer[:len(observer)-1], native},
		} {
			t.Run(fault.name+map[bool]string{false: "/pass", true: "/assertion"}[assertion], func(t *testing.T) {
				got := Validate(testIdentity, seeds, fault.observer, fault.native, fault.raw, fault.converted, fault.timeout, fault.truncated)
				if got.Valid || got.Detected || got.Canonical != "" || got.Diagnostic == "" {
					t.Fatalf("invalid raw/protocol credited: %+v", got)
				}
			})
		}
		mismatchExit := 1 - exit
		if got := Validate(testIdentity, seeds, observer, native, mismatchExit, 0, false, false); got.Valid || got.Detected {
			t.Fatalf("raw/native disagreement credited: %+v", got)
		}
	}
}

func TestObserverAndNativeFieldsAreNotAdvisory(t *testing.T) {
	for _, test := range []struct {
		name   string
		mutate func([]observe.Record, []nativeEvent)
	}{
		{"untrusted-entry", func(r []observe.Record, _ []nativeEvent) { r[3].Trusted = false }},
		{"seed-replaced", func(r []observe.Record, _ []nativeEvent) { r[2].Input = []byte{0} }},
		{"wrong-seed-call", func(r []observe.Record, _ []nativeEvent) { r[3].Input = []byte{0} }},
		{"orphan-decode", func(r []observe.Record, _ []nativeEvent) { r[5].Input = []byte("81") }},
		{"return-invocation", func(r []observe.Record, _ []nativeEvent) { r[6].Invocation = 1 }},
		{"unexpected-end", func(r []observe.Record, _ []nativeEvent) { r[len(r)-2].End = "unexpected" }},
		{"assert-count", func(r []observe.Record, _ []nativeEvent) { r[len(r)-2].Assertions = 0 }},
		{"irrelevant-field", func(r []observe.Record, _ []nativeEvent) { r[0].End = "normal" }},
		{"native-package", func(_ []observe.Record, e []nativeEvent) { e[1].Package = "task.local/other-001/subject" }},
		{"assertion-hidden", func(_ []observe.Record, e []nativeEvent) { e[3].OutputType = "frame" }},
		{"native-child-pass", func(_ []observe.Record, e []nativeEvent) { e[4].Action = "pass" }},
		{"unsupported-native-action", func(_ []observe.Record, e []nativeEvent) { e[3].Action = "skip" }},
	} {
		t.Run(test.name, func(t *testing.T) {
			seeds, records, events := transcripts(t, true)
			test.mutate(records, events)
			got := Validate(testIdentity, seeds, jsonLines(t, records), jsonLines(t, events), 1, 0, false, false)
			if got.Valid || got.Detected {
				t.Fatalf("mismatched fields credited: %+v", got)
			}
		})
	}
	seeds, records, events := transcripts(t, true)
	observer, native := jsonLines(t, records), jsonLines(t, events)
	for _, bad := range [][]byte{
		bytes.Replace(observer, []byte(`"kind":"add"`), []byte(`"kind":"add","kind":"add"`), 1),
		bytes.Replace(observer, []byte(`"version":3`), []byte(`"Version":3`), 1),
		bytes.Replace(observer, []byte(`"count":0`), []byte(`"unknown":0,"count":0`), 1),
	} {
		if got := Validate(testIdentity, seeds, bad, native, 1, 0, false, false); got.Valid || got.Detected {
			t.Fatalf("noncanonical observer JSON credited: %+v", got)
		}
	}
}

func TestCanonicalValuesNormalizeOnlyRepresentationAndDiagnosticText(t *testing.T) {
	seeds, records, events := transcripts(t, true)
	original := Validate(testIdentity, seeds, jsonLines(t, records), jsonLines(t, events), 1, 0, false, false)
	events[3].Output = "    other_path.go:99: another diagnostic\n"
	changedText := Validate(testIdentity, seeds, jsonLines(t, records), jsonLines(t, events), 1, 0, false, false)
	if !original.Valid || !changedText.Valid || original.Canonical != changedText.Canonical {
		t.Fatal("diagnostic text changed semantic determinism")
	}
	records[6].OK = false
	changedValue := Validate(testIdentity, seeds, jsonLines(t, records), jsonLines(t, events), 1, 0, false, false)
	if !changedValue.Valid || changedValue.Canonical == original.Canonical {
		t.Fatal("Decode success bit disappeared from determinism")
	}
	seeds, records, events = transcripts(t, false)
	seeds[0] = nil
	for _, index := range []int{0, 2, 3} {
		records[index].Input = nil
	}
	records[4].Output = nil
	records[5].Input = nil
	records[6].Output = nil
	nilResult := Validate(testIdentity, seeds, jsonLines(t, records), jsonLines(t, events), 0, 0, false, false)
	seeds[0] = []byte{}
	for _, index := range []int{0, 2, 3} {
		records[index].Input = []byte{}
	}
	records[4].Output = []byte{}
	records[5].Input = []byte{}
	records[6].Output = []byte{}
	emptyResult := Validate(testIdentity, seeds, jsonLines(t, records), jsonLines(t, events), 0, 0, false, false)
	if !nilResult.Valid || !emptyResult.Valid || nilResult.Canonical != emptyResult.Canonical {
		t.Fatalf("nil/empty value equivalence lost: %+v %+v", nilResult, emptyResult)
	}
	if strings.TrimSpace(nilResult.Canonical) == "" {
		t.Fatal("empty canonical outcome")
	}
}
