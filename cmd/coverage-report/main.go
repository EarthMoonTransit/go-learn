package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"go-practice-bank/internal/coveragereport"
	"go-practice-bank/internal/tasking"
)

func main() { os.Exit(run(os.Args[1:], os.Stdout, os.Stderr)) }

func run(args []string, out, diagnostic io.Writer) int {
	flags := flag.NewFlagSet("coverage-report", flag.ContinueOnError)
	flags.SetOutput(diagnostic)
	root := flags.String("root", ".", "repository root")
	scope := flags.String("scope", "", "required numeric scope: full|foundation")
	checksPath := flags.String("checks", "", "required current task-check collection; no task execution")
	check := flags.Bool("check", false, "check fixed report artifacts without writing")
	if err := flags.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		return 2
	}
	if err := coveragereport.ValidateScope(*scope); err != nil {
		fmt.Fprintln(diagnostic, err)
		return 2
	}
	if *checksPath == "" || flags.NArg() != 0 {
		fmt.Fprintln(diagnostic, "-checks is required; positional arguments are unsupported")
		return 2
	}
	checks, err := tasking.LoadCheckCollection(*root, *checksPath)
	if err != nil {
		fmt.Fprintln(diagnostic, err)
		return 1
	}
	report, err := coveragereport.Generate(*root, *scope, checks, *check)
	if err != nil {
		fmt.Fprintln(diagnostic, err)
		return 1
	}
	fmt.Fprintf(out, "numeric coverage report current: %s, %d selected skills, semantic review %s\n", report.SelectedScope, len(report.SelectedSkillIDs), report.SemanticReviewStatus)
	return 0
}
