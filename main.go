package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

const GitCommand = "git for-each-ref --sort=committerdate refs/heads/ --format='%(refname:short),'"

func main() {
	hasArgument, branchName := hasArgument()
	if hasArgument {
		createANewBrach(branchName)
	} else {
		changeBranch()
	}

}

func createANewBrach(branchName string) {
	fmt.Printf("Createad a new branch called %q\n", branchName)
	confirm := false
	prompt := &survey.Confirm{
		Message: "Do you create a new branch?",
	}
	survey.AskOne(prompt, &confirm)
	if confirm {
		runCommand("git checkout -b " + branchName)
	}
}

func changeBranch() {
	commandOutput := runCommand(GitCommand)

	if len(commandOutput) == 0 {
		fmt.Println("BRANCHER -- The current repo hasn't git branches.")
		return
	}

	branches := getBranches(commandOutput)
	selectedBranch := getSelectedBranch(branches)

	runCommand("git checkout " + selectedBranch)
}

func runCommand(command string) string {
	commandOutput, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(commandOutput)
}

func getBranches(commandOutput string) []string {
	var branches []string
	for _, s := range strings.Split(commandOutput, ",") {
		branches = append(branches, strings.TrimSpace(s))
	}
	return branches
}

func getSelectedBranch(branches []string) string {

	var qs = []*survey.Question{
		{
			Name: "Branch",
			Prompt: &survey.Select{
				Message: "Choose a branch:",
				Options: branches,
				VimMode: true,
				Default: "master",
			},
		},
	}

	answers := struct {
		Branch string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return answers.Branch
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
