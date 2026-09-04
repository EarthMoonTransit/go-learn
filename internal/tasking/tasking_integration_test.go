//go:build taskintegration

package tasking

import (
	"os"
	"path/filepath"
	"testing"

	"go-practice-bank/internal/isolation"
)

func TestRealRunnerAdversarialMatrixUsesOneEnvironment(t *testing.T) {
	root := fixtureRepository(t)
	session, err := NewCheckSession(root)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()
	runner := session.runner.(*realExecutionRunner)

	assertStatus := func(name string, mutate func(), restore func(), failingGate string) {
		t.Helper()
		mutate()
		report, err := session.Check(syntheticID)
		restore()
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		assertGateFails(t, report, failingGate)
	}

	valid, err := session.Check(syntheticID)
	if err != nil || !valid.Passed || valid.ReferenceStatus != "pass" || valid.Mutants.Detected != 1 {
		t.Fatalf("valid real fixture: report=%#v err=%v", valid, err)
	}

	starterPath := taskFile(root, "starter/identity.go")
	starter := mustRead(t, starterPath)
	assertStatus("starter unexpectedly passes", func() {
		replaceInFile(t, starterPath, "return value + 1", "return value")
	}, func() { mustWrite(t, starterPath, starter) }, "starter-public")

	referencePath := taskFile(root, "solution/reference/identity.go")
	reference := mustRead(t, referencePath)
	assertStatus("reference fails", func() {
		replaceInFile(t, referencePath, "return value", "return value + 1")
	}, func() { mustWrite(t, referencePath, reference) }, "reference-full")

	mutantPath := taskFile(root, "solution/mutants/constant-zero/identity.go")
	mutant := mustRead(t, mutantPath)
	assertStatus("mutant survives", func() {
		mustWrite(t, mutantPath, []byte("package subject\n\nfunc Identity(value int) int { return value }\n"))
	}, func() { mustWrite(t, mutantPath, mutant) }, "mutant-coverage")

	publicPath := taskFile(root, "tests/public/identity_test.go")
	public := mustRead(t, publicPath)
	noOpPublic := mustRead(t, filepath.Join(repositoryRoot(t), "internal/testdata/tasking/adversarial/noop-public/identity_test.go"))
	assertStatus("no-op public", func() { mustWrite(t, publicPath, noOpPublic) }, func() { mustWrite(t, publicPath, public) }, "starter-public")

	graderPath := taskFile(root, "tests/grader/identity_test.go")
	grader := mustRead(t, graderPath)
	noOpGrader := mustRead(t, filepath.Join(repositoryRoot(t), "internal/testdata/tasking/adversarial/noop-grader/identity_test.go"))
	realisticMutant := mustRead(t, filepath.Join(repositoryRoot(t), "internal/testdata/tasking/adversarial/surviving-realistic-mutant/identity.go"))
	assertStatus("no-op grader", func() {
		mustWrite(t, graderPath, noOpGrader)
		mustWrite(t, mutantPath, realisticMutant)
	}, func() {
		mustWrite(t, graderPath, grader)
		mustWrite(t, mutantPath, mutant)
	}, "mutant-coverage")

	beforeStatic := runner.Metrics()
	referenceBinding := filepath.Join(repositoryRoot(t), "internal/testdata/tasking/adversarial/noop-reference-binding/solution_test.go")
	if err := isolation.CopyFile(referenceBinding, taskFile(root, "solution/reference/solution_test.go")); err != nil {
		t.Fatal(err)
	}
	staticReport, err := session.Check(syntheticID)
	if err != nil {
		t.Fatal(err)
	}
	assertGateFails(t, staticReport, "package-contract")
	if afterStatic := runner.Metrics(); afterStatic != beforeStatic {
		t.Fatalf("static rejection executed task code: before=%+v after=%+v", beforeStatic, afterStatic)
	}
	if err := os.Remove(taskFile(root, "solution/reference/solution_test.go")); err != nil {
		t.Fatal(err)
	}

	race := runner.RunVariant(variantRequest{
		TaskID: syntheticID, GoVersion: session.policy.GoLanguageVersion, Race: true,
		Implementation: []string{referencePath}, Tests: []string{publicPath, graderPath},
	})
	if !race.Compiled || !race.Passed {
		t.Fatal("real race variant failed")
	}

	legacyRoot := legacyFixtureRepository(t)
	legacySource := filepath.Join(legacyRoot, "tasks/ready/defined-types-001")
	legacyBefore := mustRead(t, filepath.Join(legacySource, "value.go"))
	legacyResults, err := runner.RunLegacy(legacySource, []string{"go test ./..."})
	if err != nil || len(legacyResults) != 1 || legacyResults[0].Err != nil {
		t.Fatalf("legacy execution: results=%#v err=%v", legacyResults, err)
	}
	if after := mustRead(t, filepath.Join(legacySource, "value.go")); string(after) != string(legacyBefore) {
		t.Fatal("legacy source checkout was mutated")
	}

	metrics := runner.Metrics()
	if metrics.Sessions != 1 || metrics.ChildGoRuns != 32 || metrics.ChildGoRuns != metrics.Workspaces || metrics.ChildGoRuns != metrics.VariantRuns+metrics.LegacyCommands {
		t.Fatalf("unexpected real execution topology: %+v", metrics)
	}
	entries, err := os.ReadDir(filepath.Join(runner.environment.Root, "workspaces"))
	if err != nil || len(entries) != metrics.Workspaces {
		t.Fatalf("workspace count=%d want=%d err=%v", len(entries), metrics.Workspaces, err)
	}
}

func mustRead(t *testing.T, path string) []byte {
	t.Helper()
	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return content
}

func mustWrite(t *testing.T, path string, content []byte) {
	t.Helper()
	if err := os.WriteFile(path, content, 0o644); err != nil {
		t.Fatal(err)
	}
}
