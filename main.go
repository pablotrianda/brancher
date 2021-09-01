package main

import (
	"flag"

	"github.com/pablotrianda/brancher/cmd"
)

func main() {
	cmd.Brancher(hasArgument())
}

func hasArgument() (bool, string) {
	hasArgument := false
	nameNewBranch := ""

	newBranchArg := flag.String("n", "", "Name of a new branch")
	flag.Parse()

	if *newBranchArg != "" {
		hasArgument = true
		nameNewBranch = *newBranchArg
	}

	return hasArgument, nameNewBranch
}
