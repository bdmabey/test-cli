package main

import (
	"github.com/bdmabey/test-cli/cmd"
	"github.com/bdmabey/test-cli/pkg/util/config"
)

func main() {
	config.CreateMainConfigPath()
	cmd.Execute()
}
