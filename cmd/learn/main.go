package main

import (
	"os"

	"go-practice-bank/internal/learner"
)

func main() {
	os.Exit(learner.RunCLI(learner.Config{}, os.Args[1:], os.Stdout, os.Stderr))
}
