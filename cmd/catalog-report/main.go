package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"go-practice-bank/internal/adapters/exercism"
	"go-practice-bank/internal/catalog"
	"go-practice-bank/internal/reports"
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
	if source.ID != "exercism-go" || source.Adapter != "exercism-go-v1" || source.UsageMode != "adapt" {
		fatalf("catalog-report currently requires the registered exercism-go adapt source")
	}
	cacheRoot, err := sources.DefaultCacheRoot()
	if err != nil {
		fatalf("%v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	checkout, err := (sources.Syncer{}).Open(ctx, source, cacheRoot)
	if err != nil {
		fatalf("cached source is unavailable or invalid: %v", err)
	}
	config, err := exercism.LoadConfig(checkout.Directory)
	if err != nil {
		fatalf("%v", err)
	}
	catalogPath := "catalog/candidates/" + source.ID + ".jsonl"
	candidates, err := catalog.ReadJSONL(*root, catalogPath)
	if err != nil {
		fatalf("%v", err)
	}
	for _, candidate := range candidates {
		if candidate.SourceID != source.ID || candidate.SourceCommit != source.Commit || candidate.License != source.License.Expression || candidate.UsageMode != source.UsageMode {
			fatalf("candidate %q provenance does not match registry", candidate.Slug)
		}
	}
	summary, err := reports.Generate(*root, source, candidates, config)
	if err != nil {
		fatalf("%v", err)
	}
	fmt.Printf("source=%s commit=%s candidates=%d concept=%d practice=%d starter=%d tests=%d reference=%d generated=%d exact-provenance=%d warnings=%d incomplete=%d\n", source.ID, source.Commit, summary.Candidates, summary.Concept, summary.Practice, summary.WithStarter, summary.WithTests, summary.WithReference, summary.WithGeneratedData, summary.ExactProvenanceRefs, summary.WithWarnings, summary.Incomplete)
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "catalog-report: "+format+"\n", args...)
	os.Exit(1)
}
