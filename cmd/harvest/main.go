package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"go-practice-bank/internal/adapters/exercism"
	"go-practice-bank/internal/catalog"
	"go-practice-bank/internal/sources"
)

func main() {
	root := flag.String("root", ".", "repository root")
	sourceID := flag.String("source", "", "registered source id")
	flag.Parse()
	if !sources.ValidSourceID(*sourceID) {
		fatalf("invalid or missing -source")
	}
	registry, err := sources.LoadRegistry(*root)
	if err != nil {
		fatalf("%v", err)
	}
	source, err := registry.Find(*sourceID)
	if err != nil {
		fatalf("%v", err)
	}
	if source.UsageMode == "blocked" || source.UsageMode == "concepts-only" {
		fatalf("source %q with usage mode %q cannot be harvested", source.ID, source.UsageMode)
	}
	cacheRoot, err := sources.DefaultCacheRoot()
	if err != nil {
		fatalf("%v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	checkout, err := (sources.Syncer{}).Open(ctx, source, cacheRoot)
	if err != nil {
		fatalf("cached source is unavailable or invalid; run source-sync first: %v", err)
	}
	var candidates []catalog.Candidate
	switch source.Adapter {
	case "exercism-go-v1":
		var companion *sources.Source
		if registered, ok := registry.ExercismCompanion(source); ok {
			companion = &registered
		}
		candidates, err = (exercism.Adapter{}).Harvest(checkout.Directory, source, companion)
	default:
		fatalf("adapter %q is not implemented", source.Adapter)
	}
	if err != nil {
		fatalf("%v", err)
	}
	output := "catalog/candidates/" + source.ID + ".jsonl"
	changed, err := catalog.WriteJSONL(*root, output, candidates)
	if err != nil {
		fatalf("%v", err)
	}
	fmt.Printf("source=%s commit=%s candidates=%d changed=%t output=%s\n", source.ID, source.Commit, len(candidates), changed, output)
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "harvest: "+format+"\n", args...)
	os.Exit(1)
}
