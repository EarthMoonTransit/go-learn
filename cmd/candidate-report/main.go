package main

import (
	"fmt"
	"os"

	"go-practice-bank/internal/catalog"
	"go-practice-bank/internal/sources"
	"go-practice-bank/internal/triage"
)

func main() {
	root := "."
	if len(os.Args) == 3 && os.Args[1] == "-root" {
		root = os.Args[2]
	} else if len(os.Args) != 1 {
		fmt.Fprintln(os.Stderr, "usage: candidate-report [-root PATH]")
		os.Exit(2)
	}

	registry, err := sources.LoadRegistry(root)
	if err != nil {
		fail(err)
	}
	source, err := registry.Find("exercism-go")
	if err != nil {
		fail(err)
	}
	candidates, err := catalog.ReadJSONL(root, "catalog/candidates/exercism-go.jsonl")
	if err != nil {
		fail(err)
	}
	bundle, err := triage.Load(root)
	if err != nil {
		fail(err)
	}
	version, skills, err := triage.LoadSkillContracts(root)
	if err != nil {
		fail(err)
	}
	if err := triage.Validate(bundle, candidates, skills, version, source); err != nil {
		fail(err)
	}
	if err := triage.GenerateReports(root, bundle, skills); err != nil {
		fail(err)
	}
	fmt.Printf("candidate triage reports: %d reviewed, %d clusters, %d selected seeds\n", len(bundle.Reviews.Reviews), len(bundle.Clusters.Clusters), len(bundle.Seeds.SelectedSeeds))
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
