package main

import (
	"flag"
	"fmt"
	"os"

	"go-practice-bank/internal/safeio"
	"go-practice-bank/internal/tasking"
)

func main() {
	root := flag.String("root", ".", "repository root")
	markdownPath := flag.String("markdown", "reports/task-quality.md", "repository-relative Markdown report path")
	yamlPath := flag.String("yaml", "reports/task-quality.yaml", "repository-relative machine-readable report path")
	check := flag.Bool("check", false, "verify committed reports are current without writing")
	flag.Parse()

	report, err := tasking.BuildAggregateReport(*root)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := tasking.ValidateAggregateReport(*root, report); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	yamlContent, err := tasking.MarshalAggregateYAML(report)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	markdownContent := tasking.MarshalAggregateMarkdown(report)
	passed := true
	for _, gate := range report.Gates {
		if gate.Status == "fail" {
			passed = false
		}
	}
	if *check {
		if err := tasking.CheckAggregateArtifacts(*root, *markdownPath, *yamlPath, report); err != nil {
			fmt.Fprintln(os.Stderr, err, "; run task-report without -check")
			os.Exit(1)
		}
		fmt.Printf("task quality report current: %d packages\n", len(report.Tasks))
		if !passed {
			os.Exit(1)
		}
		return
	}
	markdownChanged, err := safeio.AtomicWrite(*root, *markdownPath, markdownContent, 4<<20)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	yamlChanged, err := safeio.AtomicWrite(*root, *yamlPath, yamlContent, 4<<20)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("task quality report: %d packages (markdown changed=%t, yaml changed=%t)\n", len(report.Tasks), markdownChanged, yamlChanged)
	if !passed {
		os.Exit(1)
	}
}
