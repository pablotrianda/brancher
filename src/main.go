package main

import (
	"flag"
	"os"

	"github.com/pablotrianda/brancher/src/cmd"
)

func main() {
	cmd.Brancher(hasArgument())
}

func hasArgument() cmd.Cli {
	cli := cmd.Cli{}

	cli.FillDefaults()

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
		switch opt := string(os.Args[1]); opt {
		case ".":
			cli.BackToPreviousBranch = true
		case "s":
			cli.MakeStash = true
		}
	}

	return cli
}
