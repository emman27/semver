package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "v0.1.0"

var Command = &cobra.Command{
	Use:   "version",
	Short: "Prints the current version number of semver",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}
