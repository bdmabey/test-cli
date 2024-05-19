package cmd

import (
	"fmt"
	"os"

	"github.com/bdmabey/test-cli/cmd/version"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "test-cli",
	Short: "Something short to test on",
	Long: `A fast and awesome little test app.
It works a little bit.
But not really.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(version.NewVersionCommand())
}
