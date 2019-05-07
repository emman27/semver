package cli

import (
	"fmt"
	"os"

	"github.com/coreos/go-semver/semver"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "cli",
	Short: "Use semver in cli mode",
}

var greaterCmd = &cobra.Command{
	Use:   "greater",
	Short: "Compares if a version is greater or equal than another",
	Long: `Compares if a version is greater or equal than another. This is typically your usage case when you want to check if a particular version is acceptable
Effectively,
	> semver cli greater a b
will return
Usage:
	> semver cli greater 1.1.0 1.0.0
	> echo $?
	0
	> semver cli greater 1.1.0 1.2.0
	> echo $?
	1`,
	RunE: greater,
}

func init() {
	Command.AddCommand(greaterCmd)
}

func greater(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("wrong number of args given, need 2 arguments but got args: %v", args)
	}
	if semver.New(args[0]).Compare(*semver.New(args[1])) == -1 {
		fmt.Printf("%s is not greater than %s\n", args[0], args[1])
		os.Exit(1)
	}
	return nil
}
