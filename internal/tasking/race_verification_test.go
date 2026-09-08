package tasking

import (
	"path/filepath"
	"testing"
)

type raceObservingRunner struct {
	*fakeExecutionRunner
	modes []bool
}

func (r *raceObservingRunner) RunVariant(q variantRequest) variantResult {
	if !q.Vet {
		r.modes = append(r.modes, q.Race)
	}
	return r.fakeExecutionRunner.RunVariant(q)
}
func (r *raceObservingRunner) RunWorkspace(q workspaceRequest) variantResult {
	if !q.Vet {
		r.modes = append(r.modes, q.Race)
	}
	return r.fakeExecutionRunner.RunWorkspace(q)
}
func TestRaceRequiredAppliesToEveryV2TestInvocation(t *testing.T) {
	root := repositoryRoot(t)
	for _, profile := range []string{"fixed", "mutation", "workspace"} {
		for _, required := range []bool{false, true} {
			t.Run(profile+map[bool]string{false: "/ordinary", true: "/race"}[required], func(t *testing.T) {
				id := "synthetic-" + profile + "-v2-001"
				located := LocatedTask{ID: id, State: StateDraft, Path: filepath.Join(root, "internal/testdata/tasking/v2", profile, id)}
				data, problems := loadPackage(root, located)
				if len(problems) > 0 {
					t.Fatal(problems)
				}
				problems, layout := checkPackageContractV2(root, data)
				if len(problems) > 0 {
					t.Fatal(problems)
				}
				data.TaskV2.Verification.RaceRequired = required
				runner := &raceObservingRunner{fakeExecutionRunner: newFakeExecutionRunner()}
				session, err := newCheckSessionWithRunner(root, runner)
				if err != nil {
					t.Fatal(err)
				}
				defer session.Close()
				c := &checker{}
				session.executeV2(c, data, layout)
				if len(runner.modes) < 5 {
					t.Fatalf("missing starter/reference/mutant invocations: %v", runner.modes)
				}
				for i, mode := range runner.modes {
					if mode != required {
						t.Errorf("test invocation %d race=%v want %v", i, mode, required)
					}
				}
			})
		}
	}
}
func TestV2CollectionCannotCertifyRaceAwareVerification(t *testing.T) {
	root := fixtureRepository(t)
	collection, err := fakeSession(t, root).CheckAll([]State{StateDraft})
	if err != nil {
		t.Fatal(err)
	}
	collection.ContractVersion = "task-check/v2"
	if err := ValidateCheckCollection(root, collection, []State{StateDraft}); err == nil {
		t.Fatal("old execution semantics accepted")
	}
}
