package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"go-practice-bank/internal/tasking"
	"go-practice-bank/internal/validation"
)

func main() { os.Exit(run(os.Args[1:], os.Stdout, os.Stderr)) }

func run(args []string, out, diagnostic io.Writer) int {
	flags := flag.NewFlagSet("validate", flag.ContinueOnError)
	flags.SetOutput(diagnostic)
	root := flags.String("root", ".", "repository root")
	phase1 := flags.Bool("phase1", false, "allow an empty ready-task catalog")
	coverage := flags.Bool("coverage", false, "enforce curriculum minimums and 350..500 ready tasks")
	foundationCoverage := flags.Bool("foundation-coverage", false, "enforce v4 foundation encounter minimums using current checks; retain the independent full-bank target")
	runCommands := flags.Bool("run-commands", false, "run validation_commands for ready tasks")
	checksPath := flags.String("checks", "", "reuse a current task-check collection without executing task code")
	if err := flags.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		return 2
	}
	if flags.NArg() != 0 {
		fmt.Fprintln(diagnostic, "positional arguments are unsupported")
		return 2
	}

	if *coverage && *foundationCoverage {
		fmt.Fprintln(diagnostic, "-coverage and -foundation-coverage are mutually exclusive")
		return 2
	}
	if *phase1 && *coverage {
		fmt.Fprintln(diagnostic, "-phase1 and -coverage are mutually exclusive")
		return 2
	}
	if *phase1 && *foundationCoverage {
		fmt.Fprintln(diagnostic, "-phase1 and -foundation-coverage are mutually exclusive")
		return 2
	}
	if *checksPath != "" && *runCommands {
		fmt.Fprintln(diagnostic, "-checks and -run-commands are mutually exclusive")
		return 2
	}
	if *foundationCoverage && *checksPath == "" && !*runCommands {
		fmt.Fprintln(diagnostic, "-foundation-coverage requires a current -checks collection or explicit -run-commands")
		return 2
	}
	var checks *tasking.CheckCollection
	if *checksPath != "" {
		collection, err := tasking.LoadCheckCollection(*root, *checksPath)
		if err != nil {
			fmt.Fprintln(diagnostic, err)
			return 1
		}
		checks = &collection
	}
	result, err := validation.ValidateRoot(*root, validation.Options{
		AllowEmpty:                *phase1,
		EnforceCoverage:           *coverage,
		EnforceFoundationCoverage: *foundationCoverage,
		RunCommands:               *runCommands,
		Checks:                    checks,
	})
	printSummary(out, result)
	if err != nil {
		fmt.Fprintln(diagnostic, err)
		return 1
	}
	return 0
}

func printSummary(out io.Writer, result validation.Result) {
	totalPrimary, totalRequired := 0, 0
	deprecated := 0
	for _, domain := range result.Curriculum.Domains {
		for _, skill := range domain.Skills {
			if skill.Status == "deprecated" {
				deprecated++
			}
		}
	}
	for _, item := range result.Coverage {
		totalPrimary += item.RequiredPrimary
		totalRequired += item.RequiredTotal
	}
	fmt.Fprintf(out, "taxonomy: %d domains, %d assignable skills, %d deprecated stubs, minimum %d primary / %d total encounters\n", len(result.Curriculum.Domains), len(result.Coverage), deprecated, totalPrimary, totalRequired)
	fmt.Fprintf(out, "sources: %d audited entries\n", len(result.Registry.Sources))
	fmt.Fprintf(out, "candidate triage: %d reviewed, %d selected pilot seeds\n", result.ReviewedCandidates, result.SelectedSeeds)
	fmt.Fprintf(out, "tasks: %d manifests, %d ready\n", len(result.Tasks), result.ReadyTasks)
	if result.Curriculum.Version == 4 {
		fmt.Fprintf(out, "numeric selection: %s\n", result.SelectedScope)
		if result.NumericCoverage == nil {
			fmt.Fprintf(out, "numeric evaluation unavailable: %s\n", result.NumericUnavailable)
		} else {
			n := result.NumericCoverage
			for _, group := range []struct {
				name string
				data validation.NumericScopeCoverage
			}{{"foundation", n.Foundation}, {"specialist", n.Specialist}} {
				fmt.Fprintf(out, "%s minima: %s; deficits %d primary / %d total (%d skills)\n", group.name, numericStatus(group.data.MinimaMet), group.data.PrimaryDeficit, group.data.TotalDeficit, len(group.data.SkillIDs))
			}
			fmt.Fprintf(out, "full-bank numeric coverage: %s; ready target %d/%d..%d (%s)\n", numericStatus(n.FullNumericPassed), n.ReadyTarget.Actual, n.ReadyTarget.Min, n.ReadyTarget.Max, numericStatus(n.ReadyTarget.Met))
		}
		fmt.Fprintln(out, "semantic review: required modes, distinct task types and observable outcomes are not established by numeric counts")
	}
	fmt.Fprintln(out, "coverage by domain (actual/required primary, actual/required total):")
	type totals struct{ ap, rp, at, rt int }
	byDomain := map[string]totals{}
	for _, item := range result.Coverage {
		v := byDomain[item.DomainID]
		v.ap += item.ActualPrimary
		v.rp += item.RequiredPrimary
		v.at += item.ActualTotal
		v.rt += item.RequiredTotal
		byDomain[item.DomainID] = v
	}
	domains := make([]string, 0, len(byDomain))
	for domain := range byDomain {
		domains = append(domains, domain)
	}
	sort.Strings(domains)
	for _, domain := range domains {
		v := byDomain[domain]
		fmt.Fprintf(out, "  %-12s %d/%d primary, %d/%d total\n", domain, v.ap, v.rp, v.at, v.rt)
	}
}

func numericStatus(passed bool) string {
	if passed {
		return "pass"
	}
	return "fail"
}
