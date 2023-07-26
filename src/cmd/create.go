package cmd

import (
	"github.com/pablotrianda/brancher/src/pkg/repo"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new branch",
	Aliases: []string{"D"},
	Long:    `This command will create a branch`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := args[0]
		repo.SaveActualBranch()
		repo.CreateANewBrach(branchName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
