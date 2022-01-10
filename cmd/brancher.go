package cmd

import (
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func Brancher(hasArgument bool, branchName string) {
	if hasArgument {
		createANewBrach(branchName)
	} else {
		changeBranch()
	}
}

func createANewBrach(branchName string) {
	showAlert("Createad a new branch called "+branchName, 2)
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
		showAlert(ERROR_NOT_BRANCHES, FAIL_ALERT)
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
		showAlert("BRANCHER: "+errorMessage, FAIL_ALERT)
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
		showAlert(err.Error(), FAIL_ALERT)
		return ""
	}

	return answers.Branch
}
