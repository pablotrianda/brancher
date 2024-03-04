package cmd

import (
	"github.com/pablotrianda/brancher/src/pkg/repo"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "f",
	Short: "Fetch all branches from orgin",
	Long:  `This fetch all branches from orign repo`,
	Run: func(cmd *cobra.Command, args []string) {
		repo.FetchBranches()
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
