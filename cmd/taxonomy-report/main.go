package main

import (
	"flag"
	"fmt"
	"os"

	"go-practice-bank/internal/safeio"
	"go-practice-bank/internal/taxonomy"
)

func main() {
	root := flag.String("root", ".", "repository root")
	check := flag.Bool("check", false, "check without writing the generated migration report")
	flag.Parse()
	if flag.NArg() != 0 {
		fail(fmt.Errorf("usage: taxonomy-report [-root PATH] [-check]"))
	}
	if *check {
		if err := taxonomy.CheckMigrationReport(*root); err != nil {
			fail(err)
		}
		fmt.Println("taxonomy v4 migration report current")
		return
	}
	content, err := taxonomy.MigrationReport(*root)
	if err != nil {
		fail(err)
	}
	if _, err := safeio.AtomicWrite(*root, taxonomy.ReportPath, content, 1<<20); err != nil {
		fail(err)
	}
	fmt.Println("taxonomy v4 migration report generated")
}

func fail(err error) { fmt.Fprintln(os.Stderr, err); os.Exit(1) }
