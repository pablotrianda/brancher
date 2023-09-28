package commandline

import (
	"os/exec"

	"github.com/pablotrianda/brancher/src/pkg/constans"
	"github.com/pablotrianda/brancher/src/pkg/prompt"
)

func RunCommand(command string, errorMessage string) string {
	commandOutput, err := exec.Command("bash", "-c", command).CombinedOutput()
	if err != nil {
		prompt.ShowAlert("BRANCHER: "+errorMessage, constans.FAIL_ALERT)
		return ""
	}
	return string(commandOutput)
}
