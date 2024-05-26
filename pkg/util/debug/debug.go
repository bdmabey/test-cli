package debug

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Print(cmd *cobra.Command, msg string) {
	if on, err := cmd.Flags().GetBool("debug"); err != nil {
		panic(fmt.Errorf("something happened: %w", err))
	} else {
		if on {
			fmt.Println(msg)
		}
	}
}
