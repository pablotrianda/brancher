package cmd

import (
	"github.com/pablotrianda/brancher/src/pkg/repo"
	"github.com/spf13/cobra"
)

var changeCmd = &cobra.Command{
	Use:     "change",
	Short:   "Change branch",
	Aliases: []string{"c"},
	Long:    `This command will change the local branch`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := args[0]
		repo.SaveActualBranch()
		repo.ChangeToBranch(branchName)
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
}
