package cmd

import (
	"github.com/pablotrianda/brancher/src/pkg/repo"
)

func create(branchName string) {
	repo.SaveActualBranch()
	repo.CreateANewBrach(branchName)
}
