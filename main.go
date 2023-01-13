package main

import (
	"flag"
	"os"

	"github.com/pablotrianda/brancher/cmd"
)

func main() {
	cmd.Brancher(hasArgument())
}

func hasArgument() (bool, string, bool, string) {
	hasArgument := false
	nameNewBranch := ""
	nameDeleteBranch := ""
	backToPreviousBranch := false

	newBranchArg := flag.String("n", "", "Name of a new branch")
	deleteBranchArg := flag.String("D", "", "Name of branch to delete")
	flag.Parse()

	if *newBranchArg != "" {
		hasArgument = true
		nameNewBranch = *newBranchArg
		return hasArgument, nameNewBranch, backToPreviousBranch, nameDeleteBranch
	}

	if *deleteBranchArg != "" {
		hasArgument = true
		nameDeleteBranch = *deleteBranchArg
		return hasArgument, nameNewBranch, backToPreviousBranch, nameDeleteBranch
	}

	if len(os.Args) > 1 {
		backToPreviousBranch = string(os.Args[1]) == "."
	}

	return hasArgument, nameNewBranch, backToPreviousBranch, nameDeleteBranch
}
