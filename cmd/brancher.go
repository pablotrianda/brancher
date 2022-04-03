package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func Brancher(hasArgument bool, branchName string, backToPreviousBranch bool) {
	if _, err := findConfiguration(CONFIG_DIR); err != nil {
		if conf := confirmCreateConfig(); !conf {
			showAlert(ERROR_CONFIG, FAIL_ALERT)
			return
		}
		if err := CreateNewConfig(); err != nil {
			showAlert(ERROR_CRATE_CONFIG, FAIL_ALERT)
			return
		}
	}

	if backToPreviousBranch {
		toPreviousBranch()
	} else {
		if hasArgument {
			createANewBrach(branchName)
		} else {
			changeBranch()
		}
		actualName := runCommand(GIT_GET_NAME, ERROR_SAVE_BRANCH)
		saveActualBranch(actualName)
	}

}

func confirmCreateConfig() bool {
	confirm := false
	prompt := &survey.Confirm{
		Message: "Do you a file config on $HOME/.config/brancher?",
	}
	survey.AskOne(prompt, &confirm)

	return confirm
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

func saveActualBranch(branchName string) {
	repoDir := runCommand(GIT_GET_DIR, ERROR_SAVE_BRANCH)
	err := SaveBranch(getRepoName(), repoDir, branchName)
	if err != nil {
		showAlert("BRANCHER: Cant save the info", FAIL_ALERT)
	}
}

func runCommand(command string, errorMessage string) string {
	commandOutput, err := exec.Command("bash", "-c", command).CombinedOutput()
	if err != nil {
		showAlert("BRANCHER: "+errorMessage, FAIL_ALERT)
		return ""
	}
	return string(commandOutput)
}

func getRepoName() string {
	dir, _ := os.Getwd()
	return filepath.Base(dir)
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

func toPreviousBranch() {
	prevBranch, err := GetPreviousBranchName(getRepoName())
	if err != nil {
		showAlert(ERROR_NOT_BRANCHES, FAIL_ALERT)
		return
	}
	runCommand("git checkout "+prevBranch, ERROR_CHANGE)

	saveActualBranch(prevBranch)
}
