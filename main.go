package main

import (
	"flag"
	"os"

	"github.com/pablotrianda/brancher/cmd"
)

func main() {
	cmd.Brancher(hasArgument())
}

func hasArgument() (bool, string, bool) {
	hasArgument := false
	nameNewBranch := ""
	backToPreviousBranch := false

	newBranchArg := flag.String("n", "", "Name of a new branch")
	flag.Parse()

	if *newBranchArg != "" {
		hasArgument = true
		nameNewBranch = *newBranchArg
		return hasArgument, nameNewBranch, backToPreviousBranch
	}

	if len(os.Args) > 1 {
		backToPreviousBranch = string(os.Args[1]) == "."
	}

	return hasArgument, nameNewBranch, backToPreviousBranch
}
