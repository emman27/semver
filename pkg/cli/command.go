package cli

import "github.com/spf13/cobra"

var Command = &cobra.Command{
	Use:   "cli",
	Short: "Use hugo in cli mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}
