package main

import (
	"flag"
	"os"

	"github.com/pablotrianda/brancher/cmd"
)


func main() {
	cmd.Brancher(hasArgument())
}

func hasArgument() cmd.Cli {
	cli := cmd.Cli{}

	cli.Fill_defaults()

	newBranchArg := flag.String("n", "", "Name of a new branch")
	deleteBranchArg := flag.String("D", "", "Name of branch to delete")
	flag.Parse()

	if *newBranchArg != "" {
		cli.HasArgument = true
		cli.NameNewBranch = *newBranchArg
		return cli
	}

	if *deleteBranchArg != "" {
		cli.HasArgument = true
		cli.NameDeleteBranch = *deleteBranchArg
		return cli
	}

	if len(os.Args) > 1 {
		cli.BackToPreviousBranch = string(os.Args[1]) == "."
	}

	return cli
}
