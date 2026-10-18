package tasking

import (
	"encoding/hex"
	"fmt"
	"sort"

	"go-practice-bank/internal/fuzznative"
)

type fuzzRunSet struct {
	Results         []fuzznative.Result
	Problems        []string
	AllDetected     bool
	AllSurvived     bool
	RaceUnsupported bool
}

func repeatFuzz(runner executionRunner, request fuzzRequest, runs int, label string) fuzzRunSet {
	set := fuzzRunSet{AllDetected: true, AllSurvived: true}
	if runs < 2 || runs > 3 {
		return fuzzRunSet{Problems: []string{label + ": require 2..3 deterministic native runs"}}
	}
	canonical := ""
	for index := 0; index < runs; index++ {
		result := runner.RunFuzz(request)
		set.Results = append(set.Results, result)
		set.RaceUnsupported = set.RaceUnsupported || result.RaceUnsupported
		decoded, digestErr := hex.DecodeString(result.Canonical)
		if !result.Valid || digestErr != nil || len(decoded) != 32 {
			set.Problems = append(set.Problems, fmt.Sprintf("%s run %d is not eligible (%s): %s", label, index+1, result.Phase, result.Diagnostic))
		} else {
			if index == 0 {
				canonical = result.Canonical
			} else if canonical != result.Canonical {
				set.Problems = append(set.Problems, label+" has different observed native events or codec values across deterministic runs")
			}
		}
		if request.Vet && !result.Vetted {
			set.Problems = append(set.Problems, fmt.Sprintf("%s run %d did not complete declared go vet", label, index+1))
		}
		set.AllDetected = set.AllDetected && result.Valid && result.Detected && result.Assertions > 0
		set.AllSurvived = set.AllSurvived && result.Valid && !result.Detected && result.Assertions == 0
	}
	set.Problems = uniqueSorted(set.Problems)
	if len(set.Problems) > 0 {
		set.AllDetected = false
		set.AllSurvived = false
	}
	return set
}

func (set fuzzRunSet) correctProblems(label string) []string {
	problems := append([]string{}, set.Problems...)
	if len(problems) == 0 && !set.AllSurvived {
		diagnostic := ""
		for _, result := range set.Results {
			if result.Detected || result.Assertions > 0 {
				diagnostic = result.Diagnostic
				break
			}
		}
		problems = append(problems, label+" rejects valid public behavior: "+diagnostic)
	}
	return problems
}

func fuzzCorrectVariants(layout packageLayout) []correctVariantLayout {
	return append([]correctVariantLayout{{ID: "canonical", Files: layout.VerificationReference}}, layout.VerificationCorrect...)
}

func (s *CheckSession) executeV2Fuzz(c *checker, data packageData, layout packageLayout) {
	plan := data.TaskV2.Verification
	base := fuzzRequest{variantRequest: variantRequest{TaskID: data.Task.ID, GoVersion: s.policy.GoLanguageVersion, TimeoutSeconds: plan.Limits.CommandTimeoutSeconds, Race: plan.RaceRequired}, ModulePath: s.policy.ModulePrefix + data.Task.ID, Plan: *plan.Fuzz}
	request := func(implementation, tests []string, vet bool) fuzzRequest {
		value := base
		value.Implementation, value.Tests, value.Vet = implementation, tests, vet
		return value
	}
	var allSets []fuzzRunSet
	var vetProblems []string
	run := func(implementation, tests []string, vet bool, label string) fuzzRunSet {
		set := repeatFuzz(s.runner, request(implementation, tests, vet), plan.Limits.DeterminismRuns, label)
		allSets = append(allSets, set)
		if vet {
			for _, result := range set.Results {
				if !result.Vetted {
					vetProblems = append(vetProblems, label+" did not complete declared go vet")
				}
			}
		}
		return set
	}
	correct := func(tests []string, label string) []string {
		var problems []string
		for _, variant := range fuzzCorrectVariants(layout) {
			name := label + " against correct variant " + variant.ID
			problems = append(problems, run(variant.Files, tests, plan.Vet, name).correctProblems(name)...)
		}
		return uniqueSorted(problems)
	}
	starterProblems := correct(layout.Starter, "starter tests")
	c.gate("starter-public", starterProblems)
	var insufficiencyProblems []string
	if len(starterProblems) == 0 {
		survivor := false
		for _, mutant := range layout.Mutants {
			set := run(mutant.Files, layout.Starter, false, "starter tests against mutant "+mutant.ID)
			insufficiencyProblems = append(insufficiencyProblems, set.Problems...)
			survivor = survivor || set.AllSurvived
		}
		if !survivor {
			insufficiencyProblems = append(insufficiencyProblems, "starter tests need one same eligible surviving mutant across every repetition")
		}
	} else {
		insufficiencyProblems = append(insufficiencyProblems, "starter mutant checks skipped because starter tests do not satisfy public correctness")
	}
	c.gate("starter-full", uniqueSorted(insufficiencyProblems))
	refProblems := correct(layout.Reference, "reference test submission")
	c.gate("reference-full", refProblems)
	c.report.Mutants.Required = len(plan.RequiredMutants)
	c.report.Mutants.Total = len(layout.Mutants)
	var mutantProblems []string
	if len(refProblems) == 0 {
		for _, mutant := range layout.Mutants {
			set := run(mutant.Files, layout.Reference, false, "reference tests against mutant "+mutant.ID)
			if len(set.Problems) > 0 {
				mutantProblems = append(mutantProblems, set.Problems...)
			} else if !set.AllDetected {
				c.report.Mutants.Survived = append(c.report.Mutants.Survived, mutant.ID)
				mutantProblems = append(mutantProblems, "mutant "+mutant.ID+" survives the submitted native seeds and invariants")
			} else {
				c.report.Mutants.Detected++
			}
		}
	} else {
		mutantProblems = append(mutantProblems, "mutant detection skipped because reference tests do not satisfy public correctness")
	}
	sort.Strings(c.report.Mutants.Survived)
	c.gate("mutant-coverage", uniqueSorted(mutantProblems))
	if plan.Vet {
		c.gate("vet", uniqueSorted(vetProblems))
	} else {
		c.notApplicable("vet", "verification plan does not require go vet")
	}
	if plan.RaceRequired {
		var raceProblems []string
		unsupported := false
		for _, set := range allSets {
			raceProblems = append(raceProblems, set.Problems...)
			unsupported = unsupported || set.RaceUnsupported
		}
		if unsupported {
			c.infrastructureUnsupported("race", "native fuzz race execution is unsupported on the actual toolchain target")
		} else {
			c.gate("race", uniqueSorted(raceProblems))
		}
	} else {
		c.notApplicable("race", "verification plan does not require the race detector")
	}
}
