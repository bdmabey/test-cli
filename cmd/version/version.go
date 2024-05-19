package version

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Version struct {
	version int
}

func NewVersion() *Version {
	return &Version{
		version: 12,
	}
}

func NewVersionCommand() *cobra.Command {
	o := NewVersion()
	o.SetupViper()

	var cmd = &cobra.Command{
		Use:     "version",
		Short:   "Prints the version",
		Aliases: []string{"v", "ver"},
		Long: `
Prints out the version.
Can use the flag --version to set the version.`,
		Run: func(cmd *cobra.Command, args []string) {
			o.RunCmd()
		},
	}

	cmd.Flags().IntVarP(&o.version, "version", "v", o.version, "Changes the version printed.")
	cmd.MarkFlagRequired("version")

	return cmd
}

func (o *Version) RunCmd() {
	viper.Set("version", 10)
	viper.Set("version.test", 12)
	fmt.Printf("Viper version is: %d\n", viper.GetInt("version.test"))
	viper.WriteConfigAs("/Users/brade/.test-cli/version/version.yaml")
	fmt.Printf("Version is now set to: %d", o.version)
}

func (o *Version) SetupViper() {
	viper.SetConfigName("version")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/Users/brade/.test-cli/version/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("Config file not found. Attempting to create...\n")
			if err := os.WriteFile("/Users/brade/.test-cli/version/version.yaml", []byte("---"), 0666); err != nil {
				panic(fmt.Errorf("could not create the file: %w", err))
			} else {
				fmt.Println("File created successfully")
			}
		} else {
			fmt.Println("File found but something else happened.")
		}
	}
}
