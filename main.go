package main

import (
	"github.com/bdmabey/test-cli/cmd"
	"github.com/bdmabey/test-cli/pkg/cmdutils/config"
)

func main() {
	config.CreateMainConfigPath()
	cmd.Execute()
}
