package toolchainpolicy

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRepositoryPolicyAndLegacyVersions(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	policy, err := Load(root)
	if err != nil {
		t.Fatal(err)
	}
	if err := ValidateRepository(root, policy); err != nil {
		t.Fatal(err)
	}
	if got, _ := policy.VersionForTask("defined-types-001", true); got != "1.24" {
		t.Fatalf("legacy version = %q", got)
	}
	if got, _ := policy.VersionForTask("future-task-001", false); got != "1.27" {
		t.Fatalf("current version = %q", got)
	}
}

func TestTaskGoModPolicyRejectsDependenciesAndToolchain(t *testing.T) {
	content := []byte("module task.local/example-001\n\ngo 1.27\ntoolchain go1.27.1\nrequire example.com/x v1.0.0\n")
	problems := ValidateGoMod(content, "task.local/example-001", "1.27", true)
	joined := strings.Join(problems, "\n")
	for _, wanted := range []string{"toolchain directive is forbidden", "directive require is not allowed"} {
		if !strings.Contains(joined, wanted) {
			t.Fatalf("missing %q in %v", wanted, problems)
		}
	}
}

func TestLoadRejectsUnknownPolicyField(t *testing.T) {
	root := t.TempDir()
	if err := os.Mkdir(filepath.Join(root, "config"), 0o755); err != nil {
		t.Fatal(err)
	}
	content := []byte("schema_version: 1\ngo_language_version: \"1.27\"\nci_toolchain: 1.27.x\nmodule_prefix: task.local/\nforbid_toolchain_directive: true\nlegacy_tasks: []\nunknown: true\n")
	if err := os.WriteFile(filepath.Join(root, "config", "toolchain.yaml"), content, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := Load(root); err == nil {
		t.Fatal("unknown field was accepted")
	}
}
