package main

import (
	"github.com/pablotrianda/brancher/src/cmd"
	"github.com/pablotrianda/brancher/src/pkg/commandLine/utils"
	"github.com/pablotrianda/brancher/src/pkg/config"
)

func main() {
	go utils.GracefullShutdown()

	if !config.ValidateCurrentConfigurationAndAlert() {
		return
	}

	cmd.Execute()
}
