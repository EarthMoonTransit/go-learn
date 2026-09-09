package tasking

import (
	"fmt"
	"sort"
)

func (s *CheckSession) executeV2(c *checker, data packageData, layout packageLayout) {
	switch data.TaskV2.Verification.Profile {
	case "fixed-tests":
		s.executeV2Fixed(c, data, layout)
	case "mutation-tests":
		s.executeV2Mutation(c, data, layout)
	case "workspace-tests":
		s.executeV2Workspace(c, data, layout)
	}
}

func (s *CheckSession) executeV2Fixed(c *checker, data packageData, layout packageLayout) {
	plan := data.TaskV2.Verification
	base := variantRequest{TaskID: data.Task.ID, GoVersion: s.policy.GoLanguageVersion, TimeoutSeconds: plan.Limits.CommandTimeoutSeconds, Race: plan.RaceRequired}
	request := func(implementation, tests []string) variantRequest {
		value := base
		value.Implementation, value.Tests = implementation, tests
		return value
	}
	c.gate("starter-public", expectFailure("starter with public tests", s.runner.RunVariant(request(layout.Starter, layout.PublicTests))))
	c.gate("starter-full", expectFailure("starter with full tests", s.runner.RunVariant(request(layout.Starter, appendCopy(layout.PublicTests, layout.GraderTests)))))
	referenceRequest := request(layout.Reference, appendCopy(layout.PublicTests, layout.GraderTests))
	referenceProblems := repeatVariant(plan.Limits.DeterminismRuns, func() variantResult { return s.runner.RunVariant(referenceRequest) }, "reference with full tests")
	c.gate("reference-full", referenceProblems)
	c.report.Mutants.Required = len(plan.RequiredMutants)
	c.report.Mutants.Total = len(layout.Mutants)
	var mutantProblems []string
	for _, mutant := range layout.Mutants {
		result := s.runner.RunVariant(request(mutant.Files, appendCopy(layout.PublicTests, layout.GraderTests)))
		recordMutantResult(c, mutant.ID, result, &mutantProblems)
	}
	sort.Strings(c.report.Mutants.Survived)
	c.gate("mutant-coverage", mutantProblems)
	s.executeV2Auxiliary(c, data, func(race, vet bool) variantResult {
		value := referenceRequest
		value.Race, value.Vet = race, vet
		return s.runner.RunVariant(value)
	})
}

func (s *CheckSession) executeV2Mutation(c *checker, data packageData, layout packageLayout) {
	plan := data.TaskV2.Verification
	base := variantRequest{TaskID: data.Task.ID, GoVersion: s.policy.GoLanguageVersion, TimeoutSeconds: plan.Limits.CommandTimeoutSeconds, Race: plan.RaceRequired}
	request := func(implementation, tests []string) variantRequest {
		value := base
		value.Implementation, value.Tests = implementation, tests
		return value
	}
	starterReference := s.runner.RunVariant(request(layout.VerificationReference, layout.Starter))
	var starterProblems []string
	if !starterReference.Compiled || !starterReference.Passed {
		starterProblems = append(starterProblems, describeVariantFailure("starter tests against reference", starterReference))
	}
	c.gate("starter-public", starterProblems)

	survives := false
	var insufficiencyProblems []string
	for _, mutant := range layout.Mutants {
		result := s.runner.RunVariant(request(mutant.Files, layout.Starter))
		if !result.Compiled {
			insufficiencyProblems = append(insufficiencyProblems, describeVariantFailure("starter tests against mutant "+mutant.ID, result))
			continue
		}
		if result.Passed {
			survives = true
		}
	}
	if !survives {
		insufficiencyProblems = append(insufficiencyProblems, "starter tests already detect every required mutant; learner submission is not needed")
	}
	c.gate("starter-full", insufficiencyProblems)

	referenceRequest := request(layout.VerificationReference, layout.Reference)
	c.gate("reference-full", repeatVariant(plan.Limits.DeterminismRuns, func() variantResult { return s.runner.RunVariant(referenceRequest) }, "reference test submission"))
	c.report.Mutants.Required = len(plan.RequiredMutants)
	c.report.Mutants.Total = len(layout.Mutants)
	var mutantProblems []string
	for _, mutant := range layout.Mutants {
		result := s.runner.RunVariant(request(mutant.Files, layout.Reference))
		recordMutantResult(c, mutant.ID, result, &mutantProblems)
	}
	sort.Strings(c.report.Mutants.Survived)
	c.gate("mutant-coverage", mutantProblems)
	s.executeV2Auxiliary(c, data, func(race, vet bool) variantResult {
		value := referenceRequest
		value.Race, value.Vet = race, vet
		return s.runner.RunVariant(value)
	})
}

func (s *CheckSession) executeV2Workspace(c *checker, data packageData, layout packageLayout) {
	plan := data.TaskV2.Verification
	base := workspaceRequest{TaskID: data.Task.ID, GoVersion: s.policy.GoLanguageVersion, TimeoutSeconds: plan.Limits.CommandTimeoutSeconds, Race: plan.RaceRequired, Skeleton: layout.WorkspaceSkeleton}
	request := func(overlay, tests []mappedFile, checks []string) workspaceRequest {
		value := base
		value.Overlay, value.Tests, value.Checks = overlay, tests, checks
		return value
	}
	c.gate("starter-public", expectFailure("workspace starter with public tests", s.runner.RunWorkspace(request(layout.WorkspaceStarter, layout.WorkspacePublic, plan.PublicChecks))))
	c.gate("starter-full", expectFailure("workspace starter with full tests", s.runner.RunWorkspace(request(layout.WorkspaceStarter, appendMapped(layout.WorkspacePublic, layout.WorkspaceGrader), plan.GraderChecks))))
	referenceRequest := request(layout.WorkspaceReference, appendMapped(layout.WorkspacePublic, layout.WorkspaceGrader), plan.ReferenceChecks)
	c.gate("reference-full", repeatVariant(plan.Limits.DeterminismRuns, func() variantResult { return s.runner.RunWorkspace(referenceRequest) }, "workspace reference with full tests"))
	c.report.Mutants.Required = len(plan.RequiredMutants)
	c.report.Mutants.Total = len(layout.WorkspaceMutants)
	var mutantProblems []string
	for _, mutant := range layout.WorkspaceMutants {
		result := s.runner.RunWorkspace(request(mutant.Files, appendMapped(layout.WorkspacePublic, layout.WorkspaceGrader), plan.GraderChecks))
		recordMutantResult(c, mutant.ID, result, &mutantProblems)
	}
	sort.Strings(c.report.Mutants.Survived)
	c.gate("mutant-coverage", mutantProblems)
	s.executeV2Auxiliary(c, data, func(race, vet bool) variantResult {
		value := referenceRequest
		value.Race, value.Vet = race, vet
		return s.runner.RunWorkspace(value)
	})
}

func repeatVariant(runs int, run func() variantResult, label string) []string {
	var problems []string
	canonical := ""
	for index := 0; index < runs; index++ {
		result := run()
		if !result.Compiled || !result.Passed {
			problems = append(problems, describeVariantFailure(label, result))
			continue
		}
		if index == 0 {
			canonical = result.Canonical
		} else if result.Canonical != canonical {
			problems = append(problems, label+" has different observable test events across deterministic runs")
		}
	}
	return uniqueSorted(problems)
}

func recordMutantResult(c *checker, id string, result variantResult, problems *[]string) {
	if !result.Compiled {
		*problems = append(*problems, describeVariantFailure("mutant "+id, result))
		return
	}
	if result.Passed {
		c.report.Mutants.Survived = append(c.report.Mutants.Survived, id)
		*problems = append(*problems, fmt.Sprintf("mutant %s survives required verification", id))
		return
	}
	if result.FailureKind != variantFailureTest {
		*problems = append(*problems, describeVariantFailure("mutant "+id, result))
		return
	}
	c.report.Mutants.Detected++
}

func (s *CheckSession) executeV2Auxiliary(c *checker, data packageData, run func(bool, bool) variantResult) {
	if data.TaskV2.Verification.Vet {
		result := run(false, true)
		var problems []string
		if !result.Passed {
			problems = append(problems, describeVariantFailure("reference go vet", result))
		}
		c.gate("vet", problems)
	} else {
		c.notApplicable("vet", "verification plan does not require go vet")
	}
	if data.TaskV2.Verification.RaceRequired {
		result := run(true, false)
		if result.FailureKind == variantFailureRaceUnsupported {
			c.infrastructureUnsupported("race", result.Diagnostic)
		} else {
			var problems []string
			if !result.Compiled || !result.Passed {
				problems = append(problems, describeVariantFailure("reference go test -race", result))
			}
			c.gate("race", problems)
		}
	} else {
		c.notApplicable("race", "verification plan does not require the race detector")
	}
}

func appendMapped(left, right []mappedFile) []mappedFile {
	result := append([]mappedFile(nil), left...)
	return append(result, right...)
}
