package config

import (
	"fmt"
	"os"

	"github.com/bdmabey/test-cli/pkg/util/debug"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Called on run of top level command. Ensures that the configuration directory is made.
func CreateMainConfigPath() error {
	if userPath, err := os.UserHomeDir(); err != nil {
		panic(fmt.Errorf("unable to locate user home dir: %w", err))
	} else {
		mainConfigPath := userPath + "/.test-cli/"
		if err := os.MkdirAll(mainConfigPath, 0755); err != nil {
			panic(fmt.Errorf("unable to create configuration directory in the user home directory: %w", err))
		}
	}
	return nil
}

// Struct that contains the config path for the cmd directory.
type Config struct {
	configPath string
}

// Creates a new config object
func newConfig(cmd *cobra.Command) *Config {
	return &Config{
		configPath: setCmdConfigPath(cmd),
	}
}

// Attempts to set the configPath to $HOME/.test-cli/
// Each command will somehow need to check and make sure their folder is made
// If it isn't create it and create a default config file.
// If it is cool, set the configPath to that.
func setCmdConfigPath(cmd *cobra.Command) string {
	// Checks to see if we can get the home directory of the user.
	if userPath, err := os.UserHomeDir(); err != nil {
		panic(fmt.Errorf("unable to locate user home dir: %w", err))
	} else {
		// need to ensure the directory is created. if not there create it.
		configPath := userPath + "/.test-cli/" + cmd.Use + "/"
		if err := os.MkdirAll(configPath, 0755); err != nil {
			panic(fmt.Errorf("unable to create configuration directory in the home directory: %w", err))
		} else {
			msg := cmd.Use + " folder exists/created."
			debug.Print(cmd, msg)
		}
		return configPath
	}
}

// Creates a new config file if one does not exist when get config is called.
func (o *Config) createConfigFile(cmd *cobra.Command) {
	if err := os.WriteFile(o.configPath+cmd.Use+".yaml", []byte("---"), 0755); err != nil {
		panic(fmt.Errorf("could not create configuration file for %s: %w", cmd.Use, err))
	} else {
		msg := cmd.Use + " file created successfully" + ".yaml"
		debug.Print(cmd, msg)
	}
}

// Gets specific config requested by any command.
func GetConfig(filename string, cmd *cobra.Command) (*viper.Viper, error) {
	v := viper.New()
	o := newConfig(cmd)
	v.SetConfigName(filename)
	v.SetConfigType("yaml")
	v.AddConfigPath(o.configPath)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			debug.Print(cmd, "Config file not found. Attempting to create...")
			o.createConfigFile(cmd)
		} else {
			debug.Print(cmd, "File found but something happened while loading it.")
		}
	} else {
		msg := cmd.Use + " configuration file loaded successfully."
		debug.Print(cmd, msg)
	}

	return v, nil
}
