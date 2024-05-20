package version

import (
	"fmt"

	"github.com/bdmabey/test-cli/pkg/cmdutils/config"
	"github.com/spf13/cobra"
)

type Version struct {
	version int
}

func newVersion() *Version {
	return &Version{
		version: 12,
	}
}

func NewVersionCommand() *cobra.Command {
	o := newVersion()

	var cmd = &cobra.Command{
		Use:     "version",
		Short:   "Prints the version",
		Aliases: []string{"v", "ver"},
		Long: `
Prints out the version.
Can use the flag --version to set the version.`,
		Run: func(cmd *cobra.Command, args []string) {
			o.runCmd(cmd)
		},
	}

	cmd.Flags().IntVarP(&o.version, "version", "v", o.version, "Changes the version printed.")

	return cmd
}

// If the version flag is set then it will print that.
// If it is not set it will set version to what is in the viper config.
func (o *Version) runCmd(cmd *cobra.Command) {
	v, _ := config.GetConfig("version", "version")
	v.BindPFlag("version", cmd.Flags().Lookup("version"))
	if cmd.Flags().Lookup("version").Changed {
		fmt.Printf("The default version is: %d\n", o.version)
	} else {
		fmt.Printf("Version in file is: %s\n", v.GetString("version"))
	}
}
