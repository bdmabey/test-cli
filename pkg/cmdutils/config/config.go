package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Called on run of top level command. Ensures that the configuration directory is made.
func CreateMainConfigPath() error {
	if userPath, err := os.UserHomeDir(); err != nil {
		panic(fmt.Errorf("unable to locate user home dir: %w", err))
	} else {
		mainConfigPath := userPath + "/.test-cli/"
		if err := os.MkdirAll(mainConfigPath, 0666); err != nil {
			panic(fmt.Errorf("unable to create configuration directory in the user home directory: %w", err))
		} else {
			fmt.Println("Folder already created.")
		}
	}
	return nil
}

// Struct that contains the config path for the cmd directory.
type Config struct {
	configPath string
}

// Attempts to set the configPath to $HOME/.test-cli/
// Each command will somehow need to check and make sure their folder is made
// If it isn't create it and create a default config file.
// If it is cool, set the configPath to that.
func (o *Config) setCmdConfigPath(cmd string) string {
	// Checks to see if we can get the home directory of the user.
	if userPath, err := os.UserHomeDir(); err != nil {
		panic(fmt.Errorf("unable to locate user home dir: %w", err))
	} else {
		// need to ensure the directory is created. if not there create it.
		configPath := userPath + "/.test-cli/" + cmd + "/"
		if err := os.MkdirAll(configPath, 0666); err != nil {
			panic(fmt.Errorf("unable to create configuration directory in the home directory: %w", err))
		} else {
			fmt.Println("Folder already created.")
		}
		return configPath
	}
}

func (o *Config) newConfig(cmd string) *Config {
	return &Config{
		configPath: o.setCmdConfigPath(cmd),
	}
}

// Gets specific config requested by any command.
func GetConfig(filename string, cmd string) error {
	o := &Config{}
	o.newConfig(cmd)
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(o.configPath)
	return nil
}
