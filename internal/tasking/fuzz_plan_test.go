package tasking

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func fuzzSchemaTask(t *testing.T) TaskV2 {
	t.Helper()
	var task TaskV2
	path := filepath.Join(repositoryRoot(t), "internal/testdata/tasking/v2/mutation/synthetic-mutation-v2-001/task.yaml")
	if err := decodeStrictFile(path, &task); err != nil {
		t.Fatal(err)
	}
	task.Activity = "fuzz-test-writing"
	task.Verification.Profile = "fuzz-seed-tests"
	task.Verification.Fuzz = &FuzzPlan{Grammar: "hex-codec-v1", Target: "FuzzHex", CorrectVariants: []string{"loop", "nil-empty"}}
	return task
}

func TestFuzzPlanSchemaIsClosedAndProfileSpecific(t *testing.T) {
	root := repositoryRoot(t)
	for _, test := range []struct {
		name   string
		mutate func(*TaskV2)
		valid  bool
	}{
		{"finite-descriptor", func(*TaskV2) {}, true},
		{"missing-descriptor", func(v *TaskV2) { v.Verification.Fuzz = nil }, false},
		{"wrong-grammar", func(v *TaskV2) { v.Verification.Fuzz.Grammar = "arbitrary" }, false},
		{"wrong-target", func(v *TaskV2) { v.Verification.Fuzz.Target = "FuzzOther" }, false},
		{"one-correct-variant", func(v *TaskV2) { v.Verification.Fuzz.CorrectVariants = []string{"loop"} }, false},
		{"five-correct-variants", func(v *TaskV2) { v.Verification.Fuzz.CorrectVariants = []string{"one", "two", "three", "four", "five"} }, false},
		{"duplicate-correct-variant", func(v *TaskV2) { v.Verification.Fuzz.CorrectVariants = []string{"loop", "loop"} }, false},
		{"unsafe-correct-id", func(v *TaskV2) { v.Verification.Fuzz.CorrectVariants = []string{"loop", "../other"} }, false},
		{"active-fuzzing", func(v *TaskV2) { v.Verification.Limits.FuzzSeconds = 1 }, false},
		{"no-mutants", func(v *TaskV2) { v.Verification.RequiredMutants = []string{} }, false},
		{"wrong-activity", func(v *TaskV2) { v.Activity = "test-writing" }, false},
		{"wrong-submission", func(v *TaskV2) { v.Submission.Kind = "implementation" }, false},
		{"descriptor-on-mutation", func(v *TaskV2) { v.Activity = "test-writing"; v.Verification.Profile = "mutation-tests" }, false},
		{"old-mutation-omits-descriptor", func(v *TaskV2) {
			v.Activity = "test-writing"
			v.Verification.Profile = "mutation-tests"
			v.Verification.Fuzz = nil
		}, true},
		{"recursive-packages", func(v *TaskV2) { v.Verification.Packages = []string{"./..."} }, false},
		{"recursive-public", func(v *TaskV2) { v.Verification.PublicChecks = []string{"./..."} }, false},
		{"recursive-grader", func(v *TaskV2) { v.Verification.GraderChecks = []string{"./..."} }, false},
		{"recursive-reference", func(v *TaskV2) { v.Verification.ReferenceChecks = []string{"./..."} }, false},
	} {
		t.Run(test.name, func(t *testing.T) {
			task := fuzzSchemaTask(t)
			test.mutate(&task)
			err := validateSchema(root, "task-v2.schema.json", task)
			if (err == nil) != test.valid {
				t.Fatalf("valid=%v want=%v: %v", err == nil, test.valid, err)
			}
		})
	}
}

func TestFuzzPlanYAMLKeepsNullPresenceAndUnknownFieldGuards(t *testing.T) {
	for _, test := range []struct{ name, document, want string }{
		{"explicit-null", "profile: mutation-tests\nfuzz: null\n", "verification.fuzz must be an object"},
		{"aliased-null", "profile: mutation-tests\nfuzz: &none null\n", "verification.fuzz must be an object"},
		{"merged-null", "<<: {fuzz: null}\nprofile: mutation-tests\n", "verification.fuzz must be an object"},
		{"unknown-plan-field", "profile: mutation-tests\nunknown: true\n", "field unknown not found"},
		{"unknown-limit-field", "profile: mutation-tests\nlimits: {unknown: 1}\n", "field unknown not found"},
		{"unknown-fuzz-field", "profile: fuzz-seed-tests\nfuzz: {grammar: hex-codec-v1, target: FuzzHex, correct_variants: [loop, empty], command: arbitrary}\n", "field command not found"},
		{"duplicate-profile", "profile: mutation-tests\nprofile: fuzz-seed-tests\n", "already defined"},
	} {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "plan.yaml")
			if err := os.WriteFile(path, []byte(test.document), 0600); err != nil {
				t.Fatal(err)
			}
			var plan VerificationPlan
			err := decodeStrictFile(path, &plan)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("want closed YAML diagnostic %q, got %v", test.want, err)
			}
		})
	}
}
