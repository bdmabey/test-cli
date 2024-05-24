package util

import "github.com/spf13/cobra"

func AddVersionFlag(cmd *cobra.Command) {
	cmd.Flags().IntP("version", "v", 0, "Returns the version.")
}
