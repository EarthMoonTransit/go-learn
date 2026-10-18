package tasking

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-practice-bank/internal/isolation"
)

func TestFuzzInitPublishesDeclaredDraftContract(t *testing.T) {
	root := fixtureRepository(t)
	if err := os.MkdirAll(filepath.Join(root, "docs"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := isolation.CopyFile(filepath.Join(repositoryRoot(t), filepath.FromSlash(fuzzCanonicalContractPath)), filepath.Join(root, filepath.FromSlash(fuzzCanonicalContractPath))); err != nil {
		t.Fatal(err)
	}
	located, err := Init(root, InitOptions{Slug: "finite-hex", PrimarySkill: "functions.signatures-results", VerificationProfile: "fuzz-seed-tests", Actor: "fixture-author", Date: "2026-10-18"})
	if err != nil {
		t.Fatal(err)
	}
	data, problems := loadPackage(root, located)
	if len(problems) > 0 {
		t.Fatal(problems)
	}
	if data.TaskV2.Activity != "fuzz-test-writing" || data.TaskV2.Submission.Kind != "tests" || data.TaskV2.Verification.Fuzz == nil || data.TaskV2.Verification.Fuzz.Target != "FuzzHex" {
		t.Fatalf("wrong finite scaffold: %+v", data.TaskV2)
	}
	for _, path := range []string{"starter", "solution/reference", "verification/reference", "verification/correct/loop", "verification/correct/empty", "verification/mutants/required-behavior"} {
		info, err := os.Stat(filepath.Join(located.Path, filepath.FromSlash(path)))
		if err != nil || !info.IsDir() {
			t.Fatalf("missing scaffold directory %s: %v", path, err)
		}
	}
	contract, err := os.ReadFile(filepath.Join(located.Path, filepath.FromSlash(fuzzSourceContractPath)))
	if err != nil {
		t.Fatal(err)
	}
	canonical, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(fuzzCanonicalContractPath)))
	if err != nil || !bytes.Equal(contract, canonical) {
		t.Fatalf("draft public contract differs: %v", err)
	}
	if !bytes.Contains(data.Prompt, []byte("](FUZZ_SOURCE.md)")) {
		t.Fatal("draft prompt does not link its learner contract")
	}
	before, err := TaskContentDigest(located.Path)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(located.Path, filepath.FromSlash(fuzzSourceContractPath)), append(contract, '\n'), 0600); err != nil {
		t.Fatal(err)
	}
	after, err := TaskContentDigest(located.Path)
	if err != nil || before == after {
		t.Fatalf("public grammar is not bound to authored content: %v", err)
	}
	if data.ProvenanceV2.Kind != "original" || len(data.ProvenanceV2.Sources) != 0 || data.ProvenanceV2.Project.ContentDigest != before {
		t.Fatal("new scaffold lost original creation/digest binding")
	}
}

func TestFuzzPublicContractIsReadonlyAndNotAHiddenImplementation(t *testing.T) {
	directory := t.TempDir()
	files := map[string]string{"TASK.md": "# Hex source contract\n\n[Contract](FUZZ_SOURCE.md)\n", "go.mod": "module task.local/contract-hex-001\n\ngo 1.27\n", "starter/fuzz_test.go": "package subject_test", "verification/FUZZ_SOURCE.md": "public finite source contract", "verification/reference/hex.go": "package subject\nfunc Encode([]byte)[]byte{return nil}\nfunc Decode([]byte)([]byte,bool){return nil,true}\n"}
	for name, body := range files {
		path := filepath.Join(directory, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(body), 0600); err != nil {
			t.Fatal(err)
		}
	}
	loaded := learnerPackage{data: packageData{Located: LocatedTask{Path: directory}, TaskV2: TaskV2{Submission: Submission{Kind: "tests"}, Verification: VerificationPlan{Profile: "fuzz-seed-tests"}}}, layout: packageLayout{Starter: []string{filepath.Join(directory, "starter/fuzz_test.go")}, VerificationReference: []string{filepath.Join(directory, "verification/reference/hex.go")}}}
	exported, err := buildLearnerWorkspaceFiles(loaded)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, file := range exported {
		if file.Path == fuzzPublicContractPath {
			found = true
			if file.Editable || string(file.Content) != "public finite source contract" {
				t.Fatal("public grammar changed or became editable")
			}
		}
		if strings.Contains(file.Path, "verification/") || strings.Contains(file.Path, "correct/") || strings.Contains(file.Path, "mutants/") {
			t.Fatalf("private verification path exported: %s", file.Path)
		}
	}
	if !found {
		t.Fatal("learner workspace omitted public finite grammar")
	}
}
