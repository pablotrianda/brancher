package prompt

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
	"github.com/pablotrianda/brancher/src/pkg/constans"
)

func ConfirmCreateConfig() bool {
	confirm := false
	prompt := &survey.Confirm{
		Message: "Do you a file config on $HOME/.config/brancher?",
	}
	survey.AskOne(prompt, &confirm)

	return confirm
}

func GetSelectedBranch(branches []string) string {
	var qs = []*survey.Question{
		{
			Name: "Branch",
			Prompt: &survey.Select{
				Message: "Choose a branch:",
				Options: branches,
				VimMode: true,
				Default: "master",
			},
		},
	}

	answers := struct {
		Branch string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		if err.Error() != "interrupt" {
			ShowAlert(err.Error(), constans.FAIL_ALERT)
			return ""
		}
	}

	return answers.Branch
}

func ShowAlert(message string, alertCode int) {
	switch alertCode {
	case 1:
		color.Red(message)
	case 2:
		color.Cyan(message)
	default:
		color.White(message)
	}
}
