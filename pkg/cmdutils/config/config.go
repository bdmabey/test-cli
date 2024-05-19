package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	configPath string
}

// Attempts to set the configPath to $HOME/.test-cli/
func setConfigPath() string {
	// Checks to see if we can get the home directory of the user.
	if userPath, err := os.UserHomeDir(); err != nil {
		panic(fmt.Errorf("unable to locate user home dir: %w", err))
	} else {
		// need to ensure the directory is created. if not there create it.
		return userPath
	}
}

func newConfig() *Config {
	return &Config{
		configPath: setConfigPath(),
	}
}

// Gets specific config requested by any command.
func GetConfig(filename string, cmd string) error {
	o := newConfig()
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(o.configPath)
	return nil
}
