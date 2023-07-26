package cmd

import (
	"github.com/pablotrianda/brancher/src/pkg/repo"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete a branch",
	Aliases: []string{"D"},
	Long:    `This command will delete the local branch`,
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		branchName := args[0]
		repo.DeleteBranch(branchName)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
