package tasking

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLearnerAPISourceSupportsMethodsAndFields(t *testing.T) {
	for _, fixture := range []struct {
		name, imports, body string
		wantMissing         bool
	}{
		{"free function", `"task.local/example/subject"`, `if subject.Value()!=1 {t.Error("value")}`, false},
		{"receiver method", `"task.local/example/subject"`, `d:=subject.Dictionary{};if err:=d.Add("k","v");err!=nil {t.Error(err)}`, false},
		{"pointer receiver", `"task.local/example/subject"`, `w:=&subject.Wallet{};w.Deposit(3);if w.Balance()!=3 {t.Error("balance")}`, false},
		{"aliased import", `api "task.local/example/subject"`, `d:=api.Dictionary{};if err:=d.Add("k","v");err!=nil {t.Error(err)}`, false},
		{"method expression", `"task.local/example/subject"`, `add:=subject.Dictionary.Add;if err:=add(nil,"k","v");err!=nil {t.Error(err)}`, false},
		{"exported field", `"task.local/example/subject"`, `v:=subject.Record{Count:3};if v.Count!=3 {t.Error("count")}`, false},
		{"no task usage", `"strings"`, `if strings.TrimSpace(" a ")!="a" {t.Error("trim")}`, true},
		{"unused task import", `"task.local/example/subject"`, `if 1!=1 {t.Error("constant")}`, true},
	} {
		t.Run(fixture.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "submission_test.go")
			content := `package subject_test
import("testing";` + fixture.imports + `)
func TestContract(t *testing.T){` + fixture.body + `}
`
			if err := os.WriteFile(path, []byte(content), 0600); err != nil {
				t.Fatal(err)
			}
			problems := inspectLearnerTestSource(path)
			if missing := containsText(problems, "must reference the task public API"); missing != fixture.wantMissing {
				t.Fatalf("missing=%v want=%v problems=%v", missing, fixture.wantMissing, problems)
			}
			if !fixture.wantMissing && len(problems) != 0 {
				t.Fatalf("valid API usage rejected: %v", problems)
			}
		})
	}
}

func TestLearnerMethodAPIGlobalFixture(t *testing.T) {
	path := filepath.Join(t.TempDir(), "submission_test.go")
	source := `package subject_test
 import("testing";"task.local/example/subject")
 var shared=subject.Dictionary{}
 func TestAdd(t *testing.T){if err:=shared.Add("k","v");err!=nil{t.Error(err)}}
 `
	if err := os.WriteFile(path, []byte(source), 0600); err != nil {
		t.Fatal(err)
	}
	if problems := inspectLearnerTestSource(path); len(problems) != 0 {
		t.Fatalf("global task fixture rejected: %v", problems)
	}
}
