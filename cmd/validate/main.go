package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"go-practice-bank/internal/validation"
)

func main() {
	root := flag.String("root", ".", "repository root")
	phase1 := flag.Bool("phase1", false, "allow an empty ready-task catalog")
	coverage := flag.Bool("coverage", false, "enforce curriculum minimums and 350..500 ready tasks")
	runCommands := flag.Bool("run-commands", false, "run validation_commands for ready tasks")
	flag.Parse()

	if *phase1 && *coverage {
		fmt.Fprintln(os.Stderr, "-phase1 and -coverage are mutually exclusive")
		os.Exit(2)
	}
	result, err := validation.ValidateRoot(*root, validation.Options{
		AllowEmpty:      *phase1,
		EnforceCoverage: *coverage,
		RunCommands:     *runCommands,
	})
	printSummary(result)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func printSummary(result validation.Result) {
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
	fmt.Printf("taxonomy: %d domains, %d assignable skills, %d deprecated stubs, minimum %d primary / %d total encounters\n", len(result.Curriculum.Domains), len(result.Coverage), deprecated, totalPrimary, totalRequired)
	fmt.Printf("sources: %d audited entries\n", len(result.Registry.Sources))
	fmt.Printf("candidate triage: %d reviewed, %d selected pilot seeds\n", result.ReviewedCandidates, result.SelectedSeeds)
	fmt.Printf("tasks: %d manifests, %d ready\n", len(result.Tasks), result.ReadyTasks)
	fmt.Println("coverage by domain (actual/required primary, actual/required total):")
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
		fmt.Printf("  %-12s %d/%d primary, %d/%d total\n", domain, v.ap, v.rp, v.at, v.rt)
	}
}
