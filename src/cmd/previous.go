package cmd

import (
	"github.com/pablotrianda/brancher/src/pkg/repo"
	"github.com/spf13/cobra"
)

var previousCmd = &cobra.Command{
	Use:   ".",
	Short: "Previous branch",
	Long:  `This command return to previous branch`,
	Run: func(cmd *cobra.Command, args []string) {
		repo.ToPreviousBranch()
		repo.SaveActualBranch()
	},
}

func init() {
	rootCmd.AddCommand(previousCmd)
}
