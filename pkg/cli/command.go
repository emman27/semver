package cli

import "github.com/spf13/cobra"

var Command = &cobra.Command{
	Use:   "cli",
	Short: "Use semvercmp in cli mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}