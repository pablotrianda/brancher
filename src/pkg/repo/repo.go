package repo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	commandline "github.com/pablotrianda/brancher/src/pkg/commandLine"
	"github.com/pablotrianda/brancher/src/pkg/constans"
	"github.com/pablotrianda/brancher/src/pkg/db"
	"github.com/pablotrianda/brancher/src/pkg/prompt"
)

func ChangeBranch() {
	commandOutput := commandline.RunCommand(constans.GITCOMMAND, constans.ERROR_CHANGE)

	if len(commandOutput) == 0 {
		prompt.ShowAlert(constans.ERROR_NOT_BRANCHES, constans.FAIL_ALERT)
		return
	}

	var selectedBranch string

	known_branch := os.Args[1:]
	if len(known_branch) != 0 {
		selectedBranch = known_branch[0]
	} else {
		branches := getBranches(commandOutput)
		selectedBranch = prompt.GetSelectedBranch(branches)
	}

	commandline.RunCommand("git checkout "+selectedBranch, constans.ERROR_CHANGE)
}

func CreateANewBrach(branchName string) {
	prompt.ShowAlert("Createad a new branch called "+branchName, 2)
	confirm := false
	p := &survey.Confirm{
		Message: "Do you create a new branch?",
	}
	survey.AskOne(p, &confirm)
	if confirm {
		ChangeToBranch(branchName)
	}
}

func ChangeToBranch(branchName string) {
	checkoutCommand := fmt.Sprintf("git checkout %s", branchName)
	commandline.RunCommand(checkoutCommand, constans.ERROR_CHANGE)
}

func DeleteBranch(branchName string) {
	deleteMessage := fmt.Sprintf("Delete PERMANENTLY the selected branch? %s", branchName)
	prompt.ShowAlert(deleteMessage, constans.MAXIM_ALERT)

	// TODO group al  survey interactions
	confirm := false
	p := &survey.Confirm{
		Message: "Do you DELETE the branch?",
	}
	survey.AskOne(p, &confirm)
	if confirm {
		commandline.RunCommand("git branch -D "+branchName, constans.ERROR_DELETE_BRANCH)
		prompt.ShowAlert(branchName+" deleted successfully! ", constans.SUCCESS_ALERT)
	}
}

func SaveActualBranch() {
	actualName := commandline.RunCommand(constans.GIT_GET_NAME, constans.ERROR_SAVE_BRANCH)
	repoDir := commandline.RunCommand(constans.GIT_GET_DIR, constans.ERROR_SAVE_BRANCH)
	err := db.SaveBranch(getRepoName(), repoDir, actualName)
	if err != nil {
		prompt.ShowAlert("BRANCHER: Cant save the info", constans.FAIL_ALERT)
	}
}

func ToPreviousBranch() {
	prevBranch, err := db.GetPreviousBranchName(getRepoName())

	if err != nil {
		prompt.ShowAlert(constans.ERROR_NOT_BRANCHES, constans.FAIL_ALERT)
		return
	}
	commandline.RunCommand("git checkout "+prevBranch, constans.ERROR_CHANGE)

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
