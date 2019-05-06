package cmd

import (
	"fmt"
	"os"

	"github.com/emman27/semver/pkg/version"
	"github.com/emman27/semver/pkg/cli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "semver",
	Short: "semver is a comparer for semantic version numbers",
}

func Execute() {
	rootCmd.AddCommand(cli.Command)
	rootCmd.AddCommand(version.Command)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
