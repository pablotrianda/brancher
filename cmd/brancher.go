package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

const GITCOMMAND = "git for-each-ref --sort=committerdate refs/heads/ --format='%(refname:short),'"
const ERROR_CREATE = "Error when tried to create a new branch"
const ERROR_CHANGE = "Error when tried to change to another branch"

func Brancher(hasArgument bool, branchName string) {
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
		runCommand("git checkout -b "+branchName, ERROR_CREATE)
	}
}

func changeBranch() {
	commandOutput := runCommand(GITCOMMAND, ERROR_CHANGE)

	if len(commandOutput) == 0 {
		fmt.Println("BRANCHER -- The current repo hasn't git branches.")
		return
	}

	var selectedBranch string

	known_branch := os.Args[1:]
	if len(known_branch) != 0 {
		selectedBranch = known_branch[0]
	} else {
		branches := getBranches(commandOutput)
		selectedBranch = getSelectedBranch(branches)
	}

	runCommand("git checkout "+selectedBranch, ERROR_CHANGE)
}

func runCommand(command string, errorMessage string) string {
	commandOutput, err := exec.Command("bash", "-c", command).CombinedOutput()
	if err != nil {
		color.Red("BRANCHER: " + errorMessage)
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
