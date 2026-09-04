package main

import (
	"flag"
	"fmt"
	"os"

	"go-practice-bank/internal/tasking"
)

func main() {
	root := flag.String("root", ".", "repository root")
	slug := flag.String("slug", "", "stable kebab-case task slug without numeric suffix")
	primary := flag.String("primary", "", "primary taxonomy skill id")
	source := flag.String("source", "", "registered source id")
	sourcePath := flag.String("source-path", "", "exact source repository path")
	taskType := flag.String("task-type", "implement", "task type")
	transformation := flag.String("transformation", "", "transformation type; defaults from source mode")
	actor := flag.String("actor", "task-init", "author recorded in lifecycle")
	date := flag.String("date", "", "lifecycle date in YYYY-MM-DD; defaults to today")
	flag.Parse()

	created, err := tasking.Init(*root, tasking.InitOptions{
		Slug: *slug, PrimarySkill: *primary, SourceID: *source, SourcePath: *sourcePath,
		TaskType: *taskType, TransformationType: *transformation, Actor: *actor, Date: *date,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("created %s draft at %s\n", created.ID, created.Path)
}
