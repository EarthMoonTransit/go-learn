package main

import (
	"flag"
	"fmt"
	"os"

	"go-practice-bank/internal/tasking"
)

func main() {
	root := flag.String("root", ".", "repository root")
	id := flag.String("id", "", "task id")
	toValue := flag.String("to", "", "target state: draft, in-review, ready, rejected, or retired")
	actor := flag.String("actor", "", "reviewer or author responsible for the transition")
	reason := flag.String("reason", "", "reviewable transition reason")
	date := flag.String("date", "", "transition date in YYYY-MM-DD; defaults to today")
	flag.Parse()

	to, err := tasking.ParseState(*toValue)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	moved, err := tasking.Promote(*root, tasking.PromoteOptions{
		ID: *id, To: to, Actor: *actor, Reason: *reason, Date: *date,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("moved %s to %s at %s\n", moved.ID, moved.State, moved.Path)
}
