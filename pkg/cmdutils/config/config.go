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
			fmt.Printf("%s folder created/already created.\n", mainConfigPath)
		}
	}
	return nil
}

// Struct that contains the config path for the cmd directory.
type Config struct {
	configPath string
}

// Creates a new config object
func newConfig(cmd string) *Config {
	return &Config{
		configPath: setCmdConfigPath(cmd),
	}
}

// Attempts to set the configPath to $HOME/.test-cli/
// Each command will somehow need to check and make sure their folder is made
// If it isn't create it and create a default config file.
// If it is cool, set the configPath to that.
func setCmdConfigPath(cmd string) string {
	// Checks to see if we can get the home directory of the user.
	if userPath, err := os.UserHomeDir(); err != nil {
		panic(fmt.Errorf("unable to locate user home dir: %w", err))
	} else {
		// need to ensure the directory is created. if not there create it.
		configPath := userPath + "/.test-cli/" + cmd + "/"
		if err := os.MkdirAll(configPath, 0666); err != nil {
			panic(fmt.Errorf("unable to create configuration directory in the home directory: %w", err))
		} else {
			fmt.Printf("%s folder is already created.\n", configPath)
		}
		return configPath
	}
}

// Creates a new config file if one does not exist when get config is called.
func (o *Config) createConfigFile(cmd string) {
	if err := os.WriteFile(o.configPath+cmd+".yaml", []byte("---"), 0666); err != nil {
		panic(fmt.Errorf("could not create configuration file for %s: %w", cmd, err))
	} else {
		fmt.Printf("%s file created successfully\n", (cmd + ".yaml"))
	}
}

// Gets specific config requested by any command.
func GetConfig(filename string, cmd string) (*viper.Viper, error) {
	v := viper.New()
	o := newConfig(cmd)
	v.SetConfigName(filename)
	v.SetConfigType("yaml")
	v.AddConfigPath(o.configPath)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config file not found. Attempting to create...")
			o.createConfigFile(cmd)
		} else {
			fmt.Println("File found but something happened while loading it.")
		}
	}

	return v, nil
}
