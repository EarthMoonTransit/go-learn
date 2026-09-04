package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
			for _, checked := range collection.Tasks {
				if !checked.Report.Passed {
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
	} else if err := writeOutput(*root, *output, content); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if !passed {
		os.Exit(1)
	}
}

func writeOutput(root, path string, content []byte) error {
	if !filepath.IsAbs(path) {
		_, err := safeio.AtomicWrite(root, filepath.ToSlash(path), content, 2<<20)
		return err
	}
	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, 0o700); err != nil {
		return err
	}
	temporary, err := os.CreateTemp(directory, ".task-check-*")
	if err != nil {
		return err
	}
	temporaryPath := temporary.Name()
	defer os.Remove(temporaryPath)
	if err := temporary.Chmod(0o600); err != nil {
		_ = temporary.Close()
		return err
	}
	if _, err := temporary.Write(content); err != nil {
		_ = temporary.Close()
		return err
	}
	if err := temporary.Sync(); err != nil {
		_ = temporary.Close()
		return err
	}
	if err := temporary.Close(); err != nil {
		return err
	}
	return os.Rename(temporaryPath, path)
}
