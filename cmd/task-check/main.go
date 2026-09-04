package main

import (
	"flag"
	"fmt"
	"os"

	"go-practice-bank/internal/safeio"
	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

func main() {
	root := flag.String("root", ".", "repository root")
	id := flag.String("id", "", "task id in any lifecycle state")
	output := flag.String("output", "-", "repository-relative YAML output path, or - for stdout")
	flag.Parse()

	report, err := tasking.Check(*root, *id)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	content, err := yaml.Marshal(report)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *output == "-" {
		_, _ = os.Stdout.Write(content)
	} else if _, err := safeio.AtomicWrite(*root, *output, content, 2<<20); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !report.Passed {
		os.Exit(1)
	}
}
