package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/pablotrianda/brancher/src/pkg/repo"
)

var rootCmd = &cobra.Command{
	Use:   "brancher",
	Short: "Tool manage local git branches",
	Long: ` Tool to manage local git branches
                Made by @pablotrianda
                Complete documentation is available at http://github.com/pablotrianda/brancher`,
	Run: func(cmd *cobra.Command, args []string) {
		newBranchName, _ := cmd.Flags().GetString("name")

		if newBranchName != "" {
			create(newBranchName)
		} else {
			repo.ChangeBranch()
		}
	},
}

func Execute() {
	var branchName string

	rootCmd.Flags().StringVarP(&branchName, "name", "n", "", "Name of the branch")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
