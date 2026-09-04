package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"go-practice-bank/internal/sources"
)

func main() {
	root := flag.String("root", ".", "repository root containing sources/registry.yaml")
	sourceID := flag.String("source", "", "registered source id")
	allowConceptsOnly := flag.Bool("allow-concepts-only", false, "explicitly allow synchronizing a concepts-only source")
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
	cacheRoot, err := sources.DefaultCacheRoot()
	if err != nil {
		fatalf("%v", err)
	}
	syncer := sources.Syncer{}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()
	checkout, err := syncer.Sync(ctx, source, cacheRoot, *allowConceptsOnly)
	if err != nil {
		fatalf("%v", err)
	}
	printCheckout(checkout)
	if companion, ok := registry.ExercismCompanion(source); ok {
		checkout, err := syncer.Sync(ctx, companion, cacheRoot, *allowConceptsOnly)
		if err != nil {
			fatalf("sync companion %q: %v", companion.ID, err)
		}
		printCheckout(checkout)
	}
}

func printCheckout(checkout sources.Checkout) {
	fmt.Printf("source=%s commit=%s detached=%t checkout=%s manifest=%s\n", checkout.Source.ID, checkout.Manifest.GitHead, checkout.Manifest.DetachedHead, checkout.Directory, checkout.ManifestPath)
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "source-sync: "+format+"\n", args...)
	os.Exit(1)
}
