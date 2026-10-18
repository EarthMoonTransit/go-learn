package tasking

import (
	"fmt"
	"strings"
)

func verifyLearnerFuzz(runner executionRunner, loaded learnerPackage, tests []string, scope LearnerVerificationScope) LearnerVerificationResult {
	plan := loaded.data.TaskV2.Verification
	result := LearnerVerificationResult{Scope: scope, SubmissionKind: "tests"}
	if plan.Fuzz == nil {
		result.Summary = "finite native verification descriptor is missing"
		result.FailureCategory = "configuration"
		return result
	}
	// The common layout has validated go.mod against the fixed task module
	// policy. This is the same generated identity used by other learner profiles.
	base := fuzzRequest{variantRequest: variantRequest{TaskID: loaded.data.Task.ID, GoVersion: learnerGoVersion(loaded), TimeoutSeconds: plan.Limits.CommandTimeoutSeconds, Race: plan.RaceRequired, Tests: tests}, ModulePath: "task.local/" + loaded.data.Task.ID, Plan: *plan.Fuzz}
	var publicProblems []string
	publicCategory := "public-behavior"
	for index, variant := range fuzzCorrectVariants(loaded.layout) {
		request := base
		request.Implementation = variant.Files
		request.Vet = plan.Vet
		label := fmt.Sprintf("public behavior variant %d", index+1)
		set := repeatFuzz(runner, request, plan.Limits.DeterminismRuns, label)
		publicProblems = append(publicProblems, set.correctProblems(label)...)
		if len(set.Problems) > 0 {
			publicCategory = "native-eligibility"
		}
	}
	if len(publicProblems) > 0 {
		result.FailureCategory = publicCategory
		result.Summary = "submitted tests reject valid public behavior"
		if publicCategory == "native-eligibility" {
			result.Summary = "submitted tests did not complete eligible native checks"
		}
		if scope == LearnerPublicChecks {
			result.Diagnostic = sanitizeVariantDiagnostic([]byte(strings.Join(uniqueSorted(publicProblems), "\n")), nil, "", "", loaded.data.Task.ID)
		}
		return result
	}
	if scope == LearnerPublicChecks {
		result.Passed = true
		result.Summary = "public checks passed"
		return result
	}
	invalid, survivor := false, false
	for _, mutant := range loaded.layout.Mutants {
		request := base
		request.Implementation = mutant.Files
		set := repeatFuzz(runner, request, plan.Limits.DeterminismRuns, "private native checks")
		invalid = invalid || len(set.Problems) > 0
		survivor = survivor || !set.AllDetected
	}
	if invalid {
		result.Summary = "one or more private checks did not complete eligible native execution"
		result.FailureCategory = "native-eligibility"
		return result
	}
	if survivor {
		result.Summary = "one or more private defect models survived the submitted native seeds and invariants"
		result.FailureCategory = "hidden-verification"
		return result
	}
	result.Passed = true
	result.Summary = "full verification passed"
	return result
}
