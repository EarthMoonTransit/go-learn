package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"go-practice-bank/internal/safeio"
	"go-practice-bank/internal/tasking"
	"gopkg.in/yaml.v3"
)

func main() {
	root := flag.String("root", ".", "repository root")
	id := flag.String("id", "", "task id in any lifecycle state")
	all := flag.Bool("all", false, "check all tasks in deterministic id order")
	statesValue := flag.String("states", "", "comma-separated lifecycle states used with -all")
	output := flag.String("output", "-", "repository-relative YAML output path, or - for stdout")
	flag.Parse()

	if (*all && *id != "") || (!*all && *id == "") || (!*all && *statesValue != "") {
		fmt.Fprintln(os.Stderr, "use exactly one of -id or -all; -states requires -all")
		os.Exit(2)
	}
	var value any
	passed := true
	var err error
	if *all {
		var states []tasking.State
		if *statesValue != "" {
			for _, raw := range strings.Split(*statesValue, ",") {
				state, parseErr := tasking.ParseState(strings.TrimSpace(raw))
				if parseErr != nil {
					err = parseErr
					break
				}
				states = append(states, state)
			}
		}
		if err == nil {
			var collection tasking.CheckCollection
			collection, err = tasking.CheckAll(*root, states)
			value = collection
			for _, report := range collection.Reports {
				if !report.Passed {
					passed = false
				}
			}
		}
	} else {
		var report tasking.QualityReport
		report, err = tasking.Check(*root, *id)
		value = report
		passed = report.Passed
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	content, err := yaml.Marshal(value)
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
	if !passed {
		os.Exit(1)
	}
}
