package cmd

import (
	"fmt"
	"os"

	"github.com/pablotrianda/brancher/src/pkg/repo"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "brancher",
	Short: "Tool manage local git branches",
	Long: ` Toll to manage local git branches
                Made by @pablotrianda
                Complete documentation is available at http://github.com/pablotrianda/brancher`,
	Run: func(cmd *cobra.Command, args []string) {
		repo.ChangeBranch()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
